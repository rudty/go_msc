package main

import (
	"fmt"
	"runtime"
)

func main() {
	ballast := make([]byte, 10<<30)
	_ = ballast
	runtime.GC()
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	fmt.Println(mem)
	var a int
	fmt.Scan(&a)
}
