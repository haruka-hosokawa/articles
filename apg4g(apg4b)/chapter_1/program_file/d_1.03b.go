package main

import "fmt"

func main() {
	fmt.Println(-4 % 3) // -1
	fmt.Println(-3 % 3) // 0
	fmt.Println(-2 % 3) // -2
	fmt.Println(-1 % 3) // -1

	fmt.Println(0 % 3) // 0

	fmt.Println(1 % 3) // 1
	fmt.Println(2 % 3) // 2
	fmt.Println(3 % 3) // 0
	fmt.Println(4 % 3) // 1
}
