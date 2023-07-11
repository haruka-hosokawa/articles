# 本記事について

競技プログラミングサイトの１つである[AtCoder](https://atcoder.jp/home)では、プログラミング学習教材の[APG4b(AtCoder Programming Guide for beginners)](https://atcoder.jp/contests/APG4b)が公開されています。使用言語はC++です。

また、sabaさんによって、[Python版APG4b](https://qiita.com/saba/items/b9418d7b54cce4b106e4)も書かれています。さらに、その他多くのプログラミング言語を使ったAPG4bの記事が公開されています。

本記事では以上のコンテンツに則り、Go言語版APG4bを書いてみたいと思います。

今回扱う内容は以下の通りです。

- [1.00. はじめに](#100-はじめに)
- [1.01. 出力とコメント](#101-出力とコメント)
- [1.02. プログラムの書き方とエラー](#102-プログラムの書き方とエラー)
- [1.03. 四則演算と優先順位](#103-四則演算と優先順位)
- [1.04. 変数と型](#104-変数と型)
- [1.05. プログラムの実行順序と入力](#105-プログラムの実行順序と入力)

## [1.00. はじめに](https://atcoder.jp/contests/apg4b/tasks/APG4b_a)

はじめに、AtCoderにコードを提出してみます。

### 手順

1. [A - 1.00.はじめに](https://atcoder.jp/contests/apg4b/tasks/APG4b_a)のページを開く
1. 最下部の「言語」欄を「Go(1.14.1)」にする
1. 以下のコードを「ソースコード」欄にコピペし、提出ボタンを押す
1. 「AC」となることを確認する

```golang: a_1.00.go
package main

import "fmt"

func main() {
 fmt.Println("Hello, world!")
}
```

AC（Accepted）と表示されれば、問題に対して正しく答えることが出来たことになります。おめでとうございます。

WA(Wrong Answer)やRE(Runtime Error)、CE(Compile Error)の時は、

- 表示する文字列の```Hello, world!```が間違ってないか（大文字/小文字/空白など）
- 括弧が正しく閉じているか
- 言語設定が誤っていないか（PythonやC++を選択していないか）

などを確認してください。

AtCoderには提出前に、書いたコードがAtCoder上で動くかどうか確認することができる[コードテスト](https://atcoder.jp/contests/apg4b/custom_test)ページがあります。提出する前にコードテストで動かしてみることも検討してみてください。

## [1.01. 出力とコメント](https://atcoder.jp/contests/apg4b/tasks/APG4b_b)

### キーポイント

- `pakcage main`と`import "fmt"`は毎回最初に書く
- Go言語のプログラムは**main関数**から始まる
- `fmt.Println("文字列")`で文字列を出力できる
- `//`や`/* */`でコメントを書ける

| 書き方  | コメントになる場所             |
| ------- | ------------------------------ |
| `//`    | 同じ行の`//`と書いた場所の後ろ |
| `/* */` | `/*`と`*/`の間全て             |

Go言語のプログラムの基本形は次のプログラムになります。
`go run b_1.01.go`とコマンドを実行しても、このプログラムは何もしません。

```golang: b_1.01a.go
package main

func main() {
 // この部分が実行される
}
```

### main関数

Go言語のプログラムは`main関数`から始まります。上の`b_1.01.go`では以下の部分がmain関数です。

```golang: main関数
func main() {
 // この部分が実行される
}
```

Go言語のプログラムを起動すると、main関数の`{}`の中にある処理が動きます。

最初のプログラムを見てみましょう。
このプログラムは、`Hello, world!`という文字列を出力するプログラムです。

```golang: a_1.00.go
package main

import "fmt"

func main() {
 fmt.Println("Hello, world!")
}
```

```text: 実行結果
Hello, world!
```

main関数の中にある次の一行が、画面に文字列を表示させる（出力する）プログラムです。

```golang: 出力を行う部分
fmt.Println("Hello, world!")
```

Go言語では、画面に文字列などを出力する時に```fmt.Println()```を使います。`()`の中に出力させたいものを入れることで、文字列や数値などを出力させることが出来ます。

Go言語で文字列を扱う場合、`"`（ダブルクオーテーション）で囲む必要があります。
```fmt.Println()```は自動で文末で改行してくれます。

### 行の終わり

C++では行末に`;`が必要ですが、Go言語では必要ありません。Pythonと同様です。

※ 書いても問題なく動きますが、VS CodeなどでGo言語拡張機能を導入している場合やgofmtなどのコード整形ツールを使っている場合は改行に置き換えられます。

### 複数の出力

[a_1.00.go](#手順)のプログラムを使って、別の文字列を出力したい場合は`" "`の中身を変えれば良いです。

```golang: b_1.01b.go
package main

import "fmt"

func main() {
 fmt.Println("こんにちは世界")
}
```

```text: 実行結果
こんにちは世界
```

### 複数の出力

出力を複数回行うことも出来ます。
また、`()`内に`,`区切りで複数の文字列を書くことで、`c d`のように空白区切りで出力させることが出来ます。

```go: b_1.01c.go
package main

import "fmt"

func main() {
 fmt.Println("a")
 fmt.Println("b")
 fmt.Println("c", "d")
}
```

```text: 実行結果
a
b
c d
```

また、改行が必要無い場合は以下のように書くことも出来ます。

```go: b_1.01d.go
package main

import "fmt"

func main() {
 fmt.Print("a")
 fmt.Print("b")
 fmt.Print("c", "d")
}
```

`fmt.Print()`とすることで、改行せずに文字列を出力させることが出来ます。

```text: 実行結果
abcd
```

### 数値の出力

数値を出力する時は、`" "`を使わずに、直接書くことでも出力できます。

```go: b_1.01e.go
package main

import "fmt"

func main() {
 fmt.Println(2525)
}
```

```text: 実行結果
2525
```

### 全角文字

Go言語はUTF-8という文字コードを使ってプログラムが書かれることを前提としています。
UTF-8は日本語等の全角文字をサポートしているため、全角文字であってもそのまま使うことが出来ます。

しかし、多くのプログラムは半角英数字と記号で書かれており、全角文字で書くことはスタンダードではないため、半角文字で書くべきだと思います。

```go: b_1.01f.go
package main

import "fmt"

func main() {
 こんにちは世界 := "Hello, world."
 fmt.Println(こんにちは世界)
}
```

```text: 実行結果
Hello, world.
```

### コメント

コメントは人間がプログラムの動作を説明するためのメモ書きを残すための機能です。
プログラムとしての機能はありません。

Go言語でのコメントの例は以下のようになります。

```go: b_1.01g.go
package main

import "fmt"

func main() {
 fmt.Println("Hello")
 // fmt.Println("World")

 /*
  ここも
  全部
  コメントになります！

  空白行があっても大丈夫！
 */
}
```

```text: 実行結果
Hello
```

コメントには２種類の書き方があります。

| 書き方  | コメントになる場所             |
| ------- | ------------------------------ |
| `//`    | 同じ行の`//`と書いた場所の後ろ |
| `/* */` | `/*`と`*/`の間全て             |

### 練習問題

[EX1 - コードテストと出力の練習](https://atcoder.jp/contests/apg4b/tasks/APG4b_cv)を解いてみましょう。

<details><summary>解答（解けなかったら開いてください）</summary>

```go: 解答
package main

import "fmt"

func main() {
 fmt.Println("こんにちは")
 fmt.Println("AtCoder")
}
```

</details>

## [1.02. プログラムの書き方とエラー](https://atcoder.jp/contests/apg4b/tasks/APG4b_c)

### キーポイント

- スペース、改行、インデントを使ってプログラムを見やすくする
- Goは公式からプログラムの体裁を整えるツール（`gopls`や`gofumpt`が提供されている
- コンパイルエラーは文法のエラーで、プログラムは実行されない
- 実行時エラーは内容のエラーで、プログラムは強制終了される
- 論理エラーは内容のエラーで、プログラムは正しく動いているように見えてしまう

### スペースと改行

Go言語では、スペースと改行は同じ意味になります。

また、C++と同様に、Go言語はセミコロンで行末であることを明示することができ、
一つの行に１つの処理しか書かれていない場合、セミコロンを省略することが出来ます。

以下の２つのコードは同じ動きになります。

```go: c_1.02a.go
package main

import "fmt"

func main() {
 fmt.Println("a")
 fmt.Println("b")
 fmt.Println("c")
 fmt.Println("d")
}
```

```go : コード整形をしない例
package main; import "fmt"; func main() { fmt.Println("a"); fmt.Println("b"); fmt.Println("c"); fmt.Println("d"); }
```

```text: 実行結果
a
b
c
d
```

### Go言語のコード整形ツール

Go言語は公式が提供する`gopls`と`gofumpt`の２つの書式でコードをフォーマットしてくれます。
字下げ（インデント）や括弧などの記号の空白調整を行ってくれますし、何より可読性が向上するため有効にしておいて損はありません。

VS Codeでは、Go言語の公式から配布されている拡張機能をインストールすることで、セーブ時にコードを整形してくれます。

参考資料

- <https://code.visualstudio.com/docs/languages/go#_formatting>
- <https://marketplace.visualstudio.com/items?itemName=golang.Go>

### プログラムのエラー

Go言語はC++と同じコンパイル言語です。
そのため、予めプログラムを機械語に翻訳し、実行ファイルを作るコンパイル処理が必要です。

コンパイラがコードを読み解けなかった場合、コンパイルエラー（文法エラー）を起こします。

その他エラーもありますが、大まかな理解は以下の通りです。

- コンパイルエラー
  - 書いたプログラムの文法が間違えている
- 実行時エラー
  - 文法は間違えていなかったが、動かしてみたら致命的な間違いがあった
- 論理エラー
  - 文法も処理も間違えていないが、書いた人間が勘違いかタイプミスをしている

焦らずゆっくり直しましょう。

### コンパイルエラーの例

全角スペースをコード上に紛れ込ませるとこのようなエラーが出ます。

```go: 全角スペースが紛れ込んだ例
package main

import "fmt"

func main() {　// ← "{" の次に全角スペースが入っている
 fmt.Println("Hello, world!")
}

```

エラーログはこのようになります。
「`c_1.02b.go`の5行目14文字目が`\U+3000（全角スペース）`でGoでは読めない文字ですよ。」と言っています。消去すれば問題なく動きます。

※2行目はUndefined（定義されていない）と書かれているため、変数として認識されているのだと思います。

```shell: エラーログ
# command-line-arguments
./c_1.02b.go:5:14: invalid character U+3000 in identifier
./c_1.02b.go:5:14: undefined: 　
```

Go言語は使わないデータに対して敏感なので、使用しない変数が存在した場合、コンパイルエラーを起こします。

```go: 使用しない変数を入れている場合
package main

import "fmt"

func main() {
 hello := "Hello"
 world := "World"
 // fmt.Println(hello, world) // こっちは動く
 fmt.Println(hello) // Worldが使われないためコンパイルエラー
}
```

```text: エラーログ
./c_1.02c.go:7:2: world declared and not used
```

### 練習問題
  
  [EX2 - エラーの修正](https://atcoder.jp/contests/apg4b/tasks/APG4b_cu)を解いてみましょう。

```go: A君が書いたプログラム(ex2_error.go)
package main

import "fmt"

func main() {
 // \nは改行するための文字（エスケープシーケンス）です。
 fmt.Print("いつも, 2525 "\n") 
 fmt.Println("AtCoderくん"
}
```

```text: 標準エラー出力
./ex2_error.go:7:30: invalid character U+005C '\'
./ex2_error.go:7:31: syntax error: unexpected n in argument list; possibly missing comma or )
./ex2_error.go:7:35: newline in string
```

<details><summary>解答（解けなかったら開いてください）</summary>

```go: 修正したコード
package main

import "fmt"

func main() {
 // \nは改行するための文字（エスケープシーケンス）です。
 fmt.Print("いつも", 2525, "\n")
 fmt.Println("AtCoderくん")
}
```

</details>

## [1.03. 四則演算と優先順位](https://atcoder.jp/contests/apg4b/tasks/APG4b_d)

### キーポイント

| 演算子 | 計算内容     |
| ------ | ------------ |
| `+`    | 足し算       |
| `-`    | 引き算       |
| `*`    | かけ算       |
| `/`    | 割り算       |
| `%`    | 割ったあまり |

### 四則演算

Go言語のプログラムで簡単な計算をする方法を見ていきましょう。

次のプログラムは上から順に「足し算」「引き算」「掛け算」「割り算」を行っています。

```go: d_1.03a.go
package main

import "fmt"

func main() {
 fmt.Println(1 + 1) // 2　（足し算）

 fmt.Println(3 - 4) // -1　（引き算）

 fmt.Println(2 * 3) // 6　（かけ算）

 fmt.Println(7 / 3)     // 2　（割り算）
 fmt.Println(-7 / 3)    // -2　（負の値の割り算）
 fmt.Println(7.0 / 3) // 小数の割り算
}
```

```text:実行結果
2
-1
6
2
-2
2.3333333333333335
```

これらの記号（`+`,`-`,`*`,`/`）のことを**算術演算子**といいます。`+`,`-`,`*`,`/`は順に、「足し算」「引き算」「かけ算」「割り算」に使う演算子（記号）です。

引き算の例のように、（`+`,`-`,`*`,`/`）を使って負の値も計算することが出来ます。

Go言語では、整数同士で割り算（`/`）をした場合、小数点以下を切り捨てた値になります。

負の値の割り算も小数点が切り捨てられます。
切り捨ててほしくない場合、`7.0`というように`.0`をつけると、小数点以下も計算されます。

### 剰余演算

`%`を使うことで、「割った時のあまり」を計算することが出来ます。Go言語は負の値のあまりも計算することが出来ます。

```go: d_1.03b.go
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
```

```text: 実行結果
-1
0
-2
-1
0
1
2
0
1
```

### 演算子の優先順位

C++及びPythonの優先順位と変わりません。

- 最初に、かけ算、割り算（`*`,`/`）が計算される
- 次に、足し算、引き算（`+`,`-`）が計算される
- 括弧`()`で括られた計算があるなら、括弧の中の計算が優先される

### ゼロ除算（細かい話）

**Go言語ではゼロで割っても実行時エラーになりません。**
※ただし、数値を直接指定してゼロ除算をしている場合はコンパイルエラーを起こします。

```go: d_1.03c.go
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
```

```text:実行結果
0 +Inf -Inf NaN
```

このように、0や正/負の無限大、NaN(Not a Number)として計算されます。

参考資料

- <https://en.wikipedia.org/wiki/NaN>
- <https://go.dev/ref/spec#Representability>

### 練習問題

[EX3 - 計算問題](https://atcoder.jp/contests/apg4b/tasks/APG4b_ct)を解いてみましょう。

<details><summary>解答（解けなかったら開いてください）</summary>

```go: 解答
package main

import "fmt"

func main() {
 fmt.Println(100 * (100 + 1) / 2)
}
```

#### 補足

なぜ、

```math
\frac{1}{2} \times 100 \times (100 + 1)
```

をそのまま計算すると不正解となってしまうのか？を説明します。

結論から言うと、`1 / 2`は整数の割り算の切り捨て処理によって、`0`となってしまうからです。

実は、`1.0 / 2.0`とすることで回避できます。（`0.5`になる）

ですが、**割り算は最後に計算する**というプログラムを書いたほうが、数値計算として誤差が出にくくなるため、割り算は最後に書いたほうが良いと思います。

```go: d_1.03d.go
package main

import "fmt"

func main() {
 fmt.Println(1 / 2)
 fmt.Println(1.0 / 2.0 * 100 * (100 + 1))
}
```

```text: 実行結果
0
5050
```

</details>

## [1.04. 変数と型](https://atcoder.jp/contests/apg4b/tasks/APG4b_e)

### キーポイント

- 変数は「メモ」
- `=`は代入
- 「データの種類」を「型」と言う

| 型      | 書き込むデータの種類 |
| ------- | -------------------- |
| int     | 整数                 |
| float64 | 小数                 |
| string  | 文字列               |

### 変数

変数という機能について見ていきましょう。変数は「メモ」や「ふせん」というイメージがいいかもしれません。

メモはこのようなイメージです。詳しくは後述します。

- 変数名 : メモのタイトル
- 型 : メモに書いてあるデータの種類と紙の大きさ
- 値 : メモの内容

例を見てみましょう。

```go: e_1.04a.go
package main

import "fmt"

func main() {
 var variable_1 int = 10
 var variable_1a int
 variable_2 := 20

 fmt.Println(variable_1, variable_1+10)
 fmt.Println(variable_1a, variable_1a+10)
 fmt.Println(variable_2, variable_2+10)

}
```

```text:実行結果
10 20
0 10
20 30
```

#### 宣言と代入

変数を使うには最初に「宣言」を行う必要があります。
変数を宣言する方法には2種類ありますが、どちらを使っても構いません。

変数を宣言する時は、変数の名前を指定します。

```go :変数の宣言をする例
var variable_1 int = 10 // 明示的型付け
var variable_1a int // 明示的型付け その２
variable_2 := 20 // 暗黙的型付け
```

1行目の明示的型付けで宣言を行う場合、最初に`var`を付けます。
その後に変数名（`variable_1`）を書き、最後に型(`int`)を付けます。

2行目の暗黙的型付けの場合は変数名（`valiable_2`）を書き、`:=`を書いて、その後にメモの内容となる値（データ）を書きます。

※暗黙的型付けの場合、明示的型付けで行うような型の指定はコンパイル時に自動的に解釈されます。

宣言した変数にデータをメモする時は、`=`を使って代入をします。
また、宣言時に値を同時に代入することを初期化と言います。

明示的型付けの場合は初期化をする必要がありませんが、暗黙的型付けの場合は`:=`を使った初期化が必要です。

#### 変数のコピー

`変数1 = 変数2`というような形で書いた場合、変数の値そのものがコピーされます。

その後に変数の値を変えても、もう片方の変数に影響はありません。

```go: e_1.04b.go
package main

import "fmt"

func main() {
 var a int = 10
 var b = a
 a = 40
 fmt.Println(a, b)
}
```

```text:実行結果
40 10
```

#### 変数を同時に宣言

変数の宣言時に`,`を入れることで、複数の変数を同時に宣言することも出来ます。

```go: e_1.04c.go
package main

import "fmt"

func main() {
 var a, b int = 10, 5 // 明示的型付け
 c, d := 20, 25       // 暗黙的型付け

 fmt.Println(a, b)
 fmt.Println(c, d)

}
```

```text:実行結果
10 5
20 25
```

さらに、明示的宣言の場合は`()`で囲んで改行区切りで書くことで、１つの`var`で複数行に渡って変数を宣言することが出来ます。

```go: e_1.04d.go
package main

import "fmt"

func main() {
 var (
  a int
  b int
  c int = 5
  d int = 10
 )
 fmt.Println(a, b, c, d)
}
```

```text: 実行結果
0 0 5 10
```

### 変数のルール

変数は基本的に自由につけることが出来ますが、一部の名前を変数名にするとコンパイルエラーを起こします。

#### 利用できない変数名

以下の条件に該当する変数名は変数名に出来ません。

- 数字で始まる名前
- `_`以外の記号が使われている名前
- キーワード（Go言語が使っている一部の単語）

例えば、以下のコードはエラーになります。

```go: e_1.04e.go
package main

import "fmt"

func main() {
 var 100hello int
 var na+me
 var go int
 fmt.Println(100hello, na+me, go)
}
```

```text: エラーログ
./e_1.04e.go:6:6: syntax error: unexpected literal 100, expected name
./e_1.04e.go:7:8: syntax error: unexpected +, expected type
./e_1.04e.go:8:6: syntax error: unexpected go, expected name
./e_1.04e.go:9:17: syntax error: unexpected hello in argument list; possibly missing comma or )
```

以下のような変数名は大丈夫です。

```go: 問題ない変数名
var hello10 int
var hello_world int
```

#### 同じ変数名の変数は宣言できない

同じ名前の変数を複数回宣言することは出来ません。

```go: e_1.04f.go
package main

import "fmt"

func main() {
 var a = 0
 var a = 5
 fmt.Println(a)
}
```

```text: エラーログ
./e_1.04f.go:7:6: a redeclared in this block
        ./e_1.04f.go:6:6: other declaration of a
```

### 型

`int`以外にもGoには様々な型があります。
ここでは重要な３つの型を紹介します。

| 型      | 書き込むデータの種類 |
| ------- | -------------------- |
| int     | 整数                 |
| float64 | 小数                 |
| string  | 文字列               |

```go : e_1.04g.go
package main

import "fmt"

func main() {
 i := 10
 f := 0.5
 s := "Hello"

 fmt.Println(i, f, s)
}
```

```text: 実行結果
10 0.5 Hello
```

### 異なる型同士の計算

Go言語では異なる型同士の計算が出来ません。
異なる型同士の計算をしようとすると、コンパイルエラーとなります。

```go: e_1.04h.go
package main

import "fmt"

func main() {
 a := 1
 b := 1.5
 fmt.Println(a * b)
}
```

```text:エラーログ
./e_1.04h.go:9:14: invalid operation: a * b (mismatched types int and float64)
```

### 異なる型同士の代入

Go言語では異なる型同士の代入も出来ません。
[異なる型同士の計算](#異なる型同士の計算)と同様にコンパイルエラーとなります。

```go: e_1.04h.go
package main

import "fmt"

func main() {
 i := 1
 var d float64 = i
 fmt.Println(i, d)
}
```

```text:エラーログ
./e_1.04i.go:7:18: cannot use i (variable of type int) as float64 value in variable declaration
```

#### 細かい話

<details><summary>変数の細かい話です</summary>

Go言語では`int`はキーワードではないので、
`var int int`という変数宣言が可能です。
（コードを読む際に混乱することになるので、書くべきではありません）

```go:e_1.04e_ex.go
package main

import "fmt"

func main() {
 var int int // intの初期値は0
 fmt.Println(int)
}
```

```text:実行結果
0
```

Go言語のキーワードについては、[Keywords - The Go Programming Language Specification](https://go.dev/ref/spec#Keywords)を御覧ください。

</details>

### 練習問題

[EX4 - ◯年は何秒？](https://atcoder.jp/contests/apg4b/tasks/APG4b_cs)を解いてみましょう。

```go: コピーするプログラム
package main

import "fmt"

func main() {
 seconds := 365 * 24 * 60 * 60

 // 文字列の部分を消して追記する
 fmt.Println("1年は何秒か")
 fmt.Println("2年は何秒か")
 fmt.Println("5年は何秒か")
 fmt.Println("10年は何秒か")
}
```

<details><summary>解答（解けなかったら開いてください）</summary>

```go: 解答
package main

import "fmt"

func main() {
 seconds := 365 * 24 * 60 * 60

 // 文字列の部分を消して追記する
 fmt.Println(seconds)
 fmt.Println(seconds * 2)
 fmt.Println(seconds * 5)
 fmt.Println(seconds * 10)
}
```

</details>

## [1.05. プログラムの実行順序と入力](https://atcoder.jp/contests/apg4b/tasks/APG4b_f)

### キーポイント

- プログラムは上から下へ順番に実行される
- `bufio.Scanner`で入力を受け取ることが出来る

### プログラムの実行順序

基本的に、プログラムは上から順に実行されます。

```go: f_1.05a.go
package main

import "fmt"

func main() {
 var name int

 name = 10

 fmt.Println(name)

 name = 5

 fmt.Println(name)

}
```

```text: 実行結果
10
5
```

### 入力

プログラムの実行時にデータの入力を受け取る方法を見ていきましょう。
「入力機能」を使うことにより、プログラムを書き換えなくても様々な計算を行えるようになります。

入力を受け取るには、`bufio.Scanner`を使います。

次のプログラムは、入力として受け取った数値を10倍、20倍にして出力するプログラムです。

```go: f_1.05b.go
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
}
```

`bufio.NewScanner(os.Stdin)`で、標準入力(`os.Stdin`)から入力を受け取るスキャナ（読み取り機）を作成します。

今回は`input`という変数にスキャナを初期化して代入しています。

`Scan()`は一行の入力を受け取ります。
さらに、`Text()`とすることで、受け取った入力を文字列として取得することができます。

int型として受け取るためには`strconv.Atoi()`、もしくは`strconv.ParseInt()`を使ってint型に変換してあげる必要があります。

`strconv.ParseInt()`を使う場合、`Text()`の後ろにカンマ区切りで`0`を2回入れることで`int`型で入力を受け取ることが出来ます。

```text:入力
5
```

```text:実行結果
50
100
```

#### 整数以外のデータの入力

整数以外のデータを受け取りたい時は、`strconv`を使って必要な方に変換する必要があります。

実数(浮動小数)を受け取る時は、`strconv.ParseFloat`を使う必要があります。`strconv.ParseInt`の時の様に、後ろにカンマ区切りで一回だけ`64`と書いてあげてください。

```go:f_1.05.go
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
```

```text :入力
3.14
```

```text :実行結果
3.14 6.28
```

#### 空白区切りの入力

`bufio.Scanner`の`Scan()`は一行の入力を受け取り、`Text()`で文字列を取得できます。

```go: f_1.05d.go
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
```

```text :入力
10 20 30
```

```text :実行結果
10 20 30 string
```

そのため、空白区切りの入力を受け取ると、上の実行結果のように、`"10 20 30"`という文字列として代入されてしまいます。

#### 　文字列の分割

文字列を分解する場合、`strings.Split`を使って文字列を分解してあげる必要があります。

`strings.Split`を使う場合、`()`の中は`(分解される文字列, 区切り文字)`という順番で書いてください。

また、分解した結果(今回は`list`という変数)は`変数名[前から数えてn番目]`という形式で取得することが出来ます。

`前から数えてn番目`というのは`0`番目、`1`番目、...と0から始まる連続した整数です。※人間は1から数えがちですが、コンピュータでは0から数えることが一般的です。

```go: f_1.05e.go
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
```

```text :入力
10 20 30
```

```text :実行結果
10 string
20 string
30 string
```

こうすることで、`"10 20 30"`が空白区切りで分割されて、
`10`(空白区切りでの`0`番目)、
`20`(空白区切りでの`1`番目)、
`30`(空白区切りでの`2`番目)の３つにすることが出来ます。

#### 複数行の入力

`bufio.Scanner`の`Scan()`は１行の入力を受け取るものなので、
入力が複数行ある場合はその行数分書いてあげる必要があります。

```go: f_1.05g.go
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
```

```text:入力
Hello
10
```

```text:実行結果
Hello 10
```

### 細かい話

<details><summary>Goの標準入出力の細かい話です</summary>

`fmt.Scan()`でも入力を受け取ることが出来ますが、Goで競技プログラミングをする上で安易に使うとTLE（Time Limit Exceeded : 実行制限時間の超過）になります。

以下のプログラムは、[複数行の入力](#複数行の入力)のプログラムを`fmt.Scan()`を使って実装した例です。

```go:f_1.05_ex.go
package main

import (
 "fmt"
)

func main() {
 var (
  s1 string
  s2 string
 )

 fmt.Scan(&s1, &s2)
 fmt.Println(s1, s2)
}
```

実は、`fmt.Scan()`は簡潔に書くことができ便利な半面、競技プログラミングにおいて致命的になりうる**時間がかかるというデメリット**があります。そのため、文字列から色々変換しなくてはなりませんが、`bufio.Scanner`を使った方法だけを紹介しました。

自分の中では、`fmt.Scan()`は「入力が１行の時のみに使う」のが良いと考えており、**複数行の入力を受け取らなければならない場合、`bufio`を使って入力を受け取る**べきだと考えています。

詳細は、ren510rev(Ren Goto)さんの記事[【競プロ】ScanningでTLEが発生してコケる](https://qiita.com/ren510dev/items/38fe6d09831d08fde537)を御覧ください。
</details>

### 練習問題

[EX5 - A足すB問題](https://atcoder.jp/contests/apg4b/tasks/APG4b_cr)を解いてみましょう。

```go: コピーするプログラム
package main

func main() {
 // ここにプログラムを追記
}
```

<details><summary>解答（解けなかったら開いてください）</summary>

```go: 解答
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
```

</details>

## 第一章 (1/3)　おわりに

第一章の内容の1/3は以上です。以降の内容は別途投稿する予定です。
より良い書き方がありましたら適宜ご指摘お願いします。

Pythonを普段書いている身で、Goを少し勉強してこの記事を書いてみました。
Goの標準入出力って中々疲れますね（Javaより少し面倒なレベル）

蛇足 : APG4gはAtCoder Programming Guide for Gopherの略です。bを180°回転させたらgに見えませんか？

## 参考資料

- AtCoder / APG4b系
  - [APG4b](https://atcoder.jp/contests/APG4b)
  - [Python版 APG4b](https://qiita.com/saba/items/b9418d7b54cce4b106e4)
  - [AtCoder コードテスト](https://atcoder.jp/contests/apg4b/custom_test)
- Go言語ツール系
  - [Go for Visual Studio Code - VS Code Go言語拡張機能リンク](https://marketplace.visualstudio.com/items?itemName=golang.Go)
  - [Go in Visual Studio Code - VS Code Go言語拡張機能 ドキュメント](https://code.visualstudio.com/docs/languages/go#_formatting)
- Go言語の機能系
  - [Keywords - The Go Programming Language Specification](https://go.dev/ref/spec#Keywords)
  - [Golangでの文字列・数値変換 - 小野マトペの納豆ペペロンチーノ日記](https://matope.hatenablog.com/entry/2014/04/22/101127)
  - [Integer literals - The Go Programming Language Specification](https://go.dev/ref/spec#Integer_literals)
  - [【競プロ】ScanningでTLEが発生してコケる](https://qiita.com/ren510dev/items/38fe6d09831d08fde537)
  - [Representability - The Go Programming Language Specification](https://go.dev/ref/spec#Representability)
- その他
  - [NaN - Wikipedia](https://ja.wikipedia.org/wiki/NaN)
