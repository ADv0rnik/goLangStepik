package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Group struct {
	Students []struct {
		Rating []float64 `json:"Rating"`
	} `json:"Students"`
}

type Output struct {
	Average float64
}

func main() {
	const root = "students.json"
	var group Group
	var totalRates float64
	var count int

	file, err := os.ReadFile(root)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal([]byte(file), &group); err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%+v\n", group)
	for _, val := range group.Students {
		totalRates += getSum(val.Rating)
		count++
	}
	averageGroupRate := totalRates / float64(count)
	fmt.Println(totalRates, count)
	fmt.Printf("%s\n", toJson(averageGroupRate))
}

func getSum(arr []float64) float64 {
	var sum float64
	for _, val := range arr {
		sum += val
	}
	return sum
}

func toJson(value float64) []byte {
	output := Output{
		Average: value,
	}
	data, err := json.MarshalIndent(output, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	return data
}
