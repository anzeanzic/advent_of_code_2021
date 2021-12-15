package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	x     int
	y     int
	value int
}
type NodeQueue struct {
	node     Node
	distance int
}
type ByCost []Node

func (a ByCost) Len() int           { return len(a) }
func (a ByCost) Less(i, j int) bool { return a[i].value < a[j].value }
func (a ByCost) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	ReadFile()
}

var adjacent_nodes []Node

func ReadFile() {
	file, err := os.Open("input_test2.txt")

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

	start := Node{x: 0, y: 0, value: risk_level[0][0]}
	end := Node{x: len(risk_level[0]) - 1, y: len(risk_level) - 1, value: risk_level[len(risk_level)-1][len(risk_level[0])-1]}
	adjacent_nodes = make([]Node, 0)

	ShortestPath(risk_level, start, end)
}

func ShortestPath(risk_level [][]int, start Node, end Node) {
	distances := make([][]int, len(risk_level))

	for i := 0; i < len(distances); i++ {
		distances[i] = make([]int, len(risk_level[0]))

		for j := 0; j < len(distances[i]); j++ {
			distances[i][j] = math.MaxInt32
		}
	}

	queue := make([]Node, 0)
	visited := make([]Node, 0)

	queue = append(queue, start)
	distances[start.y][start.x] = 0

	for {
		if len(visited) == len(risk_level)*len(risk_level[0]) {
			fmt.Println(len(visited), len(risk_level)*len(risk_level[0]))
			break
		}

		currentNode := queue[0]
		queue = queue[1:]

		if !CheckIfNodeWasVisited(visited, currentNode) {
			visited = append(visited, currentNode)
		}
		distances, queue = GetAdjacentNodes(risk_level, currentNode, visited, distances, queue)
	}

	Print(distances)
	//fmt.Println(visited)
}

func GetAdjacentNodes(risk_level [][]int, currentNode Node, visited []Node, distances [][]int, queue []Node) ([][]int, []Node) {
	edgeDistance := -1
	newDistance := 1

	// adjacent nodes
	adjacent_nodes = make([]Node, 0)

	// left
	if currentNode.x-1 >= 0 {
		adjacent_nodes = append(adjacent_nodes, Node{x: currentNode.x - 1, y: currentNode.y, value: risk_level[currentNode.y][currentNode.x]})
	}

	// right
	if currentNode.x+1 < len(risk_level[0]) {
		adjacent_nodes = append(adjacent_nodes, Node{x: currentNode.x + 1, y: currentNode.y, value: risk_level[currentNode.y][currentNode.x]})
	}

	// bottom
	if currentNode.y+1 < len(risk_level) {
		adjacent_nodes = append(adjacent_nodes, Node{x: currentNode.x, y: currentNode.y + 1, value: risk_level[currentNode.y][currentNode.x]})
	}

	// top
	if currentNode.y-1 >= 0 {
		adjacent_nodes = append(adjacent_nodes, Node{x: currentNode.x, y: currentNode.y - 1, value: risk_level[currentNode.y][currentNode.x]})
	}

	for i := 0; i < len(adjacent_nodes); i++ {
		if !CheckIfNodeWasVisited(visited, adjacent_nodes[i]) {
			edgeDistance = risk_level[adjacent_nodes[i].y][adjacent_nodes[i].x]
			newDistance = distances[currentNode.y][currentNode.x] + edgeDistance

			if adjacent_nodes[i].y == 9 && adjacent_nodes[i].x == 9 {
				fmt.Println(newDistance)
			}

			if newDistance < distances[adjacent_nodes[i].y][adjacent_nodes[i].x] {
				distances[adjacent_nodes[i].y][adjacent_nodes[i].x] = newDistance
			}

			queue = append(queue, adjacent_nodes[i])
		}
	}

	sort.Sort(ByCost(queue))

	return distances, queue
}

func CheckIfNodeWasVisited(visited []Node, adjacentNode Node) bool {
	isVisited := false

	for i := 0; i < len(visited); i++ {
		if visited[i].x == adjacentNode.x && visited[i].y == adjacentNode.y {
			isVisited = true
			break
		}
	}

	return isVisited
}

func Print(risk_level [][]int) {
	for y := 0; y < len(risk_level); y++ {
		for x := 0; x < len(risk_level[y]); x++ {
			fmt.Print(risk_level[y][x], " ")
		}

		fmt.Println()
	}

	fmt.Println("-------------------")
}
