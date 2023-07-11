package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	list := strings.Split(input.Text(), " ")
	a, _ := strconv.ParseInt(list[0], 0, 0)
	b, _ := strconv.ParseInt(list[1], 0, 0)

	fmt.Println(a + b)
}
