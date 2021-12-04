package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type bits struct {
	zeros int
	ones  int
}

func main() {
	ReadFile()
}

func ReadFile() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	var counters = []bits{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if len(counters) == 0 {
			counters = make([]bits, len(line))
		}

		for i, char := range line {
			bit, err := strconv.Atoi(string(char))

			if err != nil {
				panic(err)
			}

			switch {
			case bit == 0:
				counters[i].zeros++
			case bit == 1:
				counters[i].ones++
			}
		}
	}

	var gamma string
	var epsilon string

	for i := range counters {
		if counters[i].zeros > counters[i].ones {
			gamma += "0"
			epsilon += "1"
		} else if counters[i].zeros < counters[i].ones {
			gamma += "1"
			epsilon += "0"
		}
	}

	gamma_int, err := strconv.ParseInt(gamma, 2, 64)
	epsilon_int, err := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(counters)
	fmt.Println(gamma, gamma_int, epsilon, epsilon_int, gamma_int*epsilon_int)

	file.Close()
}
