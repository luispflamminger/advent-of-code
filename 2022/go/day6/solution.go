package main

import (
	"fmt"
	"os"
	"strings"
)

// Enqueue while keeping length. Return dequeued value
func rotateQueue(queue []string, value string) string {
	dequeuedItem := queue[len(queue)-1]
	for i := len(queue) - 1; i >= 0; i-- {
		if i < len(queue)-1 {
			queue[i+1] = queue[i]
		}
	}
	queue[0] = value
	return dequeuedItem
}

// Returns true if all items in the slice are unique
func uniqueItems(slice []string) bool {
	seen := []string{}
	for _, v := range slice {
		for _, seenValue := range seen {
			if seenValue == v {
				return false
			}
		}
		seen = append(seen, v)
	}
	return true
}

func main() {

	// Use this for part 1
	bufferSize := 4

	// Use this for part 2
	//bufferSize := 14

	filePath := "input.txt"
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	stream := string(file)

	// Fill initial buffer
	var buffer []string
	for i := 0; i < bufferSize; i++ {
		buffer = append(buffer, string(stream[i]))
	}

	// Cut initial buffer off
	stream = stream[bufferSize:]

	// Iterate through stream, enqueue, and check if unique
	for charsReceived, currentChar := range stream {
		rotateQueue(buffer, string(currentChar))
		fmt.Println(strings.Join(buffer, ", "))
		if uniqueItems(buffer) {
			println("Sequence: ", fmt.Sprint(buffer))
			println(fmt.Sprint(charsReceived + bufferSize + 1))
			break
		}
	}
}
