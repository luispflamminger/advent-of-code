package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

/*
		Pseudocode:
	 1. Store input as slice of slices
	 2. Traverse from left to right and top to bottom
	 3. If we encounter a digit, store digit and starting indices, then continue traversing and storing until no longer a digit
	 4. Store number of digits and combined int.
	 5. Scan [y-1][x-1] to [y-1][x+len(number)], [y][x-1], [y][x+len(number)], [y+1][x-1] to [y+1][x+len(number)]
	 6. If any of the scanned fields contains a *, we add it to a map of the following structure:
        {"xOfStar,yOfStar": [number1, number2], ...}
*/

func isDigit(val byte) bool {
	return val >= '0' && val <= '9'
}

func constructNumber(line []byte, startingIndex int, startingValue int) int {
	digitAtIndex, _ := strconv.Atoi(string(line[startingIndex]))
	newNumber := 10*startingValue + digitAtIndex
	if startingIndex+1 > len(line)-1 {
		// Base Case
		return newNumber
	} else if !isDigit(line[startingIndex+1]) {
		// Base Case
		return newNumber
	} else {
		// Recursive Call
		return constructNumber(line, startingIndex+1, newNumber)
	}
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// Returns all coordinates of adjacent gears as []string
func getGearCoordinates(data [][]byte, initialX int, initialY int, number int) []string {
	numberOfDigits := len(strconv.Itoa(number))
	gears := make([]string, 0)

	var startY, stopY int
	if initialY == 0 {
		startY = 0
		stopY = 1
	} else if initialY == len(data)-1 {
		startY = len(data) - 2
		stopY = len(data) - 1
	} else {
		startY = initialY - 1
		stopY = initialY + 1
	}

	var startX int
	if initialX == 0 {
		startX = initialX
	} else {
		startX = initialX - 1
	}
	stopX := Min(len(data[initialY])-1, initialX+numberOfDigits)

	for y := startY; y <= stopY; y++ {
		for x := startX; x <= stopX; x++ {
			if data[y][x] == '*' {
				gears = append(gears, fmt.Sprintf("%d,%d", x, y))
			}
		}
	}
	return gears
}

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error while reading file: ", err)
	}
	data := bytes.Split(input, []byte("\n")) // Data representation is a [][]byte

	gearMap := make(map[string][]int, 0)

	for y := 0; y < len(data); y++ {
		line := data[y]
		for x := 0; x < len(data[y]); x++ {
			currentChar := data[y][x]
			if isDigit(currentChar) {
				currentNumber := constructNumber(line, x, 0)
				gears := getGearCoordinates(data, x, y, currentNumber)
				for _, gear := range gears {
					if gearMap[gear] == nil {
						gearMap[gear] = []int{currentNumber}
					} else {
						gearMap[gear] = append(gearMap[gear], currentNumber)
					}

				}
				// Skip forward to next character outside of number
				x += len(strconv.Itoa(currentNumber))
			}
		}
	}

	sum := 0
	for _, gear := range gearMap {
		if len(gear) == 2 {
			sum += gear[0] * gear[1]
		}
	}

	fmt.Println("Sum of gear ratios: ", sum)

}
