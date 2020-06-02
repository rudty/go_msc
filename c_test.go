package main

import (
	"fmt"
	"testing"
)

func TestCLang(t *testing.T) {
	f := foo()
	fmt.Println(f)
}

func TestInt(t *testing.T) {
	printInt(3)
}
