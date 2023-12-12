package main

import (
	"fmt"
	"os"
	"strings"
)

const gridSize = 140

func parseInput() (galaxies map[[2]int]bool) {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := string(input)

	galaxies = map[[2]int]bool{}
	for x, line := range strings.Split(inputStr, "\n") {
		if line == "" {
			continue
		}

		for y, char := range line {
			if char == '#' {
				galaxies[[2]int{x, y}] = true
			}
		}
	}

	return galaxies
}

func galaxyContains(galaxies map[[2]int]bool, coord [2]int) bool {
	if _, ok := galaxies[coord]; ok {
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func expand(galaxies map[[2]int]bool, multiple int) (newGalaxies map[[2]int]bool, expandedGridSize int) {
	rowsHasGalaxies := [gridSize]bool{}
	colsHasGalaxies := [gridSize]bool{}
	for coord, galaxy := range galaxies {
		if galaxy {
			rowsHasGalaxies[coord[0]] = true
			colsHasGalaxies[coord[1]] = true
		}
	}

	rowsToExpand := 0
	colsToExpand := 0
	newGalaxies = map[[2]int]bool{}
	for x := 0; x < gridSize; x++ {
		if !rowsHasGalaxies[x] {
			rowsToExpand += multiple - 1
		}
		for y := 0; y < gridSize; y++ {
			if !colsHasGalaxies[y] {
				colsToExpand += multiple - 1
			}
			if contains := galaxyContains(galaxies, [2]int{x, y}); contains {
				newGalaxies[[2]int{x + rowsToExpand, y + colsToExpand}] = true
			}
		}
		if x != gridSize-1 {
			colsToExpand = 0
		}
	}

	expandedGridSize = gridSize + max(colsToExpand, rowsToExpand)

	return newGalaxies, expandedGridSize
}

func printSpace(galaxies map[[2]int]bool, gridSize int) {
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			if contains := galaxyContains(galaxies, [2]int{x, y}); contains {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func calcSumOfDistances(galaxies map[[2]int]bool) (sum int) {
	completed := map[[2]int]bool{}
	for a := range galaxies {
		for b := range galaxies {
			if _, ok := completed[b]; ok {
				continue
			}
			if a == b {
				continue
			}
			sum += abs(a[0]-b[0]) + abs(a[1]-b[1])
		}
		completed[a] = true
	}
	return sum
}

func main() {
	galaxies := parseInput()
	// printSpace(galaxies, gridSize)
	// fmt.Println()
	part1Galaxies, _ := expand(galaxies, 2)
	// printSpace(galaxies, expandedGridSize)
	fmt.Println("Part 1:", calcSumOfDistances(part1Galaxies))

	part2Galaxies, _ := expand(galaxies, 1000000)
	fmt.Println("Part 2:", calcSumOfDistances(part2Galaxies))
}
