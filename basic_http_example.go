package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
	Author  string `json:"-"`

	// 값이 비어있으면 출력안함
	Date string `json:",omitempty"`

	// 문자열 출력 + 이름을 id 로 바꿈
	ID int `json:"id,string"`
}

func main() {
	port := 8080
	http.HandleFunc("/helloworld", helloworldHandler)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloworldHandler(w http.ResponseWriter, r *http.Request) {
	res := helloWorldResponse{
		Message: "HelloWorld",
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(res)
}
