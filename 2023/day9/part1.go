package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// First dimension contains sequences, second contains values
func parseInput() [][]int {
	file, _ := os.ReadFile("input")
	lines := strings.Split(string(file), "\n")
	data := make([][]int, len(lines))
	for i, v := range lines {
		strs := strings.Split(v, " ")
		for _, w := range strs {
			num, _ := strconv.Atoi(w)
			data[i] = append(data[i], num)
		}
	}

	return data
}

// Returns a new []int containing the differences between
// each of the values of the input slice
func calculateExtrapolation(sequence []int) []int {
	extrapolated := make([]int, len(sequence)-1)
	for i := 1; i < len(sequence); i++ {
		extrapolated[i-1] = sequence[i] - sequence[i-1]
	}

	return extrapolated
}

// Checks if all elements of an []int are zero
func allZeroes(slice []int) bool {
	zeros := true
	for _, v := range slice {
		if v != 0 {
			zeros = false
			break
		}
	}
	return zeros
}

func main() {
	data := parseInput()

	// We store all of the forecasted numbers here
	forecasts := make([]int, len(data))

	for i, seq := range data {
		// We store all extrapolated sequences from the current sequence here
		extraps := [][]int{seq}
		for {
			// Append a new extrapolation until all values in the latest one are zeroes
			extraps = append(extraps, calculateExtrapolation(extraps[len(extraps)-1]))
			if allZeroes(extraps[len(extraps)-1]) {
				break
			}
		}

		// Move backwards through the extrapolations and append the sum of the last elements
		// of the current and previous extrapolation to the current extrapolation
		for i := len(extraps) - 2; i >= 0; i-- {
			extraps[i] = append(extraps[i], extraps[i][len(extraps[i])-1]+extraps[i+1][len(extraps[i+1])-1])
		}

		// Grab the latest value of the first sequence
		forecasts[i] = extraps[0][len(extraps[0])-1]
	}

	sum := 0
	for _, v := range forecasts {
		sum += v
	}

	fmt.Println("Sum of forecasts: ", sum)
}
