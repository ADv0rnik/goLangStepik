package main

import (
	"fmt"
	"strings"
)

func main() {
	var line string
	line_array := make([]string, 0)

	const SPLITER = "*"

	fmt.Scan(&line)

	for _, elem := range line {
		line_array = append(line_array, string(elem))
	}

	fmt.Println(strings.Join(line_array, SPLITER))
}
