package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	ReadFile()
}

func ReadFile() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	// 1: 2
	// 4: 4
	// 7: 3
	// 8: 7

	scanner := bufio.NewScanner(file)
	counter := 0

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, "|")
		output := strings.Fields(line_arr[1])

		for _, value := range output {
			if len(value) == 2 || len(value) == 4 || len(value) == 3 || len(value) == 7 {
				counter++
			}
		}
	}

	file.Close()
	fmt.Println(counter)
}
