/*
На вход дается строка, из нее нужно сделать другую строку,
оставив только нечетные символы (считая с нуля)
*/

package main

import (
	"fmt"
)

func main() {
	var s, x string

	fmt.Scan(&s)
	for i, char := range s {
		if i%2 != 0 {
			x += string(char)
		}
	}
	fmt.Println(x)
}
