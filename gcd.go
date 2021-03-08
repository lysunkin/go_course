package main

import (
	"fmt"
)

func GCD(a, b int) int {
	if b == 0 {
		return a
	}

	rem := a%b
	return GCD(b, rem)
}

func main() {
	fmt.Println(GCD(357,234))
}
