package main

import (
	"fmt"
	"os"
	"sort"
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

func main() {
	file, err := os.ReadFile("./input")

	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	content := string(file)

	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1]

	list1 := make([]int, len(lines))
	list2 := make([]int, len(lines))
	var distance_sum int

	for count, line := range lines {
		nums := strings.Split(line, "   ")
		list1[count] = unsafe(strconv.Atoi(nums[0]))
		list2[count] = unsafe(strconv.Atoi(nums[1]))
	}

	sort.Ints(list1)
	sort.Ints(list2)

	for i := 0; i < len(list1); i++ {
		dis := abs(list1[i] - list2[i])
		distance_sum += dis
	}

	fmt.Println("Sum of distances (part1): ", distance_sum)

	sim_score := 0
	for _, num1 := range list1 {

		for _, num2 := range list2 {
			if num1 == num2 {
				sim_score += num2
			}
		}
	}

	fmt.Println("Sim score (part2): ", sim_score)
}
