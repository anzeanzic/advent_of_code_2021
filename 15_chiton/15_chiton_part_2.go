package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	x        int
	y        int
	value    int
	distance int
	index    int
}
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]

	return item
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func main() {
	ReadFile()
}

var adjacent_nodes []Node

func ReadFile() {
	file, err := os.Open("input.txt")

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
	risk_level = ExtendRiskLevelMap(risk_level)
	//Print(risk_level)

	start := Node{x: 0, y: 0, value: risk_level[0][0], distance: 0}
	end := Node{x: len(risk_level[0]) - 1, y: len(risk_level) - 1, value: risk_level[len(risk_level)-1][len(risk_level[0])-1], distance: 0}
	adjacent_nodes = make([]Node, 0)

	ShortestPath(risk_level, start, end)
}

func ExtendRiskLevelMap(risk_level [][]int) [][]int {
	new_risk_level := make([][]int, len(risk_level)*5)

	for i := 0; i < len(new_risk_level); i++ {
		new_risk_level[i] = make([]int, len(risk_level[0])*5)
	}

	for y := 0; y < len(new_risk_level); y++ {
		for x := 0; x < len(new_risk_level[y]); x++ {
			map_y := int(math.Floor(float64(y) / float64(len(risk_level))))
			map_x := int(math.Floor(float64(x) / float64(len(risk_level[0]))))
			new_y := y % len(risk_level)
			new_x := x % len(risk_level[0])
			new_value := risk_level[new_y][new_x] + map_x + map_y

			if new_value > 9 {
				new_value = new_value % 9
			}

			new_risk_level[y][x] = new_value
		}
	}

	return new_risk_level
}

func ShortestPath(risk_level [][]int, start Node, end Node) {
	// dijkstra
	distances := make([][]int, len(risk_level))

	for i := 0; i < len(distances); i++ {
		distances[i] = make([]int, len(risk_level[0]))

		for j := 0; j < len(distances[i]); j++ {
			distances[i][j] = math.MaxInt32
		}
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Node{x: start.x, y: start.y, value: start.value, distance: start.distance})

	visited := make(map[string]int, 0)
	distances[start.y][start.x] = 0

	for {
		/*if pq.Len()%100 == 0 {
			fmt.Println(pq.Len())
		}*/
		fmt.Println(len(visited))

		//if len(visited) == len(risk_level)*len(risk_level[0]) {
		if pq.Len() == 0 {
			fmt.Println(len(visited), len(risk_level)*len(risk_level[0]))
			break
		}

		currentNode := heap.Pop(&pq).(*Node)

		if currentNode.x == end.x && currentNode.y == end.y {
			break
		}

		if !CheckIfNodeWasVisited(visited, *currentNode) {
			visited[strconv.Itoa(currentNode.y)+strconv.Itoa(currentNode.x)] = 1
			distances = GetAdjacentNodes(risk_level, currentNode, visited, distances, &pq)
		}
	}

	Print(distances)
}

func GetAdjacentNodes(risk_level [][]int, currentNode *Node, visited map[string]int, distances [][]int, pq *PriorityQueue) [][]int {
	edgeDistance := -1
	newDistance := 1

	// adjacent nodes
	adjacent_nodes = make([]Node, 0)

	// left
	if currentNode.x-1 >= 0 {
		adjacent_nodes = append(adjacent_nodes, Node{x: currentNode.x - 1, y: currentNode.y, value: risk_level[currentNode.y][currentNode.x], distance: 0})
	}

	// right
	if currentNode.x+1 < len(risk_level[0]) {
		adjacent_nodes = append(adjacent_nodes, Node{x: currentNode.x + 1, y: currentNode.y, value: risk_level[currentNode.y][currentNode.x], distance: 0})
	}

	// bottom
	if currentNode.y+1 < len(risk_level) {
		adjacent_nodes = append(adjacent_nodes, Node{x: currentNode.x, y: currentNode.y + 1, value: risk_level[currentNode.y][currentNode.x], distance: 0})
	}

	// top
	if currentNode.y-1 >= 0 {
		adjacent_nodes = append(adjacent_nodes, Node{x: currentNode.x, y: currentNode.y - 1, value: risk_level[currentNode.y][currentNode.x], distance: 0})
	}

	for i := 0; i < len(adjacent_nodes); i++ {
		if !CheckIfNodeWasVisited(visited, adjacent_nodes[i]) {
			edgeDistance = risk_level[adjacent_nodes[i].y][adjacent_nodes[i].x]
			newDistance = distances[currentNode.y][currentNode.x] + edgeDistance
			adjacent_nodes[i].distance = newDistance

			if newDistance < distances[adjacent_nodes[i].y][adjacent_nodes[i].x] {
				distances[adjacent_nodes[i].y][adjacent_nodes[i].x] = newDistance
			}

			heap.Push(pq, &Node{x: adjacent_nodes[i].x, y: adjacent_nodes[i].y, value: adjacent_nodes[i].value, distance: adjacent_nodes[i].distance})
		}
	}

	return distances
}

func CheckIfNodeWasVisited(visited map[string]int, adjacentNode Node) bool {
	_, isVisited := visited[strconv.Itoa(adjacentNode.y)+strconv.Itoa(adjacentNode.x)]
	//fmt.Println("adding", adjacentNode, isVisited)

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
