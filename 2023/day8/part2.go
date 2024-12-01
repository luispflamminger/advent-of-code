// This is the simple brute force solution, but it takes too long...
package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput() (string, map[string][]string) {
	file, _ := os.ReadFile("input")
	rawData := strings.Split(string(file), "\n\n")

	rl := rawData[0]

	nodes := make(map[string][]string, 0)
	for _, line := range strings.Split(rawData[1], "\n") {
		splitLine := strings.Split(line, "=")
		splitDirs := strings.Split(splitLine[1], ",")
		left := strings.Trim(splitDirs[0], "( ")
		right := strings.Trim(splitDirs[1], ") ")
		nodes[strings.Trim(splitLine[0], " ")] = []string{left, right}
	}

	// rl: Just the RLRLRLRLR string
	// nodes: keys are nodes, values are slices where [0] is the left and [1] is the right instruction
	return rl, nodes

}
func main() {

	inst, nodes := parseInput()
	instCounter := 0
	totalCounter := 0
	currentNodes := make([]string, 0)

	// Find all nodes that end with A
	for k, _ := range nodes {
		if strings.HasSuffix(k, "A") {
			currentNodes = append(currentNodes, k)
		}
	}

	for {
		// Reset to first instruction if end is reached
		if instCounter == len(inst) {
			instCounter = 0
		}

		end := true
		for i, node := range currentNodes {
			if inst[instCounter] == 'R' {
				currentNodes[i] = nodes[node][1]
			} else {
				currentNodes[i] = nodes[node][0]
			}

			if !strings.HasSuffix(currentNodes[i], "Z") {
				end = false
			}
		}

		totalCounter++
		instCounter++

		if end {
			break
		}
	}

	fmt.Println("Total Steps: ", totalCounter)
}
