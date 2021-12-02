package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	scanner := bufio.NewScanner(file)
	horizontal_pos := 0
	depth := 0

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Fields(line)
		direction := line_arr[0]
		units, err := strconv.Atoi(line_arr[1])

		if err != nil {
			panic(err)
		}

		switch {
		case direction == "forward":
			horizontal_pos += units
		case direction == "up":
			depth -= units
		case direction == "down":
			depth += units
		}
	}

	fmt.Println("Horizontal position:", horizontal_pos, "Depth:", depth)
	fmt.Println("Multiplied:", horizontal_pos*depth)

	file.Close()
}
