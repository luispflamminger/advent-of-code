package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

	file, err := os.ReadFile("input")
	if err != nil {
		println("Error while reading file")
	}

	content := string(file)
	games := strings.Split(content, "\n")
	sets := make([][]string, len(games))
	for i, v := range games {
		sets[i] = strings.Split(v, ";")
	}

	// [[red, blue, green][...]]
	maxes := make([][]int, len(games))

	for gameCount, game := range sets {
		for _, set := range game {
			regexRed := regexp.MustCompile("(\\d+?)(?: red)")
			regexBlue := regexp.MustCompile("(\\d+?)(?: blue)")
			regexGreen := regexp.MustCompile("(\\d+?)(?: green)")
			resRed := regexRed.FindSubmatch([]byte(set))
			resBlue := regexBlue.FindSubmatch([]byte(set))
			resGreen := regexGreen.FindSubmatch([]byte(set))

			red := 0
			blue := 0
			green := 0

			if len(resRed) >= 2 {
				red, _ = strconv.Atoi(string(resRed[1]))
			}
			if len(resBlue) >= 2 {
				blue, _ = strconv.Atoi((string(resBlue[1])))
			}
			if len(resGreen) >= 2 {
				green, _ = strconv.Atoi((string(resGreen[1])))
			}

			if len(maxes[gameCount]) > 0 {
				red = Max(maxes[gameCount][0], red)
				blue = Max(maxes[gameCount][1], blue)
				green = Max(maxes[gameCount][2], green)
			}

			maxes[gameCount] = []int{red, blue, green}
		}
	}

	sum := 0
	for _, v := range maxes {
		sum += v[0] * v[1] * v[2]
	}

	fmt.Println(maxes)

	fmt.Println("Sum of powers: ", sum)
}
