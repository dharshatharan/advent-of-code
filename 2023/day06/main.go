package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLinePart1(line string) (items []int) {
	itemStrs := strings.Fields(line)[1:]
	for _, itemStr := range itemStrs {
		item, _ := strconv.Atoi(itemStr)
		items = append(items, item)
	}
	return items
}

func parseLinePart2(line string) (item []int) {
	line = strings.ReplaceAll(line, " ", "")
	itemStr := strings.Split(line, ":")[1]
	itemVal, _ := strconv.Atoi(itemStr)
	return []int{itemVal}
}

func parseInput(parseLine func(string) []int) (times, distances []int) {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := string(input)

	split := strings.Split(inputStr, "\n")

	timeStr := split[0]
	times = parseLine(timeStr)

	distanceStr := split[1]
	distances = parseLine(distanceStr)

	return times, distances
}

func calculateWaysToWin(times, distances []int) int {
	countProduct := 1
	for i := 0; i < len(times); i++ {
		count := 0
		for j := 0; j < times[i]; j++ {
			if (times[i]-j)*j > distances[i] {
				count++
			}
		}
		countProduct *= count
	}
	return countProduct
}

func main() {
	times, distances := parseInput(parseLinePart1)
	fmt.Println("Part 1:", calculateWaysToWin(times, distances))
	times, distances = parseInput(parseLinePart2)
	fmt.Println("Part 2:", calculateWaysToWin(times, distances))
}
