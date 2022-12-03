package main

import "fmt"

func main() {
	var num int32

	fmt.Scan(&num)
	first := num / 100
	second := num / 10 % 10
	third := num % 10

	if first != second && second != third && first != third {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
