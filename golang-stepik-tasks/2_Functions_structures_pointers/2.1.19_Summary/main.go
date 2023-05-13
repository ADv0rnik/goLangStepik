package main

import (
	"fmt"
	"math"
)

var k, p, v float64 = 1296, 6, 6

func main() {
	fmt.Println(T())
}

func T() float64 {
	const variable = 6
	t := variable / W()
	return t
}

func W() float64 {
	w := math.Sqrt(k / M())
	return w
}

func M() float64 {
	m := p * v
	return m
}
