package main
import (
	"fmt"
	"os"
)
/*
   subject : Printf format %
*/
type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)
	fmt.Printf("%b %b %b %b\n", 14, 15, 16, 17)
	fmt.Printf("%c\n", 33)
	fmt.Printf("%x\n\n", 456)

	fmt.Printf("%f\n", 78.9)
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n\n", 123400000.0)

	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "hex this")
	fmt.Printf("%p\n\n", &p)

	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/08misc/" -*-
// Compilation started at Sat Oct  9 17:33:27
//  
// go run misc07.go
// {1 2}
// {x:1 y:2}
// main.point{x:1, y:2}
// main.point
// true
// 123
// 1110 1111 10000 10001
// !
// 1c8
//  
// 78.900000
// 1.234000e+08
// 1.234000E+08
//  
// "string"
// "\"string\""
// 6865782074686973
// 0xc0000ae010
//  
// |    12|   345|
// |  1.20|  3.45|
// |1.20  |3.45  |
// |   foo|     b|
// |foo   |b     |
// a string
// an error
//  
// Compilation finished at Sat Oct  9 17:33:27
