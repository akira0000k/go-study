package main
import (
	"fmt"
	//"math/rand"
	"sync"
	//"sync/atomic"
	"time"
)

/*
   subject: detect channel receive error
*/
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	fmt.Println("start main")

	//reads := make(chan readOp)
	//writes := make(chan writeOp)
	
	var wg sync.WaitGroup
	if true {
		ints := make(chan int)
		wg.Add(1)
		go func() {
			defer wg.Done()
			for irec:=range ints {
				fmt.Println(irec)
			}
			fmt.Println("irec stop")
		}()

		for i:=3; i>=-3; i-- {
			time.Sleep(time.Second)
			ints<-i
		}
		
		fmt.Println("close(ints)")
		close(ints)
		wg.Wait()
		fmt.Println("go ints ended")

		time.Sleep(time.Second * 5)
	}

	if true {
		ints := make(chan int)
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i:=0; i<10; i++ {
				select {
				case irec, ok := <-ints:
					if !ok {
						fmt.Println("not ok")
						fmt.Println("go func end")
						return
					}
					fmt.Println(irec)
				}
			}
		}()

		for i:=3; i>=-3; i-- {
			time.Sleep(time.Second)
			ints<-i
		}
		
		fmt.Println("close(ints)")
		close(ints)
		wg.Wait()
		fmt.Println("go ints ended")

		time.Sleep(time.Second * 5)
	}

	fmt.Println("end main")
}
// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/02sync/" -*-
// Compilation started at Fri Sep 24 19:29:22
//  
// go run sync13-1.go 
// start main
// 3
// 2
// 1
// 0
// -1
// -2
// close(ints)
// -3
// irec stop
// go ints ended
// 3
// 2
// 1
// 0
// -1
// -2
// close(ints)
// -3
// not ok
// go func end
// go ints ended
// end main
//  
// Compilation finished at Fri Sep 24 19:29:47
