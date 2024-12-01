package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput() [][][]rune {
	data, _ := os.ReadFile("input")
	notes := strings.Split(string(data), "\n\n")

	res := make([][][]rune, len(notes))
	for i, note := range notes {
		res[i] = make([][]rune, len(strings.Split(string(note), "\n")))
		for j, line := range strings.Split(string(note), "\n") {
			res[i][j] = make([]rune, len(line))
			for _, c := range line {
				res[i][j] = append(res[i][j], c)
			}
		}

	}
	return res
}

func isEqual(a []rune, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func printRunes(line []rune) {
	for _, c := range line {
		fmt.Print(string(c))
	}
	fmt.Println()
}

// returns nil if no symmetry can be found
func findStart(note [][]rune) [][]rune {

	printRunes(note[0])
	printRunes(note[1])
	printRunes(note[len(note)-2])
	printRunes(note[len(note)-1])

	if isEqual(note[0], note[len(note)-1]) {
		println("equal")
		return note
	}

	if isEqual(note[0], note[len(note)-2]) {
		println("start")
		return note[:len(note)-1]
	}

	if isEqual(note[1], note[len(note)-1]) {
		println("end")
		return note[1:]
	}

	println("nil")
	return nil
}

func isSymmetric(note [][]rune) bool {
	j := len(note) - 1
	for i := 0; i < len(note); i++ {
		j--

		if i == j {
			break
		}

		if !isEqual(note[i], note[j]) {
			return false
		}

	}
	return true
}

func transformNote(note [][]rune) [][]rune {

	// initialize res
	res := make([][]rune, len(note[0]))
	for i := range res {
		res[i] = make([]rune, len(note))
	}

	// populate res
	for i, line := range note {
		for j, c := range line {
			res[j][i] = c
		}
	}

	fmt.Println("transformed")
	for _, line := range res {
		printRunes(line)
	}
	println()

	return res
}

func main() {

	data := parseInput()

	for _, note := range data {

		var res = 0
		truncNote := findStart(note)

		if truncNote != nil {
			if isSymmetric(truncNote) {
				res += len(truncNote) / 2 * 100
			}
		}

		transformedNote := transformNote(note)

		truncNote = findStart(transformedNote)

		if truncNote != nil {
			if isSymmetric(truncNote) {
				res += len(truncNote) / 2 * 100
			}
		}

		fmt.Println(res)
	}
}
