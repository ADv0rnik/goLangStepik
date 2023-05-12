package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	res, err := Devide(a, b)
	if err != nil {
		fmt.Println("An error occured durin the operation")
	} else {
		fmt.Println(res)
	}
}

func Devide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't divide '%d' by zero", a)
	}
	return a / b, nil
}
