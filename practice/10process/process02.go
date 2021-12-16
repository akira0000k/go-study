package main
import (
	"os"
	"os/exec"
	"syscall"
)
/*
   subject : Exec'ing Process
*/
func main() {

	//この例では ls を exec する。 実行したいバイナリの絶対パスを Go は必要とするので、 exec.LookPath を使う
	//（おそらく /bin/ls だろう）。
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	//Exec の引数はスライスで表現する（ひとつの大きな文字列ではない）。
	//ここでは ls によく使う引数を渡してみる。 なお、最初の引数はプログラム名であることに注意する。
	args := []string{"ls", "-a", "-l", "-h"}

	//Exec には環境変数も渡す必要がある。 ここでは、現在の環境変数をそのまま渡す。
	env := os.Environ()

	//syscall.Exec を呼ぶ。
	//呼び出しが成功すると、このプロセスはここで終わり、/bin/ls -a -l -h を実行するプロセスに置き換わる。
	//もしエラーがあれば、それを表す値が返ってくる。
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
	//ここへは来ない
}
// -*- mode: compilation; default-directory: "~/go/src/practice/10process/" -*-
// Compilation started at Fri Oct 15 12:43:54
//  
// go run process02.go
// total 24
// drwxr-xr-x   5 Akira  staff   160B 10 15 12:43 .
// drwxr-xr-x  15 Akira  staff   480B 10 15 11:41 ..
// -rw-r--r--   1 Akira  staff   2.7K 10 15 12:39 process01.go
// -rw-r--r--   1 Akira  staff   2.1K 10 15 11:45 process01.go~
// -rw-r--r--   1 Akira  staff   1.1K 10 15 12:43 process02.go
//  
// Compilation finished at Fri Oct 15 12:43:55
