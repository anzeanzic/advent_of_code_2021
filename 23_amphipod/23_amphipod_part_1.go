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
	file, err := os.Open("input_test.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	locations := make([][]string, 0)
	row_ndx := 0

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, "")

		locations = append(locations, line_arr)

		row_ndx++
	}

	file.Close()

	fmt.Println(locations)
}

func GetEnergyPoints(amphipod string) int {
	switch {
	case amphipod == "A":
		return 1
	case amphipod == "B":
		return 10
	case amphipod == "C":
		return 100
	case amphipod == "D":
		return 1000
	}

	return 0
}
