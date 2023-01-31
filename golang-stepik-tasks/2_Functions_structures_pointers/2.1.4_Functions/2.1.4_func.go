package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
	var fib1, fib2 = 1, 1
	for i := 1; i < n; i++ {
		fib1, fib2 = fib2, fib2+fib1
	}
	return fib1
}
