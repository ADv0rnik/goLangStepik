package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number string
	var firstNum int

	fmt.Scan(&number)
	firstNum, _ = strconv.Atoi(number[0:1])
	fmt.Println(firstNum)
}
