package main

import (
	"fmt"
	"os"
	"regexp"
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

	input := strings.ReplaceAll(content, "\n", "")
	rx := regexp.MustCompile(`mul\(\d+,\d+\)`)
	muls := rx.FindAllString(input, -1)

	sum := 0
	for _, mul := range muls {

		rx2 := regexp.MustCompile(`\d+`)
		strs := rx2.FindAllString(mul, -1)

		x := unsafe(strconv.Atoi(strs[0]))
		y := unsafe(strconv.Atoi(strs[1]))

		sum += x * y

	}

	fmt.Println("Sum all (part1): ", sum)

	rx = regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)
	muls = rx.FindAllString(input, -1)

	res := make([]string, 0)
	enabled := true
	for _, mul := range muls {
		if mul == "do()" {
			enabled = true
		} else if mul == "don't()" {
			enabled = false
		} else if enabled {
			res = append(res, mul)
		}
	}

	sum = 0
	for _, mul := range res {

		rx2 := regexp.MustCompile(`\d+`)
		strs := rx2.FindAllString(mul, -1)

		x := unsafe(strconv.Atoi(strs[0]))
		y := unsafe(strconv.Atoi(strs[1]))

		sum += x * y
	}

	fmt.Println("Sum enabled (part2): ", sum)
}
