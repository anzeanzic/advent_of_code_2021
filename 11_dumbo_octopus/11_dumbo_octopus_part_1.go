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

var flashCounter = 0

func ReadFile() {
	file, err := os.Open("input_test.txt")

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
	fmt.Println(flashCounter)
}

func AsStepsGoBy(octopuses [][]int) {
	for step := 0; step < 100; step++ {
		for y := 0; y < len(octopuses); y++ {
			for x := 0; x < len(octopuses[y]); x++ {
				octopuses[y][x] += 1

				CheckIfFlashes(octopuses, y, x, y, x)
			}
		}

		for y := 0; y < len(octopuses); y++ {
			for x := 0; x < len(octopuses[y]); x++ {
				if octopuses[y][x] > 9 {
					flashCounter++
					octopuses[y][x] = 0
				}
			}
		}

		Print(octopuses)
	}
}

func CheckIfFlashes(octopuses [][]int, y int, x int, current_y int, current_x int) {
	if octopuses[y][x] > 9 {
		//isPrevious := y < current_y || (y == current_y && x < current_x)
		//fmt.Println(y, x, octopuses[y][x])

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

		//octopuses[y][x] = 0
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
