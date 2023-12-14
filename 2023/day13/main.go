package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput() []string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := strings.TrimSuffix(string(input), "\n")
	return strings.Split(inputStr, "\n\n")
}

type Tile struct {
	rows []string
	cols []string
}

func inputToTile(input string) Tile {
	rows := strings.Split(input, "\n")
	cols := make([]string, len(rows[0]))
	for i := 0; i < len(rows[0]); i++ {
		for j := 0; j < len(rows); j++ {
			cols[i] += string(rows[j][i])
		}
	}
	return Tile{rows, cols}
}

func equalsButOne(a, b string, diffLeft *int) bool {
	if *diffLeft <= 0 {
		return false
	}
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}
	if diff == 1 {
		*diffLeft--
		return true
	}
	return false
}

func customEquals(a, b string, alt bool, diffLeft *int) bool {
	if alt {
		return a == b || equalsButOne(a, b, diffLeft)
	}
	return a == b
}

func findReflection(input []string, alt bool) int {
	for i := 0; i < len(input)-1; i++ {
		diffLeft := 1
		if customEquals(input[i], input[i+1], alt, &diffLeft) {
			l := i - 1
			r := i + 1 + 1
			reflective := true
			for l >= 0 && r < len(input) {
				if !customEquals(input[l], input[r], alt, &diffLeft) {
					reflective = false
					break
				}
				l--
				r++
			}
			if reflective {
				if alt && diffLeft > 0 {
					continue
				}
				return i + 1
			}
		}
	}

	return -1
}

func findReflectionValue(input string, alt bool) int {
	tile := inputToTile(input)
	reflection := findReflection(tile.cols, alt)
	if reflection == -1 {
		reflection = findReflection(tile.rows, alt) * 100
		if reflection == -100 {
			panic("No reflection found")
		}
	}
	return reflection
}

func part1(inputs []string) int {
	total := 0
	for _, input := range inputs {
		total += findReflectionValue(input, false)
	}
	return total
}

func part2(inputs []string) int {
	total := 0
	for _, input := range inputs {
		total += findReflectionValue(input, true)
	}
	return total
}

func main() {
	inputs := parseInput()
	fmt.Println("Part 1:", part1(inputs))
	fmt.Println("Part 2:", part2(inputs))
}
