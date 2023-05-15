package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   []string{"Село", "Деревня", "ПГТ"},  // города с населением 10-99 тыс. человек
		100:  []string{"Город", "Станица"},        // города с населением 100-999 тыс. человек
		1000: []string{"Мегаполис", "Миллионник"}, // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"Город":     101,
		"Станица":   102,
		"Село":      103,
		"Мегаполис": 104,
	}

	for _, city := range groupCity[10] {
		delete(cityPopulation, city)
	}

	for _, city := range groupCity[1000] {
		delete(cityPopulation, city)
	}

	/*
		for key := range cityPopulation {
			isInMap := false
			for _, val := range groupCity[100] {
				if key == val {
					isInMap = true
					break
				}
			}
			if !isInMap {
				delete(cityPopulation, key)
			}
		}
	*/
	for key, val := range cityPopulation {
		fmt.Println(key, val)
	}

}
