package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var timeStr string
	const rushHour = 13

	timeStr, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	timeStr = strings.TrimRight(timeStr, "\n")
	tTime, err := time.Parse(time.DateTime, timeStr)
	if err != nil {
		log.Fatal(err)
	}
	hour := tTime.Hour()

	if hour < rushHour {
		fmt.Println(tTime.Format(time.DateTime))
	} else {
		newTime := tTime.AddDate(0, 0, 1)
		fmt.Println(newTime.Format(time.DateTime))
	}
}
