// Was stumpped on how to solve this problem. I had to look up the solution.
// Credits to https://www.youtube.com/watch?v=g3Ms5e7Jdqo
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	cfg  string
	nums []int
}

func parseInput() (input []Line) {
	inputStr, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range strings.Split(string(inputStr), "\n") {
		if line == "" {
			continue
		}
		split := strings.Split(line, " ")
		cfg := split[0]
		nums := []int{}
		for _, numStr := range strings.Split(split[1], ",") {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}
		input = append(input, Line{cfg, nums})
	}
	return input
}

var cache = map[string]int{}

func getKey(line Line) string {
	return fmt.Sprintf("%v %v", line.cfg, line.nums)
}

func count(line Line) int {
	// fmt.Println(len(cache))
	key := getKey(line)
	if val, ok := cache[key]; ok {
		return val
	}

	result := 0

	if line.cfg == "" {
		if len(line.nums) == 0 {
			return 1
		}
		return 0
	}

	if len(line.nums) == 0 {
		if strings.Contains(line.cfg, "#") {
			// fmt.Println(line.cfg)
			return 0
		}
		return 1
	}

	if line.cfg[0] == '.' || line.cfg[0] == '?' {
		result += count(Line{line.cfg[1:], line.nums})
	}

	if line.cfg[0] == '#' || line.cfg[0] == '?' {
		if line.nums[0] <= len(line.cfg) &&
			!strings.Contains(line.cfg[:line.nums[0]], ".") &&
			(line.nums[0] == len(line.cfg) || line.cfg[line.nums[0]] != '#') {

			if line.nums[0] == len(line.cfg) {
				result += count(Line{line.cfg[line.nums[0]:], line.nums[1:]})
			} else {
				result += count(Line{line.cfg[line.nums[0]+1:], line.nums[1:]})
			}
		}
	}

	cache[key] = result
	return result
}

func part1(input []Line) int {
	total := 0
	for _, line := range input {
		total += count(line)
	}
	return total
}

func part2(input []Line) int {
	total := 0
	for _, line := range input {
		// duplicate cfg 5 times
		line.cfg = fmt.Sprintf("%v?%v?%v?%v?%v", line.cfg, line.cfg, line.cfg, line.cfg, line.cfg)

		// duplicate nums 5 times
		nums := make([]int, len(line.nums))
		copy(nums, line.nums)
		for i := 1; i < 5; i++ {
			line.nums = append(line.nums, nums...)
		}

		total += count(line)
	}
	return total
}

func main() {
	input := parseInput()
	fmt.Println("Part 1:", part1(input))
	cache = map[string]int{}
	fmt.Println("Part 2:", part2(input))
}
