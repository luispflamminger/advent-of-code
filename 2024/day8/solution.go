package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("./input")
	content := string(file)
	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1]

	antennas := make([][]rune, 0)
	antennas = make([][]rune, len(lines))
	for i, line := range lines {
		antennas[i] = []rune(line)
	}

	antinodes := make([][]bool, len(antennas))
	for i := range antinodes {
		antinodes[i] = make([]bool, len(antennas[0]))
	}

	for i1, row1 := range antennas {
		for j1, pos1 := range row1 {
			if pos1 != '.' {
				for i2, row2 := range antennas {
					for j2, pos2 := range row2 {
						if pos2 == pos1 && i1 != i2 && j1 != j2 {
							di := i2 - i1
							dj := j2 - j1
							// if inbounds
							if i1-di >= 0 && i1-di < len(antennas) && j1-dj >= 0 && j1-dj < len(antennas[0]) {
								antinodes[i1-di][j1-dj] = true
							}
							if i2+di >= 0 && i2+di < len(antennas) && j2+dj >= 0 && j2+dj < len(antennas[0]) {
								antinodes[i2+di][j2+dj] = true
							}
						}
					}
				}

			}
		}
	}

	// count trues
	count := 0
	for _, row := range antinodes {
		for _, val := range row {
			if val {
				count++
			}
		}
	}

	fmt.Println("1: ", count)

	// reset antinodes
	for i := range antinodes {
		for j := range antinodes[i] {
			antinodes[i][j] = false
		}
	}

	for i1, row1 := range antennas {
		for j1, pos1 := range row1 {
			if pos1 != '.' {
				for i2, row2 := range antennas {
					for j2, pos2 := range row2 {
						if pos2 == pos1 && i1 != i2 && j1 != j2 {
							di := i2 - i1
							dj := j2 - j1
							// if inbounds
							mul := 0
							for i1-mul*di >= 0 && i1-mul*di < len(antennas) && j1-mul*dj >= 0 && j1-mul*dj < len(antennas[0]) {
								antinodes[i1-mul*di][j1-mul*dj] = true
								mul++
							}
							mul = 0
							for i2+mul*di >= 0 && i2+mul*di < len(antennas) && j2+mul*dj >= 0 && j2+mul*dj < len(antennas[0]) {
								antinodes[i2+mul*di][j2+mul*dj] = true
								mul++
							}
						}
					}
				}

			}
		}
	}

	count = 0
	for _, row := range antinodes {
		for _, val := range row {
			if val {
				count++
			}
		}
	}

	fmt.Println("2: ", count)

}
