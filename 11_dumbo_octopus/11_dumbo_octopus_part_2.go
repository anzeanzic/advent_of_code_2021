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
	var octopuses [][]int

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, "")

		var temp []int
		for _, num_str := range line_arr {
			number, _ := strconv.Atoi(num_str)
			temp = append(temp, number)
		}

		octopuses = append(octopuses, temp)
	}

	file.Close()

	Print(octopuses)

	AsStepsGoBy(octopuses)
}

func AsStepsGoBy(octopuses [][]int) {
	step := 0
	octopusFlashCounter := 0

	for ok := true; ok; ok = (octopusFlashCounter != len(octopuses)*len(octopuses[0])) {
		for y := 0; y < len(octopuses); y++ {
			for x := 0; x < len(octopuses[y]); x++ {
				octopuses[y][x] += 1

				CheckIfFlashes(octopuses, y, x, y, x)
			}
		}

		octopusFlashCounter = 0

		for y := 0; y < len(octopuses); y++ {
			for x := 0; x < len(octopuses[y]); x++ {
				if octopuses[y][x] > 9 {
					octopusFlashCounter++
					octopuses[y][x] = 0
				}
			}
		}

		if octopusFlashCounter == len(octopuses)*len(octopuses[0]) {
			fmt.Println("step", step+1)
		}

		Print(octopuses)
		step++
	}
}

func CheckIfFlashes(octopuses [][]int, y int, x int, current_y int, current_x int) {
	if octopuses[y][x] > 9 {
		if octopuses[y][x] > 10 {
			return
		}

		// horizontally
		if x+1 < len(octopuses[y]) {
			octopuses[y][x+1]++

			if octopuses[y][x+1] > 9 {
				CheckIfFlashes(octopuses, y, x+1, current_y, current_x)
			}
		}
		if x-1 >= 0 {
			octopuses[y][x-1]++

			if octopuses[y][x-1] > 9 {
				CheckIfFlashes(octopuses, y, x-1, current_y, current_x)
			}
		}
		// vertically
		if y+1 < len(octopuses) {
			octopuses[y+1][x]++

			if octopuses[y+1][x] > 9 {
				CheckIfFlashes(octopuses, y+1, x, current_y, current_x)
			}
		}
		if y-1 >= 0 {
			octopuses[y-1][x]++

			if octopuses[y-1][x] > 9 {
				CheckIfFlashes(octopuses, y-1, x, current_y, current_x)
			}
		}
		// diagonals
		if x+1 < len(octopuses[y]) && y+1 < len(octopuses) {
			octopuses[y+1][x+1]++

			if octopuses[y+1][x+1] > 9 {
				CheckIfFlashes(octopuses, y+1, x+1, current_y, current_x)
			}
		}
		if x-1 >= 0 && y+1 < len(octopuses) {
			octopuses[y+1][x-1]++

			if octopuses[y+1][x-1] > 9 {
				CheckIfFlashes(octopuses, y+1, x-1, current_y, current_x)
			}
		}
		if x-1 >= 0 && y-1 >= 0 {
			octopuses[y-1][x-1]++

			if octopuses[y-1][x-1] > 9 {
				CheckIfFlashes(octopuses, y-1, x-1, current_y, current_x)
			}
		}
		if x+1 < len(octopuses[y]) && y-1 >= 0 {
			octopuses[y-1][x+1]++

			if octopuses[y-1][x+1] > 9 {
				CheckIfFlashes(octopuses, y-1, x+1, current_y, current_x)
			}
		}
	}
}

func Print(octopuses [][]int) {
	for y := 0; y < len(octopuses); y++ {
		for x := 0; x < len(octopuses[y]); x++ {
			fmt.Print(octopuses[y][x], "")
		}

		fmt.Println()
	}

	fmt.Println("-------------------")
}
