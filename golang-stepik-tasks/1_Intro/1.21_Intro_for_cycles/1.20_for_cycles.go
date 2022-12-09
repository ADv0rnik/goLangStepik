package main

import "fmt"

func main() {
	var x, p, y int

	fmt.Scan(&x, &p, &y)
	for year := 1; x < y; year++ {
		x = x + (x * p / 100)
		if x >= y {
			fmt.Println(year)
		}
	}
}
