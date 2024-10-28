package main

import (
	"AOC2015/util"
	"fmt"
	"strings"
)

func main() {

	//	part1((*util.GetFileAsLines("./day5/input")))
	part2((*util.GetFileAsLines("./day5/input")))

}

func part2(list []string) {
	count := 0
	for _, s := range list {
		if checkString2(s) {
			count = count + 1
		}
	}
	fmt.Printf("part2 count = %d\n", count)
}

func checkString2(stringToCheck string) bool {

	if twoLettersAppearTwice(stringToCheck) && repeatWithOneLetterBetween(stringToCheck) {
		return true
	}
	return false
}

func repeatWithOneLetterBetween(input string) bool {
	for i := 0; i < len(input)-2; i++ {
		if input[i] == input[i+2] {
			fmt.Printf("found 2 chars with one letter between %c %c %c\n", input[i], input[i+1], input[i+2])
			return true
		}
	}
	return false
}

func twoLettersAppearTwice(input string) bool {
	inputLength := len(input) - 2
	fmt.Printf("input length = %d\n", inputLength)
	for i := 0; i < inputLength; i++ {
		for j := i + 2; j < len(input)-1; j++ {
			fmt.Printf("i = %d j = %d\n", i, j)
			if input[i] == input[j] && input[i+1] == input[j+1] {
				fmt.Printf("found 2 chars twice in a row %c%c %c%c\n", input[i], input[i+1], input[j], input[j+1])
				return true
			}
		}

	}
	return false
}

func part1(list []string) {
	count := 0
	for _, s := range list {
		if checkString1(s) {
			count = count + 1
		}
	}
	fmt.Printf("part1 count = %d\n", count)

}

func checkString1(stringToCheck string) bool {

	// Initialize an empty set
	doNotContain := make(map[string]struct{})

	// Add elements to the set
	doNotContain["ab"] = struct{}{}
	doNotContain["cd"] = struct{}{}
	doNotContain["pq"] = struct{}{}
	doNotContain["xy"] = struct{}{}

	if countVowels(stringToCheck) > 2 && checkIf2CharsInARow(stringToCheck) && doesNotContain(doNotContain, stringToCheck) {
		return true
	}

	return false

}

func doesNotContain(doNotContain map[string]struct{}, stringtocheck string) bool {
	for s, _ := range doNotContain {
		if strings.Contains(stringtocheck, s) {
			fmt.Printf("found %s in %s\n", s, stringtocheck)
			return false
		}
	}
	fmt.Printf("did not find any of the strings in %s\n", stringtocheck)
	return true
}

func countVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count = count + 1
		}
	}
	fmt.Printf("count of vowels in %s = %d\n", s, count)
	return count
}

func checkIf2CharsInARow(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			fmt.Printf("found 2 chars in a row %c %c\n", input[i], input[i+1])
			return true
		}
	}
	return false
}
