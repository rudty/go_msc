package main

import (
	"fmt"
	"runtime"
)

func main() {
	ballast := make([]byte, 10<<30)
	_ = ballast
	runtime.GC()
	var a int
	fmt.Scan(&a)
}
