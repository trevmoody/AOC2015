package main

import (
	"AOC2015/util"
	"fmt"
	"regexp"
	"strconv"
)

var pattern = regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)

func main() {
	part1(*util.GetFileAsLines("./day6/input"))
	part2(*util.GetFileAsLines("./day6/input"))
}

func part1(input []string) {

	var grid [1000][1000]bool
	var count = setAndCountLitLights(grid, input)
	fmt.Printf("Part1: Number of lit lights: %d\n", count)
}

func part2(input []string) {

	var grid [1000][1000]int
	var count = setAndCountBrightness(grid, input)
	fmt.Printf("Part2: Total Brightness of lit lights: %d\n", count)
}

func setAndCountBrightness(grid [1000][1000]int, instructions []string) int {
	for _, line := range instructions {
		matches := pattern.FindStringSubmatch(line)
		if matches == nil {
			fmt.Println("Invalid line")
			continue
		}

		action := matches[1]
		startX, _ := strconv.Atoi(matches[2])
		startY, _ := strconv.Atoi(matches[3])
		endX, _ := strconv.Atoi(matches[4])
		endY, _ := strconv.Atoi(matches[5])

		for x := startX; x <= endX; x++ {
			for y := startY; y <= endY; y++ {
				switch action {
				case "turn on":
					grid[x][y] = grid[x][y] + 1
				case "turn off":
					if grid[x][y] != 0 {
						grid[x][y] = grid[x][y] - 1
					}
				case "toggle":
					grid[x][y] = grid[x][y] + 2
				}
			}
		}
	}

	total := 0
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			total = total + grid[x][y]
		}
	}
	return total
}

func setAndCountLitLightsNew(grid [][]bool, instructions []string) [][]bool {
	return grid
}

func setAndCountLitLights(grid [1000][1000]bool, instructions []string) int {

	for _, line := range instructions {
		matches := pattern.FindStringSubmatch(line)
		if matches == nil {
			fmt.Println("Invalid line")
			continue
		}

		action := matches[1]
		startX, _ := strconv.Atoi(matches[2])
		startY, _ := strconv.Atoi(matches[3])
		endX, _ := strconv.Atoi(matches[4])
		endY, _ := strconv.Atoi(matches[5])

		// Apply the instruction to the grid
		for x := startX; x <= endX; x++ {
			for y := startY; y <= endY; y++ {
				switch action {
				case "turn on":
					grid[x][y] = true
				case "turn off":
					grid[x][y] = false
				case "toggle":
					grid[x][y] = !grid[x][y]
				}
			}
		}
	}

	count := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] {
				count++
			}
		}
	}
	return count
}
