package main

import (
	"fmt"
	"time"
)

/*
 subject: Goroutines buffered channel. recv/send type
 */ 


func main() {
	_ = time.Second
	//messages := make(chan string)
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"
	//fmt.Println(<-messages)
	//fmt.Println(<-messages)
	mstr := <- messages
	fmt.Println(mstr)
	fmt.Println(<- messages)

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "MESSAGe2")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
func ping(pings chan <- string, msg string) {
	pings <- msg
}
func pong(pings <- chan string, pongs chan <- string) {
	msg := <- pings
	pongs <- msg
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Fri Oct 29 18:17:36
//  
// go run sync02.go
// buffered
// channel
// MESSAGe2
//  
// Compilation finished at Fri Oct 29 18:17:36
