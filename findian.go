package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Input a string: ")
	fmt.Scanln(&input) 

	if strings.Contains(input, "i") &&
		strings.Contains(input, "a") &&
		strings.Contains(input, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
