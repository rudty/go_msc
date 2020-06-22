package threading

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
)

// Thread OS thread
type Thread struct {
	wg sync.WaitGroup
	fn interface{}
}

// NewThread create thread with call function
func NewThread(fn interface{}) *Thread {
	t := Thread{
		fn: fn,
	}
	return &t
}

func (t *Thread) onThreadPanic() {
	r := recover()
	if r != nil {
		fmt.Println(r)
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

	reflect.ValueOf(t.fn).Call(in)
}

// Start start os thread with argument
func (t *Thread) Start(args ...interface{}) {
	t.wg.Add(1)
	go t.start(args)
}

// Join  blocks until end thread
func (t *Thread) Join() {
	t.wg.Wait()
}
