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
 subject: isClosed関数を無理矢理実装   キューの順序が狂うがなんとか動いている
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

	time.Sleep(time.Millisecond * 3500)

	log.Printf("close(ch)")
	close(ch)

	log.Printf("waiting wg....")
	wg.Wait()

	log.Println("end")
}
// -*- mode: compilation; default-directory: "~/go/src/go-study/practice/03anti-pattern/" -*-
// Compilation started at Thu Dec 16 00:27:09
//  
// go run patn05-41.go
// 2021/12/16 00:27:12 start
// ch<- 9 sending.. sent
// 2021/12/16 00:27:12 read v= 8
// 2021/12/16 00:27:12 read v= 10
// 2021/12/16 00:27:12 read v= 9
// default:
// 2021/12/16 00:27:13 read v= 11
// default:
// 2021/12/16 00:27:14 read v= 12
// default:
// 2021/12/16 00:27:15 read v= 13
// default:
// 2021/12/16 00:27:15 read v= 14
// default:
// 2021/12/16 00:27:15 read v= 15
// 2021/12/16 00:27:16 close(ch)
// 2021/12/16 00:27:16 waiting wg....
// 2021/12/16 00:27:16 v, ok<- 0 false
// 2021/12/16 00:27:16 end
//  
// Compilation finished at Thu Dec 16 00:27:16
