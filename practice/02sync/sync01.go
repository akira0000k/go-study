package main

import (
	"fmt"
	"time"
)

/*
 subject: Goroutines receive channel
 */ 

func f(from string, m chan string) {
	for i:=0; i<30; i++ {
		fmt.Println(from, ":", i)
	}
	fmt.Println("end f")
	m <-"endf"
}

func main() {
	//channel := make(chan string, 2)
	channel := make(chan string)
	_ = time.Second
	//f("Direct")

	//go func(from string) {
	// 	for i:=0; i<30; i++ {
	// 		fmt.Println(from, ":", i)
	// 	}
	// 	fmt.Println("end f")
	// 	channel <- "endf"
	//}("GO F()")
	go f("GO F()", channel)
	time.Sleep(time.Millisecond / 100)

	
	go func(msg string) {
		fmt.Println(msg)
		channel <- "went"
	}("go-------ing")
	//time.Sleep(time.Millisecond / 10)
	//time.Sleep(time.Second * 5)

	for i:=0; i< 2; i++ {
		msg := <-channel
		fmt.Println("received:", msg)
	}
	fmt.Println("end main")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Fri Oct 29 18:11:34
//  
// go run sync01.go
// go-------ing
// received: went
// GO F() : 0
// GO F() : 1
// GO F() : 2
// GO F() : 3
// GO F() : 4
// GO F() : 5
// GO F() : 6
// GO F() : 7
// GO F() : 8
// GO F() : 9
// GO F() : 10
// GO F() : 11
// GO F() : 12
// GO F() : 13
// GO F() : 14
// GO F() : 15
// GO F() : 16
// GO F() : 17
// GO F() : 18
// GO F() : 19
// GO F() : 20
// GO F() : 21
// GO F() : 22
// GO F() : 23
// GO F() : 24
// GO F() : 25
// GO F() : 26
// GO F() : 27
// GO F() : 28
// GO F() : 29
// end f
// received: endf
// end main
//  
// Compilation finished at Fri Oct 29 18:11:37
