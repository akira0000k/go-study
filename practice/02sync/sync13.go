package main
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
   subject: Stateful Goroutines
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
	
	var readOps uint64
	var writeOps uint64

	var stopflag = false
	
	reads := make(chan readOp)
	writes := make(chan writeOp)

	var wg sync.WaitGroup
	
	go func() {
		var state = make(map[int]int)
		for {
			//if stopflag { return }
			select {
			case read, ok := <-reads:
				// closed: receive {0, nil}
				if !ok {
					fmt.Println("reads end")
					return
				}
				val, ok := state[read.key]
				if ok {
					read.resp <- val
				} else {
					read.resp <- -1
				}
			case write, ok := <-writes:
				if !ok {
					fmt.Println("writes end")
					return
				}
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < nrwk; r++ {
		wg.Add(1)
		go func(wk int) {
			defer wg.Done()
			for {
				if stopflag { return }
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				val := <-read.resp
				fmt.Printf("%d read[%d] = %d\n", wk, read.key, val)
				atomic.AddUint64(&readOps, 1)
				time.Sleep(readintv)
			}
		}(r)
	}

	for w := 0; w < nwwk; w++ {
		wg.Add(1)
		go func(wk int) {
			defer wg.Done()
			for {
				if stopflag { return }
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				fmt.Printf("%d write[%d]:%d response %v\n", wk, write.key, write.val, <-write.resp)
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(writeintv)
			}
		}(w)
	}
	
	time.Sleep(time.Second * 2)
	stopflag = true
	wg.Wait() //これがないと、panic: send on closed channel
	
	fmt.Println("close(reads)")
	close(reads)
	fmt.Println("close(writes)")
	close(writes)
	
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	fmt.Println("============")
	time.Sleep(time.Second * 3)
	readOpsFinal = atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal = atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/02sync/" -*-
// Compilation started at Fri Sep 24 19:17:04
//  
// go run sync13.go 
// 1 read[1] = -1
// 0 write[2]:47 response true
// 1 write[4]:81 response true
// 0 read[3] = -1
// 0 read[0] = -1
// 1 read[0] = -1
// 1 read[1] = -1
// 0 read[0] = -1
// 1 write[4]:11 response true
// 0 write[2]:89 response true
// 1 read[3] = -1
// 0 read[4] = 11
// 0 read[1] = -1
// 1 read[0] = -1
// 1 read[2] = 89
// 0 read[1] = -1
// close(reads)
// close(writes)
// readOps: 12
// writeOps: 4
// ============
// reads end
