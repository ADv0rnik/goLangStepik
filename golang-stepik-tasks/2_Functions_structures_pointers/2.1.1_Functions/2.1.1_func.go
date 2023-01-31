/*
Напишите функцию f(), которая будет принимать строку text и выводить ее (печатать на экране).
*/

package main

import "fmt"

func main() {
	var text string

	fmt.Scan(&text)
	f(text)

}

func f(text string) {
	fmt.Println(text)
}
