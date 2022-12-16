package main

import (
	"fmt"
)

func Main() {
	test := 7
	fmt.Println(fib(test))
}
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
