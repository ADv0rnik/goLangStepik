/*
Найдите количество минимальных элементов в последовательности
*/

package main

import "fmt"

func main() {
	var n, count int

	fmt.Scan(&n)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	min := a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
	}
	for _, value := range a {
		if value == min {
			count++
		}
	}
	fmt.Println(count)
}
