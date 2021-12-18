package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type pair struct {
	ndx   int
	start int
	end   int
}

func main() {
	ReadFile()
}

func ReadFile() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var snailfish_numbers [][]string

	for scanner.Scan() {
		line := scanner.Text()
		stack := make([]string, 0)

		for _, char := range line {
			if string(char) != "," {
				stack = append(stack, string(char))
			}
		}

		snailfish_numbers = append(snailfish_numbers, stack)
	}

	file.Close()

	fmt.Println(snailfish_numbers)
	fmt.Println("---------------")
	Calculate(snailfish_numbers)
	//CheckTheMagnitude(snailfish_numbers[0])
}

func Calculate(snailfish_numbers [][]string) {
	last_added_ndx := 0

	for {
		fmt.Println("--------")
		stack := snailfish_numbers[0]

		// infinite loop until the snailfish number is reduced
		for {
			// counters
			exploded_counter := 0
			splitted_counter := 0

			// go through the whole snailfish number
			nested_in_pairs := 0
			exploding_pair := false
			exploding_pair_start := 0
			new_stack := make([]string, 0)

			isPairGonnaExplode := CheckIfPairIsGonnaExplode(stack)
			fmt.Println("beginning", stack)

			for i, _ := range stack {
				current := stack[i]
				next := "["
				next_next := "["

				if i+1 < len(stack) {
					next = stack[i+1]
				}
				if i+2 < len(stack) {
					next_next = stack[i+2]
				}

				pairComesUpNext := next != "[" && next != "]" && next_next != "[" && next_next != "]"

				if current == "[" {
					nested_in_pairs++
				} else if current == "]" {
					nested_in_pairs--
				}

				// explode left
				if exploded_counter == 0 && splitted_counter == 0 && current == "[" && pairComesUpNext && nested_in_pairs >= 5 {
					exploding_pair = true
					exploding_pair_start = i
				} else if exploding_pair && current == "]" {
					new_stack, stack = ExplodePair(stack, new_stack, exploding_pair_start, i)
					exploding_pair = false
					exploded_counter++
				} else if !exploding_pair {
					if current != "[" && current != "]" {
						current_num, err := strconv.Atoi(current)

						if err != nil {
							panic(err)
						}

						if exploded_counter == 0 && splitted_counter == 0 && !isPairGonnaExplode && current_num >= 10 {
							new_stack = SplitARegularNumber(new_stack, current_num)
							splitted_counter++
						} else {
							new_stack = append(new_stack, current)
						}
					} else {
						new_stack = append(new_stack, current)
					}
				}
			}

			fmt.Println("new stack", new_stack)

			stack = new_stack

			if exploded_counter == 0 && splitted_counter == 0 {
				fmt.Println("breaking")
				break
			}
		}

		// add two numbers
		if last_added_ndx+1 < len(snailfish_numbers) {
			last_added_ndx = last_added_ndx + 1
			snailfish_numbers[0] = Addition(stack, snailfish_numbers[last_added_ndx])
			fmt.Println("Addition:", snailfish_numbers[0])
		} else {
			snailfish_numbers[0] = stack
			break
		}
	}

	CheckTheMagnitude(snailfish_numbers[0])
}

func CheckIfPairIsGonnaExplode(stack []string) bool {
	nested_in_pairs := 0

	for _, current := range stack {
		if current == "[" {
			nested_in_pairs++
		} else if current == "]" {
			nested_in_pairs--
		} else if nested_in_pairs > 4 {
			return true
		}
	}

	return false
}

func Addition(stack1 []string, stack2 []string) []string {
	new_stack := make([]string, 0)
	new_stack = append(new_stack, "[")
	new_stack = append(new_stack, stack1...)
	new_stack = append(new_stack, stack2...)
	new_stack = append(new_stack, "]")

	return new_stack
}

