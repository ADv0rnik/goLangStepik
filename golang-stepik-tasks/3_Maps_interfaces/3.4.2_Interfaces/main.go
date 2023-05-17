package main

import (
	"fmt"
	"strings"
)

type Battery struct {
	capacity string
}

func (b *Battery) String() string {
	countZero, countOne := 0, 0

	for _, val := range b.capacity {
		if val == 49 {
			countZero++
		} else if val == 48 {
			countOne++
		}
	}
	output := "[" + strings.Repeat(" ", countZero) + strings.Repeat("X", countOne) + "]"
	return output
}

func main() {
	var seq string = "100010111"

	//fmt.Scan(&seq)
	batteryForTest := Battery{capacity: seq}
	fmt.Println(batteryForTest.String())
}
