package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, x string

	fmt.Scan(&x, &s)
	fmt.Println(strings.Index(x, s))
}
