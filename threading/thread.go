package threading

import (
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
)

// ThreadStatus indicates the state of the thread
type ThreadStatus int32

const (
	// NEW create thread
	NEW ThreadStatus = iota

	// STARTED thread is started. Start() function was called
	STARTED

	// TERMINATED thread is end
	TERMINATED
)

// Thread OS thread
type Thread struct {
	runPanic     interface{}
	threadStatus int32
	wg           sync.WaitGroup
	fn           reflect.Value
}

// NewThread create thread with call function
func NewThread(fn interface{}) *Thread {
	r := reflect.ValueOf(fn)
	if r.Kind() != reflect.Func {
		panic("must function")
	}

	t := Thread{
		fn: r,
	}
	return &t
}

func (t *Thread) onThreadPanic() {
	r := recover()
	if r != nil {
		t.runPanic = r
	}
}

func mapToReflectValue(a []interface{}) []reflect.Value {
	r := make([]reflect.Value, len(a))
	for i := 0; i < len(a); i++ {
		r[i] = reflect.ValueOf(a[i])
	}
	return r
}

func (t *Thread) start(args []interface{}) {
	defer t.wg.Done()
	defer t.onThreadPanic()

	in := mapToReflectValue(args)

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	defer atomic.StoreInt32(&t.threadStatus, int32(TERMINATED))

	t.fn.Call(in)
}

// Start start os thread with argument
func (t *Thread) Start(args ...interface{}) {
	if t.threadStatus != int32(NEW) {
		panic("thread already started")
	}
	atomic.StoreInt32(&t.threadStatus, int32(STARTED))
	t.wg.Add(1)
	go t.start(args)
}

// Join blocks until end thread
func (t *Thread) Join() {
	t.wg.Wait()
	if t.runPanic != nil {
		panic(t.runPanic)
	}
}

// GetStatus return state this thread
func (t *Thread) GetStatus() ThreadStatus {
	return ThreadStatus(t.threadStatus)
}
