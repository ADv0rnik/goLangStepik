package main

import "fmt"

func main() {
	var a, b float32
	var avg float32

	fmt.Scan(&a, &b)
	avg = (a + b) / 2.0
	fmt.Println(avg)
}
