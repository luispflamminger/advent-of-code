package main

import (
	"fmt"
	"os"
	"strings"
)

var PART_TWO = true

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

func contains(s []rune, e rune) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func expandColumns(u [][]rune, factor int) [][]rune {
	expansion := make([]rune, factor)
	for i := range expansion {
		expansion[i] = '.'
	}

	for col := 0; col < len(u[0]); col++ {
		column := make([]rune, len(u))
		for i, v := range u {
			column[i] = v[col]
		}

		if !contains(column, '#') {
			for line := range u {
				u[line] = append(u[line][:col], append(expansion, u[line][col:]...)...)
			}
			col += factor
		}
	}

	return u
}

/*
As an optimization, this function also returns a set of integers that represent y positions
marking the beginning of an expanded section.
When searching for galaxies later, we can skip these portions entirely as walking
an empty section for part 2 will require (10^6)^2 array accesses.
We could even further optimize by doing the same thing in the expandColumns function, but this
will result in a far lower performance gain.
*/
func expandLines(u [][]rune, factor int) ([][]rune, map[int]bool) {
	expandedSections := make(map[int]bool, 0)
	emptyLines := make([][]rune, factor)
	emptyLine := make([]rune, len(u[0]))
	for i := range emptyLine {
		emptyLine[i] = '.'
	}

	for i := range emptyLines {
		emptyLines[i] = emptyLine
	}

	for i := 0; i < len(u); i++ {
		if !contains(u[i], '#') {
			u = append(u[:i], append(emptyLines, u[i:]...)...)
			expandedSections[i] = true
			i += factor
		}
	}

	return u, expandedSections
}

func expand(u [][]rune, factor int) ([][]rune, map[int]bool) {

	return expandLines(expandColumns(u, factor), factor)
}

type Coordinate struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func calculateCombinedXYOffset(a Coordinate, b Coordinate) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func main() {
	var expansionFactor int
	if PART_TWO {
		expansionFactor = 999999
	} else {
		expansionFactor = 1
	}

	expanded, expandedSections := expand(parseInput(), expansionFactor)

	// We use this map as a replacement for a set data structure
	distances := make(map[Coordinate]bool, 0)

	var sum int
	for y := 0; y < len(expanded); y++ {

		// Optimization to skip areas with empty lines
		if expandedSections[y] {
			y += expansionFactor
			continue
		}

		line := expanded[y]
		for x, c := range line {
			current := Coordinate{x, y}
			if c == '#' && !distances[current] {
				distances[current] = true

				/* We can calculate the distances right here,
				   which ensures that we won't have double values
				   and only need to walk the map once */
				for coord := range distances {
					if coord != current {
						sum += calculateCombinedXYOffset(coord, current)
					}
				}
			}
		}
	}
	fmt.Println("Sum: ", sum)
}
