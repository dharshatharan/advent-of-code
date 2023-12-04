package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

const gridSize = 140
const inputFile = "./input.txt"

func readInput() [][]rune {
	// read input file to string
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	grid := make([][]rune, gridSize)
	for i, line := range strings.Split(string(input), "\n") {
		if i == gridSize {
			break
		}
		grid[i] = make([]rune, gridSize)
		for j, char := range line {
			if j == gridSize {
				break
			}
			grid[i][j] = char
		}
	}
	return grid
}

func parseNumberAt(input [][]rune, x, y int) (number int, start int, length int) {
	if !unicode.IsDigit(input[x][y]) {
		return 0, y, 0
	}

	number = int(input[x][y]) - '0'
	start = y
	length = 1

	// search left
	for i := y - 1; i >= 0 && unicode.IsDigit(input[x][i]); i-- {
		number = (int(input[x][i])-'0')*int(math.Pow10(y-i)) + number
		start = i
		length++
	}

	// search right
	for i := y + 1; i < gridSize && unicode.IsDigit(input[x][i]); i++ {
		number = number*10 + int(input[x][i]) - '0'
		length++
	}

	return number, start, length
}

func isPartNumber(input [][]rune, x, y, length int) bool {
	if !unicode.IsDigit(input[x][y]) {
		return false
	}

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+length; j++ {
			if i < 0 || i >= gridSize || j < 0 || j >= gridSize {
				continue
			}
			if i == x && j >= y && j < y+length {
				continue
			}
			char := input[i][j]
			if char != '.' && !unicode.IsDigit(char) {
				return true
			}
		}
	}
	return false

}

func part1(input [][]rune) (result int) {
	totalSum := 0
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			char := input[x][y]
			if unicode.IsDigit(char) {
				number, start, length := parseNumberAt(input, x, y)
				if isPartNumber(input, x, y, length) {
					totalSum += number
				}
				if length > 1 {
					y = start + length - 1
				}
			}
		}
	}
	return totalSum
}

func buildStarMap(input [][]rune, x, y, length int, starMap map[string][]int, number int) {
	if !unicode.IsDigit(input[x][y]) {
		return
	}

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+length; j++ {
			if i < 0 || i >= gridSize || j < 0 || j >= gridSize {
				continue
			}
			if i == x && j >= y && j < y+length {
				continue
			}
			char := input[i][j]
			if char == '*' {
				starId := fmt.Sprintf("%d:%d", i, j)
				if _, ok := starMap[starId]; !ok {
					starMap[starId] = []int{number}
				} else {
					starMap[starId] = append(starMap[starId], number)
				}
			}
		}
	}
}

func part2(input [][]rune) (result int) {
	totalSum := 0
	starMap := make(map[string][]int)
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			char := input[x][y]
			if unicode.IsDigit(char) {
				number, start, length := parseNumberAt(input, x, y)
				buildStarMap(input, x, y, length, starMap, number)
				if length > 1 {
					y = start + length - 1
				}
			}
		}
	}
	for _, star := range starMap {
		if len(star) == 2 {
			totalSum += star[0] * star[1]
		}
	}
	return totalSum
}

func main() {
	input := readInput()
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
