/*
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.
*/

package main

import "fmt"

func main() {
	fmt.Println(minimumFromFour())
}

func minimumFromFour() int {
	var a int
	arr := make([]int, 0)

	for i := 0; i < 4; i++ {
		fmt.Scan(&a)
		arr = append(arr, a)
	}

	min := arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min
}
