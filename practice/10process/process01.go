package main
import (
	"fmt"
	"io/ioutil"
	"os/exec"
)
/*
   subject : Spawning Process
*/
func main() {

	//まずは引数も入力も取らず、標準出力に何かを書くだけの単純なコマンドから始めよう。
	//ヘルパーコマンド exec.Command は外部プロセスを表すオブジェクトを作る。
	dateCmd := exec.Command("date")

	//他に .Output というヘルパーもあり、これはコマンドを実行し、その終了を待ちながら、出力を集める。
	//エラーが起きなければ、dateOut は日時情報を表すバイト列を保持する。
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	
	//続いて、パイプを使って外部プロセスの標準入力にデータを渡し、標準出力から結果を収集する方法を見てみる。
	grepCmd := exec.Command("grep", "hello")

	//ここでは明示的に入力・出力パイプを得て、プロセスを開始し、入力を書き込み、結果を読み出し、プロセスが終了するのを待っている。
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	//上の例ではエラーチェックを省いたが、いつものように if err != nil のパターンでエラーチェックをしてもよい。
	//また、StdoutPipe だけを見ていたが、StderrPipe も全く同じやり方で読み出せる。
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	
	//コマンドを spawn するときは、コマンドラインに打ち込む文字列を丸ごと渡すのではなく、
	//コマンドとその引数の配列を明示的に並べなければならない。
	//もしコマンド全体を単に文字列で書きたいなら、bash の -c オプションを使えばよい。
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
// -*- mode: compilation; default-directory: "~/go/src/practice/10process/" -*-
// Compilation started at Fri Oct 15 11:46:01
//  
// go run process01.go
// > date
// 2021年 10月15日 金曜日 11時46分03秒 JST
//  
// > grep hello
// hello grep
//  
// > ls -a -l -h
// total 16
// drwxr-xr-x   4 Akira  staff   128B 10 15 11:46 .
// drwxr-xr-x  15 Akira  staff   480B 10 15 11:41 ..
// -rw-r--r--   1 Akira  staff   2.1K 10 15 11:46 process01.go
// -rw-r--r--   1 Akira  staff   2.1K 10 15 11:45 process01.go~
//  
//  
// Compilation finished at Fri Oct 15 11:46:03
