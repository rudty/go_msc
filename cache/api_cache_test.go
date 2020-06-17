package cache

import (
	"fmt"
	"testing"
)

func TestFunctionCacheArg0(t *testing.T) {
	fc := NewFunctionCache(func() int {
		return 3
	})

	fmt.Println(fc.Call())
}

func TestFunctionCacheArg1(t *testing.T) {
	mustCallOnce := 0
	fc := NewFunctionCache(func(a int) int {
		mustCallOnce++
		return a + 1
	})

	if fc.Call(1) != 2 {
		t.Error("return 2")
	}

	if fc.Call(1) != 2 {
		t.Error("return 2")
	}

	if mustCallOnce != 1 {
		t.Error("cache error")
	}
}
