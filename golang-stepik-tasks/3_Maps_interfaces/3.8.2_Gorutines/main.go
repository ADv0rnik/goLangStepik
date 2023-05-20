package main

import "fmt"

func main() {
	var s string
	c := make(chan string, 5)

	fmt.Scan(&s)
	go task2(c, s)

	fmt.Print(<-c)

}

func task2(c chan string, n string) {
	for i := 0; i < 5; i++ {
		c <- n + " "
	}
}
