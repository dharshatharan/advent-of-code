package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		re := regexp.MustCompile("[0-9]")
		numbers := re.FindAllString(line, -1)
		if config, err := strconv.Atoi(numbers[0] + numbers[len(numbers)-1]); err == nil {
			sum += config
		}
	}
	return sum
}

func part2(input string) int {
	sum := 0
	numberMap := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		for k, v := range numberMap {
			line = strings.ReplaceAll(line, k, v)
		}

		re := regexp.MustCompile("[0-9]")
		numbers := re.FindAllString(line, -1)
		if config, err := strconv.Atoi(numbers[0] + numbers[len(numbers)-1]); err == nil {
			sum += config
		}
	}
	return sum
}

func readInput() string {
	// read input file to string
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return string(input)
}

func main() {
	input := readInput()
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
