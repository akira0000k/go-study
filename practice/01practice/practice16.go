package main

import (
	"fmt"
)

/*
   subject : function pparam
 */ 
func  countup(n int, np *int) {
	n++
	*np += 2
}
func  cntup(n int, np *int) int {
	*np++
	return n + 3
}
func  countac(n int, n2 int) (int, int) {
	return n+10, n2+100
}

type n1n2 struct {
	n1 int
	n2 int
}
func (p *n1n2) strup() (int, int) {
	p.n1 += 2
	p.n2 += 4
	return p.n1, p.n2
}

func main() {

	fmt.Println("start")
	var n1, n2 int=0, 0

	countup(n1, &n2)
	fmt.Println(n1, n2)

	countup(n1, &n2)
	fmt.Println(n1, n2)

	countup(n1, &n2)
	fmt.Println(n1, n2)

	n1 = cntup(n1, &n2)
	fmt.Println(n1, n2)
	
	n1 = cntup(n1, &n2)
	fmt.Println(n1, n2)
	
	n1 = cntup(n1, &n2)
	fmt.Println(n1, n2)

	n1, n2 = countac(n1, n2)
	fmt.Println(n1, n2)
	n1, n2 = countac(n1, n2)
	fmt.Println(n1, n2)
	n1, n2 = countac(n1, n2)
	fmt.Println(n1, n2)

	qq := n1n2{ 0, 0 }
	fmt.Println(qq.strup())
	fmt.Println(qq.strup())
	fmt.Println(qq.strup())
	fmt.Println(qq.strup())
	qq.strup()
	fmt.Println(qq.n1, qq.n2)
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sat Oct 30 15:29:47
//  
// go run practice16.go
// start
// 0 2
// 0 4
// 0 6
// 3 7
// 6 8
// 9 9
// 19 109
// 29 209
// 39 309
// 2 4
// 4 8
// 6 12
// 8 16
// 10 20
//  
// Compilation finished at Sat Oct 30 15:29:48
