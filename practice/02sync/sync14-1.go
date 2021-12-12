package main
import (
	"fmt"
	//"math/rand"
	"sync"
	//"sync/atomic"
	"time"
)

/*
   subject: for range use channel at start time
*/

func main() {
	fmt.Println("main start")

	var waitg sync.WaitGroup
	ch1 := make(chan int)

	waitg.Add(1)
	go func() {
		defer waitg.Done()
		//  rangeが使う ch1は forループが始まった時のものがずっと使われる
		for i := range ch1 {
			fmt.Printf("i=%d  ch1=%v\n", i, ch1) //ch1の値は外部で書き換えられるかもしれない
		}
		fmt.Println("go func exit.<-------")
	}()
	// go func(rcvchan chan int) {
	//  	defer waitg.Done()
	//  	//             rcvchanは引数なので外部のch1変更の影響は受けない
	//  	for i := range rcvchan {
	//  		fmt.Printf("i=%d\n", i)
	//  	}
	//  	fmt.Println("go func exit.<-------")
	// }(ch1)

	time.Sleep(time.Second)

	i := 0
	for ; i<5; i++ {
		ch1 <-i
		time.Sleep(time.Second)
	}

	// チャネルの実体は同じ
	fmt.Println("ch1 = nil")
	ch2 := ch1
	ch1 = nil
	time.Sleep(time.Second * 3)
	//fmt.Println("ch1 = ch2")
	//ch1 = ch2
	
	for ; i<10; i++ {
		//ch1 <-i
		ch2 <-i
		time.Sleep(time.Second)
	}
	
	fmt.Println("close(ch2)")
	close(ch2)
	fmt.Println("waitg.Wait()")
	waitg.Wait()
	fmt.Println("waitg.Wait()<---exit")
	
	fmt.Println("main end")
}
// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/02sync/" -*-
// Compilation started at Tue Sep 21 22:19:05
//  
// go run sync14.go
// main start
// i=0  ch1=0xc000056060
// i=1  ch1=0xc000056060
// i=2  ch1=0xc000056060
// i=3  ch1=0xc000056060
// i=4  ch1=0xc000056060
// ch1 = nil
// i=5  ch1=<nil>
// i=6  ch1=<nil>
// i=7  ch1=<nil>
// i=8  ch1=<nil>
// i=9  ch1=<nil>
// close(ch2)
// waitg.Wait()
// go func exit.<-------
// waitg.Wait()<---exit
// main end
//  
// Compilation finished at Tue Sep 21 22:19:20
