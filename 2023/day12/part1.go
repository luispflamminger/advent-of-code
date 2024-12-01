package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput() ([][]rune, [][]int) {
	data, _ := os.ReadFile("input")
	splitData := strings.Split(string(data), "\n")
	conditions := make([][]rune, len(splitData))
	groupings := make([][]int, len(splitData))
	for i, v := range splitData {
		s := strings.Split(v, " ")
		for _, c := range s[0] {
			conditions[i] = append(conditions[i], c)
		}

		for _, c := range strings.Split(s[1], ",") {
			j, _ := strconv.Atoi(c)
			groupings[i] = append(groupings[i], j)
		}

	}

	return conditions, groupings
}

func printConditions(c [][]rune) {
	for _, l := range c {
		for _, r := range l {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}

func isValid(conditions []bool, groupings []int) bool {

	currentGrouping := 0
	for i := 0; i < len(conditions); i++ {
		v := conditions[i]
		if v {
			var groupLength int
			for j := i; j < len(conditions); j++ {
				if !conditions[j] {
					fmt.Println("hit group ending")
					fmt.Println("j: ", j)
					fmt.Println("i: ", i)
					groupLength = j - i
					i = j + 1
					break
				}
			}
			fmt.Println(groupLength)
			fmt.Println(groupings[currentGrouping])
			if groupLength > groupings[currentGrouping] {
				return false
			} else {
				currentGrouping++
			}
		}
	}
	return true

}

func main() {
	conditions, groupings := parseInput()
	fmt.Println("Conditions:")
	printConditions(conditions)
	fmt.Println("Groupings:")
	fmt.Println(groupings)

	testc := []bool{true, true, true, false, true, false}
	testg := []int{3, 1}
	fmt.Println(isValid(testc, testg))

}
