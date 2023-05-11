package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string

	fmt.Scan(&s)
	if len(s) >= 5 {
		for i, char := range s {
			if !IsValid(char) {
				fmt.Println("Wrong password")
				break
			} else if i == len(s)-1 && IsValid(char) {
				fmt.Println("Ok")
			}
		}
	} else {
		fmt.Println("Wrong password")
	}
}

func IsValid(char rune) bool {
	return unicode.IsDigit(char) || unicode.Is(unicode.Latin, char)
}
