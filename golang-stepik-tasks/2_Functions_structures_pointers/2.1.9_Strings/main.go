/*
На вход подается строка. Нужно определить, является ли она палиндромом.
Если строка является палиндромом - вывести Палиндром иначе - вывести Нет.
Все входные строчки в нижнем регистре.
*/

package main

import "fmt"

func main() {
	var text string
	fmt.Scan(&text)
	s := IsPalindrome(text)
	if s {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}

func IsPalindrome(s string) bool {
	var reversStr string

	for _, char := range s {
		reversStr = string(char) + reversStr
	}

	return s == reversStr
}
