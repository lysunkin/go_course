package main

import (
	"fmt"
	"strconv"
)

func InputAndTruncate(sampleName string) {
	var input string
	fmt.Printf("Input %s number: ", sampleName)
	fmt.Scanln(&input) 
	var converted float64
	converted, _ = strconv.ParseFloat(input, 64)
	fmt.Printf("truncated %s: %d\n", sampleName, int(converted))
}

func main() {
	InputAndTruncate("first")
	InputAndTruncate("second") 
}
