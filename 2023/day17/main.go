package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	gridSize = 0
)

func parseInput() [][]int {
	input, err := os.ReadFile("./input-example.txt")
	if err != nil {
		panic(err)
	}
	inputStr := strings.TrimSuffix(string(input), "\n")
	lines := strings.Split(inputStr, "\n")
	gridSize = len(lines)
	grid := make([][]int, gridSize)
	for y, line := range lines {
		grid[y] = make([]int, gridSize)
		for x, numStr := range line {
			grid[y][x] = int(numStr - '0')
		}
	}
	return grid
}

func printGrid(grid [][]int, path [][2]int) {
	for y, row := range grid {
		for x, num := range row {
			// if num == 1 {
			// 	fmt.Printf("#")
			// 	continue
			// }
			found := false
			if path != nil {
				for _, coord := range path {
					if coord[0] == x && coord[1] == y {
						fmt.Printf(".")
						found = true
						continue
					}
				}
			}
			if !found {
				fmt.Printf("%d", num)
			}
		}
		fmt.Println()
	}
}

func distance(to, from [2]int) int {
	return (to[0] - from[0]) + (to[1] - from[1])
}

func heuristic(grid [][]int, coord [2]int) int {
	return grid[coord[1]][coord[0]]
}

func reconstructPath(cameFrom map[[2]int][2]int, current [2]int) [][2]int {
	totalPath := [][2]int{current}
	for {
		if _, ok := cameFrom[current]; !ok {
			break
		}
		current = cameFrom[current]
		totalPath = append([][2]int{current}, totalPath...)
	}
	return totalPath
}

func aStar(grid [][]int, start, goal [2]int) [][2]int {
	openSet := map[[2]int]bool{start: true}
	cameFrom := make(map[[2]int][2]int)
	gScore := make(map[[2]int]int)
	fScore := make(map[[2]int]int)
	gScore[start] = 0
	fScore[start] = heuristic(grid, start)
	for len(openSet) > 0 {
		var current [2]int
		var currentFScore int
		for k, v := range openSet {
			if !v {
				continue
			}
			if currentFScore == 0 || fScore[k] < currentFScore {
				current = k
				currentFScore = fScore[k]
			}
		}
		if current == goal {
			return reconstructPath(cameFrom, current)
		}
		openSet[current] = false
		for _, neighbor := range [][2]int{
			{current[0] - 1, current[1]},
			{current[0] + 1, current[1]},
			{current[0], current[1] - 1},
			{current[0], current[1] + 1},
		} {
			if neighbor[0] < 0 || neighbor[0] >= gridSize || neighbor[1] < 0 || neighbor[1] >= gridSize {
				continue
			}
			// if grid[neighbor[1]][neighbor[0]] == 1 {
			// 	continue
			// }
			tentativeGScore := gScore[current] + distance(current, neighbor)
			if _, ok := gScore[neighbor]; !ok || tentativeGScore < gScore[neighbor] {
				cameFrom[neighbor] = current
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = tentativeGScore + heuristic(grid, neighbor)
				if _, ok := openSet[neighbor]; !ok {
					openSet[neighbor] = true
				}
			}
		}
	}
	return nil
}

func main() {
	grid := parseInput()
	printGrid(grid, nil)
	fmt.Println()
	path := aStar(grid, [2]int{0, 0}, [2]int{gridSize - 1, gridSize - 1})
	printGrid(grid, path)
}
