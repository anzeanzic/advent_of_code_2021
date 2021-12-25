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
	instructions := make([][]string, 0)
	instruction_ndx := 0
	row_ndx := 0

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, " ")

		instructions = append(instructions, line_arr)

		row_ndx++
	}

	file.Close()

	fmt.Println(instructions)
	// 99999971020077
	// 13579246899999
	// 42221199523782
	for i := 13579246899999; i > 0; i = i - 100000 {
		if i%10 == 0 {
			continue
		}

		// 42221099584782

		z := ExecuteInstructions(instructions, i)
		fmt.Println(i, z)
		break
		if i%99999 == 0 {
			fmt.Println(i, z)
		}

		if z == 0 {
			fmt.Println("Largest num: ", i)
			break
		}
	}

	/*number_arr := []int{4, 2, 2, 2, 1, 1, 1, 1, 2, 1, 1, 1, 2, 2}

	for i := 0; i < len(number_arr); i++ {
		min_val := math.MaxInt16
		min_num := number_arr[i]
		fmt.Println(number_arr, number_arr[i])

		for num := number_arr[i]; num > 0; num-- {
			new_num, _ := join(number_arr)

			z := ExecuteInstructions(instructions, new_num)
			fmt.Println(i, num, z)

			if z <= min_val {
				min_val = z
				min_num = num
			}

			if z == 0 {
				fmt.Println("Largest num: ", i)
				break
			}
		}

		number_arr[i] = min_num
	}

	fmt.Println(number_arr)*/
}

func ExecuteInstructions(instructions [][]string, number int) int {
	m := make(map[string]int)
	number_ndx_counter := 0
	number_array := strings.Split(strconv.Itoa(number), "")

	for _, instruction := range instructions {
		instruction_name := instruction[0]

		switch {
		case instruction_name == "inp":
			var_name := instruction[1]
			num, err := strconv.Atoi(number_array[number_ndx_counter])

			if err != nil {
				panic(err)
			}

			m[var_name] = num
			number_ndx_counter++
		case instruction_name == "add":
			param_1 := instruction[1]
			addition := 0

			if val, err := strconv.Atoi(instruction[2]); err == nil {
				addition = val
			} else {
				addition = m[instruction[2]]
			}

			m[param_1] = m[param_1] + addition
		case instruction_name == "mul":
			param_1 := instruction[1]
			multiplier := 0

			if val, err := strconv.Atoi(instruction[2]); err == nil {
				multiplier = val
			} else {
				multiplier = m[instruction[2]]
			}

			m[param_1] = m[param_1] * multiplier
		case instruction_name == "div":
			param_1 := instruction[1]
			param_2, err := strconv.Atoi(instruction[2])

			if err != nil {
				panic(err)
			}

			if param_2 > 0 {
				m[param_1] = int(m[param_1] / param_2)
			}
		case instruction_name == "mod":
			param_1 := instruction[1]
			mod, err := strconv.Atoi(instruction[2])

			if err != nil {
				panic(err)
			}

			if m[param_1] >= 0 && mod > 0 {
				m[param_1] = m[param_1] % mod
			}
		case instruction_name == "eql":
			param_1 := instruction[1]
			eql := 0

			if val, err := strconv.Atoi(instruction[2]); err == nil {
				eql = val
			} else {
				eql = m[instruction[2]]
			}

			if m[param_1] == eql {
				m[param_1] = 1
			} else {
				m[param_1] = 0
			}
		}
	}

	//fmt.Println(m)

	return m["z"]
}

func join(nums []int) (int, error) {
	var str string
	for i := range nums {
		str += strconv.Itoa(nums[i])
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	} else {
		return num, nil
	}
}
