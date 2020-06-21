package main

import (
	"fmt"
)

func main() {
	ballast := make([]byte, 10<<30)
	_ = ballast
	var a int
	fmt.Scan(&a)
}
