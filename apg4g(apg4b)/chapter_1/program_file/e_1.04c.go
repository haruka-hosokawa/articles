package main

import "fmt"

func main() {
	var a, b int = 10, 5 // 明示的型付け
	c, d := 20, 25       // 暗黙的型付け

	fmt.Println(a, b)
	fmt.Println(c, d)

}
