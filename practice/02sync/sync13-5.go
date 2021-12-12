package main
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
   subject: channel=nil and close backup. needs timeout case.
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
					fmt.Println("reads end") //nilを入れて無視される前にこちらに来ることがある
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
					fmt.Println("writes end") //nilを入れて無視される前にこちらに来ることがある
					break
				}
				state[write.key] = write.val
				write.resp <- true
			case to := <-time.After(1 * time.Second):
				// これがないとdeadlockとか出る
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
					//ここに割り込まれたらと考えると本当はマズイ
					reads <- read
					val := <-read.resp
					fmt.Printf("%d read[%d] = %d\n", wk, read.key, val)
					atomic.AddUint64(&readOps, 1)
				} else {
					fmt.Println("reads == nil..skip")
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

	time.Sleep(time.Second * 1)

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

	//再びread中止
	fmt.Println("reads=nil")
	reads = nil
	time.Sleep(time.Second * 2)

	fmt.Println("close(reads2)")
	close(reads2) //mainで一度だけ感知できる。

	time.Sleep(time.Second * 3)

	//read終了
	fmt.Println("stopr = true")
	stopr = true
	wgr.Wait()
	fmt.Println("wgr.Wait() exit")
	//

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
// Compilation started at Tue Sep 21 14:01:47
//  
// go run sync13-41.go
// 1 write[2]:59 response true
// 1 read[1] = -1
// 0 write[1]:18 response true
// 0 read[2] = 59
// 1 read[0] = -1
// 0 read[0] = -1
// 1 read[0] = -1
// 0 read[1] = 18
// 1 write[2]:89 response true
// 0 write[4]:11 response true
// 0 read[3] = -1
// 1 read[4] = 11
// 1 read[1] = 18
// 0 read[0] = -1
// 0 read[2] = 89
// 1 read[1] = 18
// stopw = true
// 1 write[0]:66 response true
// 0 write[3]:58 response true
// 1 read[2] = 89
// 0 read[2] = 89
// 0 read[2] = 89
// 1 read[3] = 58
// 1 read[0] = 66
// 0 read[0] = 66
// 0 read[1] = 18
// write func exit <--------------www
// write func exit <--------------www
// wgw.Wait() exit
// 1 read[3] = 58
// writes=nil
// 0 read[2] = 89
// 1 read[1] = 18
// 0 read[1] = 18
// 1 read[4] = 11
// close(writes2)
// 1 read[2] = 89
// 0 read[1] = 18
// 1 read[0] = 66
// 0 read[1] = 18
// 0 read[3] = 58
// 1 read[0] = 66
// 1 read[4] = 11
// 0 read[3] = 58
// 1 read[2] = 89
// 0 read[3] = 58
// 0 read[3] = 58
// 1 read[4] = 11
// 1 read[4] = 11
// 0 read[3] = 58
// 0 read[2] = 89
// 1 read[1] = 18
// 0 read[4] = 11
// 1 read[4] = 11
// reads=nil
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// main timeout 2021-09-21 14:01:55.414545 +0900 JST m=+7.684376280
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// main timeout 2021-09-21 14:01:56.415448 +0900 JST m=+8.685330743
// reads == nil..skip
// reads == nil..skip
// reads=reads2
// main timeout 2021-09-21 14:01:57.416103 +0900 JST m=+9.686035628
// 0 read[1] = 18
// 1 read[2] = 89
// 1 read[3] = 58
// 0 read[1] = 18
// 0 read[3] = 58
// 1 read[1] = 18
// 0 read[3] = 58
// 1 read[2] = 89
// reads=nil
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// main timeout 2021-09-21 14:01:59.417589 +0900 JST m=+11.687623101
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// main timeout 2021-09-21 14:02:00.417762 +0900 JST m=+12.687846693
// reads == nil..skip
// reads == nil..skip
// close(reads2)
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// main timeout 2021-09-21 14:02:01.41799 +0900 JST m=+13.688126118
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// main timeout 2021-09-21 14:02:02.41811 +0900 JST m=+14.688296198
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// reads == nil..skip
// main timeout 2021-09-21 14:02:03.418186 +0900 JST m=+15.688422771
// reads == nil..skip
// reads == nil..skip
// stopr = true
// read func exit <-------------rrr
// read func exit <-------------rrr
// wgr.Wait() exit
// main timeout 2021-09-21 14:02:04.41845 +0900 JST m=+16.688738206
// stopm = true
// main timeout 2021-09-21 14:02:05.418515 +0900 JST m=+17.688854099
// func main end <=============
// wgm.Wait() exit
// readOps: 50
// writeOps: 6
// ============
// readOps: 50
// writeOps: 6
//  
// Compilation finished at Tue Sep 21 14:02:11
