/*
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n,
что φn=A. Если А не является числом Фибоначчи, выведите число -1.
*/

package main

import "fmt"

func main() {
	var fib1, fib2 int = 1, 1
	var val int

	fmt.Scan(&val)

	for i := 1; i <= val; i++ {
		fib1, fib2 = fib2, fib1+fib2

		if fib2 == val {
			fmt.Println(i + 2)
			return
		}
	}
	fmt.Println(-1)
}
