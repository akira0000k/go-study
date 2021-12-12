package main

import (
	"fmt"
	//"math/rand"
	//"sync"
	//"sync/atomic"
	"time"
)

/*
 subject: 終了eventを送るのは良くなくて、closeすべきと言うことなんだけど。。
*/

func A(done chan struct{}) {
	defer close(done)
	// do something
	fmt.Println("A start")
	time.Sleep(time.Second * 2)

	// do something
	//fmt.Println("illegal return A"); return


	fmt.Println("done <-struct{}{}"); done <-struct{}{}
	//fmt.Println("close(done)"); close(done)
	fmt.Println("A return")
}

func B() {
	fmt.Println("B start")

	done := make(chan struct{})
	go A(done)


	time.Sleep(time.Second * 4)


	fmt.Println("<-done")
	val, ok := <-done
	fmt.Println(val, ok)
	fmt.Println("B end")
}

func main() {
	fmt.Println("start")

	B()

	fmt.Println("end")
}
// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/" -*-
// Compilation started at Sat Sep 25 19:07:18
//  
// go run patn01.go
// start
// B start
// A start
// done <-struct{}{}
// <-done
// {} true
// B end
// end
//  
// Compilation finished at Sat Sep 25 19:07:22
