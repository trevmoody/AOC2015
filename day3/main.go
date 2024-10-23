package main

import (
	"AOC2015/util"
	"fmt"
)

func main() {
	part1((*util.GetFileAsLines("./day3/input"))[0])
	part2((*util.GetFileAsLines("./day3/input"))[0])
}

func part1(line string) int {

	housesVisited := make(map[House]bool)
	currentHouse := House{0, 0}
	housesVisited[currentHouse] = true

	instructions := []rune(line)

	housesVisited = getHousesVisited(instructions, housesVisited)
	fmt.Printf("Part1: housesVisited is  %d\n", len(housesVisited))

	return len(housesVisited)
}

func part2(line string) int {
	housesVisited := make(map[House]bool)
	currentHouse := House{0, 0}
	housesVisited[currentHouse] = true

	instructions := []rune(line)

	santaInstructions, robotInstructions := util.SplitSliceOddEvenIndex(instructions)

	housesVisited = getHousesVisited(santaInstructions, housesVisited)
	housesVisited = getHousesVisited(robotInstructions, housesVisited)

	fmt.Printf("Part2: housesVisited is  %d\n", len(housesVisited))

	return len(housesVisited)

}

func getHousesVisited(instructions []rune, visited map[House]bool) map[House]bool {

	currentHouse := House{0, 0}

	for _, char := range instructions {
		var nextHouse House
		switch char {
		case '>':
			nextHouse = currentHouse.MoveRight()
		case '<':
			nextHouse = currentHouse.MoveLeft()
		case '^':
			nextHouse = currentHouse.MoveUp()
		case 'v':
			nextHouse = currentHouse.MoveDown()
		}

		if _, ok := visited[nextHouse]; !ok {
			visited[nextHouse] = true
		}

		currentHouse = nextHouse
	}

	return visited

}
