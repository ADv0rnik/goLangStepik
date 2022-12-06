/*
Последовательность состоит из натуральных чисел и завершается числом 0.
Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.
*/

package main

import "fmt"

func main() {
	var a int
	max := 1
	count := 0

	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		if a == max {
			count += 1
		} else if a > max {
			max = a
			count = 1
		}
	}
	fmt.Println(count)
}
