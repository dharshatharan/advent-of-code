package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseInput() []string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := strings.TrimSuffix(string(input), "\n")
	return strings.Split(inputStr, ",")
}

func hash(input string) int {
	h := 0
	for _, char := range input {
		ascii := int(char)
		h += ascii
		h *= 17
		h %= 256
	}
	return h
}

func part1(input []string) int {
	hashTotal := 0
	for _, str := range input {
		hashTotal += hash(str)
	}
	return hashTotal
}

type Lens struct {
	label       string
	focalLength int
}

func part2(input []string) int {
	hashmap := make(map[int][]Lens)
	addRe := regexp.MustCompile(`(\w+)=(\d)`)
	removeRe := regexp.MustCompile(`(\w+)-`)
	for _, str := range input {
		addMatches := addRe.FindStringSubmatch(str)
		if len(addMatches) > 0 {
			label := addMatches[1]
			focalLength, _ := strconv.Atoi(addMatches[2])
			hashVal := hash(label)
			if _, ok := hashmap[hashVal]; ok {
				lenses := hashmap[hashVal]
				found := false
				for idx, lens := range lenses {
					if lens.label == label {
						lenses[idx].focalLength = focalLength
						found = true
						break
					}

				}
				if !found {
					hashmap[hashVal] = append(hashmap[hashVal], Lens{label, focalLength})
				}
			} else {
				hashmap[hashVal] = []Lens{{label, focalLength}}
			}
		} else {
			removeMatches := removeRe.FindStringSubmatch(str)
			label := removeMatches[1]
			hashVal := hash(label)
			if _, ok := hashmap[hash(label)]; ok {
				lenses := hashmap[hashVal]
				for idx, lens := range lenses {
					if lens.label == label {
						lenses = append(lenses[:idx], lenses[idx+1:]...)
						hashmap[hashVal] = lenses
						break
					}
				}
			}
		}
	}

	focusPower := 0
	for box, lenses := range hashmap {
		for slot, lens := range lenses {
			focusPower += (box + 1) * (slot + 1) * lens.focalLength
		}
	}

	return focusPower
}

func printHashmap(hashmap map[int][]Lens) {
	for key, lenses := range hashmap {
		fmt.Println(key, lenses)
	}
}

func main() {
	input := parseInput()
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
