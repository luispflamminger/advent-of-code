package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func determineStrength(hand string) int {
	occurences := make(map[rune]int, 0)

	for _, c := range hand {
		occurences[c]++
	}

	// Find the max number of equal cards
	max := 0
	for _, v := range occurences {
		if v > max {
			max = v
		}
	}

	if max == 5 {
		return 6 // Five of a kind
	} else if max == 4 {
		return 5 // Four of a kind
	} else if max == 3 { // Could be either full house or three of a kind
		fullHouse := false
		for _, v := range occurences {
			if v == 2 {
				fullHouse = true
			}
		}
		if fullHouse {
			return 4 // Full house
		} else {
			return 3 // Three of a kind
		}
	} else if max == 1 {
		return 0 // High card
	} else {
		maxCounter := 0
		for _, v := range occurences {
			if v == max {
				maxCounter++
			}
		}
		if maxCounter == 2 { // Could be either two pair or one pair
			return 2 // Two pair
		} else {
			return 1 // One pair
		}
	}
}

// Returns the stronger hand
func breakTie(hand1 string, hand2 string) string {
	cardStrengths := map[rune]int{'A': 13, 'K': 12, 'Q': 11, 'J': 10, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1}

	hand1Wins := false
	for i := 0; i < 5; i++ {
		if cardStrengths[rune(hand1[i])] < cardStrengths[rune(hand2[i])] {
			hand1Wins = false
			break
		} else if cardStrengths[rune(hand1[i])] > cardStrengths[rune(hand2[i])] {
			hand1Wins = true
			break
		}
	}

	if hand1Wins {
		return hand1
	} else {
		return hand2
	}
}

func main() {
	file, _ := os.ReadFile("input")
	gameStrs := strings.Split(string(file), "\n")

	// keys are hands as string, values are slices with [0]: bet, [1]: strength
	games := make(map[string][]int, len(gameStrs))
	for _, str := range gameStrs {
		s := strings.Split(str, " ")
		betInt, _ := strconv.Atoi(s[1])
		games[s[0]] = []int{betInt, determineStrength(s[0])}
	}

	// Ordered hands are incrementally inserted
	orderedHands := make([]string, 0)
	for hand, game := range games {

		// If nothing inserted yet, insert the first hand
		if len(orderedHands) == 0 {
			orderedHands = append(orderedHands, hand)
			continue
		}

		placed := false
		for i := range orderedHands {
			// Check where to insert. We insert before the hand that is either stronger by type or, if type is equal, stronger by rank
			if game[1] < games[orderedHands[i]][1] || (game[1] == games[orderedHands[i]][1] && breakTie(hand, orderedHands[i]) != hand) {

				// Insert value into slice
				orderedHands = append(orderedHands[:i+1], orderedHands[i:]...)
				orderedHands[i] = hand

				placed = true
				break
			}
		}

		// If hand has not been placed yet, it is the strongest and needs to be inserted at the end
		if !placed {
			orderedHands = append(orderedHands, hand)
		}
	}

	total := 0
	// Sum up values
	for i, hand := range orderedHands {
		total += (i + 1) * games[hand][0]
	}

	fmt.Println("Total: ", total)
}
