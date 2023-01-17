package main

import "fmt"

func main() {
	var n, a int
	count := 0

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a == 0 {
			count++
		}
	}
	fmt.Println(count)
}

