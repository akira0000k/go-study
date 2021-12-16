package main
/*
   subject : Command-Line Flags
   Go の flag パッケージを使ってコマンドラインフラグをパースできる。 実際にコマンドラインプログラムを作って、このパッケージを使ってみよう。
*/
import (
	"flag"
	"fmt"
)
func main() {

	//文字列、整数、真偽値のオプションを受け取るフラグを宣言できる。
	//ここではフラグ word を宣言し、そのデフォルト値を "foo" とし、フラグの短い説明を与えている。
	//関数 flag.String は（文字列そのものではなく）文字列のポインタを返す。 このポインタの使い方は追って説明する。
	

	wordPtr := flag.String("word", "foo", "a string")

	//word と同様のやり方で、フラグ numb、fork を宣言している。
	

	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	//プログラム内の別の場所で宣言した変数を使い、オプションを宣言することもできる。 なお、この場合変数のポインタを渡すことに注意する。
	

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	//すべてのフラグを宣言したら、flag.Parse() を呼んでコマンドラインをパースする。
	

	flag.Parse()

	//パースしたオプションと、余りの引数を表示する。 *wordPtr のように参照を剥がしてオプション値を呼んでいることに注意する。
	

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
// Akira@MBP practice % ./process07
// word: foo
// numb: 42
// fork: false
// svar: bar
// tail: []
// Akira@MBP practice % ./process07 -word=opt
// word: opt
// numb: 42
// fork: false
// svar: bar
// tail: []
// Akira@MBP practice % ./process07 -word=opt a1 a2 a3
// word: opt
// numb: 42
// fork: false
// svar: bar
// tail: [a1 a2 a3]
// Akira@MBP practice % ./process07 -word=opt a1 a2 a3 -numb=7
// word: opt
// numb: 42
// fork: false
// svar: bar
// tail: [a1 a2 a3 -numb=7]
// Akira@MBP practice % ./process07 -word=opt  -numb=7 a1 a2 a3
// word: opt
// numb: 7
// fork: false
// svar: bar
// tail: [a1 a2 a3]
// Akira@MBP practice % ./process07 -h
// Usage of ./process07:
//   -fork
//     	a bool
//   -numb int
//     	an int (default 42)
//   -svar string
//     	a string var (default "bar")
//   -word string
//     	a string (default "foo")
// Akira@MBP practice % ./process07 -wat
// flag provided but not defined: -wat
// Usage of ./process07:
//   -fork
//     	a bool
//   -numb int
//     	an int (default 42)
//   -svar string
//     	a string var (default "bar")
//   -word string
//     	a string (default "foo")
// Akira@MBP practice %
