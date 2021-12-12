package main
import (
	"fmt"
	"time"
	"sync"
)
/*
   subject: Rate Limiting
*/

func main() {
	fmt.Println("start")
	const msec = time.Millisecond
	const okburst = 3
	const reqmax = 100 //別になくても良さげ
	requests := make(chan int, reqmax)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		bcount := 0
	forloop:
		for {
			select {
			case req := <-requests:
				if req == 0 { //closeすると0が来る
					fmt.Println("request", bcount, req, time.Now())
					break forloop
				}
				if bcount >= okburst {
					time.Sleep(msec * 200)
				} else {
					bcount++
				}
				fmt.Println("request", bcount, req, time.Now())
			case <-time.After(msec * 200):
				bcount = 0
			}
		}
		fmt.Println("End receiving")
	}()
	fmt.Println("---------------------------")
	reqid := 0
	for i:=0; i< 5; i++ {
		reqid++
		requests<-reqid
		time.Sleep(msec * 500)
	}

	for i:=0; i< 10; i++ {
		reqid++
		requests<-reqid
	}

	for i:=0; i< 20; i++ {
		reqid++
		requests<-reqid
		time.Sleep(msec * 500)
	}
	for i:=0; i< 5; i++ {
		reqid++
		requests<-reqid
	}
	for i:=0; i< 10; i++ {
		reqid++
		requests<-reqid
		//time.Sleep(msec * 500)
	}
	fmt.Println("=========================")
	//requests<- 0 //明示的に終わりを指示したり
	close(requests)//なくてもエラーにはならないが終わらない

	wg.Wait()
	fmt.Println("receive ended")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/022burst/" -*-
// Compilation started at Wed Oct  6 19:01:05
//  
// go run sync10-2.go
// start
// ---------------------------
// request 1 1 2021-10-06 19:01:06.234442 +0900 JST m=+0.000309829
// request 1 2 2021-10-06 19:01:06.739454 +0900 JST m=+0.505339315
// request 1 3 2021-10-06 19:01:07.242838 +0900 JST m=+1.008740684
// request 1 4 2021-10-06 19:01:07.747828 +0900 JST m=+1.513748013
// request 1 5 2021-10-06 19:01:08.25301 +0900 JST m=+2.018947413
// request 1 6 2021-10-06 19:01:08.758201 +0900 JST m=+2.524155188
// request 2 7 2021-10-06 19:01:08.758276 +0900 JST m=+2.524229525
// request 3 8 2021-10-06 19:01:08.758359 +0900 JST m=+2.524313366
// request 3 9 2021-10-06 19:01:08.962866 +0900 JST m=+2.728826624
// request 3 10 2021-10-06 19:01:09.168142 +0900 JST m=+2.934110320
// request 3 11 2021-10-06 19:01:09.373417 +0900 JST m=+3.139391565
// request 3 12 2021-10-06 19:01:09.577115 +0900 JST m=+3.343096672
// request 3 13 2021-10-06 19:01:09.78079 +0900 JST m=+3.546779269
// request 3 14 2021-10-06 19:01:09.986034 +0900 JST m=+3.752030410
// request 3 15 2021-10-06 19:01:10.191286 +0900 JST m=+3.957289059
// request 3 16 2021-10-06 19:01:10.396572 +0900 JST m=+4.162582025
// request 3 17 2021-10-06 19:01:10.596872 +0900 JST m=+4.362889175
// request 3 18 2021-10-06 19:01:10.798729 +0900 JST m=+4.564753127
// request 3 19 2021-10-06 19:01:11.001557 +0900 JST m=+4.767587612
// request 3 20 2021-10-06 19:01:11.206788 +0900 JST m=+4.972825465
// request 3 21 2021-10-06 19:01:11.481028 +0900 JST m=+5.247075368
// request 1 22 2021-10-06 19:01:11.781015 +0900 JST m=+5.547072684
// request 1 23 2021-10-06 19:01:12.286187 +0900 JST m=+6.052261465
// request 1 24 2021-10-06 19:01:12.787539 +0900 JST m=+6.553630363
// request 1 25 2021-10-06 19:01:13.287673 +0900 JST m=+7.053781297
// request 1 26 2021-10-06 19:01:13.789301 +0900 JST m=+7.555427033
// request 1 27 2021-10-06 19:01:14.29448 +0900 JST m=+8.060623215
// request 1 28 2021-10-06 19:01:14.799695 +0900 JST m=+8.565854809
// request 1 29 2021-10-06 19:01:15.302275 +0900 JST m=+9.068452304
// request 1 30 2021-10-06 19:01:15.803017 +0900 JST m=+9.569210944
// request 1 31 2021-10-06 19:01:16.303046 +0900 JST m=+10.069257328
// request 1 32 2021-10-06 19:01:16.808222 +0900 JST m=+10.574450402
// request 1 33 2021-10-06 19:01:17.312478 +0900 JST m=+11.078723527
// request 1 34 2021-10-06 19:01:17.815912 +0900 JST m=+11.582174666
// request 1 35 2021-10-06 19:01:18.316312 +0900 JST m=+12.082592306
// =========================
// request 1 36 2021-10-06 19:01:18.817988 +0900 JST m=+12.584284786
// request 2 37 2021-10-06 19:01:18.818022 +0900 JST m=+12.584319117
// request 3 38 2021-10-06 19:01:18.818035 +0900 JST m=+12.584331812
// request 3 39 2021-10-06 19:01:19.023242 +0900 JST m=+12.789545693
// request 3 40 2021-10-06 19:01:19.22851 +0900 JST m=+12.994820816
// request 3 41 2021-10-06 19:01:19.429854 +0900 JST m=+13.196172190
// request 3 42 2021-10-06 19:01:19.635135 +0900 JST m=+13.401459699
// request 3 43 2021-10-06 19:01:19.839412 +0900 JST m=+13.605743520
// request 3 44 2021-10-06 19:01:20.04465 +0900 JST m=+13.810988716
// request 3 45 2021-10-06 19:01:20.247615 +0900 JST m=+14.013960404
// request 3 46 2021-10-06 19:01:20.452864 +0900 JST m=+14.219216693
// request 3 47 2021-10-06 19:01:20.658114 +0900 JST m=+14.424473394
// request 3 48 2021-10-06 19:01:20.863423 +0900 JST m=+14.629789931
// request 3 49 2021-10-06 19:01:21.067726 +0900 JST m=+14.834099619
// request 3 50 2021-10-06 19:01:21.272997 +0900 JST m=+15.039378006
// request 3 0 2021-10-06 19:01:21.273084 +0900 JST m=+15.039464941
// End receiving
// receive ended
//  
// Compilation finished at Wed Oct  6 19:01:21
