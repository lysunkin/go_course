package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func Merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	// copy the remaining items from the left
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	// copy the remaining items from the right
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}

func input() []int {
	var input string
	fmt.Println("Input integers (new line delimited, any non-number input proceeds further)")
	var slc []int
	for {
		fmt.Scanln(&input)
		var converted int
		converted, err := strconv.Atoi(input)
		if err != nil {
			break
		}
		slc = append(slc, converted)
	}
	return slc
}

func generateInput(size int) []int {
	var slc []int

	for i := 0; i < size; i++ {
		slc = append(slc, rand.Intn(1000))
	}

	return slc
}

func Sort(slc []int) {
	sort.Ints(slc)
}

// SortP - sort function for parallel execution
func SortP(slc []int) {
	Sort(slc)
	wg.Done()
}

func Split(slc []int, parts int) [][]int {
	var arrayofslice [][]int

	// fix for arrays where number of elements in a slice is not equal or array is too small
	size := int(math.Ceil(float64(len(slc)) / float64(parts)))
	for i := 0; i < parts; i++ {
		// fix for arrays where number of elements in a slice is not equal or array is too small
		lowerBound := int(math.Min(float64(i*size), float64(len(slc))))
		upperBound := int(math.Min(float64((i+1)*size), float64(len(slc))))
		part := slc[lowerBound:upperBound]
		arrayofslice = append(arrayofslice, part)
	}

	return arrayofslice
}

// MergeP - merge function for parallel execution
func MergeP(left, right []int, out *[]int) {
	res := Merge(left, right)
	*out = res
	wg.Done()
}

func main() {
	slc := input()
	//slc := generateInput(100) // for random input generation
	fmt.Println("initial: ", slc)

	parts := Split(slc, 4)

	// sort each part of the array separately
	for i, a := range parts {
		fmt.Println("part", i, a)
		wg.Add(1)
		go SortP(a)
	}
	wg.Wait()

	// merge pairs of sorted parts
	var ab []int
	wg.Add(1)
	go MergeP(parts[0], parts[1], &ab)

	var cd []int
	wg.Add(1)
	go MergeP(parts[2], parts[3], &cd)

	wg.Wait()

	// merge the last pair of sorted parts
	result := Merge(ab, cd)

	fmt.Println("sorted: ", result)
}
