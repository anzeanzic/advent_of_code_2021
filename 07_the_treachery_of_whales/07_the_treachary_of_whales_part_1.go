package main

import (
	"bufio"
	"fmt"
	"math"
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
	var crab_positions []int
	max_pos := math.MinInt32
	min_pos := math.MaxInt32

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, ",")

		for _, num_str := range line_arr {
			position, _ := strconv.Atoi(num_str)
			crab_positions = append(crab_positions, position)

			if position > max_pos {
				max_pos = position
			}
			if position < min_pos {
				min_pos = position
			}
		}
	}

	file.Close()

	fmt.Println(crab_positions)
	FindOptimalPosition(crab_positions, min_pos, max_pos)
}

func FindOptimalPosition(crab_positions []int, min_pos int, max_pos int) {
	min_steps := math.MaxInt32
	fmt.Println(min_pos, max_pos)

	for i := min_pos; i < max_pos; i++ {
		sum := 0

		for crab_ndx := 0; crab_ndx < len(crab_positions); crab_ndx++ {
			sum += int(math.Abs(float64(crab_positions[crab_ndx] - i)))
		}

		if sum < min_steps {
			min_steps = sum
		}
	}

	fmt.Println(min_steps)
}
