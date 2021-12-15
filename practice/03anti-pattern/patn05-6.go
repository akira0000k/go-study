package main

import (
	"log"
	"fmt"
	"sync"
	//"sync/atomic"
	"time"
	//"errors"
)

/*
 subject: isClosed関数を無理矢理実装   isClosedのなかでパニック
*/
type Value int


func SomeValue(v int) Value {
	return Value(v)
}

func isClosed(ch chan Value) bool {
	//fmt.Println("isClosed called")
	select {
	case v, ok := <-ch:
		if !ok {
			//fmt.Println("isClosed return true")
			return true
		}
		//fmt.Println("go func()")
		go func() {
			fmt.Print("ch<- ", v, " sending.. ")
			ch <-v  //panic
			fmt.Println("sent")
		}()
	default:
		fmt.Println("default:")
	}
	//fmt.Println("isClosed return false")
	return false
}

func A(ch chan Value, v int) bool {
	if isClosed(ch) {
		return false
	} else {
		ch <-SomeValue(v)
		return true
	}
}

func main() {
	log.Println("start")

	ch := make(chan Value, 10)
	ch<-SomeValue(8)
	ch<-SomeValue(9)
	var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	// 	defer wg.Done()
	// 	for i:=0; ; i++ {
	// 		v, ok := <-ch
	// 		if ok {
	// 			log.Println("read v=", int(v))
	// 		} else {
	// 			log.Println("v, ok<-", v, ok)
	// 			return
	// 		}
	// 		time.Sleep(time.Millisecond * 1)
	// 	}
	//}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 10
		for ; i<13; i++ {
			if !A(ch, i) {
				return
			}
			time.Sleep(time.Millisecond * 1000)
		}
		for ; i<16; i++ {
			if !A(ch, i) {
				return
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()

	time.Sleep(time.Millisecond * 1)

	log.Printf("close(ch)")
	close(ch)

	log.Printf("waiting wg....")
	wg.Wait()

	log.Println("end")
}
// -*- mode: compilation; default-directory: "~/go/src/go-study/practice/03anti-pattern/" -*-
// Compilation started at Thu Dec 16 00:56:48
//  
// go run patn05-6.go
// 2021/12/16 00:56:48 start
// ch<- 8 sending.. sent
// 2021/12/16 00:56:48 close(ch)
// 2021/12/16 00:56:48 waiting wg....
// ch<- 9 sending.. panic: send on closed channel
//  
// goroutine 8 [running]:
// main.isClosed.func1(0x9, 0xc000102000)
//  	/Users/Akira/go/src/go-study/practice/03anti-pattern/patn05-6.go:33 +0xf9
// created by main.isClosed
//  	/Users/Akira/go/src/go-study/practice/03anti-pattern/patn05-6.go:31 +0x7f
// exit status 2
//  
// Compilation exited abnormally with code 1 at Thu Dec 16 00:56:50
