package main

import (
	"fmt"
	"strings"
)

func main() {
	var first, second string

	fmt.Scan(&first, &second)
	for _, ch := range first {
		if strings.Contains(second, string(ch)) {
			fmt.Printf("%c ", ch)
		}
	}
}
