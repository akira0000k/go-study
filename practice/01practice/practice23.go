package main

import (
	"fmt"
)
/*
   subject : サンプルで学ぶ Go 言語：Closures
 */

func intSeq(init int) func() int {
	i := init - 1
	return func() int {
		i++
		return i
	}
}

func intCount(init int) func(bool) int {
	ii := init
	i := init - 1
	return func(ctl bool) int {
		if ctl {
			i = ii - 1
		}
		i++
		return i
	}
}

func main() {
	fmt.Println("start")

	p := intSeq(20)
	q := intSeq(30)
	for i:=0; i<10; i++ {
		fmt.Println(p())
		fmt.Println(q())
	}
	fmt.Println("-------------------")
	pp := intCount(100)
	qq := intCount(300)
	var flag bool = false
	for i:=0; i<10; i++ {
		if i == 5 {
			fmt.Println(">>")
			flag = true
		}
		fmt.Println(pp(flag))
		fmt.Println(qq(flag))
		flag = false
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sat Oct 30 12:53:17
//  
// go run practice23.go
// start
// 20
// 30
// 21
// 31
// 22
// 32
// 23
// 33
// 24
// 34
// 25
// 35
// 26
// 36
// 27
// 37
// 28
// 38
// 29
// 39
// -------------------
// 100
// 300
// 101
// 301
// 102
// 302
// 103
// 303
// 104
// 304
// >>
// 100
// 300
// 101
// 301
// 102
// 302
// 103
// 303
// 104
// 304
//  
// Compilation finished at Sat Oct 30 12:53:17
