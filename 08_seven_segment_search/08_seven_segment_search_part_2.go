package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	// 0: 6
	// 1: 2
	// 2: 5
	// 3: 5
	// 4: 4
	// 5: 5
	// 6: 6
	// 7: 3
	// 8: 7
	// 9: 6

	scanner := bufio.NewScanner(file)
	sum := 0
	var one string
	var seven string
	var four string
	var eight string

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, "|")
		arr := strings.Fields(line_arr[0])
		output := strings.Fields(line_arr[1])
		arr = append(arr, output...)

		dict := make(map[string]int)
		undefined_keys := make(map[string]int)

		for _, key := range arr {
			if len(key) == 2 || len(key) == 4 || len(key) == 3 || len(key) == 7 {
				if _, ok := dict[key]; !ok {
					number := 0

					switch len(key) {
					case 2:
						number = 1
						one = SortStringByCharacter(key)
					case 3:
						number = 7
						seven = SortStringByCharacter(key)
					case 4:
						number = 4
						four = SortStringByCharacter(key)
					case 7:
						number = 8
						eight = SortStringByCharacter(key)
					}

					dict[SortStringByCharacter(key)] = number
				}
			} else {
				if _, ok := undefined_keys[key]; !ok {
					undefined_keys[SortStringByCharacter(key)] = 0
				}
			}
		}

		// HERE
		// find 6 (8 - one of 1s)
		six, right_up := FindKey(eight, one, undefined_keys)
		if six != "" {
			dict[six] = 6
			delete(undefined_keys, six)
		}

		// find 0 or 9
		zero, nine := FindZeroOrNine(seven, four, undefined_keys)
		if zero != "" {
			dict[zero] = 0
			delete(undefined_keys, zero)
		}
		if nine != "" {
			dict[nine] = 9
			delete(undefined_keys, nine)
		}
		// find 2
		right_down := strings.ReplaceAll(one, right_up, "")
		two := FindTwo(right_down, undefined_keys)
		if two != "" {
			dict[two] = 2
			delete(undefined_keys, two)
		}
		// find 3 and 5
		three, five := ThreeAndFive(two, undefined_keys)
		if three != "" {
			dict[three] = 3
			delete(undefined_keys, three)
		}
		if five != "" {
			dict[five] = 5
			delete(undefined_keys, five)
		}

		//fmt.Println("dict", dict)
		//fmt.Println("undef", undefined_keys)

		sum += CalculateOutput(output, dict)
	}

	file.Close()

	fmt.Println("sum", sum)
}

func StringDiff(slice1 string, slice2 string) []string {
	var diff []string

	for _, char := range slice2 {
		found := false

		for _, char2 := range slice1 {
			if char == char2 {
				found = true
			}
		}

		if !found {
			diff = append(diff, string(char))
		}
	}

	return diff
}

func StringToRuneSlice(s string) []rune {
	var r []rune

	for _, runeValue := range s {
		r = append(r, runeValue)
	}

	return r
}

func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)

	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}

func FindKey(str1 string, str2 string, undefined_keys map[string]int) (string, string) {
	new_number := ""
	char := ""

	for _, str2_char := range str2 {
		temp := strings.ReplaceAll(str1, string(str2_char), "")

		for k, _ := range undefined_keys {
			if k == temp {
				new_number = k
				char = string(str2_char)
				break
			}
		}
	}

	return new_number, char
}

func FindZeroOrNine(seven string, four string, undefined_keys map[string]int) (string, string) {
	var temp_arr []string
	var key_arr []string

	for k, _ := range undefined_keys {
		if len(k) == 6 {
			temp_arr = append(temp_arr, RemoveAllCharsFromString(k, seven))
			key_arr = append(key_arr, k)
		}
	}

	var diff []string
	zero := ""
	nine := ""

	for i := 0; i < len(temp_arr); i++ {
		if i == 0 {
			diff = StringDiff(temp_arr[i+1], temp_arr[i])
		} else {
			diff = StringDiff(temp_arr[i-1], temp_arr[i])
		}

		if len(diff) > 0 {
			if CheckIfCharInString(diff[0], four) {
				nine = key_arr[i]
			} else {
				zero = key_arr[i]
			}
		}
	}

	return zero, nine
}

func FindTwo(char string, undefined_keys map[string]int) string {
	two := ""

	for k, _ := range undefined_keys {
		if !CheckIfCharInString(char, k) {
			two = k
			break
		}
	}

	return two
}

func ThreeAndFive(two string, undefined_keys map[string]int) (string, string) {
	three := ""
	five := ""

	for k, _ := range undefined_keys {
		diff := StringDiff(k, two)

		if len(diff) == 1 {
			three = k
		} else {
			five = k
		}
	}

	return three, five
}

func RemoveAllCharsFromString(str1 string, str2 string) string {
	for _, str2_char := range str2 {
		str1 = strings.ReplaceAll(str1, string(str2_char), "")
	}

	return str1
}

func CheckIfCharInString(char string, str string) bool {
	for _, str_char := range str {
		if string(str_char) == char {
			return true
		}
	}

	return false
}

func CalculateOutput(output []string, dict map[string]int) int {
	sum := 0
	num_str := ""

	for i := 0; i < len(output); i++ {
		key := SortStringByCharacter(output[i])
		num_str += strconv.Itoa(dict[key])
	}

	num, err := strconv.Atoi(num_str)

	if err != nil {
		panic(err)
	}

	sum += num

	return sum
}
