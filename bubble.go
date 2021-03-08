package main

import (
	"fmt"
	"strconv"
)

// ascending sort order
func Asc(a, b int) bool {
	return a < b	
}

// descending sort order
func Desc(a, b int) bool {
	return a > b	
}

func Swap(sli []int, i int) {
	sli[i+1], sli[i] = sli[i], sli[i+1]
}

func BubbleSort(sli []int) {
	swapped := true
	for swapped {
		swapped = false
		for i :=0; i < len(sli)-1; i++ {
			if !Asc(sli[i], sli[i+1]) {
				Swap(sli, i)
				swapped = true
			}
		}
	}
}

func main() {
	var input string
	fmt.Println("Input up to 10 integers (new line delimited, any non-number prints the sorting result)")
	var slc []int
	maxIntegers := 10
	i := 0
	for {
		if i >= maxIntegers {
			break
		}

		fmt.Scanln(&input) 
		var converted int
		converted, err := strconv.Atoi(input)
		if err != nil {
			break
		}
		slc = append(slc, converted)
		i++
	}
	BubbleSort(slc)
	fmt.Println("sorted: ", slc)
}
