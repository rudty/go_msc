// vcpkg grpc
// 
// 
// 복사 {vcpkgdir}\installed\x86-windows\tools\grpc .\bin\
// .\bin\protoc --grpc_out=. "--plugin=protoc-gen-grpc=.\bin\grpc_cpp_plugin.exe" helloworld.proto
// .\bin\protoc --cpp_out=. helloworld.proto
//
// 빌드 후 에러 발생 시
// Windows.h include 순서가 맞지 않아서 생기는 문제
// git clone grpc 해서 헤더를 최신 git 으로 교체 
// 
#pragma warning (disable:26812)
#pragma warning (disable:26495)

#include <iostream>
#include <memory>
#include <thread>
#include <string>
#include <ppltasks.h>
#include <grpcpp/grpcpp.h>
#include <grpcpp/health_check_service_interface.h>
#include <grpcpp/ext/proto_server_reflection_plugin.h>
#include <future>
#ifdef BAZEL_BUILD
#include "examples/protos/helloworld.grpc.pb.h"
#else
#include "helloworld.grpc.pb.h"
#endif
#ifdef _MSC_VER
#pragma comment(lib, "ws2_32.lib")
#endif
using grpc::Server;
using grpc::ServerAsyncResponseWriter;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::ServerCompletionQueue;
using grpc::Status;
using helloworld::HelloRequest;
using helloworld::HelloReply;
using helloworld::Greeter;

struct RpcRequestProcessable {
    virtual ~RpcRequestProcessable() = default;
    virtual void onProceed() = 0;
};

enum CallStatus { PROCESS, FINISH };


template<
    typename _Request, 
    typename _Response>
class MessageRpcRequest: public RpcRequestProcessable{
    CallStatus status_ = PROCESS;
    _Request req_;
    _Response res_;
    ServerContext ctx_;
    ServerAsyncResponseWriter<_Response> responder_;
public:

    MessageRpcRequest() : responder_(&ctx_) { std::cout << "MessageRpcRequest" << std::endl; }

    template <typename _Service>
    void registerServiceFunction(
        _Service* svc, 
        ServerCompletionQueue* cq, 
        void(_Service::*f)(::grpc::ServerContext* context, _Request* request, ::grpc::ServerAsyncResponseWriter<_Response>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void* tag)) {
        ((*svc).*f)(&ctx_, &req_, &responder_, cq, cq, this);
    }

    virtual void onCreate() = 0;

    void onProceed() final override {
        if (status_ == PROCESS) {
            status_ = FINISH;
            auto* newIam = onClone();
            newIam->onCreate();
            concurrency::create_task([this] {
                try {
                    onRequest(&req_, &res_);
                    responder_.Finish(res_, Status::OK, this);
                }
                catch (std::invalid_argument& e) {
                    responder_.FinishWithError(Status(grpc::StatusCode::INVALID_ARGUMENT, e.what()), this);
                }
                catch (std::exception& e) {
                    responder_.FinishWithError(Status(grpc::StatusCode::INTERNAL, e.what()), this);
                }
            });
        }
        else {
            onRelease();
        }
    }

    virtual void onRequest(_Request* req, _Response* res) = 0;
    virtual MessageRpcRequest<_Request, _Response>* onClone() = 0;
    virtual void onRelease() = 0;
};

class SayHelloRequest: public MessageRpcRequest<HelloRequest, HelloReply> {
    Greeter::AsyncService* service_;
    ServerCompletionQueue* cq_;
public:
    SayHelloRequest(Greeter::AsyncService* service, ServerCompletionQueue* cq)
        : service_(service), cq_(cq) {
        std::cout << "SayHelloRequest" << std::endl;
    }

    void onCreate() override {
        registerServiceFunction(service_, cq_, &Greeter::AsyncService::RequestSayHello);
    }

    void onRequest(HelloRequest* req, HelloReply* res) override {
        std::string prefix("Hello ");
        res->set_message(prefix + req->name());
        std::cout << "dd" << " " << GetCurrentThreadId() << std::endl;
    }

    SayHelloRequest* onClone() override {
        return new SayHelloRequest(service_, cq_);
    }

    void onRelease() override {
        delete this;
    }
};

class ServerImpl final: public Greeter::AsyncService {
    using _MyBase = Greeter::AsyncService;
    std::unique_ptr<ServerCompletionQueue> cq_;
    std::unique_ptr<Server> server_;
public:
    ~ServerImpl() {
        server_->Shutdown();
        // Always shutdown the completion queue after the server.
        cq_->Shutdown();
    }

    // There is no shutdown handling in this code.
    void Run() {
        std::string server_address("0.0.0.0:50051");

        ServerBuilder builder;
        // Listen on the given address without any authentication mechanism.
        builder.AddListeningPort(server_address, grpc::InsecureServerCredentials()); 
        // Register "service_" as the instance through which we'll communicate with
        // clients. In this case it corresponds to an *asynchronous* service.
        builder.RegisterService(this);
        // Get hold of the completion queue used for the asynchronous communication
        // with the gRPC runtime.
        cq_ = builder.AddCompletionQueue();
        // Finally assemble the server.
        server_ = builder.BuildAndStart();
        std::cout << "Server listening on " << server_address << std::endl;

        // Spawn a new CallData instance to serve new clients.
        (new SayHelloRequest(this, cq_.get()))->onCreate();
        HandleRpcs();
    }

private:
    // This can be run in multiple threads if needed.
    void HandleRpcs() {
        void* tag;  // uniquely identifies a request.
        bool ok;
        while (true) {
            // Block waiting to read the next event from the completion queue. The
            // event is uniquely identified by its tag, which in this case is the
            // memory address of a CallData instance.
            // The return value of Next should always be checked. This return value
            // tells us whether there i     s any kind of event or cq_ is shutting down.
            GPR_ASSERT(cq_->Next(&tag, &ok));
         
            static_cast<RpcRequestProcessable*>(tag)->onProceed();
        }
    }
};

int main(int argc, char** argv) {
    ServerImpl server;
    server.Run();

    return 0;
}