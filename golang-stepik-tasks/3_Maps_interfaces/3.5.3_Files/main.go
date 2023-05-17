package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	f, err := os.Open("task.data")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		newArr := strings.Split(record[0], ";")

		for i, value := range newArr {
			fmt.Println(i, value)
			if value == "0" {
				fmt.Printf("%s, %d, %s", newArr[i], i+1, " <<< Over here!!!")
				break
			}
		}
	}
}
