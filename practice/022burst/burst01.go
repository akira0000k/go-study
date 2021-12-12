package main

import (
	"fmt"
	"time"
	"sync"
	"sync/atomic"
)
/*
   subject: Rating limit again. 
*/

const maxrequest = 10
const okburst = 3

type Request struct {
	id int64
	wait time.Duration
	name string
}
func (req *Request) Dotask(id int) {
	fmt.Println(id, *req, time.Now())
	time.Sleep(req.wait)
}
func MakeRequest(id int64, wait time.Duration, name string) Request {
	return Request{ id, wait, name }
}

func main() {
	fmt.Println("start main")
	fmt.Printf("%T %v\n", time.Second, time.Second)
	
	stop := false

	var wg sync.WaitGroup
	requests := make(chan Request, maxrequest)
	bursts := make(chan Request, okburst)

	for i:=0; i<okburst; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for req := range bursts {
				req.Dotask(id)
			}
		}(i)
	}

	go func() {
		defer close(bursts)
		for req := range requests {
			bursts <-req
		}
	}()

	var wg2 sync.WaitGroup
	var id int64=0
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for {
			if stop {
				return
			}
			atomic.AddInt64(&id, 1)
			dur := time.Millisecond * 100
			requests<-MakeRequest(id, dur, "normal")
			time.Sleep(time.Second * 2)
		}
	}()
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for {
			nb := 6
			for i:=0; i<nb; i++ {
				if stop {
					return
				}
				atomic.AddInt64(&id, 1)
				dur := time.Millisecond * 100
				requests<-MakeRequest(id, dur, "burst")
				//time.Sleep(time.Millisecond)
			}
			time.Sleep(time.Second * 5)
		}
	}()
	
	time.Sleep(time.Second * 30)

	stop = true
	wg2.Wait()
	close(requests)

	wg.Wait()

	fmt.Println("stop main")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/022burst/" -*-
