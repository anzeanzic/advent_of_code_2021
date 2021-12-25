package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
	z int
}
type Cuboid struct {
	on      bool
	x_start int
	x_end   int
	y_start int
	y_end   int
	z_start int
	z_end   int
}

func main() {
	ReadFile()
}

var MIN int
var MAX int

func ReadFile() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var cuboids []Cuboid
	MIN := -50
	MAX := 50

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, " ")
		isOn := line_arr[0] == "on"
		remainder := strings.Split(line_arr[1], ",")
		x_arr := strings.Split(remainder[0], "=")
		x_coord_arr := strings.Split(x_arr[1], "..")
		x_start, err1 := strconv.Atoi(x_coord_arr[0])
		x_end, err2 := strconv.Atoi(x_coord_arr[1])

		if err1 != nil {
			panic(err)
		}
		if err2 != nil {
			panic(err)
		}

		if x_start > x_end {
			temp := x_start
			x_start = x_end
			x_end = temp
		}

		y_arr := strings.Split(remainder[1], "=")
		y_coord_arr := strings.Split(y_arr[1], "..")
		y_start, err1 := strconv.Atoi(y_coord_arr[0])
		y_end, err2 := strconv.Atoi(y_coord_arr[1])

		if err1 != nil {
			panic(err)
		}
		if err2 != nil {
			panic(err)
		}

		if y_start > y_end {
			temp := y_start
			y_start = y_end
			y_end = temp
		}

		z_arr := strings.Split(remainder[2], "=")
		z_coord_arr := strings.Split(z_arr[1], "..")
		z_start, err1 := strconv.Atoi(z_coord_arr[0])
		z_end, err2 := strconv.Atoi(z_coord_arr[1])

		if err1 != nil {
			panic(err)
		}
		if err2 != nil {
			panic(err)
		}

		if z_start > z_end {
			temp := z_start
			z_start = z_end
			z_end = temp
		}

		if x_start >= MIN && x_start <= MAX && x_end >= MIN && x_end <= MAX && y_start >= MIN && y_start <= MAX && y_end >= MIN && y_end <= MAX && z_start >= MIN && z_start <= MAX && z_end >= MIN && z_end <= MAX {
			cuboids = append(cuboids, Cuboid{on: isOn, x_start: x_start, x_end: x_end, y_start: y_start, y_end: y_end, z_start: z_start, z_end: z_end})
		}
	}

	file.Close()

	fmt.Println(cuboids)

	RebootSequence(cuboids)
}

func RebootSequence(cuboids []Cuboid) {
	grid := make(map[Point]int)

	for _, cuboid := range cuboids {
		fmt.Println("cuboid", cuboid)

		for x := cuboid.x_start; x <= cuboid.x_end; x++ {
			for y := cuboid.y_start; y <= cuboid.y_end; y++ {
				for z := cuboid.z_start; z <= cuboid.z_end; z++ {
					p := Point{x: x, y: y, z: z}

					if cuboid.on {
						if _, ok := grid[p]; !ok {
							grid[p] = 1
						}
					} else {
						delete(grid, p)
					}
				}
			}
		}

	}

	fmt.Println("Final", len(grid))
}
