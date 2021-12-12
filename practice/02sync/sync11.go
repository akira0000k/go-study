package main
import (
	"fmt"
	//"time"
	"sync"
	"sync/atomic"
)
/*
   subject: Atomic AddUint64
*/

func main() {
	fmt.Println("start")
	
	var simpleCounter, c2, c3 uint64
	var ops uint64

	var wg sync.WaitGroup

	for i:=0; i<50; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for c:=0; c<1000; c++ {
				simpleCounter++
				c2++
				c3++
				atomic.AddUint64(&ops, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("ops:", ops)
	fmt.Println("cnt:", simpleCounter)
	fmt.Println("cnt:", c2)
	fmt.Println("cnt:", c3)
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Fri Oct 29 19:08:16
//  
// go run sync11.go
// start
// ops: 50000
// cnt: 32845
// cnt: 33190
// cnt: 33618
//  
// Compilation finished at Fri Oct 29 19:08:16
