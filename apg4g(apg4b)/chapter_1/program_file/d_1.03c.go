package main

import "fmt"

func main() {
	zero := 0.0
	pinf := 1.0 / zero
	ninf := -1.0 / zero
	nan := zero / zero

	fmt.Println(zero, pinf, ninf, nan)
	// fmt.Println(1.0 / 0.0) // ./d_1.03c.go:12:20: invalid operation: division by zero
}
