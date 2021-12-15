package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	x int
	y int
}
type NodeQueue struct {
	node     Node
	distance int
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

	start := Node{x: 0, y: 0}
	end := Node{x: len(risk_level[0]) - 1, y: len(risk_level) - 1}

	path := BFS(risk_level, start, end)
	fmt.Println(path)

	CalculatePath(risk_level, path)
}

func BFS(risk_level [][]int, start Node, end Node) []Node {
	tree := make(map[Node]Node, 0)
	var visited []Node
	var queue []NodeQueue

	queue = append(queue, NodeQueue{node: start, distance: risk_level[start.y][start.x]})

	for {
		if len(queue) == 0 {
			break
		}

		//fmt.Println(queue)

		currentQueueNode := queue[0]

		if currentQueueNode.node.x == end.x && currentQueueNode.node.y == end.y {
			//return currentQueueNode.distance
			fmt.Println(currentQueueNode.distance)
			fmt.Println("here")
			return buildPath(tree, end)
		}

		//queue = queue[:len(queue)-1]
		queue = queue[1:]
		visited = append(visited, currentQueueNode.node)

		// adjacent nodes
		adjacent_nodes := make([]Node, 0)

		// left
		if currentQueueNode.node.x-1 >= 0 {
			adjacent_nodes = append(adjacent_nodes, Node{x: currentQueueNode.node.x - 1, y: currentQueueNode.node.y})
		}

		// right
		if currentQueueNode.node.x+1 < len(risk_level[0]) {
			adjacent_nodes = append(adjacent_nodes, Node{x: currentQueueNode.node.x + 1, y: currentQueueNode.node.y})
		}

		// bottom
		if currentQueueNode.node.y+1 < len(risk_level) {
			adjacent_nodes = append(adjacent_nodes, Node{x: currentQueueNode.node.x, y: currentQueueNode.node.y + 1})
		}

		// top
		if currentQueueNode.node.y-1 >= 0 {
			adjacent_nodes = append(adjacent_nodes, Node{x: currentQueueNode.node.x, y: currentQueueNode.node.y - 1})
		}

		for i := 0; i < len(adjacent_nodes); i++ {
			if !CheckIfNodeWasVisited(visited, adjacent_nodes[i]) {
				tree[adjacent_nodes[i]] = currentQueueNode.node
				queue = append(queue, NodeQueue{node: adjacent_nodes[i], distance: currentQueueNode.distance + risk_level[adjacent_nodes[i].y][adjacent_nodes[i].x]})
			}
		}
	}

	temp := make([]Node, 0)
	return temp
	//return -1
}

func buildPath(tree map[Node]Node, end Node) []Node {
	path := make([]Node, 0)
	path = append(path, end)
	parent := tree[end]

	fmt.Println(tree)

	for {
		path = append(path, parent)
		val, ok := tree[parent]
		//fmt.Println(tree[parent])

		if !ok {
			break
		}

		parent = val
	}

	return path
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

func CalculatePath(risk_level [][]int, path []Node) {
	sum := 0

	for i := 0; i < len(path); i++ {
		sum += risk_level[path[i].y][path[i].x]
	}

	fmt.Println(sum)
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
