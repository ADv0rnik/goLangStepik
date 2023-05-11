package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, x string

	fmt.Scan(&s)
	for _, char := range s {
		if strings.Count(s, string(char)) == 1 {
			x += string(char)
		}
	}
	fmt.Println(x)
}
