/*
Поменяйте местами значения переменных на которые ссылаются указатели
*/
package main

import "fmt"

func test(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Println(*x1, *x2)
}

func main() {
	a := 2
	b := 3
	test(&a, &b)
}
