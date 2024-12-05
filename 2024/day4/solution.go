package main

import (
	"fmt"
	"os"
	"strings"
)

func unsafe[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func wrap(x int, offset int, length int) int {
	if x+offset >= length {
		return 0
	} else if x+offset < 0 {
		return 0
	} else {
		return x + offset
	}
}

func main() {
	file, _ := os.ReadFile("./input")

	content := string(file)

	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1]

	// split lines by char and get [][]rune
	runes := make([][]rune, len(lines))
	for i, line := range lines {
		runes[i] = []rune(line)
	}

	sum := 0
	for c_o, l := range runes {
		for c_i, c := range l {
			if c == 'X' {
				sum += search(runes, c_o, c_i)
			}
		}
	}

	fmt.Println("sum1: ", sum)

	sum = 0
	for c_o := 1; c_o < len(runes)-1; c_o++ {
		for c_i := 1; c_i < len(runes[0])-1; c_i++ {
			if runes[c_o][c_i] == 'A' {
				sum += search2(runes, c_o, c_i)
			}
		}
	}

	fmt.Println("sum2: ", sum)

}

func search(runes [][]rune, y int, x int) int {
	x_len := len(runes[0])
	y_len := len(runes)

	sum := 0
	if y < y_len-3 && x < x_len-3 &&
		runes[y+1][x+1] == 'M' &&
		runes[y+2][x+2] == 'A' &&
		runes[y+3][x+3] == 'S' {
		sum += 1
	}
	if y > 2 && x > 2 &&
		runes[y-1][x-1] == 'M' &&
		runes[y-2][x-2] == 'A' &&
		runes[y-3][x-3] == 'S' {
		sum += 1
	}
	if y > 2 && x < x_len-3 &&
		runes[y-1][x+1] == 'M' &&
		runes[y-2][x+2] == 'A' &&
		runes[y-3][x+3] == 'S' {
		sum += 1
	}
	if y < y_len-3 && x > 2 &&
		runes[y+1][x-1] == 'M' &&
		runes[y+2][x-2] == 'A' &&
		runes[y+3][x-3] == 'S' {
		sum += 1
	}
	if y < y_len-3 &&
		runes[y+1][x+0] == 'M' &&
		runes[y+2][x+0] == 'A' &&
		runes[y+3][x+0] == 'S' {
		sum += 1
	}
	if y > 2 &&
		runes[y-1][x+0] == 'M' &&
		runes[y-2][x+0] == 'A' &&
		runes[y-3][x+0] == 'S' {
		sum += 1
	}
	if x < x_len-3 &&
		runes[y][x+1] == 'M' &&
		runes[y][x+2] == 'A' &&
		runes[y][x+3] == 'S' {
		sum += 1
	}
	if x > 2 &&
		runes[y][x-1] == 'M' &&
		runes[y][x-2] == 'A' &&
		runes[y][x-3] == 'S' {
		sum += 1
	}
	return sum

}

func search2(runes [][]rune, y int, x int) int {
	if runes[y-1][x-1] == 'M' && runes[y-1][x+1] == 'M' && runes[y+1][x-1] == 'S' && runes[y+1][x+1] == 'S' {
		return 1
	}
	if runes[y-1][x-1] == 'S' && runes[y-1][x+1] == 'S' && runes[y+1][x-1] == 'M' && runes[y+1][x+1] == 'M' {
		return 1
	}
	if runes[y-1][x-1] == 'M' && runes[y-1][x+1] == 'S' && runes[y+1][x-1] == 'M' && runes[y+1][x+1] == 'S' {
		return 1
	}
	if runes[y-1][x-1] == 'S' && runes[y-1][x+1] == 'M' && runes[y+1][x-1] == 'S' && runes[y+1][x+1] == 'M' {
		return 1
	}
	return 0
}
