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

type combination struct {
	x int
	y int
	z int
	w int
}

func ReadFile() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	instructions := make([][][]string, 0)
	instruction_ndx := -1

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, " ")

		if line_arr[0] == "inp" {
			instruction_ndx++
			instructions = append(instructions, make([][]string, 0))
			//instructions[instruction_ndx] = make([][]string, 0)
		}

		instructions[instruction_ndx] = append(instructions[instruction_ndx], line_arr)
	}

	file.Close()

	//fmt.Println(instructions)

	// 99999971020077
	// 13579246899999
	// 42221199523782
	var m map[combination]int
	var previous_m map[combination]int

	for inst_ndx := len(instructions) - 1; inst_ndx >= 0; inst_ndx-- {
		//fmt.Println(inst_ndx)
		m = make(map[combination]int)
		arr := []int{0, 0, 0, 1} // x, y, z, w

		for x := 0; x < 1; x++ {
			for y := 0; y < 20; y++ {
				for z := 0; z < 20; z++ {
					for w := 1; w <= 9; w++ {
						arr = []int{x, y, z, w}
						new_m := ExecuteInstructions(instructions[inst_ndx], arr)

						if inst_ndx == len(instructions)-1 && new_m["z"] == 0 {
							m[combination{x: x, y: y, z: z, w: w}] = 1
						} else {
							if _, ok := previous_m[combination{x: new_m["x"], y: new_m["y"], z: new_m["z"], w: new_m["w"]}]; ok {
								m[combination{x: x, y: y, z: z, w: w}] = 1
							}
						}
					}
				}
			}
		}

		previous_m = m
		fmt.Println(len(m))
		fmt.Println(m)
	}
}

func ExecuteInstructions(instructions [][]string, arr []int) map[string]int {
	m := make(map[string]int)
	m["x"] = arr[0]
	m["y"] = arr[1]
	m["z"] = arr[2]
	m["w"] = arr[3]

	for _, instruction := range instructions {
		instruction_name := instruction[0]

		switch {
		case instruction_name == "inp":
			var_name := instruction[1]
			num := m["w"]

			m[var_name] = num
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

	return m
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
