package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var wins []int64
var spawned int

func main() {
	ReadFile()
}

func ReadFile() {
	file, err := os.Open("input_test.txt")

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

	wins = make([]int64, 2)
	score := []int{0, 0}
	playerCounter := 1
	spawned = 0

	Play(positions, score, playerCounter)
	fmt.Println("Wins:", wins)
	fmt.Println("Spawned:", spawned)
}

func Play(positions []int, score []int, playerIndex int) {
	spawned++
	//fmt.Println("Wins:", wins)

	for {
		sum := 0

		if score[0] >= 21 {
			wins[0]++
			return
		}
		if score[1] >= 21 {
			wins[1]++
			return
		}

		local_score := make([]int, 2)
		local_positions := make([]int, 2)

		if playerIndex == 1 {
			// roll die
			for roll_ndx := 0; roll_ndx < 3; roll_ndx++ {
				sum += roll_ndx + 1

				// spawn three universes
				for i := 1; i <= 3; i++ {
					copy(local_score, score)
					copy(local_positions, positions)

					new_position, new_score := GetNewScore(local_positions[0], i)
					local_positions[0] = new_position
					local_score[0] += new_score

					Play(local_positions, local_score, 2)
				}
			}

			new_position, new_score := GetNewScore(positions[0], sum)
			positions[0] = new_position
			score[0] += new_score
			playerIndex = 2
		} else {
			// roll die
			for roll_ndx := 0; roll_ndx < 3; roll_ndx++ {
				sum += roll_ndx + 1

				// spawn three universes
				for i := 1; i <= 3; i++ {
					copy(local_score, score)
					copy(local_positions, positions)

					new_position, new_score := GetNewScore(local_positions[1], i)
					local_positions[1] = new_position
					local_score[1] += new_score

					Play(local_positions, local_score, 1)
				}
			}

			new_position, new_score := GetNewScore(positions[1], sum)
			positions[1] = new_position
			score[1] += new_score
			playerIndex = 1
		}
	}
}

/*
1033337272
561514317
444356092776315
341960390180808
*/

func GetNewScore(position int, sum int) (int, int) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	position = (position + sum) % 10
	value := values[position]

	return position, value
}
