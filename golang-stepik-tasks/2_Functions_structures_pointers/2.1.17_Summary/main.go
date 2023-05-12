package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	nums, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//fmt.Println(nums)
	//numBytes := []byte(nums)
	//fmt.Println(numBytes)
	maxElem := nums[0]
	for i, elem := range nums {
		fmt.Println(i, elem)
		if elem > rune(maxElem) {
			maxElem = byte(elem)
		}
	}
	fmt.Println(string(maxElem))
}
