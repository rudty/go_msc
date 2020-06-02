package main

import (
	"fmt"
	"testing"
)

func TestMessagePool(t *testing.T) {
	m := pool.Get().(*Message)
	m.Value = 5
	fmt.Println(m)
	pool.Put(m)
	pool.Put(&Message{Value: 9})
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}

func TestMessage(t *testing.T) {
	m := ObtainMessage()
	m.Value = "a"
	m.Recycle()

	for i := 0; i < 10; i++ {
		ObtainMessage()
	}
}
