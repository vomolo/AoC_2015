package main

import (
	"bufio"
	"fmt"
	"os"
	"part1/functions"
	"strings"
)

func main() {
	useMes := "Expecting: go run your-program.go data.txt"
	// Length of arguments check
	if len(os.Args) != 2 {
		fmt.Println(useMes)
		fmt.Println("Check length of arguments")
		return
	}
	filePath := "../" + os.Args[1]
	// Check .txt file extension
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

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file content line by line
	scanner := bufio.NewScanner(file)
	off := [][]int
	on := [][]int
	for scanner.Scan() {
		line := scanner.Text() // Read the current line

		sli := functions.KeyWords(line, off, on) // Process the line with the Slice function
		fmt.Println(sli)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}
