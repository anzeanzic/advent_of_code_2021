package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type dot struct {
	x int
	y int
}
type fold struct {
	axis  string
	value int
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
	var dots []dot
	var fold_instructions []fold
	x_max := math.MinInt16
	y_max := math.MinInt16

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, " ")

		if len(line) > 0 {
			if line_arr[0] != "fold" {
				dot_arr := strings.Split(line_arr[0], ",")
				x, err := strconv.Atoi(dot_arr[0])

				if err != nil {
					panic(err)
				}

				y, err := strconv.Atoi(dot_arr[1])

				if err != nil {
					panic(err)
				}

				dots = append(dots, dot{x: x, y: y})

				if x > x_max {
					x_max = x
				}
				if y > y_max {
					y_max = y
				}
			} else {
				instruction_arr := strings.Split(line_arr[2], "=")
				value, err := strconv.Atoi(instruction_arr[1])

				if err != nil {
					panic(err)
				}

				fold_instructions = append(fold_instructions, fold{axis: instruction_arr[0], value: value})
			}
		}
	}

	file.Close()

	fmt.Println(dots)
	fmt.Println(x_max, y_max)
	fmt.Println(fold_instructions)

	DrawPaper(dots, fold_instructions, x_max, y_max)
}

func DrawPaper(dots []dot, fold_instructions []fold, x_max int, y_max int) {
	paper := make([][]string, y_max+1)

	for y := 0; y < y_max+1; y++ {
		paper[y] = make([]string, x_max+1)

		for x := 0; x < x_max+1; x++ {
			paper[y][x] = "."
		}
	}

	for i := 0; i < len(dots); i++ {
		paper[dots[i].y][dots[i].x] = "X"
	}

	PrintPaper(paper)

	for i := 0; i < len(fold_instructions); i++ {
		paper = Fold(paper, fold_instructions[i])
	}

	//paper = Fold(paper, fold_instructions[1])
}

func Fold(paper [][]string, fold_instructions fold) [][]string {
	fmt.Println(fold_instructions)

	if fold_instructions.axis == "y" {
		// write the line
		for x := 0; x < len(paper[fold_instructions.value]); x++ {
			paper[fold_instructions.value][x] = "-"
		}

		// translate the points
		for y := fold_instructions.value + 1; y < len(paper); y++ {
			for x := 0; x < len(paper[y]); x++ {
				if paper[y][x] == "X" {
					y_diff := y - fold_instructions.value
					paper[fold_instructions.value-y_diff][x] = "X"
				}
			}
		}

		fmt.Println()
		PrintPaper(paper)

		// delete the rows
		for y := len(paper) - 1; y >= fold_instructions.value; y-- {
			paper = append(paper[:y], paper[y+1:]...)
		}

		fmt.Println()
		PrintPaper(paper)
		CountDots(paper)
	} else if fold_instructions.axis == "x" {
		// write the line
		for y := 0; y < len(paper); y++ {
			paper[y][fold_instructions.value] = "-"
		}

		// translate the points
		for y := 0; y < len(paper); y++ {
			for x := fold_instructions.value + 1; x < len(paper[y]); x++ {
				if paper[y][x] == "X" {
					x_diff := x - fold_instructions.value
					paper[y][fold_instructions.value-x_diff] = "X"
				}
			}
		}

		//fmt.Println()
		//PrintPaper(paper)

		// delete the rows
		for y := 0; y < len(paper); y++ {
			for x := len(paper[y]) - 1; x >= fold_instructions.value; x-- {
				paper[y] = append(paper[y][:x], paper[y][x+1:]...)
			}
		}

		fmt.Println()
		PrintPaper(paper)
		CountDots(paper)
	}

	return paper
}

func PrintPaper(paper [][]string) {
	for i := 0; i < len(paper); i++ {
		fmt.Println(paper[i])
	}
}

func CountDots(paper [][]string) {
	counter := 0

	for y := 0; y < len(paper); y++ {
		for x := 0; x < len(paper[y]); x++ {
			if paper[y][x] == "X" {
				counter++
			}
		}
	}

	fmt.Println(counter)
}
