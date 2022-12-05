/*
Определите является ли билет счастливым. Счастливым считается билет,
в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
На вход подается номер билета - одно шестизначное  число.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number string
	var (
		nums1, nums2 int
		a1, a2       int
		b1, b2       int
		c1, c2       int
	)

	fmt.Scan(&number)
	nums1, _ = strconv.Atoi(number[0:3])
	nums2, _ = strconv.Atoi(number[3:])
	a1 = nums1 / 100
	b1 = nums1 / 10 % 10
	c1 = nums1 % 10
	a2 = nums2 / 100
	b2 = nums2 / 10 % 10
	c2 = nums2 % 10
	if (a1 + b1 + c1) == (a2 + b2 + c2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
