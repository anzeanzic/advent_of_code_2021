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
	var lows []int

	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			if i == 0 {
				if j == 0 {
					if heights[i][j] < heights[i][j+1] && heights[i][j] < heights[i+1][j] {
						lows = append(lows, heights[i][j])
					}
				} else if j == len(heights[i])-1 {
					if heights[i][j] < heights[i][j-1] && heights[i][j] < heights[i+1][j] {
						lows = append(lows, heights[i][j])
					}
				} else {
					if heights[i][j] < heights[i][j-1] && heights[i][j] < heights[i][j+1] && heights[i][j] < heights[i+1][j] {
						lows = append(lows, heights[i][j])
					}
				}
			} else if i == len(heights)-1 {
				if j == 0 {
					if heights[i][j] < heights[i][j+1] && heights[i][j] < heights[i-1][j] {
						lows = append(lows, heights[i][j])
					}
				} else if j == len(heights[i])-1 {
					if heights[i][j] < heights[i][j-1] && heights[i][j] < heights[i-1][j] {
						lows = append(lows, heights[i][j])
					}
				} else {
					if heights[i][j] < heights[i][j-1] && heights[i][j] < heights[i][j+1] && heights[i][j] < heights[i-1][j] {
						lows = append(lows, heights[i][j])
					}
				}
			} else {
				if j == 0 {
					if heights[i][j] < heights[i][j+1] && heights[i][j] < heights[i-1][j] && heights[i][j] < heights[i+1][j] {
						lows = append(lows, heights[i][j])
					}
				} else if j == len(heights[i])-1 {
					if heights[i][j] < heights[i][j-1] && heights[i][j] < heights[i-1][j] && heights[i][j] < heights[i+1][j] {
						lows = append(lows, heights[i][j])
					}
				} else {
					if heights[i][j] < heights[i][j-1] && heights[i][j] < heights[i][j+1] && heights[i][j] < heights[i-1][j] && heights[i][j] < heights[i+1][j] {
						lows = append(lows, heights[i][j])
					}
				}
			}
		}
	}

	fmt.Println(lows)

	sum := 0

	for i := 0; i < len(lows); i++ {
		sum += lows[i] + 1
	}

	fmt.Println(sum)
}
