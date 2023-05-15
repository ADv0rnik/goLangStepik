package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	var s string = "%^80"
	var s2 string = "hhhhh20&&&&nd"

	fmt.Println(adding(s, s2))

}

func adding(x string, y string) int64 {
	x_str := cleanString(x)
	y_str := cleanString(y)

	x1, _ := strconv.Atoi(x_str)
	y2, _ := strconv.Atoi(y_str)
	return int64(x1 + y2)
}

func cleanString(str string) string {
	var dig_str string

	for _, elem := range str {
		if unicode.IsDigit(elem) {
			dig_str += string(elem)
		}
	}
	return dig_str
}
