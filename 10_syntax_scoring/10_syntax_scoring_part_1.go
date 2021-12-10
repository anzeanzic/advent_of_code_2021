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

	scanner := bufio.NewScanner(file)
	var opened []string
	var illegal_chars []string
	line_num := 0

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, "")
		fmt.Println(line_num)
		opened = make([]string, 0)

		for _, char := range line_arr {
			if line_num == 2 {
				fmt.Println("char", char)
				fmt.Println(opened)
			}

			if char == ")" || char == "]" || char == "}" || char == ">" {
				n := len(opened) - 1
				opening_char := opened[n]
				opened = opened[:n] // pop

				if char == ")" && opening_char != "(" || char == "]" && opening_char != "[" || char == "}" && opening_char != "{" || char == ">" && opening_char != "<" {
					fmt.Println("Expected " + ExpectedChar(opening_char) + ", but found " + char + " instead.")
					fmt.Println(char)
					illegal_chars = append(illegal_chars, char)
					break
				}
			} else {
				opened = append(opened, char)
			}

			if line_num == 2 {
				fmt.Println(opened)
			}
		}

		line_num++
	}

	file.Close()

	fmt.Println(opened)

	GetIllegalCharsPoints(illegal_chars)
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

	fmt.Println("opening", opening_char)

	return closing_char
}

func GetIllegalCharsPoints(illegal_chars []string) {
	sum := 0
	fmt.Println(illegal_chars)

	for i := 0; i < len(illegal_chars); i++ {
		switch {
		case illegal_chars[i] == ")":
			sum += 3
		case illegal_chars[i] == "]":
			sum += 57
		case illegal_chars[i] == "}":
			sum += 1197
		case illegal_chars[i] == ">":
			sum += 25137
		}
	}

	fmt.Println(sum)
}
