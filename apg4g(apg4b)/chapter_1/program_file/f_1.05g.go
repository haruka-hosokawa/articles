package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	input.Scan()
	s1 := input.Text()

	input.Scan()
	s2 := input.Text()

	fmt.Println(s1, s2)

}
