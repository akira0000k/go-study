package main
import (
	"log"
	"sync"
	"time"
)

/*
   subject: channel close はキューを追い越せるか? 結論:追い越せない
*/
func main() {
	log.Println("*** start main ***")

	n0 := 10
	ch := make(chan int, n0)
	for i := 0; i< n0; i++ {
		ch <- i+1
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for k := range ch {
			log.Println(k)
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Millisecond * 3500)
	log.Println("close(ch)...")
	close(ch)
	log.Println("....closed")
	wg.Wait()

	log.Println("*** END main ***")
}
// -*- mode: compilation; default-directory: "~/go/src/go-study/practice/02sync/" -*-
// Compilation started at Tue Dec 28 19:57:45
//  
// go run sync16.go
// 2021/12/28 19:57:45 *** start main ***
// 2021/12/28 19:57:45 1
// 2021/12/28 19:57:46 2
// 2021/12/28 19:57:47 3
// 2021/12/28 19:57:48 4
// 2021/12/28 19:57:49 close(ch)...
// 2021/12/28 19:57:49 ....closed
// 2021/12/28 19:57:49 5
// 2021/12/28 19:57:50 6
// 2021/12/28 19:57:51 7
// 2021/12/28 19:57:52 8
// 2021/12/28 19:57:53 9
// 2021/12/28 19:57:54 10
// 2021/12/28 19:57:55 *** END main ***
//  
// Compilation finished at Tue Dec 28 19:57:55
