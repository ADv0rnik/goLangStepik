package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num uint

	fn := func(u uint) uint {
		newStr := ""
		iStr := strconv.Itoa(int(u))
		if u%2 != 0 || u == 0 {
			return uint(100)
		} else {
			for _, i := range iStr {
				i_, _ := strconv.Atoi(string(i))
				if i_%2 == 0 && i_ != 0 {
					newStr += string(i)
				}
			}
			res, _ := strconv.Atoi(newStr)
			return uint(res)
		}
	}

	fmt.Scan(&num)
	fmt.Println(fn(num))

}
