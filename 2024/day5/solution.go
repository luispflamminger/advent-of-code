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

func main() {
	file, _ := os.ReadFile("./input")

	content := string(file)

	inputs := strings.Split(content, "\n\n")

	rules := strings.Split(inputs[0], "\n")
	manuals := strings.Split(inputs[1], "\n")
	manuals = manuals[:len(manuals)-1]

	//process rules
	rulesMap := make(map[int][]int)

	for _, rule := range rules {
		ruleSplit := strings.Split(rule, "|")
		firstPage := unsafe(strconv.Atoi(ruleSplit[0]))
		secPage := unsafe(strconv.Atoi(ruleSplit[1]))
		if rulesMap[firstPage] == nil {
			rulesMap[firstPage] = []int{secPage}
		} else {
			rulesMap[firstPage] = append(rulesMap[firstPage], secPage)
		}
	}

	//process pages
	validSum := 0
	invalidSum := 0
	for _, manual := range manuals {
		pages := strings.Split(manual, ",")
		valid := true
		seen := make(map[int]bool)
		for _, page := range pages {
			pageInt := unsafe(strconv.Atoi(page))
			if rulesMap[pageInt] != nil {
				for _, after := range rulesMap[pageInt] {
					if seen[after] {
						valid = false
						break
					}
				}
			}
			seen[pageInt] = true
		}

		if valid {
			validSum += unsafe(strconv.Atoi(pages[len(pages)/2]))
		} else {
			valid = false
			for !valid {
				seen2 := make(map[int]bool)
				for i, page := range pages {
					valid = true
					pageInt := unsafe(strconv.Atoi(page))
					if rulesMap[pageInt] != nil {
						for _, after := range rulesMap[pageInt] {
							if seen2[after] {
								// bubble up invalid pages
								prev := pages[i-1]
								pages[i-1] = page
								pages[i] = prev
								valid = false
								break
							}
						}
						if !valid {
							break
						}
					}
					seen2[pageInt] = true
				}
			}
			invalidSum += unsafe(strconv.Atoi(pages[len(pages)/2]))
		}

	}
	fmt.Println("Part 1: ", validSum)
	fmt.Println("Part 2: ", invalidSum)
}
