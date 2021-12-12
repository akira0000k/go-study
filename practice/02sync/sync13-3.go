package main
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
   subject: test nil channel. needs timeout case. Don't return from main by Closing channel!
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
		tb := "\t"
	forloop:
		for i:=0; i<10000; i++ {
			if stopm {
				break forloop
			}
			select {
			case read, ok := <-reads:
				// closed: receive {0, nil}
				if !ok {
					fmt.Println(tb, "reads end")
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
					fmt.Println(tb, "writes end")
					break forloop
				}
				state[write.key] = write.val
				write.resp <- true
			case to := <-time.After(1 * time.Second):
				// reads=nil write=nil のときはここに来る
				fmt.Println(tb, "main timeout", to)
			}
			//fmt.Println(tb, "func loop", i)
		}
		fmt.Println(tb, "func main end <=============")
	}()

	for r := 0; r < nrwk; r++ {
		wgr.Add(1)
		go func(wk int) {
			defer wgr.Done()
			tb := "\t\t"
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
					fmt.Printf(tb+" %d read[%d] = %d\n", wk, read.key, val)
					atomic.AddUint64(&readOps, 1)
				}
				time.Sleep(readintv)
			}
			fmt.Println(tb, "read func exit <-------------rrr")
		}(r)
	}

	for w := 0; w < nwwk; w++ {
		wgw.Add(1)
		go func(wk int) {
			defer wgw.Done()
			tb := "\t\t\t"
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
				fmt.Printf(tb+" %d write[%d]:%d response %v\n", wk, write.key, write.val, <-write.resp)
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(writeintv)
			}
			fmt.Println(tb, "write func exit <--------------www")
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
// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/02sync/" -*-
// Compilation started at Sat Sep 25 13:32:11
//  
// go run sync13-3.go 
//  		 0 read[1] = -1
//  			 1 write[2]:59 response true
//  			 0 write[1]:18 response true
//  		 1 read[2] = 59
//  		 0 read[0] = -1
//  		 1 read[0] = -1
//  		 1 read[1] = 18
//  		 0 read[0] = -1
//  			 1 write[2]:89 response true
//  			 0 write[4]:11 response true
//  		 1 read[4] = 11
//  		 0 read[3] = -1
//  		 1 read[1] = 18
//  		 0 read[0] = -1
//  		 0 read[2] = 89
//  		 1 read[1] = 18
// stopw = true
//  			 write func exit <--------------www
//  			 write func exit <--------------www
// wgw.Wait() exit
// writes=nil
// close(writes2)
//  	 writes end
//  	 func main end <=============
// reads=nil
// reads=reads2
// stopr = true
// fatal error: all goroutines are asleep - deadlock!
//  
// goroutine 1 [semacquire]:
// sync.runtime_Semacquire(0xc0000140b8)
//  	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/sema.go:56 +0x45
// sync.(*WaitGroup).Wait(0xc0000140b0)
//  	/usr/local/Cellar/go/1.16.6/libexec/src/sync/waitgroup.go:130 +0x65
// main.main()
//  	/Users/Akira/Desktop/work/go/practice/02sync/sync13-3.go:159 +0x6b3
//  
// goroutine 7 [chan send]:
// main.main.func2(0xc0000140b0, 0xc0000140a8, 0xc00000e028, 0xc000014098, 0x0)
//  	/Users/Akira/Desktop/work/go/practice/02sync/sync13-3.go:96 +0x118
// created by main.main
//  	/Users/Akira/Desktop/work/go/practice/02sync/sync13-3.go:84 +0x285
//  
// goroutine 8 [chan send]:
// main.main.func2(0xc0000140b0, 0xc0000140a8, 0xc00000e028, 0xc000014098, 0x1)
//  	/Users/Akira/Desktop/work/go/practice/02sync/sync13-3.go:96 +0x118
// created by main.main
//  	/Users/Akira/Desktop/work/go/practice/02sync/sync13-3.go:84 +0x285
// exit status 2
//  
// Compilation exited abnormally with code 1 at Sat Sep 25 13:32:20
