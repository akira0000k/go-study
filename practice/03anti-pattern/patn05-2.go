package main

import (
	"log"
	//"math/rand"
	//"sync"
	//"sync/atomic"
	"time"
	//"errors"
)

/*
 subject: isClosed関数を無理矢理実装     ダメぽ
 イベント送信時に channel が closed か調べたい
 複数の送信側処理があり、受信は一つで良い場合に欲しくなるやつ。たまに欲しくなります。
 が、今のところ便利なインターフェースは提供されていません。
 受信してみて第二戻り値が false なら closed ですが、受信すると第一戻り値を普通に受け取っちゃう場合もあるのでよろしくない。
 一応こんな感じで実装してみましたが、アンチパターン臭がすごい。
*/
type Value int


func SomeValue() Value {
	var v int
	v = 3
	
	return Value(v)
}

func isClosed(ch chan Value) bool {
	log.Println("isClosed called")
	for v := range ch {
		log.Println("v =", v)
		log.Println("go func()")
		go func() {
			log.Println("ch<-", v, "send")
			ch <-v        //panic
			log.Println("send done")
		}()
		log.Println("isClosed return false")
		return false
	}
	log.Println("isClosed return true")
	return false
}

func A(ch chan Value) {
	log.Println("A start")

	if !isClosed(ch) {        //panic
		ch <-SomeValue()
	}
}

func main() {
	log.Println("start")

	ch := make(chan Value, 10)
	ch <-SomeValue()
	A(ch)
	time.Sleep(time.Second)
	close(ch)
	time.Sleep(time.Second)
	A(ch)    //panic

	log.Println("end")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/03anti-pattern/" -*-
// Compilation started at Sun Dec 12 21:03:14
//  
// go run patn05-2.go
// 2021/12/12 21:03:14 start
// 2021/12/12 21:03:14 A start
// 2021/12/12 21:03:14 isClosed called
// 2021/12/12 21:03:14 v = 3
// 2021/12/12 21:03:14 go func()
// 2021/12/12 21:03:14 isClosed return false
// 2021/12/12 21:03:14 ch<- 3 send
// 2021/12/12 21:03:14 send done
// 2021/12/12 21:03:16 A start
// 2021/12/12 21:03:16 isClosed called
// 2021/12/12 21:03:16 v = 3
// 2021/12/12 21:03:16 go func()
// 2021/12/12 21:03:16 ch<- 3 send
// 2021/12/12 21:03:16 isClosed return false
// panic: send on closed channel
//  
// goroutine 1 [running]:
// main.A(0xc00006c000)
//  	/Users/Akira/go/src/practice/03anti-pattern/patn05-2.go:51 +0xa5
// main.main()
//  	/Users/Akira/go/src/practice/03anti-pattern/patn05-2.go:64 +0xf3
// exit status 2
//  
// Compilation exited abnormally with code 1 at Sun Dec 12 21:03:16