func ExplodePair(stack []string, new_stack []string, exploding_pair_start int, exploding_pair_end int) ([]string, []string) {
	first_prev_number_ndx := FindFirstPreviousNumber(new_stack, exploding_pair_start)
	first_next_number_ndx := FindFirstNumber(stack, exploding_pair_end)

	for i := exploding_pair_start; i < exploding_pair_end; i++ {
		current := stack[i]

		if i-exploding_pair_start == 1 {
			current_num, err := strconv.Atoi(current)

			if err != nil {
				panic(err)
			}

			if first_prev_number_ndx != -1 {
				number, err := strconv.Atoi(new_stack[first_prev_number_ndx])

				if err != nil {
					panic(err)
				}

				number = number + current_num
				new_stack[first_prev_number_ndx] = strconv.Itoa(number)
			}
		}
		if i-exploding_pair_start == 2 {
			current_num, err := strconv.Atoi(current)

			if err != nil {
				panic(err)
			}

			if first_next_number_ndx != -1 {
				number, err := strconv.Atoi(stack[first_next_number_ndx])

				if err != nil {
					panic(err)
				}

				number = number + current_num
				stack[first_next_number_ndx] = strconv.Itoa(number)
			}
		}
	}

	new_stack = append(new_stack, "0")

	return new_stack, stack
}

func SplitARegularNumber(new_stack []string, number int) []string {
	left_element := int(math.Floor(float64(number) / 2))
	right_element := int(math.Ceil(float64(number) / 2))

	new_stack = append(new_stack, "[")
	new_stack = append(new_stack, strconv.Itoa(left_element))
	new_stack = append(new_stack, strconv.Itoa(right_element))
	new_stack = append(new_stack, "]")

	return new_stack
}

func FindFirstPreviousNumber(stack []string, ndx int) int {
	number_ndx := -1

	for i := ndx - 1; i >= 0; i-- {
		if stack[i] != "[" && stack[i] != "]" {
			number_ndx = i
			break
		}
	}

	return number_ndx
}

func FindFirstNumber(stack []string, exploding_pair_end int) int {
	number_ndx := -1

	for i := exploding_pair_end; i < len(stack); i++ {
		if stack[i] != "[" && stack[i] != "]" {
			number_ndx = i
			break
		}
	}

	return number_ndx
}

func CheckTheMagnitude(stack []string) int {
	PrintStack(stack)
	magnitude := 0

	pairs := CountPairs(stack)

	if len(stack) == 1 {
		number, err := strconv.Atoi(stack[0])

		if err != nil {
			panic(err)
		}

		return number
	} else if stack[0] != "[" && stack[1] != "[" {
		current_num, err1 := strconv.Atoi(stack[0])
		next_num, err2 := strconv.Atoi(stack[1])

		if err1 != nil {
			panic(err1)
		}
		if err2 != nil {
			panic(err2)
		}

		fmt.Println(current_num, next_num, 3*current_num+2*next_num)

		return 3*current_num + 2*next_num
	} else {
		if len(pairs) == 1 {
			if pairs[0].end == len(stack)-1 {
				new_stack := stack[pairs[0].start+1 : pairs[0].end]
				magnitude += CheckTheMagnitude(new_stack)
			} else {
				left_stack := stack[pairs[0].start+1 : pairs[0].end]
				right_stack := stack[pairs[0].end+1:]

				fmt.Println("stacks", left_stack, right_stack)

				magnitude += 3*CheckTheMagnitude(left_stack) + 2*CheckTheMagnitude(right_stack)
			}
		} else {
			left_stack := stack[pairs[0].start+1 : pairs[0].end]
			right_stack := stack[pairs[1].start+1 : pairs[1].end]

			magnitude += 3*CheckTheMagnitude(left_stack) + 2*CheckTheMagnitude(right_stack)
		}
	}

	fmt.Println(magnitude)
	return magnitude
}

func CountPairs(stack []string) []pair {
	pairs := make([]pair, 0)
	pairs_counter := 0
	nest_level := 0

	for i, str := range stack {
		if str == "[" {
			if nest_level == 0 {
				pairs_counter++
				pairs = append(pairs, pair{ndx: pairs_counter, start: i, end: -1})
			}

			nest_level++
		} else if str == "]" {
			nest_level--

			if nest_level == 0 {
				pairs[pairs_counter-1].end = i
			}
		}
	}

	return pairs
}

func PrintStack(stack []string) {
	for _, str := range stack {
		fmt.Print(str, " ")
	}
	fmt.Println()
}
