package main
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
   subject: close channel broadcast. set flag and break select (from sync13-8.go)
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
	const mainintv = 2 * time.Second//time.Millisecond / 10//
	
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

	var wgr sync.WaitGroup
	var wgw sync.WaitGroup
	var wgm sync.WaitGroup

	wgm.Add(1)
	go func() {
		defer wgm.Done()
		var state = make(map[int]int)
		var chanstop bool=false
	forloop:
		for { //i:=0; i<10000; i++ {
			if stopm || chanstop {
				break forloop
			}
			select {
			case read, ok := <-mreads:
				// closed: receive {0, nil}
				if !ok {
					fmt.Println("mreads end") //nilを入れて無視される前にこちらに来ることがある
					time.Sleep(time.Second * 1)
					//nilを入れないと、ここに無限に来る
					break
				}
				val, ok := state[read.key]
				if ok {
					read.resp <- val
				} else {
					read.resp <- -1
				}
			case write, ok := <-mwrites:
				if !ok {
					fmt.Println("mwrites end") //nilを入れて無視される前にこちらに来ることがある
					time.Sleep(time.Second * 1)
					//nilを入れないと、ここに無限に来る
					break
				}
				state[write.key] = write.val
				write.resp <- true
			//case flag, ok := <-stopchan://ここだとdeadlock
			//	if !ok {
			//		chanstop = true
			//		fmt.Println("main<-stopchan; chanstop=true")
			//		//break
			//	}
			//	fmt.Println("stopchan=", flag)
			case to := <-time.After(mainintv):
				// これがないとdeadlockとか出る
				// mreads=nil write=nil のときはここに来る
				//fmt.Print("+")//"main timeout", to)
				_ = to
			case flag, ok := <-stopchan://ここに置けばいい(まぐれ)
				if !ok {
					chanstop = true //フラグを立てるだけで戻らないのが良?
					fmt.Println("main<-stopchan; chanstop=true")
					//break forloop//deadlock
					//return//deadlock
					//break//no need
				}
				fmt.Println("stopchan=", flag)
			}
			//fmt.Println("func loop", i)
		}
		fmt.Println("func main end <=============")
	}()

	for r := 0; r < nrwk; r++ {
		wgr.Add(1)
		go func(wk int) {
			defer wgr.Done()
			var chanstop bool=false
		forloop:
			for {
				if stopr || chanstop {
					break forloop
				}
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				select {
				case flag, ok := <-stopchan:
					if !ok {
						chanstop = true
						fmt.Printf("read(%d)<-stopchan; chanstop=true\n", wk)
						break
					}
					fmt.Println("rd stopchan=", flag)
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
			var chanstop bool=false
		forloop:
			for {
				if stopw || chanstop {
					break forloop
				}
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				select {
				case flag, ok := <-stopchan:
					if !ok {
						chanstop = true
						fmt.Printf("write(%d)<-stopchan; chanstop=true\n", wk)
						break
					}
					fmt.Println("ww stopchan=", flag)
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
	//stopm, stopw, stopr = true, true, true //これなら絶対安全
	fmt.Println("close(stopchan)"); close(stopchan) //deadlock気味
	wgw.Wait(); fmt.Println("wgw.Wait() exit")
	wgr.Wait(); fmt.Println("wgr.Wait() exit")
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
// Compilation started at Sun Sep 26 20:12:57
//  
// go run sync13-8-5.go
// ..................,,....................writes=writes2
// reads=reads2
// ,,1 write[4]:53 response true
// 0 write[2]:21 response true
// ..1 read[4] = 53
// 0 read[4] = 53
// 1 read[0] = -1
// 0 read[0] = -1
// 0 read[3] = -1
// 1 read[3] = -1
// 1 read[3] = -1
// 0 read[0] = -1
// 1 read[0] = -1
// 0 read[1] = -1
// 0 read[0] = -1
// 1 read[1] = -1
// 0 read[3] = -1
// 1 read[1] = -1
// 1 read[1] = -1
// 0 read[2] = 21
// 0 read[3] = -1
// 1 read[1] = -1
// 1 write[1]:2 response true
// 1 read[3] = -1
// 0 write[2]:94 response true
// 0 read[3] = -1
// 0 read[2] = 94
// 1 read[3] = -1
// 1 read[1] = 2
// 0 read[0] = -1
// 0 read[3] = -1
// 1 read[3] = -1
// 1 read[2] = 94
// 0 read[3] = -1
// 0 read[1] = 2
// 1 read[4] = 53
// 1 read[3] = -1
// 0 read[3] = -1
// 0 read[1] = 2
// 1 read[2] = 94
// 0 read[3] = -1
// 1 read[1] = 2
// 1 read[1] = 2
// 0 read[2] = 94
// mreads=nil
// 1 write[0]:3 response true
// 0 write[2]:43 response true
// 0 write[0]:51 response true
// 1 write[0]:57 response true
// 1 write[0]:85 response true
// 0 write[2]:10 response true
// 0 write[0]:32 response true
// 1 write[3]:53 response true
// mreads=reads2
// 0 read[3] = 53
// 0 write[1]:82 response true
// 1 read[0] = 32
// 1 write[4]:97 response true
// 1 read[2] = 10
// 0 read[2] = 10
// 0 read[1] = 82
// 1 read[4] = 97
// 1 read[1] = 82
// 0 read[2] = 10
// 0 read[1] = 82
// 1 read[4] = 97
// 1 read[1] = 82
// 0 read[0] = 32
// 0 read[3] = 53
// 1 read[1] = 82
// 1 read[4] = 97
// 0 read[1] = 82
// 0 read[2] = 10
// 1 read[0] = 32
// 0 read[0] = 32
// 1 read[0] = 32
// 0 write[2]:49 response true
// 1 write[3]:18 response true
// 0 read[4] = 97
// 1 read[3] = 18
// 1 read[4] = 97
// 0 read[2] = 49
// 0 read[2] = 49
// 1 read[2] = 49
// 1 read[1] = 82
// 0 read[4] = 97
// 0 read[0] = 32
// 1 read[1] = 82
// 1 read[1] = 82
// 0 read[1] = 82
// 0 read[0] = 32
// 1 read[1] = 82
// 1 read[4] = 97
// 0 read[4] = 97
// 0 read[0] = 32
// 1 read[3] = 18
// 1 read[1] = 82
// 0 read[0] = 32
// 0 write[2]:58 response true
// 1 write[2]:31 response true
// 0 read[3] = 18
// 1 read[4] = 97
// 1 read[2] = 31
// 0 read[3] = 18
// 0 read[2] = 31
// 1 read[3] = 18
// 1 read[3] = 18
// 0 read[3] = 18
// 1 read[0] = 32
// 0 read[1] = 82
// 0 read[0] = 32
// 1 read[0] = 32
// 0 read[4] = 97
// 1 read[2] = 31
// 1 read[0] = 32
// 0 read[2] = 31
// 0 read[1] = 82
// 1 read[4] = 97
// 1 read[0] = 32
// 0 read[3] = 18
// 0 write[0]:59 response true
// 1 write[0]:83 response true
// 0 read[0] = 83
// 1 read[4] = 97
// 1 read[2] = 31
// 0 read[0] = 83
// 1 read[2] = 31
// 0 read[0] = 83
// 1 read[0] = 83
// 0 read[4] = 97
// 0 read[3] = 18
// 1 read[1] = 82
// 1 read[0] = 83
// 0 read[1] = 82
// 0 read[0] = 83
// 1 read[1] = 82
// 1 read[4] = 97
// 0 read[2] = 31
// 0 read[1] = 82
// 1 read[2] = 31
// close(stopchan)
// main<-stopchan; chanstop=true
// stopchan= false
// func main end <=============
// write(1)<-stopchan; chanstop=true
// write func exit <--------------www
// write(0)<-stopchan; chanstop=true
// write func exit <--------------www
// wgw.Wait() exit
// read(1)<-stopchan; chanstop=true
// read func exit <-------------rrr
// read(0)<-stopchan; chanstop=true
// read func exit <-------------rrr
// wgr.Wait() exit
// wgm.Wait() exit
// readOps: 116
// writeOps: 20
// ============
// readOps: 116
// writeOps: 20
//  
// Compilation finished at Sun Sep 26 20:13:16
