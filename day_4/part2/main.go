package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func findLowestNumber(secretKey string) int {
	number := 1

	for {

		input := fmt.Sprintf("%s%d", secretKey, number)

		hash := md5.Sum([]byte(input))
		hashString := hex.EncodeToString(hash[:])

		if hashString[:6] == "000000" {
			return number
		}
		number++
	}
}

func main() {
	secretKey := "bgvyzdsv"
	lowestNumber := findLowestNumber(secretKey)
	fmt.Println(lowestNumber)
}
