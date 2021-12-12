package main
import (
	"fmt"
	//"math/rand"
	//"sync"
	//"sync/atomic"
	"time"
)

/*
   subject: how to detect channel closed
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
	const nrwk = 2 //100
	const nwwk = 2 //10
	const readintv = time.Second / 3
	const writeintv = time.Second
	
	
	reads := make(chan readOp)
	writes := make(chan writeOp)


	go func() {
		var state = make(map[int]int)
		for {
			//if stopflag { return }
			select {
			case read := <-reads:
				if read.resp == nil {
					fmt.Println("goroutine exit")
					return
				}
				val, ok := state[read.key]
				if ok {
					read.resp <- val
				} else {
					read.resp <- -1
				}
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()


	time.Sleep(time.Second * 2)
	
	fmt.Println("close(reads)")
	close(reads)
	fmt.Println("close(writes)")
	close(writes)
	
	time.Sleep(time.Second * 10)
}
// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/02sync/" -*-
// Compilation started at Fri Sep 24 19:28:14
//  
// go run sync13-0.go 
// close(reads)
// close(writes)
// goroutine exit
//  
// Compilation finished at Fri Sep 24 19:28:27
