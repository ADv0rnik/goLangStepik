/*
На вход подается целое число.
Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var nums, newNums string

	fmt.Scan(&nums)

	for i := range nums {
		num, _ := strconv.Atoi(string(nums[i]))
		newNums += strconv.Itoa(int(math.Pow(float64(num), 2)))
	}
	fmt.Println(newNums)
}
