package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.ReadFile("../input.txt")

	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	content := string(file)
	var numbers []int
	var sum int

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		var digits []rune

		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, char)
			}
		}

		numString := string(digits[0]) + string(digits[len(digits)-1])
		num, _ := strconv.Atoi(numString)
		numbers = append(numbers, num)
	}

	for _, number := range numbers {
		sum += number
	}

	fmt.Println("Numbers: ", numbers)

	println("Sum: ", sum)
}
