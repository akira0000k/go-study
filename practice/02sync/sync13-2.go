package main
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
   subject: test nil channel. needs timeout case
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
	const nwwk = 3 //10
	const readintv = time.Second / 3
	const writeintv = time.Second
	
	var readOps uint64
	var writeOps uint64

	var stopr = false
	var stopw = false
	var stopm = false
	
	reads := make(chan readOp)
	writes := make(chan writeOp)

	var wgr sync.WaitGroup
	var wgw sync.WaitGroup
	var wgm sync.WaitGroup

	wgm.Add(1)
	go func() {
		defer wgm.Done()
		var state = make(map[int]int)
	forloop:
		for i:=0; i<10000; i++ {
			if stopm {
				break forloop
			}
			select {
			case read, ok := <-reads:
				// closed: receive {0, nil}
				if !ok {
					fmt.Println("reads end")
					break forloop
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
					break forloop
				}
				state[write.key] = write.val
				write.resp <- true
			case to := <-time.After(1 * time.Second):
				// reads=nil write=nil のときはここに来る
				fmt.Println("main timeout", to)
			}
			//fmt.Println("func loop", i)
		}
		fmt.Println("func main end <=============")
	}()

	for r := 0; r < nrwk; r++ {
		wgr.Add(1)
		go func(wk int) {
			defer wgr.Done()
		forloop:
			for {
				if stopr {
					break forloop
				}
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				if reads != nil {
					reads <- read
					val := <-read.resp
					fmt.Printf("%d read[%d] = %d\n", wk, read.key, val)
					atomic.AddUint64(&readOps, 1)
				}
				time.Sleep(readintv)
			}
			fmt.Println("read func exit <-------------rrr")
		}(r)
	}

	for w := 0; w < nwwk; w++ {
		wgw.Add(1)
		go func(wk int) {
			defer wgw.Done()
		forloop:
			for {
				if stopw {
					break forloop
				}
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				fmt.Printf("%d write[%d]:%d response %v\n", wk, write.key, write.val, <-write.resp)
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(writeintv)
			}
			fmt.Println("write func exit <--------------www")
		}(w)
	}
	
	time.Sleep(time.Second * 2)
	fmt.Println("stopw = true")
	stopw = true
	wgw.Wait() //これがないと、panic: send on closed channel
	fmt.Println("wgw.Wait() exit")
	
	fmt.Println("writes=nil")
	writes2 := writes
	writes = nil

	//fmt.Println("close(writes2)")
	//close(writes2) //main終了していないとdeadlock

	time.Sleep(time.Second * 3)
	
	//一時read中止
	fmt.Println("reads=nil")
	reads2 := reads
	reads = nil
	time.Sleep(time.Second * 2)

	//read再開
	fmt.Println("reads=reads2")
	reads = reads2
	time.Sleep(time.Second * 2)

	//read終了
	fmt.Println("stopr = true")
	stopr = true
	wgr.Wait() //これがないと、panic: send on closed channel
	fmt.Println("wgr.Wait() exit")
	//
	fmt.Println("reads=nil")
	reads = nil

	time.Sleep(time.Second * 3)

	fmt.Println("stopm = true")
	stopm = true
	wgm.Wait()
	fmt.Println("wgm.Wait() exit")

	
	fmt.Println("close(reads2)")
	close(reads2) //main終了していないとdeadlock
	fmt.Println("close(writes2)")
	close(writes2) //main終了していないとdeadlock

	time.Sleep(time.Second * 3)
	
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
// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/" -*-
// Compilation started at Tue Sep 21 00:38:12
//  
// go run sync13-10.go
// 1 read[1] = -1
// 1 write[1]:18 response true
// 0 write[2]:59 response true
// 2 write[0]:40 response true
// 0 read[2] = -1
// 1 read[0] = 40
// 0 read[1] = 18
// 0 read[4] = -1
// 1 read[1] = 18
// 0 write[2]:89 response true
// 2 write[1]:45 response true
// 1 write[3]:74 response true
// 0 read[2] = 89
// 1 read[1] = 45
// 1 read[0] = 40
// 0 read[1] = 45
// 1 read[3] = 74
// 0 read[3] = 74
// stopw = true
// write func exit <--------------www
// write func exit <--------------www
// write func exit <--------------www
// wgw.Wait() exit
// writes=nil
// 1 read[2] = 89
// 0 read[2] = 89
// 0 read[2] = 89
// 1 read[3] = 74
// 1 read[0] = 40
// 0 read[0] = 40
// 0 read[1] = 45
// 1 read[3] = 74
// 1 read[2] = 89
// 0 read[1] = 45
// 0 read[4] = -1
// 1 read[1] = 45
// 1 read[2] = 89
// 0 read[1] = 45
// 1 read[1] = 45
// 0 read[0] = 40
// 0 read[3] = 74
// 1 read[0] = 40
// reads=nil
// main timeout 2021-09-21 00:38:18.283871 +0900 JST m=+5.701303344
// main timeout 2021-09-21 00:38:19.289599 +0900 JST m=+6.707032182
// reads=reads2
// main timeout 2021-09-21 00:38:20.294882 +0900 JST m=+7.712315803
// 0 read[0] = 40
// 1 read[0] = 40
// 0 read[3] = 74
// 1 read[3] = 74
// 1 read[3] = 74
// 0 read[0] = 40
// 0 read[1] = 45
// 1 read[0] = 40
// stopr = true
// read func exit <-------------rrr
// read func exit <-------------rrr
// wgr.Wait() exit
// reads=nil
// main timeout 2021-09-21 00:38:22.311336 +0900 JST m=+9.728770276
// main timeout 2021-09-21 00:38:23.316685 +0900 JST m=+10.734120258
// main timeout 2021-09-21 00:38:24.321988 +0900 JST m=+11.739423322
// stopm = true
// main timeout 2021-09-21 00:38:25.322773 +0900 JST m=+12.740209025
// func main end <=============
// wgm.Wait() exit
// close(reads2)
// close(writes2)
// readOps: 38
// writeOps: 6
// ============
// readOps: 38
// writeOps: 6
//  
// Compilation finished at Tue Sep 21 00:38:31
