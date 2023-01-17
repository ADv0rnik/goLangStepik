package main

import "fmt"

func main() {
	var n, sum int

	fmt.Scan(&n)
	res := getSum(n, sum)
	fmt.Println(getSum(res, sum))
}

func getSum(num int, sum int) int {
	if num < 10 {
		return sum + num
	} else {
		sum += num % 10
		return getSum(num/10, sum)
	}
}
