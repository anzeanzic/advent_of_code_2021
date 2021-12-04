package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bingo_board struct {
	board [5][5]int
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

	var numbers_drawn []int
	var bingo_boards []bingo_board

	line_counter := 0
	bingo_board_index := -1
	bingo_board_row_index := 0
	var current_bingo_board bingo_board

	for scanner.Scan() {
		line := scanner.Text()

		// this are the numbers drawn
		if line_counter == 0 {
			line_arr := strings.Split(line, ",")

			for i := 0; i < len(line_arr); i++ {
				number, err := strconv.Atoi(line_arr[i])

				if err != nil {
					panic(err)
				}

				numbers_drawn = append(numbers_drawn, number)
			}
		} else {
			if len(line) == 0 {
				bingo_board_index++
				bingo_board_row_index = 0

				if bingo_board_index > 0 {
					bingo_boards = append(bingo_boards, current_bingo_board)
				}
			} else {
				line_arr := strings.Fields(line)

				for i := 0; i < len(line_arr); i++ {
					number, err := strconv.Atoi(line_arr[i])

					if err != nil {
						panic(err)
					}

					current_bingo_board.board[bingo_board_row_index][i] = number
				}

				bingo_board_row_index++
			}
		}

		line_counter++
	}

	file.Close()

	bingo_boards = append(bingo_boards, current_bingo_board)

	PlayBingo(numbers_drawn, bingo_boards)
}

func PlayBingo(numbers_drawn []int, bingo_boards []bingo_board) {
	boardAlreadyWon := make([]bool, len(bingo_boards))
	lastWinnerBoardNdx := -1
	lastWinningNumber := -1

	for i := 0; i < len(numbers_drawn); i++ {
		current_number := numbers_drawn[i]

		// go through all bingo boards
		for bingo_ndx := 0; bingo_ndx < len(bingo_boards); bingo_ndx++ {
			if boardAlreadyWon[bingo_ndx] {
				continue
			}

			current_bingo_board := &bingo_boards[bingo_ndx].board

			// go through all rows and columns horizontally
			for row_ndx := 0; row_ndx < len(current_bingo_board); row_ndx++ {
				marked_count := 0

				for col_ndx := 0; col_ndx < len(current_bingo_board[row_ndx]); col_ndx++ {
					if current_bingo_board[row_ndx][col_ndx] == current_number {
						current_bingo_board[row_ndx][col_ndx] = -1
					}

					if current_bingo_board[row_ndx][col_ndx] == -1 {
						marked_count++
					}
				}

				if marked_count == 5 {
					fmt.Println("he wins")
					lastWinnerBoardNdx = bingo_ndx
					lastWinningNumber = current_number
					boardAlreadyWon[bingo_ndx] = true
					break
				}
			}

			// check if all boards won
			wins := CheckIfAllBoardsWon(boardAlreadyWon)

			if wins == len(boardAlreadyWon) {
				break
			}

			// go through all rows and columns vertically
			for col_ndx := 0; col_ndx < len(current_bingo_board); col_ndx++ {
				marked_count := 0

				for row_ndx := 0; row_ndx < len(current_bingo_board); row_ndx++ {
					if current_bingo_board[row_ndx][col_ndx] == current_number {
						current_bingo_board[row_ndx][col_ndx] = -1
					}

					if current_bingo_board[row_ndx][col_ndx] == -1 {
						marked_count++
					}
				}

				if marked_count == 5 {
					fmt.Println("he wins")
					lastWinnerBoardNdx = bingo_ndx
					lastWinningNumber = current_number
					boardAlreadyWon[bingo_ndx] = true
					break
				}
			}

			wins = CheckIfAllBoardsWon(boardAlreadyWon)

			if wins == len(boardAlreadyWon) {
				break
			}

		}

		PrintAllBoards(bingo_boards)
		fmt.Println("----------------------")
	}

	fmt.Println(lastWinnerBoardNdx, lastWinningNumber)
	sum := CalculateSumOfAllUnmarked(bingo_boards[lastWinnerBoardNdx].board)
	fmt.Println(sum * lastWinningNumber)
}

func CheckIfAllBoardsWon(boardAlreadyWon []bool) int {
	// check if all boards won
	wins := 0

	for i := 0; i < len(boardAlreadyWon); i++ {
		if boardAlreadyWon[i] {
			wins++
		}
	}

	fmt.Println(wins)

	return wins
}

func CalculateSumOfAllUnmarked(board [5][5]int) int {
	sum := 0

	for j := 0; j < len(board); j++ {
		for k := 0; k < len(board[j]); k++ {
			if board[j][k] != -1 {
				sum += board[j][k]
			}
		}
	}

	return sum
}

func PrintAllBoards(bingo_boards []bingo_board) {
	for i := 0; i < len(bingo_boards); i++ {
		for j := 0; j < len(bingo_boards[i].board); j++ {
			for k := 0; k < len(bingo_boards[i].board[j]); k++ {
				fmt.Print(strconv.Itoa(bingo_boards[i].board[j][k]) + " ")
			}

			fmt.Println()
		}

		fmt.Println()
	}
}
