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

func TestCreateThreadWithArgument(t *testing.T) {
	a := NewThread(func(a int, b string) {
		fmt.Println(a, b)
	})
	a.Start(1, "2")
	a.Join()
}

func TestCreateThreadWithNoFunc(t *testing.T) {
	defer func() {
		recover()
	}()
	NewThread(3)
	t.Error("must panic")
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

func TestThreadStartCallTwice(t *testing.T) {
	defer func() {
		recover()
	}()
	a := NewThread(func() {
		time.Sleep(100)
	})
	a.Start()
	a.Start()
	t.Error("must panic")
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
