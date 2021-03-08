package main

import (
	"fmt"
	"time"
)

var i int

func racer1() {
	for c := 0; c < 100; c++ {
		i = i + c + 1
	}
	fmt.Println(i)
}

func racer2() {
	for c := 0; c < 100; c++ {
		i = i - c
	}
	fmt.Println(i)
}

func main() {
	go racer1()
	go racer2()
	time.Sleep(2 * time.Second)
}
