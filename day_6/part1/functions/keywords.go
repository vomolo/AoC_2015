package functions

import (
	"fmt"
	"regexp"
	"strconv"
)

// KeyWords processes the input string for keywords and coordinates.
func KeyWords(str string, off, on [][]int) [][]int {
	var coordinates [][]int
	keywords := []string{"toggle", "off", "on"}
	tasks := make(map[string]func())

	// Define tasks for each keyword
	tasks["toggle"] = func() {
		fmt.Println("Toggling the switch")
	}

	tasks["off"] = func() {
		// Create a map for fast lookup of elements in coordinates
		smallMap := make(map[string]struct{})
		for _, elem := range coordinates {
			key := fmt.Sprintf("%d,%d", elem[0], elem[1])
			smallMap[key] = struct{}{}
		}

		// Create a new slice to hold the result
		result := [][]int{}

		// Iterate through the 'on' slice and append elements that are not in 'coordinates'
		for _, elem := range on {
			key := fmt.Sprintf("%d,%d", elem[0], elem[1])
			if _, found := smallMap[key]; !found {
				result = append(result, elem)
			}
		}

		// Update the 'off' slice with elements from 'coordinates' that are not already in 'off'
		for _, itemA := range coordinates {
			if !contains(off, itemA) {
				off = append(off, itemA)
			}
		}

		// Print the updated off
		fmt.Println("Updated off:", off)
	}

	tasks["on"] = func() {
		fmt.Println("Turning the device on")
	}

	// Range through the str string for keywords
	for i := 0; i < len(str); i++ {
		for _, keyword := range keywords {
			if i+len(keyword) <= len(str) && str[i:i+len(keyword)] == keyword {
				tasks[keyword]()
				i += len(keyword) - 1
				break
			}
		}
	}

	// Identify the range "0,0 through 999,999"
	re := regexp.MustCompile(`(\d+),(\d+) through (\d+),(\d+)`)
	matches := re.FindStringSubmatch(str)

	if len(matches) == 5 {
		x1, _ := strconv.Atoi(matches[1])
		y1, _ := strconv.Atoi(matches[2])
		x2, _ := strconv.Atoi(matches[3])
		y2, _ := strconv.Atoi(matches[4])

		// Create a slice of slices to hold all 2D coordinates in the range
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				coordinates = append(coordinates, []int{x, y})
			}
		}
	}

	return coordinates
}

// contains checks if a slice contains a specific item.
func contains(slice [][]int, item []int) bool {
	for _, s := range slice {
		if len(s) == len(item) {
			matched := true
			for i := range s {
				if s[i] != item[i] {
					matched = false
					break
				}
			}
			if matched {
				return true
			}
		}
	}
	return false
}
