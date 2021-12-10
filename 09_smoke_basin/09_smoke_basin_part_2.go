package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type low struct {
	num int
	x   int
	y   int
}
type coordinate struct {
	x int
	y int
}

var coordinates []coordinate

func main() {
	ReadFile()
}

func ReadFile() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var heights [][]int

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, "")

		var temp []int
		for _, num_str := range line_arr {
			num, err := strconv.Atoi(num_str)

			if err != nil {
				panic(err)
			}

			temp = append(temp, num)
		}

		heights = append(heights, temp)
	}

	file.Close()

	fmt.Println(heights)

	FindLows(heights)
}

func FindLows(heights [][]int) {
	var lows []low

	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			if i == 0 {
				if j == 0 {
					if heights[i][j] < heights[i][j+1] && heights[i][j] < heights[i+1][j] {
						lows = append(lows, low{num: heights[i][j], x: j, y: i})
					}
				} else if j == len(heights[i])-1 {
					if heights[i][j] < heights[i][j-1] && heights[i][j] < heights[i+1][j] {
						lows = append(lows, low{num: heights[i][j], x: j, y: i})
					}
				} else {
					if heights[i][j] < heights[i][j-1] && heights[i][j] < heights[i][j+1] && heights[i][j] < heights[i+1][j] {
						lows = append(lows, low{num: heights[i][j], x: j, y: i})
					}
				}
			} else if i == len(heights)-1 {
				if j == 0 {
					if heights[i][j] < heights[i][j+1] && heights[i][j] < heights[i-1][j] {
						lows = append(lows, low{num: heights[i][j], x: j, y: i})
					}
				} else if j == len(heights[i])-1 {
					if heights[i][j] < heights[i][j-1] && heights[i][j] < heights[i-1][j] {
						lows = append(lows, low{num: heights[i][j], x: j, y: i})
					}
				} else {
					if heights[i][j] < heights[i][j-1] && heights[i][j] < heights[i][j+1] && heights[i][j] < heights[i-1][j] {
						lows = append(lows, low{num: heights[i][j], x: j, y: i})
					}
				}
			} else {
				if j == 0 {
					if heights[i][j] < heights[i][j+1] && heights[i][j] < heights[i-1][j] && heights[i][j] < heights[i+1][j] {
						lows = append(lows, low{num: heights[i][j], x: j, y: i})
					}
				} else if j == len(heights[i])-1 {
					if heights[i][j] < heights[i][j-1] && heights[i][j] < heights[i-1][j] && heights[i][j] < heights[i+1][j] {
						lows = append(lows, low{num: heights[i][j], x: j, y: i})
					}
				} else {
					if heights[i][j] < heights[i][j-1] && heights[i][j] < heights[i][j+1] && heights[i][j] < heights[i-1][j] && heights[i][j] < heights[i+1][j] {
						lows = append(lows, low{num: heights[i][j], x: j, y: i})
					}
				}
			}
		}
	}

	fmt.Println(lows)

	/*sum := 0

	for i := 0; i < len(lows); i++ {
		sum += lows[i] + 1
	}

	fmt.Println(sum)*/
	fmt.Println(lows)
	FindLargestBasins(lows, heights)
}

func FindLargestBasins(lows []low, heights [][]int) {
	var sums []int

	for i := 0; i < len(lows); i++ {
		previous_num := lows[i].num
		sum := 0

		sum += Recursive(heights, lows[i].x, lows[i].y, lows[i].x, lows[i].y, previous_num)

		fmt.Println("sum", lows[i].num, sum)
		fmt.Println("-----------")
		sums = append(sums, sum)
	}

	FindMaxMultiplier(sums)
}

func Recursive(heights [][]int, x int, y int, previous_x int, previous_y int, previous_num int) int {
	if x < 0 || y < 0 || x >= len(heights[y]) || y >= len(heights) || heights[y][x] == 9 {
		return 0
	}

	sum := 0

	if heights[y][x] >= previous_num && !CheckIfCoordinateAlreadyExists(x, y) {
		//fmt.Println("compare", heights[y][x], previous_num, "+1", y, x, previous_y, previous_x)
		sum += 1
		coordinates = append(coordinates, coordinate{x: x, y: y})
	} else {
		return 0
	}

	// find right
	if x+1 < len(heights[y]) && x+1 != previous_x {
		//fmt.Println("right")
		sum += Recursive(heights, x+1, y, x, y, heights[y][x])
	}
	// find left
	if x-1 >= 0 && x-1 != previous_x {
		//fmt.Println("left")
		sum += Recursive(heights, x-1, y, x, y, heights[y][x])
	}
	// find up
	if y-1 >= 0 && y-1 != previous_y {
		//fmt.Println("up")
		sum += Recursive(heights, x, y-1, x, y, heights[y][x])
	}
	// find down
	if y+1 < len(heights) && y+1 != previous_y {
		//fmt.Println("down")
		sum += Recursive(heights, x, y+1, x, y, heights[y][x])
	}

	return sum
}

func CheckIfCoordinateAlreadyExists(x int, y int) bool {
	found := false

	for _, coor := range coordinates {
		if coor.x == x && coor.y == y {
			found = true
			break
		}
	}

	return found
}

func FindMaxMultiplier(sums []int) {
	sort.Ints(sums)
	fmt.Println(sums)
	fmt.Println(sums[len(sums)-1] * sums[len(sums)-2] * sums[len(sums)-3])
}
