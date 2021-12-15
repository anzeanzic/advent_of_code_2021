package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type part_rule struct {
	adjacent string
	inserted string
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
	row_ndx := 0
	var template string
	var part_insertion_rules []part_rule

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")

		if row_ndx == 0 {
			template = line
		} else if len(line) > 0 {
			line_arr := strings.Split(line, "->")
			part_insertion_rules = append(part_insertion_rules, part_rule{adjacent: line_arr[0], inserted: line_arr[1]})
		}

		row_ndx++
	}

	file.Close()

	fmt.Println(template)
	fmt.Println(part_insertion_rules)

	adjacent_map := make(map[string]int64)
	adjacent_map2 := make(map[string]int64)

	for i, _ := range template {
		if i < len(template)-2 {
			current_rune := template[i]
			next_rune := template[i+1]
			adjacent_runes := string(current_rune) + string(next_rune)

			if _, ok := adjacent_map[adjacent_runes]; ok {
				adjacent_map[adjacent_runes]++
			} else {
				adjacent_map[adjacent_runes] = 1
			}
		}
	}
	last := string([]rune(template)[len(template)-2 : len(template)])

	fmt.Println(adjacent_map)

	for step := 0; step < 40; step++ {
		fmt.Println(step)
		adjacent_map, adjacent_map2, last = AsStepsGoBy(adjacent_map, adjacent_map2, last, part_insertion_rules)
		//fmt.Println("template", template)
		//fmt.Println(adjacent_map, last)
	}

	fmt.Println("---------")
	fmt.Println(len(template))
	fmt.Println(adjacent_map)

	CountChars(adjacent_map, last)
}

func AsStepsGoBy(primary_map map[string]int64, secondary_map map[string]int64, last string, part_insertion_rules []part_rule) (map[string]int64, map[string]int64, string) {
	primary := make(map[string]int64)
	secondary := make(map[string]int64)

	for key, count := range primary_map {
		inserted := CheckIfAnyRuleApplies(part_insertion_rules, key)

		if len(inserted) > 0 {
			key_arr := strings.Split(key, "")

			//fmt.Println("inserting", key_arr[0]+inserted, inserted+key_arr[1])

			if _, ok := primary[key_arr[0]+inserted]; ok {
				primary[key_arr[0]+inserted] += count * 1
			} else {
				primary[key_arr[0]+inserted] = count
			}

			if _, ok := secondary[inserted+key_arr[1]]; ok {
				secondary[inserted+key_arr[1]] += count * 1
			} else {
				secondary[inserted+key_arr[1]] = count
			}
		} else {
			if _, ok := primary[key]; ok {
				primary[key] += count * 1
			} else {
				primary[key] = count
			}
		}
	}

	for key, count := range secondary_map {
		inserted := CheckIfAnyRuleApplies(part_insertion_rules, key)

		if len(inserted) > 0 {
			key_arr := strings.Split(key, "")

			//fmt.Println("inserting", key_arr[0]+inserted, inserted+key_arr[1])

			if _, ok := primary[key_arr[0]+inserted]; ok {
				primary[key_arr[0]+inserted] += count * 1
			} else {
				primary[key_arr[0]+inserted] = count
			}

			if _, ok := secondary[inserted+key_arr[1]]; ok {
				secondary[inserted+key_arr[1]] += count * 1
			} else {
				secondary[inserted+key_arr[1]] = count
			}
		} else {
			if _, ok := primary[key]; ok {
				primary[key] += count * 1
			} else {
				primary[key] = count
			}
		}
	}

	// last
	if len(last) == 2 {
		inserted := CheckIfAnyRuleApplies(part_insertion_rules, last)

		if len(inserted) > 0 {
			key_arr := strings.Split(last, "")

			//fmt.Println("inserting last", key_arr[0]+inserted, inserted+key_arr[1])
			last = key_arr[0] + inserted + key_arr[1]
		}
	} else {
		for i, _ := range last {
			if i < len(last)-2 {
				current_rune := last[i]
				next_rune := last[i+1]
				adjacent_runes := string(current_rune) + string(next_rune)

				inserted := CheckIfAnyRuleApplies(part_insertion_rules, adjacent_runes)

				if len(inserted) > 0 {
					fmt.Println("inserting last", string(current_rune)+inserted, inserted+string(next_rune))

					if _, ok := primary[string(current_rune)+inserted]; ok {
						primary[string(current_rune)+inserted]++
					} else {
						primary[string(current_rune)+inserted] = 1
					}

					if _, ok := secondary[inserted+string(next_rune)]; ok {
						secondary[inserted+string(next_rune)]++
					} else {
						secondary[inserted+string(next_rune)] = 1
					}
				}
			} else if i < len(last)-1 {
				current_rune := last[i]
				next_rune := last[i+1]
				adjacent_runes := string(current_rune) + string(next_rune)

				inserted := CheckIfAnyRuleApplies(part_insertion_rules, adjacent_runes)

				if len(inserted) > 0 {
					last = string(current_rune) + inserted + string(next_rune)
				}
			}
		}
	}

	return primary, secondary, last
}

func CheckIfAnyRuleApplies(part_insertion_rules []part_rule, adjacent_runes string) string {
	inserted := ""

	for i := 0; i < len(part_insertion_rules); i++ {
		if part_insertion_rules[i].adjacent == adjacent_runes {
			inserted = part_insertion_rules[i].inserted
			break
		}
	}

	return inserted
}

func CountChars(adjacent_map map[string]int64, last string) {
	m := make(map[string]int64)

	for key, count := range adjacent_map {
		key_arr := strings.Split(key, "")

		if _, ok := m[key_arr[0]]; ok {
			m[key_arr[0]] += count
		} else {
			m[key_arr[0]] = count
		}

		if _, ok := m[key_arr[1]]; ok {
			m[key_arr[1]] += count
		} else {
			m[key_arr[1]] = count
		}
	}

	for _, rune := range last {
		if _, ok := m[string(rune)]; ok {
			m[string(rune)] += 1
		} else {
			m[string(rune)] = 1
		}
	}

	// find min and max
	var min int64
	var max int64

	min = math.MaxInt64
	max = math.MinInt64

	for _, count := range m {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}

	fmt.Println(m)

	fmt.Println(min, max)
	fmt.Println("diff", max-min)
}
