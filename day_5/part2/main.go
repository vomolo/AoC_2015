package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"part2/functions"
)

func main() {
	useMes := "Expecting: go run your-program.go data.txt"

	if len(os.Args) != 2 {
		fmt.Println(useMes)
		fmt.Println("Check length of arguments")
		return
	}
	filePath := os.Args[1]

	if !strings.HasSuffix(filePath, ".txt") {
		fmt.Println(useMes)
		fmt.Println("Make sure it's a .txt file extension")
		os.Exit(1)
	}
	fileCheck, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			return
		}
		fmt.Println("Error checking file:", err)
		return
	}
	if fileCheck.Size() == 0 {
		fmt.Println("File is empty")
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arr []string
	for scanner.Scan() {
		line := scanner.Text()
		pai := functions.NonOverlappingPair(line)
		pair := functions.RepeatedLetterWithOneInBetween(pai)
		fmt.Println(pair)

		if pair != "" {
			arr = append(arr, pair)
		}

	}

	fmt.Println(len(arr))
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}
