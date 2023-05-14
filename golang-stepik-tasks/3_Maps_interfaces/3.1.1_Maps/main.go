package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	s := make(map[int]int)

	for i := 10; i > 0; i-- {
		fmt.Scan(&n)
		if value, ok := s[n]; ok {
			fmt.Print(value, " ")
		} else {
			s[n] = work(n)
			fmt.Print(s[n], " ")
		}
	}
}

func work(n int) int {
	if n > 3 {
		time.Sleep(time.Millisecond * 500)
		return n + 1
	} else {
		time.Sleep(time.Millisecond * 500)
		return n - 1
	}
}
