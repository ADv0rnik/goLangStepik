/*
Дано натуральное число, выведите его последнюю цифру.
*/

package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	fmt.Println(a % 10)
}