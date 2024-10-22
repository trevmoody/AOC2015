package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1(*GetFileAsLines("input"))
	part2(*GetFileAsLines("input"))
}

func part1(lines []string) {
	floor := processLine1(lines[0])
	fmt.Printf("Part1: floor is  %d\n", floor)
}

func part2(lines []string) {
	position, floor := processLine2(lines[0])
	fmt.Printf("Part2: position: %d, floor is %d", position, floor)
}

func processLine2(line string) (int, int) {
	floor := 0
	position := 0
	instructions := []rune(line)

	for i, char := range instructions {
		switch char {
		case '(':
			floor = floor + 1

		case ')':
			floor = floor - 1
		default:
		}

		if floor == -1 {
			position = i + 1
			break
		}
	}

	return floor, position

}
func processLine1(line string) int {
	floor := 0
	instructions := []rune(line)

	for _, char := range instructions {
		switch char {
		case '(':
			floor = floor + 1

		case ')':
			floor = floor - 1
		default:
		}

	}

	return floor

}

func GetFileAsLines(fileName string) *[]string {

	currentDir, _ := os.Getwd()
	fmt.Printf("Current DIR: %s\n", currentDir)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic("Error opening file")
		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return &lines
}
