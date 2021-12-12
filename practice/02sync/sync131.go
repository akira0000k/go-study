package main
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
   subject: Stateful Goroutines stop with flag.全然問題ない
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
	const readintv = 0 //time.Millisecond //Second / 3
	const writeintv = 0 //time.Millisecond
	
	var readOps uint64
	var writeOps uint64

	var stopflag = false
	
	reads := make(chan readOp)
	writes := make(chan writeOp)

	var wg sync.WaitGroup
	
	go func() {
		var state = make(map[int]int)
		for {
			if stopflag { return }
			select {
			case read, ok := <-reads:
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
	
	time.Sleep(time.Millisecond * 1)
	stopflag = true

	time.Sleep(time.Second * 2)

	//wg.Wait() //これがないと、panic: send on closed channel
	
	//fmt.Println("close(reads)")
	//close(reads)
	//fmt.Println("close(writes)")
	//close(writes)
	
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
// Compilation started at Fri Sep 24 19:23:27
//  
// go run sync131.go 
// 1 read[1] = -1
// 0 write[1]:18 response true
// 0 read[2] = -1
// 1 read[0] = -1
// 1 write[2]:59 response true
// 0 write[0]:56 response true
// 0 read[0] = 56
// 1 read[4] = -1
// 0 read[4] = -1
// 1 write[1]:62 response true
// 1 read[1] = 62
// 0 read[0] = 56
// 1 read[0] = 56
// 0 read[1] = 62
// 1 read[3] = -1
// 0 read[3] = -1
// 1 read[2] = 59
// 0 read[2] = 59
// 1 read[2] = 59
// 0 read[3] = -1
// 0 write[4]:28 response true
// 1 write[2]:6 response true
// 0 write[1]:8 response true
// 1 read[0] = 56
// 0 read[0] = 56
// 1 write[2]:31 response true
// 0 write[4]:56 response true
// 1 read[2] = 31
// 1 write[0]:26 response true
// 0 write[3]:90 response true
// 0 read[1] = 8
// 1 write[3]:33 response true
// 0 write[2]:78 response true
// 1 read[4] = 56
// 0 read[4] = 56
// 1 read[4] = 56
// 0 read[4] = 56
// 1 write[4]:53 response true
// 0 write[2]:21 response true
// 1 read[0] = 26
// 1 write[3]:38 response true
// 0 read[0] = 26
// 0 write[3]:55 response true
// 1 read[1] = 8
// 1 write[0]:5 response true
// 0 read[1] = 8
// 1 read[1] = 8
// 0 write[1]:28 response true
// 0 read[1] = 28
// 1 read[3] = 55
// 1 read[2] = 83
// 1 read[4] = 53
// 0 write[1]:2 response true
// 0 write[3]:96 response true
// 0 write[0]:23 response true
// 0 write[3]:37 response true
// 0 write[3]:41 response true
// 0 write[4]:33 response true
// 0 write[3]:91 response true
// 0 write[2]:78 response true
// 0 write[1]:46 response true
// 0 read[3] = 55
// 0 read[3] = 91
// 0 write[2]:40 response true
// 1 read[2] = 83
// 1 write[2]:83 response true
// 1 write[0]:98 response true
// 0 read[2] = 40
// 1 write[0]:57 response true
// 0 read[2] = 40
// 1 write[0]:10 response true
// 1 read[3] = 91
// 1 write[0]:32 response true
// 0 read[0] = 10
// 0 write[0]:51 response true
// 1 write[3]:91 response true
// 1 read[3] = 91
// 0 write[4]:97 response true
// 1 write[2]:37 response true
// 0 write[4]:26 response true
// 1 read[1] = 46
// 0 write[4]:66 response true
// 1 write[2]:81 response true
// 0 read[2] = 37
// 0 write[0]:93 response true
// 0 read[2] = 81
// 1 write[1]:19 response true
// 1 read[1] = 46
// 1 write[2]:49 response true
// 0 read[0] = 85
// 1 write[3]:84 response true
// 1 read[3] = 84
// 0 write[0]:85 response true
// 1 read[2] = 49
// 0 read[3] = 84
// 1 write[4]:47 response true
// 0 read[0] = 85
// 1 read[4] = 47
// 0 write[2]:16 response true
// 1 write[1]:51 response true
// readOps: 51
// writeOps: 49
// ============
// readOps: 51
// writeOps: 49
//  
// Compilation finished at Fri Sep 24 19:23:33
