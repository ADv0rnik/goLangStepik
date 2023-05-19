package main

import (
	"fmt"
	"time"
)

const now = 1589570165

func main() {
	var min, sec int

	fmt.Scanf("%d мин. %d сек.", &min, &sec)

	currentTime := time.Unix(now+toSec(min, sec), 0)
	fmt.Println(currentTime.Format(time.UnixDate))

}

func toSec(m int, s int) int64 {
	return int64(m*60 + s)
}
