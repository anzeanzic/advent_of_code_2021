package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	var opened []string
	var sums []int
	line_num := 0

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, "")
		fmt.Println(line_num)
		opened = make([]string, 0)
		faulty := false

		for _, char := range line_arr {
			if char == ")" || char == "]" || char == "}" || char == ">" {
				n := len(opened) - 1
				opening_char := opened[n]
				opened = opened[:n] // pop

				if char == ")" && opening_char != "(" || char == "]" && opening_char != "[" || char == "}" && opening_char != "{" || char == ">" && opening_char != "<" {
					fmt.Println("Expected " + ExpectedChar(opening_char) + ", but found " + char + " instead.")
					faulty = true
					fmt.Println(char)
					break
				}
			} else {
				opened = append(opened, char)
			}
		}

		line_num++

		// add lines to the end
		if !faulty {
			sum := 0

			for i := len(opened) - 1; i >= 0; i-- {
				expected_char := ExpectedChar(opened[i])
				fmt.Print(expected_char)
				sum = sum*5 + GetScoreForChar(expected_char)
			}
			sums = append(sums, sum)
			fmt.Println()
		}
	}

	file.Close()

	fmt.Println(opened)
	fmt.Println(sums)

	sort.Ints(sums)
	middle_score := sums[int(math.Ceil(float64(len(sums)/2)))]

	fmt.Println(middle_score)
	//GetIllegalCharsPoints(illegal_chars)
}

// 2, 4, 5, 7, 8

func ExpectedChar(opening_char string) string {
	closing_char := ""

	switch {
	case opening_char == "(":
		closing_char = ")"
	case opening_char == "[":
		closing_char = "]"
	case opening_char == "{":
		closing_char = "}"
	case opening_char == "<":
		closing_char = ">"
	}

	return closing_char
}

func GetScoreForChar(char string) int {
	sum := 0

	for i := 0; i < len(char); i++ {
		switch {
		case char == ")":
			sum += 1
		case char == "]":
			sum += 2
		case char == "}":
			sum += 3
		case char == ">":
			sum += 4
		}
	}

	return sum
}
