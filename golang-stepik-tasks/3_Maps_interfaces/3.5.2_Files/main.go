/*
Task uncompleted
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	const root = "/home/alex-dvornik/Projects/GoLang/goLangStepik/golang-stepik-tasks/3_Maps_interfaces/3.5.2_Files/task"
	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}

	if info.IsDir() {
		return nil // Проигнорируем директории
	}

	if strings.Contains(path, "csv") {
		fmt.Printf("Name: %s\tSize: %d byte\tPath: %s\n", info.Name(), info.Size(), path)
	}

	return nil
}
