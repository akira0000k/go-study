package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
/*
   subject : サンプルで学ぶ Go 言語：Line Filters
*/
func main() {
	//バッファの付いていない os.Stdin をバッファ付きのスキャナでラップすると、
	//便利な Scan メソッドを使ってトークンごとに入力を読み進められる。
	//デフォルトのスキャナは Stdin を行ごとに読み出すのだ。
	scanner := bufio.NewScanner(os.Stdin)

	//Text は現在のトークンを返す。 今の場合は、入力から読み出した次の行である。
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		//大文字にした行を書き出す。
		fmt.Println(ucl)
	}

	//Scan でエラーが無かったか確認する。 なお、終端記号（EOF）が見つかっても、Scan はエラーとして扱わない。
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
