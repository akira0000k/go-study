package main

import (
	"fmt"
	"time"
)

/*
 subject: Goroutines channel select with timeout
 */ 

func f(from string, sendchan chan<-string) {
	for i:=0; i<30; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Second / 10)
	}
	fmt.Println("end f")
	sendchan<-"endf"
}

func main() {
	//channel := make(chan string, 2)
	channel1 := make(chan string)
	channel2 := make(chan string)
	_ = time.Second
	//f("Direct")

	//go func(from string) {
	// 	for i:=0; i<30; i++ {
	// 		fmt.Println(from, ":", i)
	// 	}
	// 	fmt.Println("end f")
	// 	channel <- "endf"
	//}("GO F()")
	go f("GO F()", channel1)
	time.Sleep(time.Millisecond / 100)

	
	go func(msg string) {
		fmt.Println(msg)
		channel2 <- "went"
	}("go-------ing")
	//time.Sleep(time.Millisecond / 10)
	//time.Sleep(time.Second * 5)

	//for i:=0; i< 3; i++ {   //fatal error: all goroutines are asleep - deadlock!
waiting:
	for i:=0; i< 2; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("received f:", msg1)
		case msg2 := <-channel2:
			fmt.Println("received func:", msg2)

		case to := <-time.After(3 * time.Second):
			fmt.Println("timeout", to)
			break waiting
		//default:
		// 	fmt.Println("default:") no default in select channel
		}
	}
	fmt.Println("end main")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Fri Oct 29 18:14:38
//  
// go run sync01-2.go
// GO F() : 0
// go-------ing
// received func: went
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
// timeout 2021-10-29 18:14:42.209647 +0900 JST m=+3.002378822
// end main
//  
// Compilation finished at Fri Oct 29 18:14:42
