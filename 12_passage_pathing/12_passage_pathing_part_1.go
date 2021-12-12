package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var paths [][]string

func main() {
	ReadFile()
}

func ReadFile() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	cave_nodes := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")
		line_arr := strings.Split(line, "-")

		if _, ok := cave_nodes[line_arr[0]]; !ok {
			cave_nodes[line_arr[0]] = make([]string, 0)
		}

		cave_nodes[line_arr[0]] = append(cave_nodes[line_arr[0]], line_arr[1])

		if _, ok := cave_nodes[line_arr[1]]; !ok {
			cave_nodes[line_arr[1]] = make([]string, 0)
		}

		cave_nodes[line_arr[1]] = append(cave_nodes[line_arr[1]], line_arr[0])
	}

	file.Close()

	fmt.Println(cave_nodes)

	visited := make([]string, 0)
	FindAllPaths(cave_nodes, "start", visited)

	fmt.Println("--------------")
	fmt.Println(len(paths))
	fmt.Println("--------------")

	for i := 0; i < len(paths); i++ {
		//fmt.Println(paths[i])
	}
}

func FindAllPaths(cave_nodes map[string][]string, node_name string, visited []string) {
	start_node := cave_nodes[node_name]

	if node_name == "end" {
		visited = append(visited, node_name)
		paths = append(paths, visited)
		return
	}

	visited = append(visited, node_name)

	// iterate over following nodes
	for i := 0; i < len(start_node); i++ {
		alreadyVisited := false

		// check if we already visited this node
		if len(visited) > 0 {
			for j := 0; j < len(visited); j++ {
				if visited[j] == start_node[i] && IsLower(start_node[i]) {
					alreadyVisited = true
				}
			}
		}

		if !alreadyVisited {
			FindAllPaths(cave_nodes, start_node[i], visited)
		}
	}
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}

	return true
}
