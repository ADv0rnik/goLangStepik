/*
Напишите "функцию голосования", возвращающую то значение (0 или 1), которое среди значений ее аргументов x, y, z встречается чаще.
*/

package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println(vote(x, y, z))
}

func vote(x int, y int, z int) int {
	sum := x + y + z
	if sum == 0 || sum == 1 {
		return 0
	}
	return 1
}
