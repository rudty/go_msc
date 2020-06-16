package cache

import (
	"fmt"
	"testing"
)

func TestFunctionCache(t *testing.T) {
	fc := NewFunctionCache(func() int {
		return 3
	})

	fmt.Println(fc.Call())
}
