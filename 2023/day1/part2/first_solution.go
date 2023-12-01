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

	replaceMap := map[string]string{"one": "o1e", "two": "t2o", "three": "t3e", "four": "f4r", "five": "f5e", "six": "s6x", "seven": "s7n", "eight": "e8t", "nine": "n9e"}
	var numbers []int
	var sum int

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		var digits []rune

		// Convert written out numbers to digits
		for i := 0; i < len(line); i++ {
			substr := line[i:]
			for k, v := range replaceMap {
				if strings.HasPrefix(substr, k) {
					line = strings.Replace(line, k, v, 1)
					i = 0
					break
				}
			}
		}

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
