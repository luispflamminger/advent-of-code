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

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func remove(slice []int, s int) []int {
	cpy := make([]int, len(slice))
	copy(cpy, slice)
	return append(cpy[:s], cpy[s+1:]...)
}

func validateReport(r []int) bool {
	prev_level := 0
	dec := false
	safe := false

	for count, level := range r {
		if count == 0 {
			prev_level = level
			continue
		}

		if abs(level-prev_level) <= 3 && abs(level-prev_level) > 0 {
			if count == 1 {
				if prev_level < level {
					dec = false
				} else {
					dec = true
				}
				prev_level = level
				continue
			}

			if (dec && prev_level > level) || (!dec && prev_level < level) {
				safe = true
				prev_level = level
				continue
			} else {
				safe = false
				break
			}
		} else {
			safe = false
			break
		}
	}

	return safe
}

func main() {
	file, err := os.ReadFile("./input")

	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	content := string(file)

	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1]

	reports := make([][]int, len(lines))

	for count, line := range lines {
		nums := strings.Split(line, " ")
		report := make([]int, len(nums))
		for i, num := range nums {
			report[i] = unsafe(strconv.Atoi(num))
		}

		reports[count] = report
	}

	safe_count := 0

	for _, report := range reports {
		if validateReport(report) {
			safe_count++
		}
	}

	fmt.Println("Safe count (undampened): ", safe_count)
	safe_count = 0

	for c, report := range reports {
		if validateReport(report) {
			safe_count++
		} else {
			fmt.Println("unsafe report: ", c, " : ", report)
			for i := 0; i < len(report); i++ {
				new := remove(report, i)
				fmt.Println("trying: ", new)
				if validateReport(new) {
					safe_count++
					break
				}
			}

		}
	}

	fmt.Println("Safe count (dampened): ", safe_count)
}
