package main
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
   subject: test nil channel. needs timeout case. Closing channel cause error!
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
					//break forloop
					break
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
					//break forloop
					break
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
	//time.Sleep(time.Second * 1)
	
	fmt.Println("close(writes2)")
	close(writes2) //main終了していないとdeadlock

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

	fmt.Println("close(reads2)")
	close(reads2) //main終了していないとdeadlock

	time.Sleep(time.Second * 1)

	fmt.Println("stopm = true")
	stopm = true
	wgm.Wait()
	fmt.Println("wgm.Wait() exit")

	
	//fmt.Println("close(reads2)")
	//close(reads2) //main終了していないとdeadlock
	////fmt.Println("close(writes2)")
	////close(writes2) //main終了していないとdeadlock

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
// Compilation started at Tue Sep 21 00:50:39
//  
// go run sync13-12.go
// 1 write[4]:81 response true
// 0 write[2]:47 response true
// 1 read[1] = -1
// 0 read[3] = -1
// 0 read[0] = -1
// 1 read[0] = -1
// 1 read[1] = -1
// 0 read[0] = -1
// 1 write[4]:11 response true
// 0 write[2]:89 response true
// 0 read[3] = -1
// 1 read[4] = 11
// 1 read[1] = -1
// 0 read[0] = -1
// 1 read[2] = 89
// 0 read[1] = -1
// stopw = true
// write func exit <--------------www
// write func exit <--------------www
// wgw.Wait() exit
// writes=nil
// close(writes2)
// writes end
// 0 read[0] = -1
// 1 read[1] = -1
// 1 read[3] = -1
// 0 read[3] = -1
// 0 read[2] = 89
// 1 read[2] = 89
// 1 read[2] = 89
// 0 read[3] = -1
// 0 read[0] = -1
// 1 read[0] = -1
// 1 read[1] = -1
// 0 read[3] = -1
// 0 read[2] = 89
// 1 read[1] = -1
// 0 read[1] = -1
// 1 read[4] = 11
// 1 read[2] = 89
// 0 read[1] = -1
// reads=nil
// main timeout 2021-09-21 00:50:46.291031 +0900 JST m=+5.697199939
// main timeout 2021-09-21 00:50:47.296688 +0900 JST m=+6.702857465
// reads=reads2
// main timeout 2021-09-21 00:50:48.301971 +0900 JST m=+7.708141634
// 0 read[2] = 89
// 1 read[1] = -1
// 1 read[4] = 11
// 0 read[4] = 11
// 0 read[0] = -1
// 1 read[0] = -1
// 1 read[3] = -1
// 0 read[3] = -1
// stopr = true
// read func exit <-------------rrr
// read func exit <-------------rrr
// wgr.Wait() exit
// reads=nil
// main timeout 2021-09-21 00:50:50.319442 +0900 JST m=+9.725613106
// main timeout 2021-09-21 00:50:51.322426 +0900 JST m=+10.728598001
// main timeout 2021-09-21 00:50:52.327706 +0900 JST m=+11.733878723
// close(reads2)
// main timeout 2021-09-21 00:50:53.332984 +0900 JST m=+12.739156888
// stopm = true
// main timeout 2021-09-21 00:50:54.338275 +0900 JST m=+13.744448080
// func main end <=============
// wgm.Wait() exit
// readOps: 38
// writeOps: 4
// ============
// readOps: 38
// writeOps: 4
//  
// Compilation finished at Tue Sep 21 00:51:00
