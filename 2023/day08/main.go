package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Instruction struct {
	left  string
	right string
}

func parseInput() (nodeMap map[string]Instruction, instructions string) {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := string(input)

	split := strings.Split(inputStr, "\n\n")
	instructions = split[0]
	nodeStr := split[1]

	nodeMap = map[string]Instruction{}
	re := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)
	for _, line := range strings.Split(nodeStr, "\n") {
		if line == "" {
			continue
		}
		match := re.FindStringSubmatch(line)
		nodeMap[match[1]] = Instruction{match[2], match[3]}
	}

	return nodeMap, instructions
}

func part1(nodeMap map[string]Instruction, instructions string) int {
	currentNode := "AAA"
	steps := 0
	for i := 0; i < len(instructions); i++ {
		dir := instructions[i]
		if dir == 'L' {
			currentNode = nodeMap[currentNode].left
		} else if dir == 'R' {
			currentNode = nodeMap[currentNode].right
		} else {
			panic("Invalid direction")
		}
		steps++
		if currentNode == "ZZZ" {
			break
		}
		if i == len(instructions)-1 {
			i = -1
		}
	}

	return steps
}

func part2(nodeMap map[string]Instruction, instructions string) int {
	startingNodes := []string{}
	for k := range nodeMap {
		if k[2] == 'A' {
			startingNodes = append(startingNodes, k)
		}
	}

	allSteps := []int{}
	for _, node := range startingNodes {
		steps := 0
		currentNode := node
		for i := 0; i < len(instructions); i++ {
			dir := instructions[i]
			if dir == 'L' {
				currentNode = nodeMap[currentNode].left
			} else if dir == 'R' {
				currentNode = nodeMap[currentNode].right
			} else {
				panic("Invalid direction")
			}
			steps++
			if currentNode[2] == 'Z' {
				break
			}
			if i == len(instructions)-1 {
				i = -1
			}
		}
		allSteps = append(allSteps, steps)
	}

	return LCM(allSteps[0], allSteps[1], allSteps[2:]...)
}

func main() {
	nodeMap, instructions := parseInput()
	fmt.Println("Part 1:", part1(nodeMap, instructions))
	fmt.Println("Part 2:", part2(nodeMap, instructions))
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
