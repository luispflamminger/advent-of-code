package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!\n\nAnd here's a random quote of the day:")
	fmt.Println("> " + quote.Go())
}
