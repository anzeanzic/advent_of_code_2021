package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}
type pairs struct {
	x1y1 point
	x2y2 point
}

func main() {
	ReadFile()
}

func ReadFile() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var array_of_pairs []pairs
	x_max := -1
	y_max := -1

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Fields(line)

		x1y1 := ParsePoint(line_arr[0])
		x2y2 := ParsePoint(line_arr[2])

		if x1y1.x > x_max {
			x_max = x1y1.x
		}
		if x2y2.x > x_max {
			x_max = x2y2.x
		}
		if x1y1.y > y_max {
			y_max = x1y1.y
		}
		if x2y2.y > y_max {
			y_max = x2y2.y
		}

		array_of_pairs = append(array_of_pairs, pairs{x1y1: x1y1, x2y2: x2y2})
	}

	file.Close()

	DrawDiagram(array_of_pairs, x_max, y_max)
}

func ParsePoint(point_str string) point {
	point_str_arr := strings.Split(point_str, ",")
	x, err1 := strconv.Atoi(point_str_arr[0])
	y, err2 := strconv.Atoi(point_str_arr[1])

	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}

	return point{x: x, y: y}
}

func DrawDiagram(array_of_pairs []pairs, x_max int, y_max int) {
	var diagram = make([][]int, y_max+1)

	for i := 0; i < y_max+1; i++ {
		diagram[i] = make([]int, x_max+1)
	}

	for i := 0; i < len(array_of_pairs); i++ {
		x1y1 := array_of_pairs[i].x1y1
		x2y2 := array_of_pairs[i].x2y2

		if x1y1.x == x2y2.x {
			internal_min := 0
			internal_max := 0

			if x1y1.y > x2y2.y {
				internal_min = x2y2.y
				internal_max = x1y1.y
			} else {
				internal_min = x1y1.y
				internal_max = x2y2.y
			}

			for y := internal_min; y <= internal_max; y++ {
				diagram[y][x1y1.x]++
			}
		} else if x1y1.y == x2y2.y {
			internal_min := 0
			internal_max := 0

			if x1y1.x > x2y2.x {
				internal_min = x2y2.x
				internal_max = x1y1.x
			} else {
				internal_min = x1y1.x
				internal_max = x2y2.x
			}

			for x := internal_min; x <= internal_max; x++ {
				diagram[x1y1.y][x]++
			}
		} else {
			internal_x_min := 0
			internal_x_max := 0
			start_y := 0
			end_y := 0
			direction := 1

			if x1y1.x < x2y2.x {
				internal_x_min = x1y1.x
				internal_x_max = x2y2.x

				if x1y1.y < x2y2.y {
					// e.g.: (1,1) -> (3,3)
					direction = 1
				} else if x1y1.y > x2y2.y {
					// e.g.: (1,9) -> (5,5)
					direction = -1
				}

				start_y = x1y1.y
				end_y = x2y2.y
			} else {
				internal_x_min = x2y2.x
				internal_x_max = x1y1.x

				if x1y1.y < x2y2.y {
					// e.g.: (4,1) -> (2,3)
					direction = -1
				} else if x1y1.y > x2y2.y {
					//e.g.: (4,8) -> (2,6)
					direction = 1
				}

				start_y = x2y2.y
				end_y = x1y1.y
			}

			for x := internal_x_min; x <= internal_x_max; x++ {
				if (direction == 1 && start_y > end_y) || (direction == -1 && start_y < end_y) {
					break
				}

				diagram[start_y][x]++
				start_y += direction
			}
		}
	}

	// count bigger than 2
	number_of_points := 0

	for i := 0; i < len(diagram); i++ {
		for j := 0; j < len(diagram[i]); j++ {
			if diagram[i][j] >= 2 {
				number_of_points++
			}
		}
	}

	fmt.Println(number_of_points)
}
