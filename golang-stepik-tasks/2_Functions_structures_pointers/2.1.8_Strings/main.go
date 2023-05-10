/*
На вход подается строка. Нужно определить, является ли она правильной или нет.
Правильная строка начинается с заглавной буквы и заканчивается точкой.
Если строка правильная - вывести Right иначе - вывести Wrong
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Print("Enter text: ") // Remove this line when apply your answer
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	rs := []rune(text)
	fmt.Println(string(rs)) // Remove this line when apply your answer

	if unicode.IsUpper(rs[0]) && string(rs[len(rs)-2]) == "." { // Change 2 to 1 when apply your answer
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
