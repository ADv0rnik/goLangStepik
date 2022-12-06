/*
Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
Программа в первой строке получает на вход число n - количество чисел в последовательности, во второй строке -- n чисел,
входящих в данную последовательность.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, c int
	var b string
	fmt.Scan(&a)

	sum := 0
	for i := 1; i <= a; i++ {
		fmt.Scan(&b)
		c, _ = strconv.Atoi(b)
		if len(b) == 2 && c%8 == 0 {
			sum += c
		}
	}
	fmt.Println(sum)
}
