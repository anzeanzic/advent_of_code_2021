package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	x     int
	y     int
	value int
}

var paths map[string]int

func main() {
	ReadFile()
}

func ReadFile() {
	file, err := os.Open("input_test.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var risk_level [][]int

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, "")

		var temp []int
		for _, num_str := range line_arr {
			number, _ := strconv.Atoi(num_str)
			temp = append(temp, number)
		}

		risk_level = append(risk_level, temp)
	}

	file.Close()

	Print(risk_level)

	paths = make(map[string]int)
	var visited []Node
	FindAllPaths(risk_level, 0, 0, visited)

	fmt.Println(paths)
}

func FindAllPaths(risk_level [][]int, x int, y int, visited []Node) {
	//fmt.Println(x, y)

	if x == len(risk_level[0])-1 && y == len(risk_level)-1 {
		visited = append(visited, Node{x: x, y: y, value: risk_level[y][x]})
		values := ""
		sum := 0

		for i := 0; i < len(visited); i++ {
			values += strconv.Itoa(visited[i].value)
			sum += visited[i].value
		}

		paths[values] = sum
		return
	}

	visited = append(visited, Node{x: x, y: y, value: risk_level[y][x]})

	// iterate over following nodes
	// left
	if x-1 >= 0 && !CheckIfNodeWasVisited(visited, x-1, y) {
		FindAllPaths(risk_level, x-1, y, visited)
	}

	// right
	if x+1 < len(risk_level[0]) && !CheckIfNodeWasVisited(visited, x+1, y) {
		FindAllPaths(risk_level, x+1, y, visited)
	}

	// bottom
	if y+1 < len(risk_level) && !CheckIfNodeWasVisited(visited, x, y+1) {
		FindAllPaths(risk_level, x, y+1, visited)
	}

	// top
	if y-1 >= 0 && !CheckIfNodeWasVisited(visited, x, y-1) {
		FindAllPaths(risk_level, x, y-1, visited)
	}
}

func CheckIfNodeWasVisited(visited []Node, x int, y int) bool {
	isVisited := false

	for i := 0; i < len(visited); i++ {
		if visited[i].x == x && visited[i].y == y {
			isVisited = true
			break
		}
	}

	return isVisited
}

func Print(risk_level [][]int) {
	for y := 0; y < len(risk_level); y++ {
		for x := 0; x < len(risk_level[y]); x++ {
			fmt.Print(risk_level[y][x], "")
		}

		fmt.Println()
	}

	fmt.Println("-------------------")
}
