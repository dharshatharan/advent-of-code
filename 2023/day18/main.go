package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Instruction struct {
	direction string
	distance  int
	color     string
}

func parseInput() []Instruction {
	input, err := os.ReadFile("./input-example.txt")
	if err != nil {
		panic(err)
	}
	inputStr := strings.TrimSuffix(string(input), "\n")
	lines := strings.Split(inputStr, "\n")

	re := regexp.MustCompile(`([RDLU]) (\d+) \((#[\d\w]{6})\)`)

	instructions := []Instruction{}
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		direction := matches[1]
		distance := int(matches[2][0] - '0')
		color := matches[3]
		instructions = append(instructions, Instruction{direction, distance, color})
	}

	return instructions
}

type Grid struct {
	grid map[[2]int]string
	xLen int
	yLen int
	xMin int
	yMin int
	xMax int
	yMax int
}

func generateGrid(instructions []Instruction) Grid {
	Grid := Grid{}
	Grid.grid = make(map[[2]int]string)

	curX := 0
	curY := 0
	minX := 0
	minY := 0
	maxX := 0
	maxY := 0
	for _, instruction := range instructions {
		switch instruction.direction {
		case "R":
			for i := 0; i < instruction.distance; i++ {
				curX++
				Grid.grid[[2]int{curX, curY}] = instruction.color
			}
		case "L":
			for i := 0; i < instruction.distance; i++ {
				curX--
				Grid.grid[[2]int{curX, curY}] = instruction.color
			}
		case "U":
			for i := 0; i < instruction.distance; i++ {
				curY--
				Grid.grid[[2]int{curX, curY}] = instruction.color
			}
		case "D":
			for i := 0; i < instruction.distance; i++ {
				curY++
				Grid.grid[[2]int{curX, curY}] = instruction.color
			}
		}
		if curX < minX {
			minX = curX
		}
		if curY < minY {
			minY = curY
		}
		if curX > maxX {
			maxX = curX
		}
		if curY > maxY {
			maxY = curY
		}
	}

	Grid.xLen = maxX - minX + 1
	Grid.yLen = maxY - minY + 1
	Grid.xMin = minX
	Grid.yMin = minY
	Grid.xMax = maxX
	Grid.yMax = maxY

	return Grid
}

func printGrid(Grid Grid) {
	for y := Grid.yMin; y <= Grid.yMax; y++ {
		for x := Grid.xMin; x <= Grid.xMax; x++ {
			if _, ok := Grid.grid[[2]int{x, y}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func calculateArea(Grid Grid) int {
	area := 0
	for y := Grid.yMin; y <= Grid.yMax; y++ {
		for x := Grid.xMin; x <= Grid.xMax; x++ {
			if _, ok := Grid.grid[[2]int{x, y}]; ok {
				area++
			}
		}
	}
	return area
}

func main() {
	instructions := parseInput()
	printGrid(generateGrid(instructions))
	fmt.Println(calculateArea(generateGrid(instructions)))
}
