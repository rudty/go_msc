package threading

import (
	"fmt"
	"testing"
)

func TestCreateThread(t *testing.T) {
	a := NewThread(func() {
		fmt.Println("a")
	})
	a.Start()
	a.Join()
}
