package main
import "fmt"
/*
   subject : 初心者に送りたいinterfaceの使い方[Golang] 1. 何でもはいる型としてのinterface
*/
func main() {
	var i interface{}
	i = 4
	fmt.Println(i)
	i = 4.5
	fmt.Println(i)
	i = "文字列だってはいるんだ"
	fmt.Println(i)
}
// -*- mode: compilation; default-directory: "~/go/src/practice/13interface/" -*-
// Compilation started at Wed Oct 20 20:38:52
//  
// go run inter01.go
// 4
// 4.5
// 文字列だってはいるんだ
//  
// Compilation finished at Wed Oct 20 20:38:55
