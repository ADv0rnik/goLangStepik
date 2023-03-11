/*
Напишите функцию, которая умножает значения на которые ссылаются два указателя и выводит получившееся произведение в консоль
*/

package main

import "fmt"

func test(x1 *int, x2 *int) {
	fmt.Println(*x1 * *x2)
}

func main() {
	a := 2
	b := 2
	test(&a, &b)

}
