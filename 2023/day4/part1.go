package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		println("Error while reading file.")
	}
	unprocessed := strings.Split(string(content), "\n")
	// Layer 1: Cards, Layer 2: [0]=winning, [1]=mine, Layer 3: Numbers
	cards := make([][][]int, len(unprocessed))

	for i := 0; i < len(unprocessed); i++ {
		line := unprocessed[i]
		rawCards := strings.Split(strings.Split(line, ":")[1], "|")
		mine := make([]int, 0)
		winning := make([]int, 0)
		for _, v := range strings.Split(rawCards[0], " ") {
			str, err := strconv.Atoi(v)
			if err == nil {
				winning = append(winning, str)
			}
		}

		for _, v := range strings.Split(rawCards[1], " ") {
			str, err := strconv.Atoi(v)
			if err == nil {
				mine = append(mine, str)
			}
		}
		cards[i] = [][]int{winning, mine}
	}

	sum := 0
	for cardCount, card := range cards {
		cardValue := 0

		for _, number := range card[1] {
			if contains(card[0], number) {
				fmt.Println("Card ", cardCount+1, " has winning number ", number)
				if cardValue < 1 {
					cardValue = 1
				} else {
					cardValue *= 2
				}
			}
		}

		fmt.Println("Value of card ", cardCount+1, ": ", cardValue)
		sum += cardValue
	}

	fmt.Println("Total value of pile: ", sum)
}
