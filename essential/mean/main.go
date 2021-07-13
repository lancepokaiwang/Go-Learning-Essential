package main

import "fmt"

func main() {
	var x float64
	var y float64

	x = 1
	y = 2

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("x=%v, type of %T\n", y, y)

	mean := (x + y) / 2
	fmt.Printf("Result: %v, type of %T\n", mean, mean)
}
