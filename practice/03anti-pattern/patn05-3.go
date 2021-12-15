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
 subject: isClosed関数を無理矢理実装  受信側で使うのは意味ない
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

func main() {
	log.Println("start")

	ch := make(chan Value, 10)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i:=0; i<10; i++ {
			if isClosed(ch) {
				log.Println("ch Closed. ------------------")
				return
			} else {
				v, ok := <-ch
				if ok {
					log.Println("read v=", int(v))
				} else {
					log.Println("v, ok<-", v, ok)
				}
			}
			time.Sleep(time.Millisecond * 1)
		}
	}()
	i := 10
	for ; i<13; i++ {
		ch <-SomeValue(i)
		time.Sleep(time.Millisecond * 1000)
	}
	for ; i<16; i++ {
		ch <-SomeValue(i)
		time.Sleep(time.Millisecond * 100)
	}
	log.Printf("close(ch)")
	close(ch)
	log.Printf("waiting wg....")
	wg.Wait()

	log.Println("end")
}
// -*- mode: compilation; default-directory: "~/go/src/go-study/practice/03anti-pattern/" -*-
// Compilation started at Wed Dec 15 21:04:05
//  
// go run patn05-31.go
// 2021/12/15 21:04:05 start
// ch<- 10 sending.. sent
// 2021/12/15 21:04:05 read v= 10
// default:
// 2021/12/15 21:04:06 read v= 11
// default:
// 2021/12/15 21:04:07 read v= 12
// default:
// 2021/12/15 21:04:08 read v= 13
// default:
// 2021/12/15 21:04:08 read v= 14
// default:
// 2021/12/15 21:04:08 read v= 15
// default:
// 2021/12/15 21:04:08 close(ch)
// 2021/12/15 21:04:08 waiting wg....
// 2021/12/15 21:04:08 v, ok<- 0 false
// 2021/12/15 21:04:08 ch Closed. ------------------
// 2021/12/15 21:04:08 end
//  
// Compilation finished at Wed Dec 15 21:04:08
