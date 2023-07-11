package main

import "fmt"

func main() {
	hello := "Hello"
	world := "World"
	// fmt.Println(hello, world) // こっちは動く
	fmt.Println(hello) // Worldが使われないためコンパイルエラー
}
