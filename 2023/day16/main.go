package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	north = [2]int{0, -1}
	south = [2]int{0, 1}
	east  = [2]int{1, 0}
	west  = [2]int{-1, 0}
)

type Tile struct {
	item        rune
	energized   bool
	enteredFrom map[[2]int]bool
}

var gridSize = 0

func parseInput() [][]Tile {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := strings.TrimSuffix(string(input), "\n")
	lines := strings.Split(inputStr, "\n")
	gridSize = len(lines)
	grid := make([][]Tile, gridSize)
	for y, line := range lines {
		grid[y] = make([]Tile, gridSize)
		for x, char := range line {
			grid[y][x] = Tile{char, false, make(map[[2]int]bool)}
		}
	}
	return grid
}

func printGrid(grid [][]Tile) {
	for _, row := range grid {
		for _, tile := range row {
			if tile.energized {
				fmt.Printf("#")
				continue
			}
			fmt.Printf("%c", tile.item)
		}
		fmt.Println()
	}
}

func energize(grid [][]Tile, x, y int, dir [2]int) {
	if x < 0 || x >= gridSize || y < 0 || y >= gridSize {
		return
	}
	tile := &grid[y][x]

	if tile.enteredFrom[dir] {
		return
	} else {
		tile.enteredFrom[dir] = true
	}

	tile.energized = true

	switch tile.item {
	case '.':
		energize(grid, x+dir[0], y+dir[1], dir)
	case '/':
		if dir == north {
			energize(grid, x+east[0], y+east[1], east)
		} else if dir == south {
			energize(grid, x+west[0], y+west[1], west)
		} else if dir == east {
			energize(grid, x+north[0], y+north[1], north)
		} else if dir == west {
			energize(grid, x+south[0], y+south[1], south)
		}
	case '\\':
		if dir == north {
			energize(grid, x+west[0], y+west[1], west)
		} else if dir == south {
			energize(grid, x+east[0], y+east[1], east)
		} else if dir == east {
			energize(grid, x+south[0], y+south[1], south)
		} else if dir == west {
			energize(grid, x+north[0], y+north[1], north)
		}
	case '|':
		if dir == north || dir == south {
			energize(grid, x+dir[0], y+dir[1], dir)
		} else {
			energize(grid, x+north[0], y+north[1], north)
			energize(grid, x+south[0], y+south[1], south)
		}
	case '-':
		if dir == east || dir == west {
			energize(grid, x+dir[0], y+dir[1], dir)
		} else {
			energize(grid, x+east[0], y+east[1], east)
			energize(grid, x+west[0], y+west[1], west)
		}
	default:
		panic("unexpected character")
	}
}

func countEnergized(grid [][]Tile) int {
	count := 0
	for _, row := range grid {
		for _, tile := range row {
			if tile.energized {
				count++
			}
		}
	}
	return count
}

func calcEnergized(x, y int, dir [2]int) int {
	grid := parseInput()
	energize(grid, x, y, dir)
	return countEnergized(grid)
}

func part1() int {
	return calcEnergized(0, 0, east)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func part2() int {
	maxCount := 0
	for z := 0; z < gridSize; z++ {
		maxCount = max(maxCount, calcEnergized(z, 0, south))
		maxCount = max(maxCount, calcEnergized(0, z, east))
		maxCount = max(maxCount, calcEnergized(gridSize-1, z, west))
		maxCount = max(maxCount, calcEnergized(z, gridSize-1, north))
	}
	return maxCount
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
