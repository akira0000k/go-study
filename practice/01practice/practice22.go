package main

import (
	"fmt"
)
/*
   subject : rune type, range
 */

func main() {
	fmt.Println("start")
	var c rune
	
	s := "event last"
	for i:=0; i<2; i++ {
		fmt.Printf("type=%T, %s\n", s, s)
		for _, c = range s {
			fmt.Printf("type=%T, %c, %v %x\n", c, c, c, c) //type=int32
			fmt.Printf("%s\n", string(c))
		}
		for i:=0; i< len(s); i++ {
			b:= s[i]
			fmt.Printf("type=%T, %c, %v %x\n", b, b, b, b) //type=uint8
		}
		s = "あいう"
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01primary/" -*-
// Compilation started at Fri Oct  8 21:33:56
//  
// go run practice22.go
// start
// type=string, event last
// type=int32, e, 101 65
// e
// type=int32, v, 118 76
// v
// type=int32, e, 101 65
// e
// type=int32, n, 110 6e
// n
// type=int32, t, 116 74
// t
// type=int32,  , 32 20
//  
// type=int32, l, 108 6c
// l
// type=int32, a, 97 61
// a
// type=int32, s, 115 73
// s
// type=int32, t, 116 74
// t
// type=uint8, e, 101 65
// type=uint8, v, 118 76
// type=uint8, e, 101 65
// type=uint8, n, 110 6e
// type=uint8, t, 116 74
// type=uint8,  , 32 20
// type=uint8, l, 108 6c
// type=uint8, a, 97 61
// type=uint8, s, 115 73
// type=uint8, t, 116 74
// type=string, あいう
// type=int32, あ, 12354 3042
// あ
// type=int32, い, 12356 3044
// い
// type=int32, う, 12358 3046
// う
// type=uint8, ã, 227 e3
// type=uint8, , 129 81
// type=uint8, , 130 82
// type=uint8, ã, 227 e3
// type=uint8, , 129 81
// type=uint8, , 132 84
// type=uint8, ã, 227 e3
// type=uint8, , 129 81
// type=uint8, , 134 86
//  
// Compilation finished at Fri Oct  8 21:33:56
