package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 標準入力から受け取るスキャナを作成
	input := bufio.NewScanner(os.Stdin)

	// 1行読み込む
	input.Scan()

	// 文字列として受け取る
	text := input.Text()

	// 実数（小数付きの数値）として受け取る
	real, _ := strconv.ParseFloat(input.Text(), 64)

	fmt.Println(text, real*2)
}
