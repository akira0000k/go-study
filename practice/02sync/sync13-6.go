package main
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
   subject: channel=nil write use select timeout
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
	const readintv = time.Millisecond * 100
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
				select {
				case reads <- read:
					val := <-read.resp
					fmt.Printf("%d read[%d] = %d\n", wk, read.key, val)
					atomic.AddUint64(&readOps, 1)

					time.Sleep(readintv)
				case <-time.After(readintv):
					fmt.Print(".")
				}
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
				select {
				case writes <- write:
					fmt.Printf("%d write[%d]:%d response %v\n", wk, write.key, write.val, <-write.resp)
					atomic.AddUint64(&writeOps, 1)
					time.Sleep(writeintv)
				case <-time.After(writeintv):
					fmt.Print(",")
				}
			}
			fmt.Println("write func exit <--------------www")
		}(w)
	}

	fmt.Println("reads=nil")
	reads2 := reads
	reads = nil
	
	time.Sleep(time.Second * 4)

	fmt.Println("writes=nil")
	writes2 := writes
	writes = nil

	time.Sleep(time.Second * 5)

	fmt.Println("stopw = true")
	stopw = true
	wgw.Wait() //これがないと、panic: send on closed channel
	fmt.Println("wgw.Wait() exit")
	
	fmt.Println("close(writes2)")
	close(writes2)

	
	time.Sleep(time.Second * 1)

	for i:=0; i< 3; i++ {
		//read再開
		fmt.Println("reads=reads2")
		reads = reads2
		time.Sleep(time.Second * 5)

		//再びread中止
		fmt.Println("reads=nil")
		reads = nil
		time.Sleep(time.Second * 3)
	}

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
// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/02sync/" -*-
// Compilation started at Sat Sep 25 13:35:22
//  
// go run sync13-6.go 
// reads=nil
// 1 write[2]:59 response true
// 0 write[1]:18 response true
// ..................main timeout 2021-09-25 13:35:24.154135 +0900 JST m=+1.005419684
// 1 write[2]:47 response true
// 0 write[2]:88 response true
// ....................0 write[4]:53 response true
// 1 write[2]:21 response true
// ....................main timeout 2021-09-25 13:35:26.159706 +0900 JST m=+3.010932057
// 1 write[2]:18 response true
// 0 write[2]:94 response true
// ..................writes=nil
// main timeout 2021-09-25 13:35:27.16117 +0900 JST m=+4.012367372
// ....................,,main timeout 2021-09-25 13:35:28.163405 +0900 JST m=+5.014572799
// ...................main timeout 2021-09-25 13:35:29.164635 +0900 JST m=+6.015773265
// .,,..................,,main timeout 2021-09-25 13:35:30.169878 +0900 JST m=+7.020986511
// ....................,,main timeout 2021-09-25 13:35:31.175204 +0900 JST m=+8.026283457
// ..................stopw = true
// ..,write func exit <--------------www
// ,write func exit <--------------www
// main timeout 2021-09-25 13:35:32.178014 +0900 JST m=+9.029063536
// wgw.Wait() exit
// close(writes2)
// ..................reads=reads2
// main timeout 2021-09-25 13:35:33.183292 +0900 JST m=+10.034312489
// ..0 read[0] = -1
// 1 read[0] = -1
// 1 read[2] = 94
// 0 read[4] = 53
// 1 read[0] = -1
// 0 read[1] = 18
// 0 read[0] = -1
// 1 read[4] = 53
// 1 read[4] = 53
// 0 read[4] = 53
// 0 read[0] = -1
// 1 read[1] = 18
// 0 read[2] = 94
// 1 read[1] = 18
// 1 read[1] = 18
// 0 read[0] = -1
// 0 read[4] = 53
// 1 read[2] = 94
// 1 read[3] = -1
// 0 read[1] = 18
// 0 read[4] = 53
// 1 read[0] = -1
// 0 read[0] = -1
// 1 read[4] = 53
// 1 read[4] = 53
// 0 read[4] = 53
// 0 read[0] = -1
// 1 read[2] = 94
// 1 read[0] = -1
// 0 read[0] = -1
// 0 read[4] = 53
// 1 read[3] = -1
// 1 read[0] = -1
// 0 read[3] = -1
// 0 read[1] = 18
// 1 read[0] = -1
// 1 read[4] = 53
// 0 read[4] = 53
// 1 read[0] = -1
// 0 read[1] = 18
// 0 read[3] = -1
// 1 read[2] = 94
// 1 read[4] = 53
// 0 read[0] = -1
// 0 read[2] = 94
// 1 read[4] = 53
// 1 read[4] = 53
// 0 read[3] = -1
// 0 read[1] = 18
// 1 read[3] = -1
// 1 read[2] = 94
// 0 read[0] = -1
// 0 read[2] = 94
// 1 read[1] = 18
// 1 read[1] = 18
// 0 read[2] = 94
// 0 read[0] = -1
// 1 read[2] = 94
// 1 read[2] = 94
// 0 read[4] = 53
// 0 read[4] = 53
// 1 read[1] = 18
// 1 read[1] = 18
// 0 read[0] = -1
// 0 read[1] = 18
// 1 read[3] = -1
// 1 read[3] = -1
// 0 read[1] = 18
// 0 read[4] = 53
// 1 read[2] = 94
// 1 read[3] = -1
// 0 read[3] = -1
// 0 read[3] = -1
// 1 read[3] = -1
// 1 read[4] = 53
// 0 read[4] = 53
// 1 read[4] = 53
// 0 read[0] = -1
// 1 read[2] = 94
// 0 read[1] = 18
// 0 read[1] = 18
// 1 read[2] = 94
// 1 read[4] = 53
// 0 read[3] = -1
// 0 read[4] = 53
// 1 read[1] = 18
// 0 read[4] = 53
// 1 read[3] = -1
// 1 read[2] = 94
// 0 read[2] = 94
// 0 read[4] = 53
// 1 read[4] = 53
// 1 read[1] = 18
// 0 read[1] = 18
// 1 read[3] = -1
// 0 read[3] = -1
// 0 read[1] = 18
// 1 read[1] = 18
// reads=nil
// ................main timeout 2021-09-25 13:35:39.137662 +0900 JST m=+15.988508377
// ....................main timeout 2021-09-25 13:35:40.138488 +0900 JST m=+16.989305121
// ...................main timeout 2021-09-25 13:35:41.143821 +0900 JST m=+17.994608535
// .reads=reads2
// ..................main timeout 2021-09-25 13:35:42.145028 +0900 JST m=+18.995785689
// 1 read[4] = 53
// 0 read[1] = 18
// 0 read[4] = 53
// 1 read[4] = 53
// 1 read[0] = -1
// 0 read[4] = 53
// 0 read[4] = 53
// 1 read[2] = 94
// 1 read[4] = 53
// 0 read[2] = 94
// 0 read[3] = -1
// 1 read[4] = 53
// 0 read[0] = -1
// 1 read[0] = -1
// 1 read[4] = 53
// 0 read[4] = 53
// 0 read[0] = -1
// 1 read[4] = 53
// 0 read[1] = 18
// 1 read[1] = 18
// 1 read[1] = 18
// 0 read[0] = -1
// 0 read[4] = 53
// 1 read[3] = -1
// 0 read[0] = -1
// 1 read[2] = 94
// 1 read[2] = 94
// 0 read[0] = -1
// 0 read[2] = 94
// 1 read[2] = 94
// 1 read[3] = -1
// 0 read[2] = 94
// 0 read[4] = 53
// 1 read[2] = 94
// 1 read[0] = -1
// 0 read[0] = -1
// 1 read[1] = 18
// 0 read[4] = 53
// 0 read[4] = 53
// 1 read[3] = -1
// 1 read[4] = 53
// 0 read[2] = 94
// 1 read[3] = -1
// 0 read[4] = 53
// 0 read[0] = -1
// 1 read[0] = -1
// 1 read[0] = -1
// 0 read[4] = 53
// 0 read[3] = -1
// 1 read[0] = -1
// 1 read[4] = 53
// 0 read[2] = 94
// 0 read[0] = -1
// 1 read[1] = 18
// 0 read[3] = -1
// 1 read[3] = -1
// 1 read[4] = 53
// 0 read[4] = 53
// 0 read[4] = 53
// 1 read[1] = 18
// 0 read[1] = 18
// 1 read[2] = 94
// 1 read[0] = -1
// 0 read[1] = 18
// 1 read[1] = 18
// 0 read[2] = 94
// 0 read[4] = 53
// 1 read[3] = -1
// 1 read[0] = -1
// 0 read[4] = 53
// 0 read[1] = 18
// 1 read[0] = -1
// 1 read[1] = 18
// 0 read[3] = -1
// 0 read[1] = 18
// 1 read[0] = -1
// 0 read[0] = -1
// 1 read[2] = 94
// reads=nil
// ................main timeout 2021-09-25 13:35:47.098007 +0900 JST m=+23.948620379
// ....................main timeout 2021-09-25 13:35:48.103385 +0900 JST m=+24.953968527
// ....................main timeout 2021-09-25 13:35:49.108636 +0900 JST m=+25.959189786
// ..reads=reads2
// ..................main timeout 2021-09-25 13:35:50.109374 +0900 JST m=+26.959898529
// 0 read[3] = -1
// 1 read[0] = -1
// 1 read[4] = 53
// 0 read[0] = -1
// 0 read[1] = 18
// 1 read[3] = -1
// 0 read[1] = 18
// 1 read[3] = -1
// 0 read[4] = 53
// 1 read[2] = 94
// 1 read[2] = 94
// 0 read[3] = -1
// 0 read[1] = 18
// 1 read[0] = -1
// 1 read[0] = -1
// 0 read[0] = -1
// 0 read[2] = 94
// 1 read[2] = 94
// 1 read[3] = -1
// 0 read[1] = 18
// 0 read[1] = 18
// 1 read[2] = 94
// 0 read[0] = -1
// 1 read[2] = 94
// 1 read[1] = 18
// 0 read[2] = 94
// 0 read[3] = -1
// 1 read[0] = -1
// 1 read[4] = 53
// 0 read[0] = -1
// 1 read[3] = -1
// 0 read[4] = 53
// 0 read[0] = -1
// 1 read[2] = 94
// 1 read[2] = 94
// 0 read[4] = 53
// 0 read[0] = -1
// 1 read[2] = 94
// 1 read[3] = -1
// 0 read[1] = 18
// 0 read[1] = 18
// 1 read[1] = 18
// 1 read[0] = -1
// 0 read[3] = -1
// 0 read[2] = 94
// 1 read[3] = -1
// 0 read[1] = 18
// 1 read[1] = 18
// 1 read[3] = -1
// 0 read[4] = 53
// 0 read[3] = -1
// 1 read[0] = -1
// 0 read[1] = 18
// 1 read[2] = 94
// 1 read[1] = 18
// 0 read[0] = -1
// 1 read[0] = -1
// 0 read[3] = -1
// 0 read[1] = 18
// 1 read[3] = -1
// 0 read[2] = 94
// 1 read[3] = -1
// 0 read[0] = -1
// 1 read[0] = -1
// 1 read[1] = 18
// 0 read[3] = -1
// 0 read[1] = 18
// 1 read[2] = 94
// 1 read[4] = 53
// 0 read[4] = 53
// 0 read[3] = -1
// 1 read[1] = 18
// 1 read[1] = 18
// 0 read[0] = -1
// 0 read[2] = 94
// 1 read[4] = 53
// 0 read[4] = 53
// 1 read[3] = -1
// 1 read[3] = -1
// 0 read[0] = -1
// reads=nil
// ................main timeout 2021-09-25 13:35:55.106672 +0900 JST m=+31.957050159
// ....................main timeout 2021-09-25 13:35:56.108085 +0900 JST m=+32.958433656
// ....................main timeout 2021-09-25 13:35:57.111408 +0900 JST m=+33.961727271
// ..close(reads2)
// ................main timeout 2021-09-25 13:35:58.11194 +0900 JST m=+34.962230023
// ....................main timeout 2021-09-25 13:35:59.11582 +0900 JST m=+35.966080987
// ...................main timeout 2021-09-25 13:36:00.121195 +0900 JST m=+36.971425881
// .stopr = true
// .read func exit <-------------rrr
// .read func exit <-------------rrr
// wgr.Wait() exit
// main timeout 2021-09-25 13:36:01.12303 +0900 JST m=+37.973231731
// stopm = true
// main timeout 2021-09-25 13:36:02.127213 +0900 JST m=+38.977385293
// func main end <=============
// wgm.Wait() exit
// readOps: 256
// writeOps: 8
// ============
// readOps: 256
// writeOps: 8
//  
// Compilation finished at Sat Sep 25 13:36:08
