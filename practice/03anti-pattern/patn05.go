package main

import (
	"fmt"
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
	fmt.Println("isClosed called")
	select {
	case v, ok := <-ch:
		if !ok {
			fmt.Println("isClosed return true")
			return true
		}
		//読み出した時点で close(ch) されるらしい
		fmt.Println("go func()")
		go func() {
			fmt.Println("ch<-", v)
			ch <-v        //panic
			fmt.Println("send done")
		}()
	default:
		fmt.Println("default:")
	}
	fmt.Println("isClosed return false")
	return false
}

func A(ch chan Value) {
	fmt.Println("A start")

	if !isClosed(ch) {        //panic
		ch <-SomeValue()
	}
}

func main() {
	fmt.Println("start")

	ch := make(chan Value, 10)
	ch <-SomeValue()
	A(ch)
	time.Sleep(time.Second)
	close(ch)
	time.Sleep(time.Second)
	A(ch)    //panic

	fmt.Println("end")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/03anti-pattern/" -*-
// Compilation started at Sun Dec 12 20:54:25
//  
// go run patn05.go
// start
// A start
// isClosed called
// go func()
// isClosed return false
// ch<- 3
// send done
// A start
// isClosed called
// go func()
// isClosed return false
// ch<- 3
// panic: send on closed channel
//  
// goroutine 17 [running]:
// main.isClosed.func1(0x3, 0xc000074000)
//  	/Users/Akira/go/src/practice/03anti-pattern/patn05.go:36 +0xcd
// created by main.isClosed
//  	/Users/Akira/go/src/practice/03anti-pattern/patn05.go:34 +0x165
// exit status 2
//  
// Compilation exited abnormally with code 1 at Sun Dec 12 20:54:27
