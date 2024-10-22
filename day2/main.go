package main

import (
	"AOC2015/util"
	"fmt"
	"sort"
)

func main() {
	part1(*util.GetFileAsLines("./day2/input"))
	part2(*util.GetFileAsLines("./day2/input"))
}

func part1(lines []string) {

	sum := 0
	for _, dimension := range lines {
		area := getAreaFromDimension(dimension)
		sum += area
	}

	fmt.Printf("Part1: area is  %d\n", sum)
}

func getAreaFromDimension(dimension string) int {
	numbers := make([]int, 3)
	_, err := fmt.Sscanf(dimension, "%dx%dx%d", &numbers[0], &numbers[1], &numbers[2])

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		sort.Ints(numbers)
		fmt.Println(numbers)
	}

	area := 3*(numbers[0]*numbers[1]) + 2*((numbers[0]*numbers[2])+(numbers[1]*numbers[2]))

	fmt.Printf("dimension: %s,  area is  %d\n", dimension, area)

	return area

}

func part2(lines []string) {
	sum := 0
	for _, dimension := range lines {
		area := getRibbonLengthFromDimension(dimension)
		sum += area
	}

	fmt.Printf("Part2: length is  %d\n", sum)
}

func getRibbonLengthFromDimension(dimension string) int {
	numbers := make([]int, 3)
	_, err := fmt.Sscanf(dimension, "%dx%dx%d", &numbers[0], &numbers[1], &numbers[2])

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		sort.Ints(numbers)
		fmt.Println(numbers)
	}

	length := 2*(numbers[0]+numbers[1]) + (numbers[0] * numbers[1] * numbers[2])

	fmt.Printf("dimension: %s,  length is  %d\n", dimension, length)

	return length
}
