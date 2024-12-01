package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput() [][]rune {
	data, _ := os.ReadFile("input")
	splitData := strings.Split(string(data), "\n")
	res := make([][]rune, len(splitData))
	for i, v := range splitData {
		for _, c := range v {
			res[i] = append(res[i], c)
		}

	}
	return res
}

type Coordinate struct {
	x int
	y int
}

// Returns coordinates for start and a possible next tile
func findStart(area [][]rune) (Coordinate, Coordinate) {
	var start Coordinate
	for y, line := range area {
		for x, point := range line {
			if point == 'S' {
				fmt.Println("Start: x ", x, ", y ", y)
				start = Coordinate{x: x, y: y}
			}
		}
	}

	// Find possible next tile
	var next Coordinate
	if start.y != 0 && (area[start.y-1][start.x] == '|' || area[start.y-1][start.x] == '7' || area[start.y-1][start.x] == 'F') {
		next = Coordinate{start.x, start.y - 1}
	} else if start.y != len(area[0])-1 && (area[start.y+1][start.x] == '|' || area[start.y+1][start.x] == 'L' || area[start.y+1][start.x] == 'J') {
		next = Coordinate{start.x, start.y + 1}
	} else if start.x != 0 && (area[start.y][start.x-1] == '-' || area[start.y][start.x-1] == 'L' || area[start.y][start.x-1] == 'F') {
		next = Coordinate{start.x - 1, start.y}
	} else {
		next = Coordinate{start.x + 1, start.y}
	}

	return start, next

}

func findLength(area [][]rune, start Coordinate, previous Coordinate, current Coordinate, currentLength int) int {

	currentTile := area[current.y][current.x]

	// Check which direction is possible to visit next
	var checkNorth, checkSouth, checkEast, checkWest bool
	switch currentTile {
	case '|':
		checkNorth, checkSouth, checkEast, checkWest = true, true, false, false
	case '-':
		checkNorth, checkSouth, checkEast, checkWest = false, false, true, true
	case 'L':
		checkNorth, checkSouth, checkEast, checkWest = true, false, true, false
	case 'J':
		checkNorth, checkSouth, checkEast, checkWest = true, false, false, true
	case '7':
		checkNorth, checkSouth, checkEast, checkWest = false, true, false, true
	case 'F':
		checkNorth, checkSouth, checkEast, checkWest = false, true, true, false
	}

	// Set previous tile false
	switch {
	case previous.x == current.x-1:
		checkWest = false
		fmt.Println("west false")
	case previous.x == current.x+1:
		checkEast = false
		fmt.Println("east false")
	case previous.y == current.y-1:
		checkNorth = false
		fmt.Println("north false")
	case previous.y == current.y+1:
		checkSouth = false
		fmt.Println("south false")
	}

	// Now only one direction should be true, set that coordinate to be next
	var next Coordinate
	if checkNorth {
		next = Coordinate{current.x, current.y - 1}
	} else if checkSouth {
		next = Coordinate{current.x, current.y + 1}
	} else if checkWest {
		next = Coordinate{current.x - 1, current.y}
	} else if checkEast {
		next = Coordinate{current.x + 1, current.y}
	}

	if next == start {
		//If next == start, we now have the full length of the loop and can return
		return currentLength
	} else {
		// Otherwise we call findLength for the next
		if current.x == 0 && current.y == 0 {
			return 0
		}
		return findLength(area, start, current, next, currentLength+1)
	}
}

func main() {
	area := parseInput()
	// Get the start tile and a possible next tile
	start, next := findStart(area)
	// Call findLength with an initial length of 2, because we start at the second tile
	length := findLength(area, start, start, next, 2)
	fmt.Println("Length: ", length)
	fmt.Println("Half: ", length/2)
}
