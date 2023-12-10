package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	north = iota
	south = iota
	east  = iota
	west  = iota
)

type Pipe struct {
	dirs [2]int
}

func (p Pipe) contains(dir int) bool {
	return p.dirs[0] == dir || p.dirs[1] == dir
}

func getOpposite(dir int) int {
	oppMap := map[int]int{
		north: south,
		south: north,
		east:  west,
		west:  east,
	}
	return oppMap[dir]
}

func identifyStartPipe(nodeMap map[[2]int]Pipe, start [2]int) (pipe Pipe) {
	x, y := start[0], start[1]
	dirCount := 0
	pipe = Pipe{}

	checkMap := map[[2]int]int{
		{x, y - 1}: east,
		{x, y + 1}: west,
		{x - 1, y}: south,
		{x + 1, y}: north,
	}

	for coord, dir := range checkMap {
		if dirCount == 2 {
			break
		}
		if _, ok := nodeMap[coord]; ok {
			if nodeMap[coord].contains(dir) {
				pipe.dirs[dirCount] = getOpposite(dir)
				dirCount++
			}
		}
	}
	return pipe
}

func parseInput() (nodeMap map[[2]int]Pipe, start [2]int, gridSize [2]int) {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := string(input)

	nodeMap = map[[2]int]Pipe{}
	split := strings.Split(inputStr, "\n")
	gridSize = [2]int{len(split) - 1, len(split[0])}
	for x, line := range split {
		for y, char := range line {
			if char == '.' {
				continue
			}
			pipe := Pipe{[2]int{}}
			if char == '|' {
				pipe.dirs = [2]int{north, south}
			} else if char == '-' {
				pipe.dirs = [2]int{east, west}
			} else if char == 'L' {
				pipe.dirs = [2]int{north, east}
			} else if char == 'J' {
				pipe.dirs = [2]int{north, west}
			} else if char == 'F' {
				pipe.dirs = [2]int{south, east}
			} else if char == '7' {
				pipe.dirs = [2]int{south, west}
			} else if char == 'S' {
				start = [2]int{x, y}
			} else {
				panic("Invalid char")
			}
			nodeMap[[2]int{x, y}] = pipe
		}
	}

	nodeMap[start] = identifyStartPipe(nodeMap, start)

	return nodeMap, start, gridSize
}

func getNextNode(nodeMap map[[2]int]Pipe, current [2]int, fromDir int) (next [2]int, toDir int) {
	x, y := current[0], current[1]
	toDir = -1
	next = [2]int{}

	for _, dir := range nodeMap[current].dirs {
		if dir != fromDir {
			toDir = dir
			break
		}
	}

	if toDir == north {
		next = [2]int{x - 1, y}
	} else if toDir == south {
		next = [2]int{x + 1, y}
	} else if toDir == east {
		next = [2]int{x, y + 1}
	} else if toDir == west {
		next = [2]int{x, y - 1}
	} else {
		panic("Invalid direction")
	}
	return next, toDir
}

func calculateFurtherestNode(nodeMap map[[2]int]Pipe, start [2]int) int {
	steps := 1
	aNode, aDir := getNextNode(nodeMap, start, nodeMap[start].dirs[0])
	bNode, bDir := getNextNode(nodeMap, start, nodeMap[start].dirs[1])

	for {
		if aNode == bNode {
			break
		}
		if _, ok := nodeMap[aNode]; ok {
			aNode, aDir = getNextNode(nodeMap, aNode, getOpposite(aDir))
		} else {
			panic("Invalid node")
		}
		if _, ok := nodeMap[bNode]; ok {
			bNode, bDir = getNextNode(nodeMap, bNode, getOpposite(bDir))
		} else {
			panic("Invalid node")
		}
		steps++
	}

	return steps
}

func getNumberOfNodesInsideLoop(nodeMap map[[2]int]Pipe, start [2]int, gridSize [2]int) int {

	loopMap := map[[2]int]Pipe{}
	node, dir := getNextNode(nodeMap, start, nodeMap[start].dirs[0])
	loopMap[node] = nodeMap[node]
	for {
		if node == start {
			break
		}
		if _, ok := nodeMap[node]; ok {
			node, dir = getNextNode(nodeMap, node, getOpposite(dir))
			loopMap[node] = nodeMap[node]
		} else {
			panic("Invalid node")
		}
	}

	count := 0

	for x := 0; x < gridSize[0]; x++ {
		intersections := 0
		for y := 0; y < gridSize[1]; y++ {
			if _, ok := loopMap[[2]int{x, y}]; ok {
				n := loopMap[[2]int{x, y}]
				if n.contains(south) {
					intersections++
				}
			} else {
				if intersections%2 == 1 {
					count++
				}
			}
		}
	}

	return count
}

func printNodeMap(nodeMap map[[2]int]Pipe, gridSize [2]int) {
	mapChars := map[[2]int]string{
		{north, south}: "|",
		{east, west}:   "-",
		{north, east}:  "L",
		{north, west}:  "J",
		{south, east}:  "F",
		{south, west}:  "7",
	}

	for x := 0; x < gridSize[0]; x++ {
		for y := 0; y < gridSize[1]; y++ {
			if _, ok := nodeMap[[2]int{x, y}]; ok {
				fmt.Print(mapChars[nodeMap[[2]int{x, y}].dirs])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	nodeMap, start, gridSize := parseInput()
	fmt.Println("Part 1:", calculateFurtherestNode(nodeMap, start))
	fmt.Println("Part 2:", getNumberOfNodesInsideLoop(nodeMap, start, gridSize))
}
