package threading

import "testing"

func TestThreadID(t *testing.T) {
	if GetThreadID() <= 0 {
		t.Error("thread id must > 0")
	}
}
