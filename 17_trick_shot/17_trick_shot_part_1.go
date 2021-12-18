package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type part_rule struct {
	adjacent string
	inserted string
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
	var str string
	x_start := 0
	x_end := 0
	y_start := 0
	y_end := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")
		line_arr := strings.Split(line, "targetarea:")
		coordinate_arr := strings.Split(line_arr[1], ",")
		x_arr := strings.Split(strings.Split(coordinate_arr[0], "x=")[1], "..")
		y_arr := strings.Split(strings.Split(coordinate_arr[1], "y=")[1], "..")
		x_start, _ = strconv.Atoi(x_arr[0])
		x_end, _ = strconv.Atoi(x_arr[1])
		y_start, _ = strconv.Atoi(y_arr[0])
		y_end, _ = strconv.Atoi(y_arr[1])
		str = line
	}

	file.Close()
	fmt.Println(str, x_start, x_end, y_start, y_end)
	DrawDiagram(x_start, x_end, y_start, y_end)
	ShootTheProbe(x_start, x_end, y_start, y_end)
}

func DrawDiagram(x_start int, x_end int, y_start int, y_end int) {
	for y := 0; y <= int(math.Abs(float64(y_start))); y++ {
		for x := 0; x <= x_end; x++ {
			if x >= x_start && x <= x_end && y >= int(math.Abs(float64(y_end))) && y <= int(math.Abs(float64(y_start))) {
				fmt.Print("T")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func ShootTheProbe(x_start int, x_end int, y_start int, y_end int) {
	new_x := 0
	new_y := 0
	velocity_x := 1
	velocity_y := 1
	vel_x := 1
	vel_y := 1
	max_y := 0
	global_max_y := 0

	// over x
	for {
		velocity_x = vel_x

		// over y
		for {
			new_x = 0
			new_y = 0
			velocity_x = vel_x
			//fmt.Println("velocity", velocity_x, velocity_y)

			for {
				/*if vel_x == 6 && vel_y == 9 {
					fmt.Println(new_x, new_y)
				}*/

				if new_x >= x_start && new_x <= x_end && new_y >= int(math.Abs(float64(y_end))) && new_y <= int(math.Abs(float64(y_start))) {
					fmt.Println("hit", new_x, new_y, max_y)
					if max_y > global_max_y {
						global_max_y = max_y
					}
					break
				} else if velocity_x == 0 && new_x > x_end {
					max_y = 0
					break
				} else if velocity_x == 0 && new_y > int(math.Abs(float64(y_start))) {
					max_y = 0
					break
				}

				new_x = new_x + velocity_x
				new_y = new_y - velocity_y

				if velocity_x > 0 {
					velocity_x--
				}
				velocity_y--

				/*if vel_x == 6 && vel_y == 9 {
					fmt.Println("max", int(math.Abs(float64(new_y))), max_y)
				}*/

				if int(math.Abs(float64(new_y))) > max_y {
					max_y = int(math.Abs(float64(new_y)))
				}
			}

			vel_y++
			velocity_y = vel_y

			if vel_y > 2000 {
				vel_y = 0
				break
			}
		}

		vel_x += 1
		velocity_x = vel_x

		if vel_x > 2000 {
			break
		}
	}

	fmt.Println("max", global_max_y)
}
