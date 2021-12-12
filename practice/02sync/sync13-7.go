package main
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
   subject: channel close signal gets infinitly
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
	go func(reads chan readOp) {
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
					time.Sleep(time.Second * 1)
					//引数にしてnilを入れられないので、ここに無限に来る
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
	}(reads)

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
// Compilation started at Tue Sep 21 23:20:48
//  
// go run sync13-7.go
// reads=nil
// 1 write[4]:81 response true
// 0 write[2]:47 response true
// ..................1 write[2]:47 response true
// 0 write[2]:88 response true
// ....................main timeout 2021-09-21 23:20:51.241964 +0900 JST m=+2.010582867
// 0 write[4]:53 response true
// 1 write[2]:21 response true
// ....................1 write[2]:18 response true
// 0 write[2]:94 response true
// ..................writes=nil
// main timeout 2021-09-21 23:20:53.243515 +0900 JST m=+4.012124661
// ....................,,main timeout 2021-09-21 23:20:54.24862 +0900 JST m=+5.017224720
// ....................main timeout 2021-09-21 23:20:55.252372 +0900 JST m=+6.020971981
// ,,....................,,main timeout 2021-09-21 23:20:56.257697 +0900 JST m=+7.026292587
// ....................main timeout 2021-09-21 23:20:57.259691 +0900 JST m=+8.028282302
// ,,..................stopw = true
// ,write func exit <--------------www
// ,write func exit <--------------www
// wgw.Wait() exit
// close(writes2)
// main timeout 2021-09-21 23:20:58.264405 +0900 JST m=+9.032990757
// ....................reads=reads2
// main timeout 2021-09-21 23:20:59.269731 +0900 JST m=+10.038312875
// ..0 read[2] = 94
// 1 read[4] = 53
// 1 read[1] = -1
// 0 read[0] = -1
// 0 read[0] = -1
// 1 read[4] = 53
// 1 read[4] = 53
// 0 read[4] = 53
// 1 read[0] = -1
// 0 read[1] = -1
// 0 read[1] = -1
// 1 read[2] = 94
// 1 read[1] = -1
// 0 read[0] = -1
// 1 read[4] = 53
// 0 read[2] = 94
// 0 read[3] = -1
// 1 read[1] = -1
// 1 read[4] = 53
// 0 read[0] = -1
// 1 read[4] = 53
// 0 read[0] = -1
// 0 read[4] = 53
// 1 read[4] = 53
// 0 read[0] = -1
// 1 read[2] = 94
// 1 read[0] = -1
// 0 read[0] = -1
// 0 read[3] = -1
// 1 read[4] = 53
// 1 read[0] = -1
// 0 read[3] = -1
// 1 read[1] = -1
// 0 read[0] = -1
// 0 read[4] = 53
// 1 read[4] = 53
// 1 read[0] = -1
// 0 read[1] = -1
// 1 read[2] = 94
// 0 read[3] = -1
// 0 read[4] = 53
// 1 read[0] = -1
// 1 read[4] = 53
// 0 read[2] = 94
// 0 read[4] = 53
// 1 read[3] = -1
// 0 read[3] = -1
// 1 read[1] = -1
// 1 read[2] = 94
// 0 read[0] = -1
// 0 read[2] = 94
// 1 read[1] = -1
// 1 read[2] = 94
// 0 read[1] = -1
// 0 read[0] = -1
// 1 read[2] = 94
// 1 read[4] = 53
// 0 read[2] = 94
// 0 read[4] = 53
// 1 read[1] = -1
// 1 read[1] = -1
// 0 read[0] = -1
// 0 read[1] = -1
// 1 read[3] = -1
// 0 read[3] = -1
// 1 read[1] = -1
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
// 0 read[1] = -1
// 1 read[2] = 94
// 0 read[1] = -1
// 0 read[4] = 53
// 1 read[3] = -1
// 1 read[4] = 53
// 0 read[1] = -1
// 0 read[3] = -1
// 1 read[4] = 53
// 1 read[2] = 94
// 0 read[2] = 94
// 1 read[4] = 53
// 0 read[4] = 53
// 0 read[1] = -1
// 1 read[1] = -1
// 1 read[3] = -1
// 0 read[3] = -1
// 0 read[1] = -1
// 1 read[1] = -1
// 1 read[2] = 94
// 0 read[2] = 94
// 0 read[1] = -1
// 1 read[1] = -1
// reads=nil
// ..................main timeout 2021-09-21 23:21:05.265729 +0900 JST m=+16.034282551
// ..................main timeout 2021-09-21 23:21:06.267657 +0900 JST m=+17.036205554
// ....................reads=reads2
// main timeout 2021-09-21 23:21:07.271038 +0900 JST m=+18.039582133
// ..1 read[1] = -1
// 0 read[0] = -1
// 0 read[2] = 94
// 1 read[2] = 94
// 0 read[1] = -1
// 1 read[4] = 53
// 1 read[3] = -1
// 0 read[3] = -1
// 0 read[1] = -1
// 1 read[0] = -1
// 1 read[2] = 94
// 0 read[0] = -1
// 1 read[1] = -1
// 0 read[4] = 53
// 0 read[4] = 53
// 1 read[4] = 53
// 1 read[0] = -1
// 0 read[4] = 53
// 0 read[4] = 53
// 1 read[2] = 94
// 0 read[4] = 53
// 1 read[2] = 94
// 1 read[4] = 53
// 0 read[3] = -1
// 0 read[0] = -1
// 1 read[0] = -1
// 0 read[4] = 53
// 1 read[4] = 53
// 1 read[0] = -1
// 0 read[4] = 53
// 1 read[1] = -1
// 0 read[1] = -1
// 0 read[1] = -1
// 1 read[0] = -1
// 1 read[4] = 53
// 0 read[3] = -1
// 1 read[2] = 94
// 0 read[0] = -1
// 0 read[2] = 94
// 1 read[0] = -1
// 0 read[2] = 94
// 1 read[2] = 94
// 0 read[2] = 94
// 1 read[3] = -1
// 1 read[4] = 53
// 0 read[2] = 94
// 1 read[0] = -1
// 0 read[0] = -1
// 0 read[1] = -1
// 1 read[4] = 53
// 1 read[4] = 53
// 0 read[3] = -1
// 0 read[4] = 53
// 1 read[2] = 94
// 1 read[3] = -1
// 0 read[4] = 53
// 0 read[0] = -1
// 1 read[0] = -1
// 0 read[4] = 53
// 1 read[0] = -1
// 1 read[3] = -1
// 0 read[0] = -1
// 1 read[2] = 94
// 0 read[4] = 53
// 0 read[0] = -1
// 1 read[1] = -1
// 1 read[3] = -1
// 0 read[3] = -1
// 0 read[4] = 53
// 1 read[4] = 53
// 1 read[4] = 53
// 0 read[1] = -1
// 0 read[2] = 94
// 1 read[1] = -1
// 1 read[0] = -1
// 0 read[1] = -1
// 1 read[1] = -1
// 0 read[2] = 94
// 0 read[4] = 53
// 1 read[3] = -1
// 1 read[0] = -1
// 0 read[4] = 53
// 0 read[1] = -1
// 1 read[0] = -1
// 1 read[1] = -1
// 0 read[3] = -1
// 0 read[1] = -1
// 1 read[0] = -1
// 1 read[0] = -1
// 0 read[2] = 94
// 0 read[3] = -1
// 1 read[4] = 53
// 0 read[0] = -1
// 1 read[4] = 53
// 1 read[1] = -1
// 0 read[1] = -1
// reads=nil
// ................main timeout 2021-09-21 23:21:13.219699 +0900 JST m=+23.988214787
// ....................main timeout 2021-09-21 23:21:14.221077 +0900 JST m=+24.989588414
// ....................main timeout 2021-09-21 23:21:15.225336 +0900 JST m=+25.993841876
// reads=reads2
// .0 read[1] = -1
// .1 read[4] = 53
// 1 read[4] = 53
// 0 read[1] = -1
// 0 read[4] = 53
// 1 read[4] = 53
// 1 read[3] = -1
// 0 read[2] = 94
// 0 read[4] = 53
// 1 read[4] = 53
// 0 read[0] = -1
// 1 read[4] = 53
// 0 read[3] = -1
// 1 read[0] = -1
// 1 read[0] = -1
// 0 read[4] = 53
// 0 read[1] = -1
// 1 read[3] = -1
// 0 read[1] = -1
// 1 read[3] = -1
// 1 read[2] = 94
// 0 read[4] = 53
// 0 read[2] = 94
// 1 read[3] = -1
// 1 read[1] = -1
// 0 read[0] = -1
// 0 read[0] = -1
// 1 read[0] = -1
// 1 read[2] = 94
// 0 read[2] = 94
// 0 read[1] = -1
// 1 read[3] = -1
// 1 read[2] = 94
// 0 read[1] = -1
// 0 read[2] = 94
// 1 read[0] = -1
// 1 read[1] = -1
// 0 read[2] = 94
// 1 read[0] = -1
// 0 read[3] = -1
// 1 read[0] = -1
// 0 read[4] = 53
// 0 read[4] = 53
// 1 read[3] = -1
// 1 read[0] = -1
// 0 read[2] = 94
// 0 read[2] = 94
// 1 read[4] = 53
// 0 read[2] = 94
// 1 read[0] = -1
// 0 read[3] = -1
// 1 read[1] = -1
// 1 read[1] = -1
// 0 read[1] = -1
// 0 read[0] = -1
// 1 read[3] = -1
// 1 read[2] = 94
// 0 read[3] = -1
// 0 read[1] = -1
// 1 read[1] = -1
// 1 read[3] = -1
// 0 read[4] = 53
// 1 read[0] = -1
// 0 read[3] = -1
// 1 read[1] = -1
// 0 read[2] = 94
// 0 read[1] = -1
// 1 read[0] = -1
// 0 read[3] = -1
// 1 read[0] = -1
// 1 read[1] = -1
// 0 read[3] = -1
// 0 read[3] = -1
// 1 read[2] = 94
// 0 read[0] = -1
// 1 read[0] = -1
// 0 read[3] = -1
// 1 read[1] = -1
// 1 read[1] = -1
// 0 read[2] = 94
// 0 read[4] = 53
// 1 read[4] = 53
// 1 read[3] = -1
// 0 read[1] = -1
// 1 read[1] = -1
// 0 read[0] = -1
// 0 read[2] = 94
// 1 read[4] = 53
// 1 read[4] = 53
// 0 read[3] = -1
// 1 read[0] = -1
// 0 read[3] = -1
// 0 read[3] = -1
// 1 read[0] = -1
// 1 read[0] = -1
// 0 read[2] = 94
// 0 read[0] = -1
// 1 read[1] = -1
// reads=nil
// ................main timeout 2021-09-21 23:21:21.262841 +0900 JST m=+32.031319024
// ....................main timeout 2021-09-21 23:21:22.266931 +0900 JST m=+33.035403881
// ..................main timeout 2021-09-21 23:21:23.269309 +0900 JST m=+34.037777380
// ..close(reads2)
// reads end
// ..................reads end
// ....................reads end
// ...................stopr = true
// .read func exit <-------------rrr
// reads end
// .read func exit <-------------rrr
// wgr.Wait() exit
// reads end
// stopm = true
// func main end <=============
// wgm.Wait() exit
// readOps: 294
// writeOps: 8
// ============
// readOps: 294
// writeOps: 8
//  
// Compilation finished at Tue Sep 21 23:21:34
