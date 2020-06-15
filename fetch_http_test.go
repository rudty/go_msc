package main

import (
	"testing"
)

func TestFetchResponseNil(t *testing.T) {
	defer func() {
		r := recover()
		if r != nil {
			t.Error("must not error")
		}
	}()
	var f *FetchResponse = nil
	if f.String() != "" {
		t.Error("must not nil")
	}
}

func TestJSONMap(t *testing.T) {
	// {
	// 	"userId": 1,
	// 	"id": 1,
	// 	"title": "delectus aut autem",
	// 	"completed": false
	//   }
	res, err := FetchBody("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		t.Error(err)
	}
	m := res.JSONMap()

	if v, ok := m["userId"].(float64); !ok || v != 1 {
		t.Error("userid: 1")
	}

	if v, ok := m["title"].(string); !ok || v != "delectus aut autem" {
		t.Error("title: delectus aut autem")
	}

	if v, ok := m["completed"].(bool); !ok || v != false {
		t.Error("completed: false")
	}

	s := res.String()
	if len(s) == 0 {
		t.Error("length > 0")
	}
}
