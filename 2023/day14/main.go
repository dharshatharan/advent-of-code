package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	empty = iota
	round = iota
	cube  = iota

	north = iota
	south = iota
	east  = iota
	west  = iota
)

var gridSize = 0

func parseInput() map[[2]int]int {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := strings.TrimSuffix(string(input), "\n")
	lines := strings.Split(inputStr, "\n")
	m := make(map[[2]int]int)
	gridSize = 0
	for y, line := range lines {
		gridSize++
		for x, char := range line {
			if char == '#' {
				m[[2]int{x, y}] = cube
			} else if char == '.' {
				m[[2]int{x, y}] = empty
			} else if char == 'O' {
				m[[2]int{x, y}] = round
			} else {
				panic("unexpected character")
			}
		}
	}
	return m
}

func getCoords(x, y, dir int) (int, int) {
	switch dir {
	case north:
		if y == 0 {
			return x, y
		}
		return x, y - 1
	case south:
		if y == gridSize-1 {
			return x, y
		}
		return x, y + 1
	case east:
		if x == gridSize-1 {
			return x, y
		}
		return x + 1, y
	case west:
		if x == 0 {
			return x, y
		}
		return x - 1, y
	default:
		panic("unexpected direction")
	}
}

func moveRoundedRock(grid map[[2]int]int, X, Y, dir int) {
	moved := false
	for moved == false {
		x, y := getCoords(X, Y, dir)
		if x == X && y == Y {
			return
		}
		if grid[[2]int{x, y}] == cube {
			return
		}
		if grid[[2]int{x, y}] == round {
			if moved {
				return
			}
			moveRoundedRock(grid, x, y, dir)
			moved = true
		}
		if grid[[2]int{x, y}] == empty {
			grid[[2]int{x, y}] = round
			grid[[2]int{X, Y}] = empty
			X, Y = x, y
		}
	}
}

func tilt(grid map[[2]int]int, dir int) {
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if grid[[2]int{x, y}] == round {
				moveRoundedRock(grid, x, y, dir)
			}
		}
	}
}

func calcLoad(grid map[[2]int]int) int {
	load := 0
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if grid[[2]int{x, y}] == round {
				load += gridSize - y
			}
		}
	}
	return load
}

var cache = make(map[string]int)

func detectCycle(grid map[[2]int]int) (bool, int, int) {
	gridStr := getGrid(grid)
	if _, ok := cache[gridStr]; ok {
		return true, cache[gridStr], (len(cache) - cache[gridStr])
	}
	cache[gridStr] = len(cache)
	return false, -1, -1
}

func getGrid(grid map[[2]int]int) string {
	gridStr := ""
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if grid[[2]int{x, y}] == cube {
				gridStr += "#"
			} else if grid[[2]int{x, y}] == round {
				gridStr += "O"
			} else {
				gridStr += "."
			}
		}
		gridStr += "\n"
	}
	return gridStr
}

func printGrid(grid map[[2]int]int) {
	fmt.Println(getGrid(grid))
}

func part1(grid map[[2]int]int) int {
	tilt(grid, north)
	return calcLoad(grid)
}

func part2(grid map[[2]int]int) int {
	results := []int{}
	for i := 0; i < 1000000000; i++ {
		tilt(grid, north)
		tilt(grid, west)
		tilt(grid, south)
		tilt(grid, east)
		cycle, cycleStart, cycleLen := detectCycle(grid)
		if cycle {
			return results[cycleStart+(1000000000-cycleStart)%cycleLen-1]
		}
		results = append(results, calcLoad(grid))
	}
	printGrid(grid)
	return calcLoad(grid)
}

func main() {
	grid := parseInput()
	fmt.Println("Part 1:", part1(grid))
	grid = parseInput()
	fmt.Println("Part 2:", part2(grid))
}