// Compilation started at Wed Oct  6 22:50:13
//  
// go run burst01.go
// start main
// time.Duration 1s
// 2 {3 100000000 burst} 2021-10-06 22:50:14.1563 +0900 JST m=+0.000219942
// 0 {1 100000000 burst} 2021-10-06 22:50:14.156288 +0900 JST m=+0.000207555
// 1 {2 100000000 burst} 2021-10-06 22:50:14.156289 +0900 JST m=+0.000208911
// 2 {5 100000000 burst} 2021-10-06 22:50:14.258247 +0900 JST m=+0.102169403
// 1 {4 100000000 burst} 2021-10-06 22:50:14.258198 +0900 JST m=+0.102119956
// 0 {6 100000000 burst} 2021-10-06 22:50:14.258262 +0900 JST m=+0.102183952
// 2 {7 100000000 normal} 2021-10-06 22:50:14.363497 +0900 JST m=+0.207421403
// 0 {8 100000000 normal} 2021-10-06 22:50:16.157899 +0900 JST m=+2.001865200
// 1 {9 100000000 normal} 2021-10-06 22:50:18.159815 +0900 JST m=+4.003827271
// 2 {10 100000000 burst} 2021-10-06 22:50:19.161378 +0900 JST m=+5.005412367
// 1 {12 100000000 burst} 2021-10-06 22:50:19.161363 +0900 JST m=+5.005398019
// 0 {11 100000000 burst} 2021-10-06 22:50:19.161368 +0900 JST m=+5.005402855
// 0 {13 100000000 burst} 2021-10-06 22:50:19.265177 +0900 JST m=+5.109214666
// 2 {14 100000000 burst} 2021-10-06 22:50:19.265245 +0900 JST m=+5.109282233
// 1 {15 100000000 burst} 2021-10-06 22:50:19.265266 +0900 JST m=+5.109303453
// 1 {16 100000000 normal} 2021-10-06 22:50:20.1619 +0900 JST m=+6.005958027
// 0 {17 100000000 normal} 2021-10-06 22:50:22.163864 +0900 JST m=+8.007967952
// 0 {20 100000000 burst} 2021-10-06 22:50:24.165834 +0900 JST m=+10.009983344
// 1 {19 100000000 burst} 2021-10-06 22:50:24.165878 +0900 JST m=+10.010027552
// 2 {18 100000000 burst} 2021-10-06 22:50:24.165861 +0900 JST m=+10.010010674
// 1 {22 100000000 burst} 2021-10-06 22:50:24.265959 +0900 JST m=+10.110110405
// 0 {23 100000000 burst} 2021-10-06 22:50:24.265996 +0900 JST m=+10.110147698
// 2 {21 100000000 burst} 2021-10-06 22:50:24.265946 +0900 JST m=+10.110097768
// 2 {24 100000000 normal} 2021-10-06 22:50:24.368454 +0900 JST m=+10.212608199
// 0 {25 100000000 normal} 2021-10-06 22:50:26.167849 +0900 JST m=+12.012044641
// 1 {26 100000000 normal} 2021-10-06 22:50:28.169687 +0900 JST m=+14.013928703
// 1 {29 100000000 burst} 2021-10-06 22:50:29.170928 +0900 JST m=+15.015192741
// 0 {28 100000000 burst} 2021-10-06 22:50:29.170972 +0900 JST m=+15.015236133
// 2 {27 100000000 burst} 2021-10-06 22:50:29.170943 +0900 JST m=+15.015207382
// 2 {30 100000000 burst} 2021-10-06 22:50:29.276255 +0900 JST m=+15.120521925
// 0 {31 100000000 burst} 2021-10-06 22:50:29.276297 +0900 JST m=+15.120563292
// 1 {32 100000000 burst} 2021-10-06 22:50:29.276315 +0900 JST m=+15.120581872
// 1 {33 100000000 normal} 2021-10-06 22:50:30.170699 +0900 JST m=+16.014986463
// 0 {34 100000000 normal} 2021-10-06 22:50:32.172623 +0900 JST m=+18.016956046
// 0 {37 100000000 burst} 2021-10-06 22:50:34.174355 +0900 JST m=+20.018734429
// 2 {35 100000000 burst} 2021-10-06 22:50:34.174382 +0900 JST m=+20.018761306
// 1 {36 100000000 normal} 2021-10-06 22:50:34.174576 +0900 JST m=+20.018955044
// 1 {38 100000000 burst} 2021-10-06 22:50:34.279621 +0900 JST m=+20.124002141
// 2 {39 100000000 burst} 2021-10-06 22:50:34.279684 +0900 JST m=+20.124065434
// 0 {40 100000000 burst} 2021-10-06 22:50:34.279687 +0900 JST m=+20.124068020
// 2 {41 100000000 burst} 2021-10-06 22:50:34.384872 +0900 JST m=+20.229255756
// 1 {42 100000000 normal} 2021-10-06 22:50:36.175218 +0900 JST m=+22.019642299
// 0 {43 100000000 normal} 2021-10-06 22:50:38.176817 +0900 JST m=+24.021288105
// 0 {46 100000000 burst} 2021-10-06 22:50:39.178156 +0900 JST m=+25.022649441
// 2 {44 100000000 burst} 2021-10-06 22:50:39.178171 +0900 JST m=+25.022664559
// 1 {45 100000000 burst} 2021-10-06 22:50:39.178222 +0900 JST m=+25.022715742
// 2 {47 100000000 burst} 2021-10-06 22:50:39.283426 +0900 JST m=+25.127922545
// 1 {48 100000000 burst} 2021-10-06 22:50:39.28345 +0900 JST m=+25.127946468
// 0 {49 100000000 burst} 2021-10-06 22:50:39.283482 +0900 JST m=+25.127977670
// 1 {50 100000000 normal} 2021-10-06 22:50:40.178729 +0900 JST m=+26.023245684
// 2 {51 100000000 normal} 2021-10-06 22:50:42.180621 +0900 JST m=+28.025183436
// stop main
//  
// Compilation finished at Wed Oct  6 22:50:44
