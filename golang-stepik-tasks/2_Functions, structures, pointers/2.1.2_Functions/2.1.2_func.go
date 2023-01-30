/*
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.
*/

package main

import "fmt"

func main() {
	fmt.Println(minimumFromFour())
}

func minimumFromFour() int {
	var arr [4]int

	fmt.Scan(&arr[0], &arr[1], &arr[2], &arr[3])

	for _, value := range arr {
		if value < arr[0] {
			arr[0] = value
		}
	}
	return arr[0]
}
