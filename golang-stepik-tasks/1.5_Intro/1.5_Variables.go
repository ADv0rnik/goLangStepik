/*
Петя торопился в школу и неправильно написал программу,
которая сначала находит квадраты двух чисел, а затем их суммирует. Исправьте его программу.
*/

package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println(a*a + b*b)

}