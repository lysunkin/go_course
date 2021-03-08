package main

import "fmt"

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return 1.0/2.0*a*t*t+v*t+s
	}
}

func main() {

	var a float64
	fmt.Println("Enter acceleration:")
	fmt.Scan(&a)
	var v float64
	fmt.Println("Enter velocity:")
	fmt.Scan(&v)
	var s float64
	fmt.Println("Enter displacement:")
	fmt.Scan(&s)

	fn := GenDisplaceFn(a, v, s)

	var t float64
	fmt.Println("Enter time:")
	fmt.Scan(&t)

	fmt.Println(fn(t))
}
