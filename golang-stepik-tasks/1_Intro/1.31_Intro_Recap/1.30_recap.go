package main

import "fmt"

func main() {
	var n, h, m int

	fmt.Scan(&n)

	h, m = n/3600, n/60
	min := m - h*60
	fmt.Printf("It is %d hours %d minutes.", h, min)
}
