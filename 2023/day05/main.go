package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Mapping struct {
	source int
	dest   int
	length int
}

type Almanac struct {
	seeds []int

	seedToSoilMappings            []Mapping
	soilToFertilizerMappings      []Mapping
	fertilizerToWaterMappings     []Mapping
	waterToLightMappings          []Mapping
	lightToTemperatureMappings    []Mapping
	temperatureToHumidityMappings []Mapping
	humidityToLocationMappings    []Mapping
}

func parseMap(input string, mapToParse string) []Mapping {
	regex := regexp.MustCompile(fmt.Sprintf(`%s map:\n((?:\d+\s*)+)`, mapToParse))
	mapStr := regex.FindStringSubmatch(input)

	mapStr[1] = strings.TrimSpace(mapStr[1])
	rows := strings.Split(mapStr[1], "\n")

	parsedMap := []Mapping{}
	for _, row := range rows {
		nums := strings.Split(row, " ")
		dest, _ := strconv.Atoi(nums[0])
		source, _ := strconv.Atoi(nums[1])
		length, _ := strconv.Atoi(nums[2])

		parsedMap = append(parsedMap, Mapping{source, dest, length})
	}

	return parsedMap
}

func parseInput() Almanac {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := string(input)

	almanac := Almanac{}

	// parse seeds
	seedsRegex := regexp.MustCompile(`^seeds: (.*)\n`)
	seeds := seedsRegex.FindStringSubmatch(inputStr)[1]
	for _, seed := range strings.Split(seeds, " ") {
		seed, _ := strconv.Atoi(seed)
		almanac.seeds = append(almanac.seeds, seed)
	}

	almanac.seedToSoilMappings = parseMap(inputStr, "seed-to-soil")
	almanac.soilToFertilizerMappings = parseMap(inputStr, "soil-to-fertilizer")
	almanac.fertilizerToWaterMappings = parseMap(inputStr, "fertilizer-to-water")
	almanac.waterToLightMappings = parseMap(inputStr, "water-to-light")
	almanac.lightToTemperatureMappings = parseMap(inputStr, "light-to-temperature")
	almanac.temperatureToHumidityMappings = parseMap(inputStr, "temperature-to-humidity")
	almanac.humidityToLocationMappings = parseMap(inputStr, "humidity-to-location")

	return almanac
}

func getMapValue(m []Mapping, key int) int {
	for _, mapping := range m {
		if key >= mapping.source && key < mapping.source+mapping.length {
			return mapping.dest + (key - mapping.source)
		}
	}

	return key
}

func (almanac Almanac) getLocation(seed int) int {
	soil := getMapValue(almanac.seedToSoilMappings, seed)
	fertilizer := getMapValue(almanac.soilToFertilizerMappings, soil)
	water := getMapValue(almanac.fertilizerToWaterMappings, fertilizer)
	light := getMapValue(almanac.waterToLightMappings, water)
	temperature := getMapValue(almanac.lightToTemperatureMappings, light)
	humidity := getMapValue(almanac.temperatureToHumidityMappings, temperature)
	location := getMapValue(almanac.humidityToLocationMappings, humidity)

	return location
}

func part1(almanac Almanac) int {
	smallestLocation := almanac.getLocation(almanac.seeds[0])
	for _, seed := range almanac.seeds[1:] {
		if location := almanac.getLocation(seed); location < smallestLocation {
			smallestLocation = location
		}
	}

	return smallestLocation
}

func part2(almanac Almanac) int {
	smallestLocation := almanac.getLocation(almanac.seeds[0])
	for i := 0; i < len(almanac.seeds); i = i + 2 {
		seed := almanac.seeds[i]
		seedRange := almanac.seeds[i+1]
		for j := seed; j < seed+seedRange; j++ {
			if location := almanac.getLocation(j); location < smallestLocation {
				smallestLocation = location
			}
		}
	}

	return smallestLocation
}

func main() {
	almanac := parseInput()
	fmt.Println("Part 1:", part1(almanac))
	fmt.Println("Part 2:", part2(almanac))
}
