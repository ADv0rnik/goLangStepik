/*
Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.
Выведите на экран фразу:
It is ... hours ... minutes.
Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом.
*/

package main

import "fmt"

func main() {
	var degree int32

	fmt.Scan(&degree)
	h := degree / 30
	m := (degree % 30) * 60 / 30

	fmt.Println("It is", h, "hours", m, "minutes.")
}
