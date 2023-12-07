package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var jokerEnabled = false

type Hand struct {
	cards []string
	bid   int
}

func (h Hand) String() string {
	return fmt.Sprintf("%v %v\n", h.cards, h.bid)
}

func parseInput() (hands []Hand) {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := string(input)

	for _, line := range strings.Split(inputStr, "\n") {
		if line == "" {
			continue
		}
		hand := Hand{}
		split := strings.Split(line, " ")
		cards := split[0]
		bid := split[1]

		hand.bid, _ = strconv.Atoi(bid)

		for _, card := range cards {
			hand.cards = append(hand.cards, string(card))
		}

		hands = append(hands, hand)
	}

	return hands
}

func getHandType(hand Hand) int {
	groupingMap := map[string]int{}

	for _, card := range hand.cards {
		if _, ok := groupingMap[card]; ok {
			groupingMap[card]++
		} else {
			groupingMap[card] = 1
		}
	}

	if jokerEnabled {
		if _, ok := groupingMap["J"]; ok {

			if groupingMap["J"] == 5 {
				// Five of a kind
				return 6
			}

			mostFreq := ""
			for card, count := range groupingMap {
				if card == "J" {
					continue
				}
				if mostFreq == "" || count > groupingMap[mostFreq] {
					mostFreq = card
				}
			}
			groupingMap[mostFreq] += groupingMap["J"]
			delete(groupingMap, "J")
		}
	}

	if len(groupingMap) == 1 {
		// Five of a kind
		return 6
	} else if len(groupingMap) == 2 {
		for _, count := range groupingMap {
			if count == 3 {
				// Full house
				return 4
			}
		}
		// Four of a kind
		return 5
	} else if len(groupingMap) == 3 {
		for _, count := range groupingMap {
			if count == 3 {
				// Three of a kind
				return 3
			}
		}
		// Two pair
		return 2
	} else if len(groupingMap) == 4 {
		// One pair
		return 1
	}
	// High card
	return 0
}

func getCardValue(card string) int {
	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		if jokerEnabled {
			return 1
		}
		return 11
	case "T":
		return 10
	default:
		cardVal, _ := strconv.Atoi(card)
		return cardVal
	}
}

func cmpHands(hand1, hand2 Hand) int {
	hand1Type := getHandType(hand1)
	hand2Type := getHandType(hand2)
	if hand1Type != hand2Type {
		return hand1Type - hand2Type
	}
	for i := 0; i < len(hand1.cards); i++ {
		hand1CardVal := getCardValue(hand1.cards[i])
		hand2CardVal := getCardValue(hand2.cards[i])
		if hand1CardVal != hand2CardVal {
			return hand1CardVal - hand2CardVal
		}
	}
	return 0
}

func calculateWinnings(hands []Hand) int {
	totalWinnings := 0
	slices.SortFunc(hands, cmpHands)
	for i := 0; i < len(hands); i++ {
		totalWinnings += hands[i].bid * (i + 1)
	}
	return totalWinnings
}

func main() {
	hands := parseInput()
	fmt.Println("Part 1:", calculateWinnings(hands))
	jokerEnabled = true
	fmt.Println("Part 2:", calculateWinnings(hands))
}
