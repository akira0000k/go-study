package main
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
   subject: close channel broadcast. stop wk, stop main.
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

	var readbuf int = 10
	var writebuf int = 0
	
	var stopr = false
	var stopw = false
	var stopm = false

	var reads, reads2 chan readOp = nil, make(chan readOp, readbuf)
	var writes, writes2 chan writeOp = nil, make(chan writeOp, writebuf)
	// reads := make(chan readOp)
	// writes := make(chan writeOp)
	mreads := reads2
	mwrites := writes2
	stopchan := make(chan bool)
	stopchan2 := make(chan bool)

	var wgr sync.WaitGroup
	var wgw sync.WaitGroup
	var wgm sync.WaitGroup

	wgm.Add(1)
	go func() {
		defer wgm.Done()
		var state = make(map[int]int)
	returnloop:
		for i:=0; i<10000; i++ {
			if stopm {
				break returnloop
			}
			select {
			case flag, ok := <-stopchan2:
				if !ok {
					fmt.Println("main<-stopchan2")
					break returnloop
				}
				fmt.Println("stopchan2=", flag)
			case read, ok := <-mreads:
				// closed: receive {0, nil}
				if !ok {
					fmt.Println("mreads end") //nilを入れて無視される前にこちらに来ることがある
					time.Sleep(time.Second * 1)
					//nilを入れないと、ここに無限に来る
					break
				}
				val, ok := state[read.key]
				if !ok {
					val = -1
				}
				select {
				case read.resp <- val:
				case flag, ok := <-stopchan2:
					if !ok {
						fmt.Println("main<-stopchan2(rret)")
						break returnloop
					}
					fmt.Println("stopchan2=", flag)
				}
			case write, ok := <-mwrites:
				if !ok {
					fmt.Println("mwrites end") //nilを入れて無視される前にこちらに来ることがある
					time.Sleep(time.Second * 1)
					//nilを入れないと、ここに無限に来る
					break
				}
				state[write.key] = write.val
				select {
				case write.resp <- true:
				case flag, ok := <-stopchan2:
					if !ok {
						fmt.Println("main<-stopchan2(wret)")
						break returnloop
					}
					fmt.Println("stopchan2=", flag)
				}
			case to := <-time.After(2 * time.Second):
				// これがないとdeadlockとか出る
				// mreads=nil write=nil のときはここに来る
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
				case flag, ok := <-stopchan:
					if !ok {
						fmt.Printf("read(%d)<-stopchan\n", wk)
						return
					}
					fmt.Println("rd stopchan=", flag)
				case reads <- read:
					select {
					case val, ok := <-read.resp:
						fmt.Printf("%d read[%d] = %d %v\n", wk, read.key, val, ok)
						atomic.AddUint64(&readOps, 1)
						time.Sleep(readintv)
					//case _, ok := <-stopchan:
					// 	if !ok {
					// 		fmt.Printf("read(%d)<-stopchan(rcv)\n", wk)
					// 		return
					// 	}
					}
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
				case flag, ok := <-stopchan:
					if !ok {
						fmt.Printf("write(%d)<-stopchan\n", wk)
						return
					}
					fmt.Println("ww stopchan=", flag)
				case writes <- write:
					select {
					case res, ok := <-write.resp:
						fmt.Printf("%d write[%d]:%d response %v %v\n",
							wk, write.key, write.val, res, ok)
						atomic.AddUint64(&writeOps, 1)
						time.Sleep(writeintv)
					//case _, ok := <-stopchan:
					// 	if !ok {
					// 		fmt.Printf("write(%d)<-stopchan(rcv)\n", wk)
					// 		return
					// 	}
					}
				case <-time.After(writeintv):
					fmt.Print(",")
				}
			}
			fmt.Println("write func exit <--------------www")
		}(w)
	}
	// // stopchan broadcast test=NG
	// time.Sleep(time.Second * 4)
	// fmt.Println("stopchan<-true");stopchan<-true //届くのはひとつ
	// fmt.Println("stopchan<-false");stopchan<-false //届くのはひとつ
	// // stopchan<-true
	// // stopchan<-false
	// // ww stopchan= false
	// // ww stopchan= true

	time.Sleep(time.Second * 2)
	
	fmt.Println("writes=writes2");writes=writes2
	fmt.Println("reads=reads2");reads=reads2

	time.Sleep(time.Second * 2)

	fmt.Println("mreads=nil");mreads=nil
	
	time.Sleep(time.Second * 4)

	fmt.Println("mreads=reads2");mreads=reads2

	time.Sleep(time.Second * 4)

	//ALL CLOSE broadcast
	fmt.Println("close(stopchan)"); close(stopchan)
	wgw.Wait(); fmt.Println("wgw.Wait() exit")
	wgr.Wait(); fmt.Println("wgr.Wait() exit")
	
	fmt.Println("close(stopchan2)"); close(stopchan2)
	wgm.Wait(); fmt.Println("wgm.Wait() exit")
	// close(stopchan)
	// write(0)<-stopchan
	// main<-stopchan
	// write(1)<-stopchan
	// wgw.Wait() exit
	// read(1)<-stopchan
	// read(0)<-stopchan
	// wgr.Wait() exit
	// wgm.Wait() exit
	
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
// Compilation started at Fri Sep 24 18:35:49
//  
// go run sync13-8-4.go 
// ..................,,....................writes=writes2
// reads=reads2
// main timeout 2021-09-24 18:35:52.396302 +0900 JST m=+2.002560891
// ,,1 write[2]:21 response true true
// 0 write[4]:53 response true true
// .0 read[4] = 53 true
// .1 read[4] = 53 true
// 1 read[0] = -1 true
// 0 read[0] = -1 true
// 0 read[3] = -1 true
// 1 read[3] = -1 true
// 1 read[3] = -1 true
// 0 read[0] = -1 true
// 0 read[1] = -1 true
// 1 read[0] = -1 true
// 1 read[1] = -1 true
// 0 read[0] = -1 true
// 0 read[1] = -1 true
// 1 read[3] = -1 true
// 1 read[1] = -1 true
// 0 read[2] = 21 true
// 0 read[3] = -1 true
// 1 read[1] = -1 true
// 1 read[3] = -1 true
// 0 read[1] = -1 true
// 0 write[2]:18 response true true
// 1 write[2]:94 response true true
// 0 read[2] = 94 true
// 1 read[3] = -1 true
// 1 read[1] = -1 true
// 0 read[0] = -1 true
// 1 read[3] = -1 true
// 0 read[3] = -1 true
// 0 read[2] = 94 true
// 1 read[3] = -1 true
// 1 read[1] = -1 true
// 0 read[4] = 53 true
// 0 read[3] = -1 true
// 1 read[3] = -1 true
// 0 read[2] = 94 true
// 1 read[1] = -1 true
// 1 read[3] = -1 true
// 0 read[1] = -1 true
// 0 read[1] = -1 true
// 1 read[2] = 94 true
// mreads=nil
// 1 write[0]:3 response true true
// 0 write[2]:43 response true true
// 0 write[0]:51 response true true
// 1 write[0]:57 response true true
// 1 write[2]:10 response true true
// 0 write[0]:85 response true true
// 0 write[0]:32 response true true
// 1 write[3]:53 response true true
// mreads=reads2
// 0 write[4]:97 response true true
// 1 read[0] = 32 true
// 1 write[1]:82 response true true
// 0 read[3] = 53 true
// 0 read[2] = 10 true
// 1 read[2] = 10 true
// 0 read[1] = 82 true
// 1 read[4] = 97 true
// 0 read[2] = 10 true
// 1 read[1] = 82 true
// 1 read[1] = 82 true
// 0 read[4] = 97 true
// 0 read[1] = 82 true
// 1 read[0] = 32 true
// 0 read[3] = 53 true
// 1 read[1] = 82 true
// 1 read[4] = 97 true
// 0 read[1] = 82 true
// 0 read[2] = 10 true
// 1 read[0] = 32 true
// 1 read[0] = 32 true
// 0 read[0] = 32 true
// 1 write[2]:49 response true true
// 0 write[3]:18 response true true
// 0 read[4] = 97 true
// 1 read[3] = 18 true
// 0 read[4] = 97 true
// 1 read[2] = 49 true
// 0 read[2] = 49 true
// 1 read[2] = 49 true
// 1 read[1] = 82 true
// 0 read[4] = 97 true
// 0 read[0] = 32 true
// 1 read[1] = 82 true
// 1 read[1] = 82 true
// 0 read[1] = 82 true
// 0 read[0] = 32 true
// 1 read[1] = 82 true
// 1 read[4] = 97 true
// 0 read[4] = 97 true
// 0 read[0] = 32 true
// 1 read[3] = 18 true
// 1 read[1] = 82 true
// 0 read[0] = 32 true
// 0 write[2]:58 response true true
// 1 write[2]:31 response true true
// 0 read[3] = 18 true
// 1 read[4] = 97 true
// 1 read[2] = 31 true
// 0 read[3] = 18 true
// 0 read[2] = 31 true
// 1 read[3] = 18 true
// 1 read[3] = 18 true
// 0 read[3] = 18 true
// 0 read[1] = 82 true
// 1 read[0] = 32 true
// 1 read[0] = 32 true
// 0 read[0] = 32 true
// 0 read[4] = 97 true
// 1 read[2] = 31 true
// 0 read[0] = 32 true
// 1 read[2] = 31 true
// 1 read[1] = 82 true
// 0 read[4] = 97 true
// 1 read[0] = 32 true
// 0 read[3] = 18 true
// 0 write[0]:59 response true true
// 1 write[0]:83 response true true
// 0 read[4] = 97 true
// 1 read[0] = 83 true
// 1 read[2] = 31 true
// 0 read[0] = 83 true
// 0 read[0] = 83 true
// 1 read[2] = 31 true
// 1 read[4] = 97 true
// 0 read[0] = 83 true
// 0 read[3] = 18 true
// 1 read[1] = 82 true
// 0 read[0] = 83 true
// 1 read[1] = 82 true
// 1 read[0] = 83 true
// 0 read[1] = 82 true
// 1 read[2] = 31 true
// 0 read[4] = 97 true
// 0 read[1] = 82 true
// 1 read[2] = 31 true
// close(stopchan)
// write(1)<-stopchan
// write(0)<-stopchan
// wgw.Wait() exit
// read(0)<-stopchan
// 1 read[2] = 31 true
// 1 read[3] = 18 true
// read(1)<-stopchan
// wgr.Wait() exit
// close(stopchan2)
// main<-stopchan2
// func main end <=============
// wgm.Wait() exit
// readOps: 118
// writeOps: 20
// ============
// readOps: 118
// writeOps: 20
//  
// Compilation finished at Fri Sep 24 18:36:08
