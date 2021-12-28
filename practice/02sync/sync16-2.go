package main
import (
	"log"
	"sync"
	"time"
)

/*
   subject: 受信待ちと待機中の両方でchstopを見る。
*/
func main() {
	log.Println("*** start main ***")

	n0 := 10
	chstop := make(chan bool)
	ch := make(chan int, n0)
	for i := 0; i< n0; i++ {
		ch <- i+1
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case k, ok := <-ch:
				if !ok {
					return
				}
				log.Println(k)
			case <-chstop:
				return
			}

			select {
			case <-time.After(time.Second):
			case <-chstop:
				return
			}
		}
	}()
	time.Sleep(time.Millisecond * 3500)
	//time.Sleep(time.Millisecond * 13500)
	log.Println("close(chstop)...")
	//close(ch)
	close(chstop)
	log.Println("....closed")
	wg.Wait()

	log.Println("*** END main ***")
}
// *- mode: compilation; default-directory: "~/go/src/go-study/practice/02sync/" -*-
// ompilation started at Tue Dec 28 20:48:15
// 
// o run sync16-2.go
// 021/12/28 20:48:15 *** start main ***
// 021/12/28 20:48:15 1
// 021/12/28 20:48:16 2
// 021/12/28 20:48:17 3
// 021/12/28 20:48:18 4
// 021/12/28 20:48:19 close(chstop)...
// 021/12/28 20:48:19 ....closed
// 021/12/28 20:48:19 *** END main ***
// 
// ompilation finished at Tue Dec 28 20:48:19

// -*- mode: compilation; default-directory: "~/go/src/go-study/practice/02sync/" -*-
// Compilation started at Tue Dec 28 20:45:29
//  
// go run sync16-2.go
// 2021/12/28 20:45:30 *** start main ***
// 2021/12/28 20:45:30 1
// 2021/12/28 20:45:31 2
// 2021/12/28 20:45:32 3
// 2021/12/28 20:45:33 4
// 2021/12/28 20:45:34 5
// 2021/12/28 20:45:35 6
// 2021/12/28 20:45:36 7
// 2021/12/28 20:45:37 8
// 2021/12/28 20:45:38 9
// 2021/12/28 20:45:39 10
// 2021/12/28 20:45:44 close(chstop)...
// 2021/12/28 20:45:44 ....closed
// 2021/12/28 20:45:44 *** END main ***
//  
// Compilation finished at Tue Dec 28 20:45:44
