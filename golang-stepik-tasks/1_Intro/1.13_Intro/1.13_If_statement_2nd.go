package main

import "fmt"

func main() {
	var nums string

	fmt.Scan(&nums)
	if nums[0]+nums[1]+nums[2] == nums[3]+nums[4]+nums[5] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
