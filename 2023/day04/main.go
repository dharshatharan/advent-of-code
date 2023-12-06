package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	id             int
	winningNumbers []int
	cardNumbers    []int
}

func parseInput() []Card {

	// read input file to string
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := string(input)

	cards := []Card{}

	cardIdRegex := regexp.MustCompile(`^Card +(\d+):`)
	for _, line := range strings.Split(inputStr, "\n") {
		card := Card{}
		if cardIdRegex.MatchString(line) {
			cardId := cardIdRegex.FindStringSubmatch(line)[1]
			card.id, _ = strconv.Atoi(cardId)

			numbers := strings.Split(line, ":")[1]
			numbers = strings.TrimSpace(numbers)
			numbers = strings.ReplaceAll(numbers, "  ", " ")

			splitNumbers := strings.Split(numbers, " | ")
			winningNumbers := strings.Split(splitNumbers[0], " ")
			cardNumbers := strings.Split(splitNumbers[1], " ")

			for _, num := range winningNumbers {
				num, _ := strconv.Atoi(num)
				card.winningNumbers = append(card.winningNumbers, num)
			}
			for _, num := range cardNumbers {
				num, _ := strconv.Atoi(num)
				card.cardNumbers = append(card.cardNumbers, num)
			}

			cards = append(cards, card)
		}

	}

	return cards
}

func part1(cards []Card) int {
	totalPoints := 0

	for _, card := range cards {
		numberMatches := 0
		for _, num := range card.winningNumbers {
			if slices.Contains(card.cardNumbers, num) {
				numberMatches++
			}
		}
		cardPoints := math.Pow(2, float64(numberMatches-1))
		totalPoints += int(cardPoints)
	}
	return totalPoints
}

func part2(cards []Card) int {
	cardsCountMap := map[int]int{}

	for _, card := range cards {
		numberMatches := 0
		for _, num := range card.winningNumbers {
			if slices.Contains(card.cardNumbers, num) {
				numberMatches++
			}
		}
		cardsCountMap[card.id]++
		for i := 1; i <= numberMatches; i++ {
			cardsCountMap[card.id+i] += cardsCountMap[card.id]
		}
	}

	totalScratchCards := 0
	for _, count := range cardsCountMap {
		totalScratchCards += count
	}

	return totalScratchCards
}

func main() {
	cards := parseInput()
	fmt.Println("Part 1:", part1(cards))
	fmt.Println("Part 2:", part2(cards))
}
