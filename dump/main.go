package main

import (
	"fmt"
	"time"
)

//set GOTRACEBACK=crash
//task manager => create dump file
//
var value int32

func addValue() {
	for {
		value++
		time.Sleep(10 * time.Millisecond)
	}
}
func subValue() {
	for {
		value--
		time.Sleep(10 * time.Millisecond)
	}
}
func main() {
	go subValue()
	go addValue()
	time.Sleep(1000 * time.Millisecond)
	//debug.PrintStack()
	//panic("hello world")

	var a int
	fmt.Scan(&a)
}
