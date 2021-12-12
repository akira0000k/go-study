package main
import (
	"fmt"
	"time"
)
/*
   subject: Tickers
*/

func main() {
	//_ = time.Second

	fmt.Println("start")

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println(time.Now(), "return func")
				return
			case t := <-ticker.C:
				fmt.Println(t, "Tick at")
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	fmt.Println(time.Now(), "Ticker stop")
	ticker.Stop()

	time.Sleep(2*time.Second)
	
	fmt.Println(time.Now(), "send done")
	done<-true

	time.Sleep(time.Second)
	fmt.Println("end")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Fri Oct 29 18:40:14
//  
// go run sync07.go
// start
// 2021-10-29 18:40:15.137698 +0900 JST m=+0.505306342 Tick at
// 2021-10-29 18:40:15.637292 +0900 JST m=+1.004884865 Tick at
// 2021-10-29 18:40:16.137146 +0900 JST m=+1.504723479 Tick at
// 2021-10-29 18:40:16.232717 +0900 JST m=+1.600292205 Ticker stop
// 2021-10-29 18:40:18.236577 +0900 JST m=+3.604091542 send done
// 2021-10-29 18:40:18.236612 +0900 JST m=+3.604126592 return func
// end
//  
// Compilation finished at Fri Oct 29 18:40:19
