package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	previous_depth := -1
	increase_counter := 0

	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())

		if err != nil {
			panic(err)
		}

		if previous_depth == -1 {
			fmt.Println(strconv.Itoa(depth) + " (N/A - no previous measurement)")
		} else {
			diff := previous_depth - depth

			if diff > 0 {
				fmt.Println(strconv.Itoa(depth) + " (decreased)")
			} else if diff < 0 {
				fmt.Println(strconv.Itoa(depth) + " (increased)")
				increase_counter++
			}
		}

		previous_depth = depth
	}

	file.Close()
	fmt.Println("---------------------------------------")
	fmt.Println("Increased counter: " + strconv.Itoa(increase_counter))
}
