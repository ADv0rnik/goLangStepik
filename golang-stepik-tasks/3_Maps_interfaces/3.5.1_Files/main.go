package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	var sum int

	scaner := bufio.NewScanner(os.Stdin)
	for scaner.Scan() {
		numStr := scaner.Text()
		num, err := strconv.Atoi(numStr)
		if err != nil {
			break
		}
		sum += num
	}
	io.WriteString(os.Stdout, strconv.Itoa(sum))
}
