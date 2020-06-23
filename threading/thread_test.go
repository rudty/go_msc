package threading

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateThread(t *testing.T) {
	a := NewThread(func() {
		fmt.Println("a")
	})
	a.Start()
	a.Join()
}

func TestJoin(t *testing.T) {
	end := false
	a := NewThread(func() {
		time.Sleep(100)
		end = true
	})
	a.Start()
	a.Join()

	if !end {
		t.Error("must end")
	}
}

func TestStatus(t *testing.T) {
	a := NewThread(func() {
		time.Sleep(100)
	})

	if a.GetStatus() != NEW {
		t.Error("status new")
	}
	a.Start()
	if a.GetStatus() != STARTED {
		t.Error("status start")
	}
	a.Join()
	if a.GetStatus() != TERMINATED {
		t.Error("status terminated")
	}
}

func TestPanic(t *testing.T) {
	a := NewThread(func() {
		panic("?")
	})

	a.Start()
	{
		defer func() {
			recover()
		}()
		a.Join()
	}

	if a.GetStatus() != TERMINATED {
		t.Error("status terminated")
	}
}
