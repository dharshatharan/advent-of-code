package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Set struct {
	red   int
	blue  int
	green int
}

type Game struct {
	id    int
	sets  []*Set
	valid bool
}

func (g *Game) isValid() bool {
	if g.valid {
		return true
	}

	for _, set := range g.sets {
		if set.red > 12 || set.blue > 14 || set.green > 13 {
			return false
		}
	}
	g.valid = true
	return true
}

func (g *Game) getPowerOfFewestNumberOfDie() int {
	fewestRed := 0
	fewestBlue := 0
	fewestGreen := 0

	for _, set := range g.sets {
		if set.red > fewestRed {
			fewestRed = set.red
		}
		if set.blue > fewestBlue {
			fewestBlue = set.blue
		}
		if set.green > fewestGreen {
			fewestGreen = set.green
		}
	}

	return fewestRed * fewestBlue * fewestGreen
}

func (g *Game) String() string {
	return fmt.Sprintf("Game %d:\n%v", g.id, g.sets)
}

func (s *Set) String() string {
	return fmt.Sprintf("\tSet: red=%d, blue=%d, green=%d\n", s.red, s.blue, s.green)
}

func stringToInt(str string, def int) int {
	if valInt, err := strconv.Atoi(str); err == nil {
		return valInt
	}
	return def
}

func stringSubmatchToInt(matches []string, index int, def int) int {
	if len(matches) > index {
		return stringToInt(matches[index], def)
	} else {
		return def
	}
}

func parseGames(input string) []*Game {
	gameIdRegex := regexp.MustCompile(`Game (\d+):(.+)`)

	redRegex := regexp.MustCompile(`(\d+) red`)
	blueRegex := regexp.MustCompile(`(\d+) blue`)
	greenRegex := regexp.MustCompile(`(\d+) green`)

	var games []*Game

	for _, line := range strings.Split(input, "\n") {
		var game Game

		gameMatch := gameIdRegex.FindStringSubmatch(line)
		if len(gameMatch) == 0 {
			continue
		}

		game.id = stringToInt(gameMatch[1], 0)

		sets := strings.Split(gameMatch[2], ";")
		for _, setStr := range sets {
			var set Set
			set.red = stringSubmatchToInt(redRegex.FindStringSubmatch(setStr), 1, 0)
			set.blue = stringSubmatchToInt(blueRegex.FindStringSubmatch(setStr), 1, 0)
			set.green = stringSubmatchToInt(greenRegex.FindStringSubmatch(setStr), 1, 0)
			game.sets = append(game.sets, &set)
		}

		games = append(games, &game)
	}
	return games
}

func readInput() string {
	// read input file to string
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return string(input)
}

func part1(games []*Game) {
	idCount := 0
	for _, game := range games {
		if game.isValid() {
			idCount += game.id
		}
	}
	fmt.Printf("Part 1: %d\n", idCount)
}

func part2(games []*Game) {
	totalPower := 0
	for _, game := range games {
		totalPower += game.getPowerOfFewestNumberOfDie()
	}
	fmt.Printf("Part 2: %d\n", totalPower)
}

func main() {
	games := parseGames(readInput())
	part1(games)
	part2(games)
}
