// NOT SOLVED
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
	// [[0]: start of range, [1]: length of range ]
	seedRanges := make([][]int, 0)

	for i := 0; i < len(seedStrings); i += 2 {
		start, _ := strconv.Atoi(seedStrings[i])
		length, _ := strconv.Atoi(seedStrings[i+1])
		seedRange := []int{start, length}
		seedRanges = append(seedRanges, seedRange)
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
	for _, seedRange := range seedRanges {
		for _, mapp := range maps {
			newSeedRanges := make([][]int, 0)
			for _, mapping := range mapp {
				mappingRangeStart := mapping[1]
				mappingRangeEnd := mappingRangeStart + mapping[2] - 1
				seedRangeStart := seedRange[0]
				fmt.Println("Seed range start: ", seedRangeStart)
				seedRangeEnd := seedRangeStart + seedRange[1] - 1
				fmt.Println("Seed range end: ", seedRangeEnd)

				// Range lies completely in mapping -> transfer complete range
				if mappingRangeStart <= seedRangeStart && mappingRangeEnd >= seedRangeEnd {
					fmt.Println("Case 1")
					offset := seedRangeStart - mappingRangeStart
					newStart := mapping[0] + offset
					newSeedRange := []int{newStart, seedRange[1]}
					newSeedRanges = append(newSeedRanges, newSeedRange)
					fmt.Println("New seed ranges: ", newSeedRanges)
					break
				}

				// Mapping lies completely in range -> Split range into 3 ranges and transfer the middle one
				if seedRangeStart < mappingRangeStart && seedRangeEnd > mappingRangeEnd {
					fmt.Println("Case 2")
					newSeedRange1 := []int{seedRangeStart, mappingRangeStart - seedRangeStart}
					newSeedRange2 := []int{mapping[0], mapping[2]}
					newSeedRange3 := []int{mappingRangeEnd + 1, seedRangeEnd - (mappingRangeEnd + 1)}
					newSeedRanges = append(newSeedRanges, newSeedRange1, newSeedRange2, newSeedRange3)
					fmt.Println("New seed range: ", newSeedRanges)
					break
				}

				// Range is partially lower than mapping -> Keep range that's outside, transform the other
				if seedRangeStart < mappingRangeStart && seedRangeEnd >= mappingRangeStart && seedRangeEnd <= mappingRangeEnd {
					fmt.Println("Case 3")
					newSeedRange1 := []int{seedRangeStart, mappingRangeStart - 1 - seedRangeStart}
					fmt.Println("New seed range 1: ", newSeedRange1)
					newSeedRange2 := []int{mapping[0], mappingRangeEnd - mappingRangeStart}
					fmt.Println("New seed range 2: ", newSeedRange2)
					newSeedRanges = append(newSeedRanges, newSeedRange1, newSeedRange2)
					fmt.Println("New seed ranges: ", newSeedRanges)
					break
				}

				// Range is partially higher than mapping -> Keep range that's outside, transform the other
				if seedRangeStart >= mappingRangeStart && seedRangeStart <= mappingRangeEnd && seedRangeEnd > mappingRangeEnd {
					fmt.Println("Case 4")
					newSeedRange1 := []int{mapping[0] + (mappingRangeStart - seedRangeStart), mappingRangeEnd - seedRangeStart}
					newSeedRange2 := []int{mappingRangeEnd + 1, seedRangeEnd - (mappingRangeEnd + 1)}
					newSeedRanges = append(newSeedRanges, newSeedRange1, newSeedRange2)
					fmt.Println("New seed range: ", newSeedRanges)
					break
				}

				// complete mismatch
				fmt.Println("Not touched")
				newSeedRanges = append(newSeedRanges, seedRange)
			}
			seedRanges = newSeedRanges
			break
		}
	}

	min := seedRanges[0][0]
	for _, seedRange := range seedRanges {
		seed := seedRange[0]
		if seed < min {
			min = seed
		}
	}

	fmt.Println("Lowest location number: ", min)

}
