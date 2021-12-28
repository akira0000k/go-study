package main
import (
	"log"
	"sync"
	"time"
)

/*
   subject: 終了信号はchstopを閉じることにする。wait中にchstopを見る。
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
		select {
		case <-chstop:
			return
		default:
		}
		for k := range ch {
			log.Println(k)
			select {
			case <-chstop:
				return
			case <-time.After(time.Second):
			}
		}
	}()
	time.Sleep(time.Millisecond * 3500)
	log.Println("close(chstop)...")
	close(chstop)
	log.Println("....closed")
	wg.Wait()

	log.Println("*** END main ***")
}
// -*- mode: compilation; default-directory: "~/go/src/go-study/practice/02sync/" -*-
// Compilation started at Tue Dec 28 20:16:33
//  
// go run sync16-1.go
// 2021/12/28 20:16:33 *** start main ***
// 2021/12/28 20:16:33 1
// 2021/12/28 20:16:34 2
// 2021/12/28 20:16:35 3
// 2021/12/28 20:16:36 4
// 2021/12/28 20:16:37 close(chstop)...
// 2021/12/28 20:16:37 ....closed
// 2021/12/28 20:16:37 *** END main ***
//  
// Compilation finished at Tue Dec 28 20:16:37
