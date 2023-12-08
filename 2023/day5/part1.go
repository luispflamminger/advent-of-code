package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// We split the input by hand so we don't need to parse as much
	seedInput, _ := os.ReadFile("seeds")
	// Need to remove descriptions from map input, not needed
	mapInput, _ := os.ReadFile("maps")
	seedStrings := strings.Split(string(seedInput), " ")
	seeds := make([]int, len(seedStrings))
	for i, seedString := range seedStrings {
		seeds[i], _ = strconv.Atoi(seedString)
	}
	rawMaps := strings.Split(string(mapInput), "\n\n")

	// First level: maps, Second level: mappings within a map, third level: [0]: dst, [1]: src, [2]: rng
	maps := make([][][]int, len(rawMaps))

	// Names don't matter, we just treat each map as one step and iterate over them
	for mapCount, mapp := range rawMaps {
		maps[mapCount] = make([][]int, 0)
		for _, mapping := range strings.Split(mapp, "\n") {
			splitMapping := strings.Split(mapping, " ")
			dst, _ := strconv.Atoi(string(splitMapping[0]))
			src, _ := strconv.Atoi(string(splitMapping[1]))
			rng, _ := strconv.Atoi(string(splitMapping[2]))
			mapping := []int{dst, src, rng}
			maps[mapCount] = append(maps[mapCount], mapping)
		}
	}

	// Iterate over seeds, then maps and then mappings
	for seedCount, seed := range seeds {
		for _, mapp := range maps {
			for _, mapping := range mapp {
				// Check if source seed is within mapping
				if seed >= mapping[1] && seed <= mapping[1]+mapping[2] {
					// If it's in the range, we can calculate the offset from the
					// src range start and add it to the dest range start to get the dest seed
					offset := seed - mapping[1]
					seed = mapping[0] + offset
					break
				}
			}
		}
		// We just store the seed in the same slice, we don't need to keep track of revisions
		seeds[seedCount] = seed
	}

	// At this point, all seed values are converted to location numbers. We just find the min.
	minRes := seeds[0]
	for _, seed := range seeds {
		if seed < minRes {
			minRes = seed
		}
	}

	fmt.Println("Lowest location number: ", minRes)

}
