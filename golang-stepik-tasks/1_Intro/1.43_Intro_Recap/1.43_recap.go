//Из натурального числа удалить заданную цифру.

package main

import "fmt"

func main() {
	var value, to_find int
	result := make([]int, 0)

	fmt.Scan(&value)
	fmt.Scan(&to_find)

	for value > 0 {
		last := value % 10
		if last != to_find {
			result = append(result, last)
		}
		value /= 10
	}
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i])
	}
}
