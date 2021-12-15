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
 subject: isClosed関数を無理矢理実装   タイミングによっては送信時にパニック
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
			ch <-v
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
		ch <-SomeValue(v)  //panic
		return true
	}
}

func main() {
	log.Println("start")

	ch := make(chan Value, 10)
	ch<-SomeValue(8)
	ch<-SomeValue(9)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i:=0; ; i++ {
			v, ok := <-ch
			if ok {
				log.Println("read v=", int(v))
			} else {
				log.Println("v, ok<-", v, ok)
				return
			}
			time.Sleep(time.Millisecond * 1)
		}
	}()

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

	time.Sleep(time.Millisecond * 3115)

	log.Printf("close(ch)")
	close(ch)

	log.Printf("waiting wg....")
	wg.Wait()

	log.Println("end")
}
// -*- mode: compilation; default-directory: "~/go/src/go-study/practice/03anti-pattern/" -*-
// Compilation started at Thu Dec 16 00:34:48
//  
// go run patn05-5.go
// 2021/12/16 00:34:49 start
// 2021/12/16 00:34:49 read v= 8
// ch<- 9 sending.. sent
// 2021/12/16 00:34:49 read v= 10
// 2021/12/16 00:34:49 read v= 9
// default:
// 2021/12/16 00:34:50 read v= 11
// default:
// 2021/12/16 00:34:51 read v= 12
// default:
// 2021/12/16 00:34:52 read v= 13
// 2021/12/16 00:34:52 close(ch)
// default:
// 2021/12/16 00:34:52 waiting wg....
// 2021/12/16 00:34:52 v, ok<- 0 false
// 2021/12/16 00:34:52 end
// panic: send on closed channel
//  
// goroutine 19 [running]:
// main.A(...)
//  	/Users/Akira/go/src/go-study/practice/03anti-pattern/patn05-5.go:47
// main.main.func2(0xc00011c060, 0xc000130000)
//  	/Users/Akira/go/src/go-study/practice/03anti-pattern/patn05-5.go:85 +0x15a
// created by main.main
//  	/Users/Akira/go/src/go-study/practice/03anti-pattern/patn05-5.go:75 +0x165
// exit status 2
//  
// Compilation exited abnormally with code 1 at Thu Dec 16 00:34:52
