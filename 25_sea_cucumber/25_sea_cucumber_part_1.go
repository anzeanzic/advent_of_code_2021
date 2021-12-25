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

var flashCounter = 0

func ReadFile() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var cucumber_map [][]string

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, "")
		cucumber_map = append(cucumber_map, line_arr)
	}

	file.Close()

	Print(cucumber_map)
	moves_counter := 0

	for step := 1; step <= 10000; step++ {
		cucumber_map, moves_counter = AsStepsGoBy(cucumber_map, step)

		if moves_counter == 0 {
			fmt.Println(step, "End:", moves_counter)
			break
		}
	}
}

func AsStepsGoBy(cucumber_map [][]string, step int) ([][]string, int) {
	fmt.Println("After", step, "step:")
	moves_counter := 0

	// first move all right
	temp_map := make([][]string, len(cucumber_map))
	for i := 0; i < len(temp_map); i++ {
		temp_map[i] = make([]string, len(cucumber_map[i]))
	}

	for y := 0; y < len(cucumber_map); y++ {
		for x := 0; x < len(cucumber_map[y]); x++ {
			if cucumber_map[y][x] == ">" {
				itMoves, new_x := CheckIfItCanMoveRight(cucumber_map, y, x)

				if itMoves {
					temp_map[y][new_x] = ">"
					temp_map[y][x] = "."
					moves_counter++
				} else {
					temp_map[y][x] = cucumber_map[y][x]
				}
			} else {
				if temp_map[y][x] == "" {
					temp_map[y][x] = cucumber_map[y][x]
				}
			}
		}
	}

	//Print(cucumber_map)
	cucumber_map = temp_map

	// secondly move all down
	temp_map = make([][]string, len(cucumber_map))
	for i := 0; i < len(temp_map); i++ {
		temp_map[i] = make([]string, len(cucumber_map[i]))
	}

	for y := 0; y < len(cucumber_map); y++ {
		for x := 0; x < len(cucumber_map[y]); x++ {
			if cucumber_map[y][x] == "v" {
				itMoves, new_y := CheckIfItCanMoveDown(cucumber_map, y, x)
				//fmt.Println(y, x, itMoves, new_y)

				if itMoves {
					temp_map[new_y][x] = "v"
					temp_map[y][x] = "."
					moves_counter++
				} else {
					temp_map[y][x] = cucumber_map[y][x]
				}
			} else {
				if temp_map[y][x] == "" {
					temp_map[y][x] = cucumber_map[y][x]
				}
			}
		}
	}

	//Print(temp_map)

	return temp_map, moves_counter
}

func CheckIfItCanMoveRight(cucumber_map [][]string, y int, x int) (bool, int) {
	// if x is inside of bounds
	if x+1 < len(cucumber_map[y]) {
		return cucumber_map[y][x+1] != ">" && cucumber_map[y][x+1] != "v", x + 1
	} else {
		return cucumber_map[y][0] != ">" && cucumber_map[y][0] != "v", 0
	}
}

func CheckIfItCanMoveDown(cucumber_map [][]string, y int, x int) (bool, int) {
	// if y is inside of bounds
	if y+1 < len(cucumber_map) {
		return cucumber_map[y+1][x] != ">" && cucumber_map[y+1][x] != "v", y + 1
	} else {
		return cucumber_map[0][x] != ">" && cucumber_map[0][x] != "v", 0
	}
}

func Print(cucumber_map [][]string) {
	for y := 0; y < len(cucumber_map); y++ {
		for x := 0; x < len(cucumber_map[y]); x++ {
			fmt.Print(cucumber_map[y][x], "")
		}

		fmt.Println()
	}

	fmt.Println("-------------------")
}
