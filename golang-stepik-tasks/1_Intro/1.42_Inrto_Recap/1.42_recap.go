//Дано натуральное число N. Выведите его представление в двоичном виде.

package main

import "fmt"

func main() {
	var value int

	fmt.Scan(&value)
	fmt.Printf("%b", value)

}
