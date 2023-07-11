package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	list := strings.Split(input.Text(), " ")
	a, b, c := list[0], list[1], list[2]

	fmt.Println(a, reflect.TypeOf(a))
	fmt.Println(b, reflect.TypeOf(b))
	fmt.Println(c, reflect.TypeOf(c))
}
