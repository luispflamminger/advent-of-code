package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func unsafe[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func seen(l int, c int) bool {
	for _, top := range seenTops {
		if top[0] == l && top[1] == c {
			return true
		}
	}
	return false
}

func trace(l int, c int) int {

	stepSum := 0

	if !part2 && grid[l][c] == 9 && !seen(l, c) {
		seenTops = append(seenTops, []int{l, c})
		return 1
	}

	if part2 && grid[l][c] == 9 {
		seenTops = append(seenTops, []int{l, c})
		return 1
	}

	// up
	if l > 0 && grid[l-1][c] == grid[l][c]+1 {
		stepSum += trace(l-1, c)
	}

	// down
	if l < len(grid)-1 && grid[l+1][c] == grid[l][c]+1 {
		stepSum += trace(l+1, c)
	}

	// left
	if c > 0 && grid[l][c-1] == grid[l][c]+1 {
		stepSum += trace(l, c-1)
	}

	// right
	if c < len(grid[0])-1 && grid[l][c+1] == grid[l][c]+1 {
		stepSum += trace(l, c+1)
	}

	return stepSum
}

var grid [][]int
var seenTops [][]int
var part2 bool

func main() {
	file, _ := os.ReadFile("./input")
	content := string(file)
	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1]

	grid = make([][]int, len(lines))
	for i := range grid {
		grid[i] = make([]int, len(lines[0]))
	}

	for i, row := range lines {
		for j, c := range row {
			grid[i][j] = unsafe(strconv.Atoi(string(c)))
		}
	}

	sum1 := 0
	sum2 := 0
	for r, row := range grid {
		for c, pos := range row {
			if pos == 0 {
				seenTops = [][]int{}
				part2 = false
				out := trace(r, c)
				//fmt.Println("trailhead at ", r, c, " has score ", out)
				sum1 += out

				part2 = true
				out = trace(r, c)
				//fmt.Println("trailhead at ", r, c, " has rating ", out)
				sum2 += out

			}
		}
	}

	fmt.Println("Sum of scores (part1): ", sum1)
	fmt.Println("Sum of ratings (part2): ", sum2)
}
