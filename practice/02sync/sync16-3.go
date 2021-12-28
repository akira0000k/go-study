package main
import (
	"log"
	"sync"
	"time"
)

/*
   subject: 一定間隔で発生するイベントを追加 time.NewTicker
*/
func main() {
	log.Println("*** start main ***")

	n0 := 5
	chstop := make(chan bool)
	ch := make(chan int, n0)
	go func() {
		i := 0
		for ; i< n0; i++ {
			ch <- i+1
		}
		time.Sleep(10*time.Second)
		for ; i< n0*2; i++ {
			ch <- i+1
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		tic := time.NewTicker(2 * time.Second)
		defer tic.Stop()
		for {
			select {
			case k, ok := <-ch:
				if !ok {
					return
				}
				log.Println(k)
				select {
				case <-time.After(time.Second):
				case <-tic.C:
					log.Println("2sec interval2")//これが来ると抜けちゃう
				case <-chstop:
					return
				}
			case <-tic.C:
				log.Println("2sec interval")
			case <-chstop:
				return
			}
		}
	}()
	//time.Sleep(time.Millisecond * 3500)
	time.Sleep(time.Millisecond * 13500)
	log.Println("close(chstop)...")
	//close(ch)
	close(chstop)
	log.Println("....closed")
	wg.Wait()

	log.Println("*** END main ***")
}
// -*- mode: compilation; default-directory: "~/go/src/go-study/practice/02sync/" -*-
// Compilation started at Tue Dec 28 21:22:54
//  
// go run sync16-3.go
// 2021/12/28 21:22:54 *** start main ***
// 2021/12/28 21:22:54 1
// 2021/12/28 21:22:55 2
// 2021/12/28 21:22:56 2sec interval2
// 2021/12/28 21:22:56 3
// 2021/12/28 21:22:57 4
// 2021/12/28 21:22:58 5
// 2021/12/28 21:22:58 2sec interval2
// 2021/12/28 21:23:00 2sec interval
// 2021/12/28 21:23:02 2sec interval
// 2021/12/28 21:23:04 6
// 2021/12/28 21:23:04 2sec interval2
// 2021/12/28 21:23:04 7
// 2021/12/28 21:23:05 8
// 2021/12/28 21:23:06 2sec interval2
// 2021/12/28 21:23:06 9
// 2021/12/28 21:23:07 10
// 2021/12/28 21:23:08 close(chstop)...
// 2021/12/28 21:23:08 ....closed
// 2021/12/28 21:23:08 *** END main ***
//  
// Compilation finished at Tue Dec 28 21:23:08
