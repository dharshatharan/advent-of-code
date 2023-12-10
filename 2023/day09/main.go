package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput() (input [][]int) {
	inputStr, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range strings.Split(string(inputStr), "\n") {
		if line == "" {
			continue
		}
		sequence := []int{}
		for _, numStr := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(numStr)
			sequence = append(sequence, num)
		}
		input = append(input, sequence)
	}
	return input
}

func createSequenceTree(sequence []int) [][]int {
	sequences := [][]int{}
	sequences = append(sequences, sequence)
	for i := 0; i < len(sequence); i++ {
		nextSequence := []int{}
		for j := 0; j < len(sequences[i])-1; j++ {
			nextSequence = append(nextSequence, sequences[i][j+1]-sequences[i][j])
		}
		sequences = append(sequences, nextSequence)
		found := false
		for j := 0; j < len(nextSequence); j++ {
			if nextSequence[j] != 0 {
				found = true
				break
			}
		}
		if !found {
			break
		}
	}
	return sequences
}

func calculateNextValue(sequence []int) int {
	sequences := createSequenceTree(sequence)

	sequences[len(sequences)-1] = append(sequences[len(sequences)-1], 0)
	for i := len(sequences) - 1; i >= 1; i-- {
		num := sequences[i][len(sequences[i])-1] + sequences[i-1][len(sequences[i-1])-1]
		sequences[i-1] = append(sequences[i-1], num)
	}
	return sequences[0][len(sequences[0])-1]
}

func calculatePreviousValue(sequence []int) int {
	sequences := createSequenceTree(sequence)

	sequences[len(sequences)-1] = append([]int{0}, sequences[len(sequences)-1]...)
	for i := len(sequences) - 1; i >= 1; i-- {
		num := sequences[i-1][0] - sequences[i][0]
		sequences[i-1] = append([]int{num}, sequences[i-1]...)
	}
	return sequences[0][0]
}

func part1(input [][]int) int {
	total := 0
	for _, sequence := range input {
		total += calculateNextValue(sequence)
	}
	return total
}

func part2(input [][]int) int {
	total := 0
	for _, sequence := range input {
		total += calculatePreviousValue(sequence)
	}
	return total
}

func main() {
	input := parseInput()
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
