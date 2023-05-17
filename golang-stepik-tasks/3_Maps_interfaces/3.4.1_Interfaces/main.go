package main

import (
	// пакет используется для проверки ответа, не удаляйте его
	"bufio"
	"fmt" // пакет используется для проверки ответа, не удаляйте его
	"os"  // пакет используется для проверки ответа, не удаляйте его
)

func main() {
	value1, value2, operation := readTask()

	var x1, x2 float64

	switch value1.(type) {
	case float64:
		x1 = value1.(float64)
	default:
		printError(value1)
		return
	}

	switch value2.(type) {
	case float64:
		x1 = value2.(float64)
	default:
		printError(value2)
		return
	}

	switch operation.(type) {
	case string:
		switch operation.(string) {
		case "+":
			fmt.Printf("%.4f", x1+x2)
			return
		case "-":
			fmt.Printf("%.4f", x1-x2)
			return
		case "*":
			fmt.Printf("%.4f", x1*x2)
			return
		case "/":
			fmt.Printf("%.4f", x1/x2)
			return
		default:
			fmt.Print("неизвестная операция")
			return
		}
	default:
		fmt.Print("неизвестная операция")
		return
	}

}

func readTask() (interface{}, interface{}, interface{}) {
	r := bufio.NewReader(os.Stdin)
	str1, _ := r.ReadString('\n')
	str2, _ := r.ReadString('\n')
	op, _ := r.ReadString('\n')
	return str1, str2, op
}

func printError(value interface{}) {
	fmt.Printf("value=%v: %T", value, value)
}
