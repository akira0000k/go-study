package main
import (
	"fmt"
	//"math/rand"
	"sync"
	//"sync/atomic"
	"time"
)

/*
   subject: range and channel1
*/

func main() {
	fmt.Println("main start")

	var waitg sync.WaitGroup
	ch1 := make(chan int)

	waitg.Add(1)
	go func(rcvchan chan int) {
		defer waitg.Done()
		
		for i := range rcvchan {
			fmt.Printf("i=%d   rcvchan=%v\n", i, rcvchan)
			//if i==3 {
			//	fmt.Println("rcvchan = nil")
			//	rcvchan = nil //影響なし
			//}
		}
		fmt.Println("go func exit.<-------")
	}(ch1)

	time.Sleep(time.Second)

	i := 0
	for ; i<5; i++ {
		ch1 <-i
		time.Sleep(time.Second)
	}
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
// Compilation started at Tue Sep 21 22:23:44
//  
// go run sync14.go
// main start
// i=0   rcvchan=0xc000056060
// i=1   rcvchan=0xc000056060
// i=2   rcvchan=0xc000056060
// i=3   rcvchan=0xc000056060
// i=4   rcvchan=0xc000056060
// ch1 = nil
// i=5   rcvchan=0xc000056060
// i=6   rcvchan=0xc000056060
// i=7   rcvchan=0xc000056060
// i=8   rcvchan=0xc000056060
// i=9   rcvchan=0xc000056060
// close(ch2)
// waitg.Wait()
// go func exit.<-------
// waitg.Wait()<---exit
// main end
//  
// Compilation finished at Tue Sep 21 22:23:59
