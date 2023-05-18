package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type MyStruct struct {
	Id int `json:"global_id"`
}

func sumIds(path string) int {
	var sum int

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	dst := []MyStruct{}
	dec.Decode(&dst)

	for _, val := range dst {
		sum += val.Id
	}
	return sum
}

func main() {
	fmt.Println(sumIds("data-20190514T0100.json"))
}
