package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	filePath := "../input.txt"
	file, err := os.ReadFile(filePath)
	content := string(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	var numbers []int
	lines := strings.Split(content, "\n")
	for lineCounter := 0; lineCounter < len(lines); lineCounter++ {
		line := lines[lineCounter]
		var digits []rune

		for charCounter := 0; charCounter < len(line); charCounter++ {
			char := rune(line[charCounter])
			if unicode.IsDigit(char) {
				digits = append(digits, char)
			}
		}
		numString := string(digits[0]) + string(digits[len(digits)-1])
		num, _ := strconv.Atoi(numString)
		numbers = append(numbers, num)
	}

	var sum int
	for numCounter := 0; numCounter < len(numbers); numCounter++ {
		sum += numbers[numCounter]
	}

	fmt.Println("Numbers: ", numbers)

	println("Sum: ", sum)
}
