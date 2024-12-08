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

func calculate(op string, prevRes int, depth int, nums []int, maxDepth int, expectedRes int) bool {
	if op == "+" {
		prevRes = prevRes + nums[depth]
	} else if op == "*" {
		prevRes = prevRes * nums[depth]
	} else if op == "||" {
		prevRes = unsafe(strconv.Atoi(fmt.Sprintf("%d%d", prevRes, nums[depth])))
	}

	if depth == maxDepth {
		return prevRes == expectedRes
	}

	if calculate("+", prevRes, depth+1, nums, maxDepth, expectedRes) {
		return true
	}

	if calculate("*", prevRes, depth+1, nums, maxDepth, expectedRes) {
		return true
	}

	if part2 && calculate("||", prevRes, depth+1, nums, maxDepth, expectedRes) {
		return true
	}

	return false
}

var part2 bool = false

func main() {

	file, _ := os.ReadFile("./input")
	content := string(file)

	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1]

	eqs := make([][]int, len(lines))

	for i, line := range lines {
		rx := regexp.MustCompile(`\d+`)
		strs := rx.FindAllString(line, -1)
		eqs[i] = make([]int, len(strs))
		for j, str := range strs {
			eqs[i][j] = unsafe(strconv.Atoi(str))
		}
	}

	res := 0
	for _, eq := range eqs {
		if calculate("", eq[1], 0, eq[1:], len(eq[1:])-1, eq[0]) {
			res += eq[0]
		}
	}

	fmt.Println("1: ", res)

	part2 = true
	res = 0
	for _, eq := range eqs {
		if calculate("", eq[1], 0, eq[1:], len(eq[1:])-1, eq[0]) {
			res += eq[0]
		}
	}

	fmt.Println("2: ", res)
}
