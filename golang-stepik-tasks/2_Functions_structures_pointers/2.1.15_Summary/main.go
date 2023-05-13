/*
На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, res float64

	fmt.Scan(&a, &b)
	a = math.Pow(a, 2.0)
	b = math.Pow(b, 2.0)
	res = math.Sqrt(a + b)
	fmt.Println(res)
}
