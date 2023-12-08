package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInputPart1() [][]int {
	file, _ := os.ReadFile("input")
	rawData := strings.Split(string(file), "\n")
	data := make([][]int, 0)

	for _, v := range strings.Split(rawData[0], " ") {
		str, err := strconv.Atoi(v)
		if err == nil {
			data = append(data, []int{str})
		}
	}

	assignmentCounter := 0
	for _, v := range strings.Split(rawData[1], " ") {
		str, err := strconv.Atoi(v)
		if err == nil {
			data[assignmentCounter] = append(data[assignmentCounter], str)
			assignmentCounter++
		}
	}

	return data
}

func parseInputPart2() [][]int {
	file, _ := os.ReadFile("input")
	rawData := strings.Split(string(file), "\n")
	data := make([][]int, 0)

	timeStr := ""
	for _, v := range rawData[0] {
		if v != ' ' {
			timeStr += string(v)
		}
	}

	distStr := ""
	for _, v := range rawData[1] {
		if v != ' ' {
			distStr += string(v)
		}
	}

	time, _ := strconv.Atoi(timeStr)
	dist, _ := strconv.Atoi(distStr)

	data = append(data, []int{time, dist})
	return data
}

func main() {
	// data := parseInputPart1() // <-- Part 1
	data := parseInputPart2() //    <-- Part 2
	fmt.Println(data)

	results := make([]int, len(data))
	for i, race := range data {
		time := race[0]
		distToBeat := race[1]
		viableSolutions := 0

		for i := 0; i <= time; i++ {
			if ((time - i) * i) > distToBeat {
				viableSolutions++
			}
		}
		results[i] = viableSolutions
	}

	fmt.Println("Results: ", results)
	final := 1
	for _, res := range results {
		final *= res
	}

	fmt.Println("Multiplied: ", final)

}
