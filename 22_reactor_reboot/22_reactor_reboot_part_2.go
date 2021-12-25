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
	file, err := os.Open("input_test.txt")

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
	Calculate(cuboids)
}

func RebootSequence(cuboids []Cuboid) {
	grid := make(map[Point]int)

	for _, cuboid := range cuboids {
		fmt.Println("cuboid", cuboid)
		//CalculateCubes(cuboid)

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

func Calculate(cuboids []Cuboid) {
	size := 0

	for i, cuboid1 := range cuboids {
		doesOverlap := false
		interlacing_cuboid := Cuboid{on: false, x_start: 0, x_end: 0, y_start: 0, y_end: 0, z_start: 0, z_end: 0}

		// check if cuboid 1 interlaces with some cuboid
		//for j, cuboid2 := range cuboids {
		for j := 0; j < i; j++ {
			if i != j {
				doesOverlap, interlacing_cuboid = CheckIfOverlaps(cuboid1, cuboids[j])
				if doesOverlap {
					break
				}
			}
		}

		fmt.Println(doesOverlap, interlacing_cuboid)

		cuboid_size := CalculateCubes(cuboid1)

		if doesOverlap {
			interlacing_cuboid_size := CalculateCubes(interlacing_cuboid)
			fmt.Println("Cuboid size", interlacing_cuboid_size)

			if cuboid1.on {
				fmt.Println("adding", cuboid_size-interlacing_cuboid_size)
				size += cuboid_size - interlacing_cuboid_size
			} else {
				fmt.Println("subtracting", interlacing_cuboid_size)
				size -= interlacing_cuboid_size
			}
		} else {
			if cuboid1.on {
				fmt.Println("adding", cuboid_size)
				size += cuboid_size
			}
		}
	}

	fmt.Println("Final", size)
}

func CheckIfOverlaps(cuboid1 Cuboid, cuboid2 Cuboid) (bool, Cuboid) {
	doesInterlace := false
	interlacing_cuboid := Cuboid{on: false, x_start: 0, x_end: 0, y_start: 0, y_end: 0, z_start: 0, z_end: 0}

	// if cube 2 x starts inside the cube 1 x
	if cuboid2.x_start >= cuboid1.x_start && cuboid2.x_start <= cuboid1.x_end {
		doesInterlace = true
		interlacing_cuboid.x_start = cuboid2.x_start

		// if cube 2 x ends inside the cube 1 x
		if cuboid2.x_end >= cuboid1.x_start && cuboid2.x_end <= cuboid1.x_end {
			interlacing_cuboid.x_end = cuboid2.x_end
		} else {
			interlacing_cuboid.x_end = cuboid1.x_end
		}
	} else {
		interlacing_cuboid.x_start = cuboid1.x_start

		// if cube 2 x ends inside the cube 1 x
		if cuboid2.x_end >= cuboid1.x_start && cuboid2.x_end <= cuboid1.x_end {
			doesInterlace = true
			interlacing_cuboid.x_end = cuboid2.x_end
		} else {
			interlacing_cuboid.x_end = cuboid1.x_end
		}
	}

	// if cube 2 y starts inside the cube 1 y
	if cuboid2.y_start >= cuboid1.y_start && cuboid2.y_start <= cuboid1.y_end {
		doesInterlace = true
		interlacing_cuboid.y_start = cuboid2.y_start

		// if cube 2 y ends inside the cube 1 y
		if cuboid2.y_end >= cuboid1.y_start && cuboid2.y_end <= cuboid1.y_end {
			interlacing_cuboid.y_end = cuboid2.y_end
		} else {
			interlacing_cuboid.y_end = cuboid1.y_end
		}
	} else {
		interlacing_cuboid.y_start = cuboid1.y_start

		// if cube 2 y ends inside the cube 1 y
		if cuboid2.y_end >= cuboid1.y_start && cuboid2.y_end <= cuboid1.y_end {
			doesInterlace = true
			interlacing_cuboid.y_end = cuboid2.y_end
		} else {
			interlacing_cuboid.y_end = cuboid1.y_end
		}
	}

	// if cube 2 z starts inside the cube 1 z
	if cuboid2.z_start >= cuboid1.z_start && cuboid2.z_start <= cuboid1.z_end {
		doesInterlace = true
		interlacing_cuboid.z_start = cuboid2.z_start

		// if cube 2 z ends inside the cube 1 z
		if cuboid2.z_end >= cuboid1.z_start && cuboid2.z_end <= cuboid1.z_end {
			interlacing_cuboid.z_end = cuboid2.z_end
		} else {
			interlacing_cuboid.z_end = cuboid1.z_end
		}
	} else {
		interlacing_cuboid.z_start = cuboid1.z_start

		// if cube 2 z ends inside the cube 1 z
		if cuboid2.z_end >= cuboid1.z_start && cuboid2.z_end <= cuboid1.z_end {
			doesInterlace = true
			interlacing_cuboid.z_end = cuboid2.z_end
		} else {
			interlacing_cuboid.z_end = cuboid1.z_end
		}
	}

	return doesInterlace, interlacing_cuboid
}

func CalculateCubes(cuboid Cuboid) int {
	calc := (cuboid.x_end - cuboid.x_start + 1) * (cuboid.y_end - cuboid.y_start + 1) * (cuboid.z_end - cuboid.z_start + 1)

	return calc
}
