package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func findLowestNumber(secretKey string) int {
	number := 1

	for {
		text := secretKey + strconv.Itoa(number)
		hash := md5.Sum([]byte(text))
		hashStr := hex.EncodeToString(hash[:])

		if hashStr[:6] == "000000" {
			return number
		}

		number++
	}
}

func main() {
	secretKey := "ckczppom"
	result := findLowestNumber(secretKey)
	fmt.Printf("The lowest number resulting in a hash starting with five zeroes is: %d\n", result)
}
