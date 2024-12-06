package main

import (
	"fmt"
	"os"
	"strings"
)

// contains the input
var runes [][]rune

// contains a track of the walked path
var track [][]rune

func trace(l int, c int, dir rune) bool {

	if track[l][c] == dir {
		// if the track already contains
		// this direction at this position, we are in a loop
		return true
	}
	// mark the track with the current direction
	track[l][c] = dir

	// check if off map
	if dir == '^' && l == 0 {
		return false
	} else if dir == '>' && c == len(runes[0])-1 {
		return false
	} else if dir == 'v' && l == len(runes)-1 {
		return false
	} else if dir == '<' && c == 0 {
		return false
	}

	// check if hitting obstacle (account for multiple turns in same spot)
	turn := true
	for turn {
		if dir == '^' && runes[l-1][c] == '#' {
			dir = '>'
		} else if dir == '>' && runes[l][c+1] == '#' {
			dir = 'v'
		} else if dir == 'v' && runes[l+1][c] == '#' {
			dir = '<'
		} else if dir == '<' && runes[l][c-1] == '#' {
			dir = '^'
		} else {
			turn = false
		}
	}

	// recursively trace the path
	if dir == '^' {
		return trace(l-1, c, dir)
	} else if dir == '>' {
		return trace(l, c+1, dir)
	} else if dir == 'v' {
		return trace(l+1, c, dir)
	} else if dir == '<' {
		return trace(l, c-1, dir)
	}

	panic("should be unreachable")
}

func main() {

	file, _ := os.ReadFile("./input")

	content := string(file)

	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1]

	// initialize track
	track = make([][]rune, len(lines))
	for i := range track {
		track[i] = make([]rune, len(lines[0]))
	}

	// initialize input runes
	runes = make([][]rune, len(lines))
	for i, line := range lines {
		runes[i] = []rune(line)
	}

	// find starting point and direction
	startL := 0
	startC := 0
	startDir := ' '

	done := false
	for cO, l := range runes {
		for cI, c := range l {
			if c == '^' || c == 'v' || c == '<' || c == '>' {
				startL = cO
				startC = cI
				startDir = c
				done = true
				break
			}
		}
		if done {
			break
		}
	}

	trace(startL, startC, startDir)

	// count seen positions
	sum := 0
	for _, l := range track {
		for _, c := range l {
			if c == '^' || c == 'v' || c == '<' || c == '>' {
				sum++
			}
		}

	}

	fmt.Println("sum1: ", sum)

	// brute force search for loops
	sum = 0
	for cO, l := range runes {
		for cI, c := range l {

			// reset track to empty
			for i := range track {
				for j := range track[i] {
					track[i][j] = ' '
				}
			}

			if c == '.' {
				// place obstacle
				runes[cO][cI] = '#'
				if trace(startL, startC, startDir) {
					sum++
				}
				// remove obstacle
				runes[cO][cI] = '.'
			}
		}
	}

	fmt.Println("sum2: ", sum)
}
