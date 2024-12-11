package main

import (
	"fmt"
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

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

var diskMap []int
var fs []int
var fs2 []int
var freeBlocks []int
var freeBlocks2 [][]int

func main() {
	file, _ := os.ReadFile("./input")
	content := string(file)
	content = strings.Split(content, "\n")[0]

	diskMap = make([]int, len(content))
	for i := range diskMap {
		diskMap[i] = unsafe(strconv.Atoi(string(content[i])))
	}

	// disk map to file system
	curId := 0
	for c, i := range diskMap {
		if c%2 == 0 {
			for j := 0; j < i; j++ {
				fs = append(fs, curId)
			}
			curId++
		} else {
			for j := 0; j < i; j++ {
				fs = append(fs, -1)
			}
		}
	}

	// copy file system for part 2
	fs2 = make([]int, len(fs))
	copy(fs2, fs)

	// PART 1

	// calculate free blocks
	for i := 0; i < len(fs); i++ {
		if fs[i] == -1 {
			freeBlocks = append(freeBlocks, i)
		}
	}

	// go backwards through the file system and fill in free blocks
	// break if there are no more free blocks before the counter
	c := 0
	for i := len(fs) - 1; i >= 0; i-- {
		if fs[i] != -1 {
			if !contains(fs[:i], -1) {
				break
			}
			fs[freeBlocks[c]] = fs[i]
			fs[i] = -1
			c++
		}
	}

	// calculate checksum (if we hit -1 we have reached the end of the defragmented block)
	cs := 0
	for i := 0; i < len(fs); i++ {
		if fs[i] == -1 {
			break
		}
		cs += i * fs[i]
	}
	fmt.Println("checksum(1): ", cs)

	// PART 2

	// free blocks should now contain the starting index and length of the block
	// [x][0] is the starting index; [x][1] is the length
	for i := 0; i < len(fs2); i++ {
		if fs2[i] == -1 {
			c := 0
			for i+c < len(fs2) && fs2[i+c] == -1 {
				c++
			}
			freeBlocks2 = append(freeBlocks2, []int{i, c})
			i += c
		}
	}

	// walk backwards through the fs
	c = 0
	for i := len(fs2) - 1; i >= 0; i-- {
		// if we hit a file, save it's id and find it's length (j)
		if fs2[i] != -1 {
			id := fs2[i]
			j := 0
			for i-j >= 0 && fs2[i-j] == id {
				j++
			}

			// find a free block that fits the file
			for _, fb := range freeBlocks2 {
				if fb[1] >= j && fb[0] <= i {
					// if a free block is found, iterate through it's positions (k)
					// and fill them with the file id. Set the original positions
					// empty
					for k := 0; k < j; k++ {
						fs2[fb[0]+k] = id
						fs2[i-k] = -1
					}

					// update free block length and starting index
					fb[1] -= j
					fb[0] += j
					break
				}
			}
			// continue walking backwards before the file that was just moved
			i = i - j + 1
		}
	}

	// calculate checksum
	cs = 0
	for i := 0; i < len(fs2); i++ {
		if fs2[i] == -1 {
			continue
		}
		cs += i * fs2[i]
	}
	fmt.Println("checksum(2): ", cs)
}
