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

	for step := 0; step < 10; step++ {
		template = AsStepsGoBy(template, part_insertion_rules)
		//fmt.Println("template", template)
	}

	fmt.Println(len(template))

	CountChars(template)
}

func AsStepsGoBy(template string, part_insertion_rules []part_rule) string {
	inserted_chars := 0

	for i, _ := range template {
		if i+1+inserted_chars == len(template) {
			break
		} else {
			current_rune := template[i+inserted_chars]
			next_rune := template[i+1+inserted_chars]
			adjacent_runes := string(current_rune) + string(next_rune)

			inserted := CheckIfAnyRuleApplies(part_insertion_rules, adjacent_runes)

			if len(inserted) > 0 {
				//fmt.Println("insert", adjacent_runes, inserted)

				new_str := string(current_rune) + inserted + string(next_rune)
				runes := []rune(template)
				old_part := string(runes[0 : i+inserted_chars])
				new_part := string(runes[i+inserted_chars : len(runes)])

				new_part = strings.Replace(new_part, adjacent_runes, new_str, 1)
				template = old_part + new_part

				inserted_chars++
			}
		}

		//fmt.Println("-----------")
	}

	return template
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

func CountChars(template string) {
	m := make(map[string]int)

	for _, rune := range template {
		if _, ok := m[string(rune)]; ok {
			m[string(rune)]++
		} else {
			m[string(rune)] = 1
		}
	}

	// find min and max
	min := math.MaxInt16
	max := math.MinInt16

	for _, count := range m {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}

	fmt.Println(min, max)
	fmt.Println(max - min)
}
