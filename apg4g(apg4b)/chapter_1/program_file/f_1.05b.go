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

	// aという変数に文字列を数値に変換した結果を保存（初期化）
	a, _ := strconv.Atoi(input.Text())
	fmt.Println(a * 10)

	// bという変数に文字列を数値に変換した結果を保存（初期化）
	b, _ := strconv.ParseInt(input.Text(), 0, 0)
	fmt.Println(b * 20)
}
