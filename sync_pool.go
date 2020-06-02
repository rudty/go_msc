package main

import (
	"sync"
)

var pool = sync.Pool{
	New: func() interface{} {
		return &Message{}
	},
}

// Message android 메세지 비슷한 메시지 전달 핸들러
type Message struct {
	Value interface{}
	What  int
	Arg1  int
	Arg2  int
}

// Recycle 해당 메세지를 메모리 풀에 넣습니다.
// 다시 사용하지 마십시오
func (m *Message) Recycle() {
	pool.Put(m)
}

// ObtainMessage 새로운 메세지를 반환합니다.
func ObtainMessage() *Message {
	for {
		msg, ok := pool.Get().(*Message)
		if ok {
			return msg
		}
	}
}
