package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type bits struct {
	zeros int
	ones  int
}

func main() {
	ReadFile()
}

func ReadFile() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	var arr []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, line)
	}

	file.Close()

	arr1 := make([]string, len(arr))
	copy(arr1, arr)

	var oxygen_gen_rating = AnalyzeNumbers(arr, true)
	var co2_scrubber_rating = AnalyzeNumbers(arr1, false)

	fmt.Println(oxygen_gen_rating * co2_scrubber_rating)
}

func AnalyzeNumbers(arr []string, isMostCommon bool) int64 {
	for i := 0; i < len(arr[0]); i++ {
		zero_counter := 0
		one_counter := 0

		if len(arr) == 1 {
			break
		}

		for j := 0; j < len(arr); j++ {
			first_bit := arr[j][i]
			bit, err := strconv.Atoi(string(first_bit))

			if err != nil {
				panic(err)
			}

			switch {
			case bit == 0:
				zero_counter++
			case bit == 1:
				one_counter++
			}
		}

		if isMostCommon {
			if zero_counter > one_counter {
				arr = RemoveItemsFromArray(arr, i, 0)
			} else if zero_counter < one_counter || zero_counter == one_counter {
				arr = RemoveItemsFromArray(arr, i, 1)
			}
		} else {
			if zero_counter < one_counter || zero_counter == one_counter {
				arr = RemoveItemsFromArray(arr, i, 0)
			} else if zero_counter > one_counter {
				arr = RemoveItemsFromArray(arr, i, 1)
			}
		}
	}

	number_int, err := strconv.ParseInt(arr[0], 2, 64)

	if err != nil {
		panic(err)
	}

	fmt.Println(number_int)

	return number_int
}

func RemoveItemsFromArray(arr []string, pos int, dominant_bit int) []string {
	for i, rlen := 0, len(arr); i < rlen; i++ {
		j := i - (rlen - len(arr))
		bit, err := strconv.Atoi(string(arr[j][pos]))

		if err != nil {
			panic(err)
		}

		if bit != dominant_bit {
			arr = RemoveIndex(arr, j)
		}
	}

	return arr
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
