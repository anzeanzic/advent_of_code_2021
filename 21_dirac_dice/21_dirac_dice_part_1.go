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
	positions := make([]int, 2)
	row_cnt := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")
		line_arr := strings.Split(line, ":")
		num, err := strconv.Atoi(line_arr[1])

		if err != nil {
			panic(err)
		}

		positions[row_cnt] = num - 1
		row_cnt++
	}

	file.Close()

	fmt.Println(positions)

	Play(positions)
}

func Play(positions []int) {
	dieRolls := 0
	dieCounter := 1
	playerCounter := 1
	score := []int{0, 0}
	losingPlayerIndex := -1

	for {
		sum := 0

		// roll die
		for i := 0; i < 3; i++ {
			sum += dieCounter
			dieCounter++
			dieRolls++
		}

		if playerCounter == 1 {
			new_position, new_score := GetNewScore(positions[0], sum)
			positions[0] = new_position
			score[0] += new_score

			//fmt.Println("Player 1 adds", sum, "to move to", positions[0]+1, "and have score", score[0])
			playerCounter = 2
		} else {
			new_position, new_score := GetNewScore(positions[1], sum)
			positions[1] = new_position
			score[1] += new_score

			//fmt.Println("Player 2 adds", sum, "to move to", positions[1]+1, "and have score", score[1])
			playerCounter = 1
		}

		if score[0] >= 1000 {
			losingPlayerIndex = 1
			break
		} else if score[1] >= 1000 {
			losingPlayerIndex = 0
			break
		}
	}

	fmt.Println(score, dieRolls)
	fmt.Println("Multiply", score[losingPlayerIndex]*dieRolls)
}

func GetNewScore(position int, sum int) (int, int) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	position = (position + sum) % 10
	value := values[position]

	return position, value
}
