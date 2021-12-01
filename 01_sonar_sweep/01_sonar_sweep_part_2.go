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
		fmt.Println("Failed to open file!")
		os.Exit(2)
	}

	scanner := bufio.NewScanner(file)
	arr := make([]int, 0, 2048)

	// go through the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Fields(line)

		depth, err := strconv.Atoi(line_arr[0])

		if err != nil {
			panic(err)
		}

		arr = append(arr, depth)
	}

	file.Close()

	// compare sliding windows in a separate function
	CompareSlidingWindows(arr)
}

func CompareSlidingWindows(arr []int) {
	previous_sum := math.MaxInt
	increased_counter := 0

	// go through all data and add the following 3 numbers
	for i := 0; i < len(arr); i++ {
		depth := arr[i]
		sum := depth

		if i < len(arr)-1 {
			sum += arr[i+1]
		}

		if i < len(arr)-2 {
			sum += arr[i+2]
		}

		if previous_sum == math.MaxInt {
			fmt.Println(strconv.Itoa(sum) + " (N/A - no previous measurement)")
		} else {
			diff := previous_sum - sum

			if diff > 0 {
				fmt.Println(strconv.Itoa(sum) + " (decreased)")
			} else if diff < 0 {
				fmt.Println(strconv.Itoa(sum) + " (increased)")
				increased_counter++
			} else {
				fmt.Println(strconv.Itoa(sum) + " (no change)")
			}
		}

		previous_sum = sum
	}

	fmt.Println("---------------------------------------")
	fmt.Println("Increased counter: " + strconv.Itoa(increased_counter))
}
