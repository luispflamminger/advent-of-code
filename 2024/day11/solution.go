package main

import (
	"flag"
	"fmt"
	"math"
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

// faster than converting to string
func lenInt(i int) int {
	if i == 0 {
		return 1
	}

	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func blinkSemiOptimized(stones []int) []int {

	// need to keep track because we are appending
	// and don't want to iterate over the new values
	initialLen := len(stones)
	numDigits := 0

	for i := 0; i < initialLen; i++ {
		if stones[i] == 0 {
			stones[i] = 1
			continue
		}

		numDigits = lenInt(stones[i])

		if numDigits%2 == 0 {
			numDigits := lenInt(stones[i])
			// using exponentials and modulo to split the number in half
			// is faster than converting to string and splitting
			left := stones[i] / int(math.Pow10(numDigits/2))
			right := stones[i] % (left * int(math.Pow10(numDigits/2)))
			stones[i] = left
			// append to the slice instead of inserting the value
			// reduces memory accesses
			stones = append(stones, right)
		} else {
			stones[i] *= 2024
		}
	}

	return stones
}

func blinkOptimized(stones []int, blinks int) int {

	sum := 0
	for _, s := range stones {
		sum += rec(s, blinks, 0)
	}

	return sum
}

func rec(stone int, maxDepth int, curDepth int) int {

	// if value in cache, return it directly
	if val, ok := seen[strconv.Itoa(stone)+","+strconv.Itoa(curDepth)]; ok {
		return val
	}

	if curDepth == maxDepth {
		return 1
	}

	numDigits := lenInt(stone)
	ret := 0

	if stone == 0 {
		ret = rec(1, maxDepth, curDepth+1)
	} else if numDigits%2 == 0 {
		numDigits := lenInt(stone)
		left := stone / int(math.Pow10(numDigits/2))
		right := stone % (left * int(math.Pow10(numDigits/2)))
		ret = rec(left, maxDepth, curDepth+1) + rec(right, maxDepth, curDepth+1)
	} else {
		ret = rec(stone*2024, maxDepth, curDepth+1)
	}

	// cache values calculated for a given stone at a given depth
	seen[strconv.Itoa(stone)+","+strconv.Itoa(curDepth)] = ret
	return ret

}

var seen = make(map[string]int)

func main() {
	flag.Parse()

	file, _ := os.ReadFile("./input")
	content := string(file)
	strings := strings.Split(strings.Split(content, "\n")[0], " ")

	// allocating a big slice to avoid reallocation
	stones := make([]int, len(strings), 1000000000)
	for i, s := range strings {
		stones[i] = unsafe(strconv.Atoi(s))
	}

	stones2 := make([]int, len(strings))
	copy(stones2, stones)

	for i := 1; i < 26; i++ {
		stones = blinkSemiOptimized(stones)
	}
	fmt.Println("sum1:", len(stones))

	fmt.Println("sum2: ", blinkOptimized(stones2, 75))
}
