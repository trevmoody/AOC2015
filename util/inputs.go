package util

import (
	"bufio"
	"fmt"
	"os"
)

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

	//for i, line := range lines {
	//	fmt.Printf("line %d: %s\n", i, line)
	//}
	//
	//fmt.Printf("line count: %d\n", len(lines))

	return &lines
}

func SplitSliceOddEvenIndex[T any](slice []T) ([]T, []T) {
	var evenIndices []T
	var oddIndices []T

	for i, value := range slice {
		if i%2 == 0 {
			evenIndices = append(evenIndices, value)
		} else {
			oddIndices = append(oddIndices, value)
		}
	}

	return evenIndices, oddIndices
}

//
//func StringsToInts(line string) []int {
//	fields := strings.Fields(line)
//	var retList []int
//	for _, f := range fields {
//		i, _ := strconv.Atoi(strings.TrimSpace(f))
//		retList = append(retList, i)
//	}
//	return retList
//}
//
//func ReverseInts(s []int) []int {
//	slices.Reverse(s)
//	return s
//
//}
