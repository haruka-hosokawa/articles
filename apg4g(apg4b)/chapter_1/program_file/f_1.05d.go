package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	a := input.Text()

	fmt.Println(a, reflect.TypeOf(a))
}
