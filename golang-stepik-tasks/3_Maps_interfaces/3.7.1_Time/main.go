package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	var timeStr string

	fmt.Scan(&timeStr)
	tTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tTime.Format(time.UnixDate))
}
