package main

import "fmt"

func main() {
	var i int64

	fmt.Scan(&i)
	fmt.Println(convert(i))
}

func convert(x int64) uint16 {
	return uint16(x)
}
