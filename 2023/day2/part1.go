package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		println("Error while reading file")
	}

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	content := string(file)
	games := strings.Split(content, "\n")
	sets := make([][]string, len(games))
	for i, v := range games {
		sets[i] = strings.Split(v, ";")
	}

	validGames := make([]int, 0)
	for gameCount, game := range sets {
		valid := true
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

			if red > maxRed || blue > maxBlue || green > maxGreen {
				valid = false
				break
			}
		}

		if valid {
			validGames = append(validGames, gameCount+1)
		}
	}
	fmt.Println("Valid Games: ", validGames)
	sum := 0
	for _, v := range validGames {
		sum += v
	}
	fmt.Println("Sum: ", sum)
}
