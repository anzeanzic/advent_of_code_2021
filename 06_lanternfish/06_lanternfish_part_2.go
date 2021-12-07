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
	file, err := os.Open("input_test.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var fish_array []int

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, ",")

		for _, num_str := range line_arr {
			number, _ := strconv.Atoi(num_str)
			fish_array = append(fish_array, number)
		}
	}

	file.Close()

	fmt.Println(fish_array)
	AsDaysGoBy(fish_array)
}

func AsDaysGoBy(fish_array []int) {
	fmt.Println("Initial state:", fish_array)
	day_no := 80

	/*for i := 0; i < day_no; i++ {
		new_fish_count := 0

		for fish_ndx := 0; fish_ndx < len(fish_array); fish_ndx++ {
			if fish_array[fish_ndx] == 0 {
				fish_array[fish_ndx] = 6
				new_fish_count++
			} else {
				fish_array[fish_ndx]--
			}
		}

		// add new fish
		if new_fish_count > 0 {
			for i := 0; i < new_fish_count; i++ {
				fish_array = append(fish_array, 8)
			}
		}

		//fmt.Println("After "+strconv.Itoa(i+1)+" days:", fish_array)
		fmt.Println(i)
	}*/

	sum := 0

	//sum = 3 * int(math.Pow(1+1/6, float64(day_no)))
	//growth_rate := 1 / 6

	for fish_ndx := 0; fish_ndx < len(fish_array); fish_ndx++ {
		//sum += growth_rate * day_no
	}

	//sum = 5 * math.Pow(float64((1+growth_rate)), float64(day_no))

	fmt.Println(sum)
	//fmt.Println(len(fish_array))
}
