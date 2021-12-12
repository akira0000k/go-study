package main
import (
	"fmt"
	"math/rand"
	"time"
	"sync"
	"sync/atomic"
)
/*
   subject: Mutexes
*/

func main() {
	fmt.Println("start")
	
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}

	var read0ps uint64
	var write0ps uint64
	var total int
	
	for r:=0; r<100; r++ {
		go func() {
			//total := 0
			for {
				key := rand.Intn(5)

				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&read0ps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w:=0; w<10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				//val := rand.Intn(100)

				mutex.Lock()
				//state[key] = val
				val, ex := state[key]
				if ex {
					state[key] = val + 1
				} else {
					state[key] = 1
				}
				mutex.Unlock()
				atomic.AddUint64(&write0ps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	for pool:=0; pool<2; pool++ {
		for loop:=0; loop<2; loop++ {
			time.Sleep(time.Second)

			read0psFinal := atomic.LoadUint64(&read0ps)
			fmt.Println("read0ps:", read0psFinal)
			write0psFinal := atomic.LoadUint64(&write0ps)
			fmt.Println("write0ps:", write0psFinal)
		}

		mutex.Lock()
		fmt.Println("state:", state)
		fmt.Println("total:", total)
		mutex.Unlock()
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Fri Oct 29 19:09:10
//  
// go run sync12.go
// start
// read0ps: 73622
// write0ps: 7362
// read0ps: 147900
// write0ps: 14790
// state: map[0:2932 1:3041 2:2914 3:2964 4:2939]
// total: 218755333
// read0ps: 220400
// write0ps: 22040
// read0ps: 294800
// write0ps: 29480
// state: map[0:5859 1:6047 2:5801 3:5910 4:5863]
// total: 869090707
//  
// Compilation finished at Fri Oct 29 19:09:15
