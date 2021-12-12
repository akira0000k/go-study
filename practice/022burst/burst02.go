package main

import (
	"fmt"
	"time"
	"sync"
	"sync/atomic"
)
/*
   subject: Rating limit again. Consider limiter channel.
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
	limiter := make(chan struct{}, okburst-1) //リミッターキュー長はバースト許容数-1 で良い。
	
	for i:=0; i<okburst; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for req := range bursts {
				req.Dotask(id)
			}
		}(i)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case _, ok := <-limiter:
				if !ok {
					fmt.Println("limiter closed")
					return
				}
				time.Sleep(time.Millisecond * 500)
			}
			//イラナイ
			//case <-time.After(time.Millisecond * 500):
			//default: time.Sleep(time.Millisecond * 10)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(limiter)
		defer close(bursts)
		for req := range requests {
			limiter <- struct{}{}
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
			time.Sleep(time.Second * 2) //normal event occurs every 2sec
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
// Compilation started at Wed Oct  6 22:38:37
//  
// go run burst02.go
// start main
// time.Duration 1s
// 0 {1 1000000 normal} 2021-10-06 22:38:37.905585 +0900 JST m=+0.000262793
// 2 {3 1000000 burst} 2021-10-06 22:38:37.905618 +0900 JST m=+0.000295561
// 1 {2 1000000 burst} 2021-10-06 22:38:37.905671 +0900 JST m=+0.000348625
// 1 {4 1000000 burst} 2021-10-06 22:38:38.410773 +0900 JST m=+0.505462141
// 2 {5 1000000 burst} 2021-10-06 22:38:38.913586 +0900 JST m=+1.008286473
// 0 {6 1000000 burst} 2021-10-06 22:38:39.41398 +0900 JST m=+1.508691685
// 1 {7 1000000 burst} 2021-10-06 22:38:39.916201 +0900 JST m=+2.010926497
// 2 {8 1000000 normal} 2021-10-06 22:38:40.421459 +0900 JST m=+2.516193548
// 0 {9 1000000 normal} 2021-10-06 22:38:41.909813 +0900 JST m=+4.004582608
// 0 {12 1000000 burst} 2021-10-06 22:38:42.909892 +0900 JST m=+5.004683766
// 2 {11 1000000 burst} 2021-10-06 22:38:42.909957 +0900 JST m=+5.004748642
// 1 {10 1000000 burst} 2021-10-06 22:38:42.909933 +0900 JST m=+5.004725209
// 2 {13 1000000 burst} 2021-10-06 22:38:43.412505 +0900 JST m=+5.507308435
// 0 {14 1000000 burst} 2021-10-06 22:38:43.913826 +0900 JST m=+6.008640858
// 1 {15 1000000 burst} 2021-10-06 22:38:44.414978 +0900 JST m=+6.509804826
// 2 {16 1000000 normal} 2021-10-06 22:38:44.916205 +0900 JST m=+7.011042847
// 0 {17 1000000 normal} 2021-10-06 22:38:45.918763 +0900 JST m=+8.013624181
// 0 {20 1000000 burst} 2021-10-06 22:38:47.914985 +0900 JST m=+10.009892203
// 1 {18 1000000 burst} 2021-10-06 22:38:47.915026 +0900 JST m=+10.009932744
// 2 {19 1000000 burst} 2021-10-06 22:38:47.91508 +0900 JST m=+10.009986572
// 0 {21 1000000 burst} 2021-10-06 22:38:48.41593 +0900 JST m=+10.510848778
// 1 {22 1000000 burst} 2021-10-06 22:38:48.9211 +0900 JST m=+11.016030167
// 2 {23 1000000 burst} 2021-10-06 22:38:49.424895 +0900 JST m=+11.519836746
// 0 {24 1000000 normal} 2021-10-06 22:38:49.925948 +0900 JST m=+12.020901102
// 1 {25 1000000 normal} 2021-10-06 22:38:50.427126 +0900 JST m=+12.522090495
// 2 {26 1000000 normal} 2021-10-06 22:38:51.925307 +0900 JST m=+14.020306133
// 0 {27 1000000 burst} 2021-10-06 22:38:52.920102 +0900 JST m=+15.015123764
// 2 {29 1000000 burst} 2021-10-06 22:38:52.920084 +0900 JST m=+15.015105849
// 1 {28 1000000 burst} 2021-10-06 22:38:52.920121 +0900 JST m=+15.015142185
// 1 {30 1000000 burst} 2021-10-06 22:38:53.425335 +0900 JST m=+15.520368127
// 0 {31 1000000 burst} 2021-10-06 22:38:53.929317 +0900 JST m=+16.024362050
// 2 {32 1000000 burst} 2021-10-06 22:38:54.434497 +0900 JST m=+16.529553671
// 1 {33 1000000 normal} 2021-10-06 22:38:54.935719 +0900 JST m=+17.030786649
// 0 {34 1000000 normal} 2021-10-06 22:38:55.931237 +0900 JST m=+18.026327704
// 2 {35 1000000 burst} 2021-10-06 22:38:57.925184 +0900 JST m=+20.020320452
// 1 {36 1000000 burst} 2021-10-06 22:38:57.925177 +0900 JST m=+20.020313224
// 0 {37 1000000 burst} 2021-10-06 22:38:57.925171 +0900 JST m=+20.020307690
// 0 {38 1000000 burst} 2021-10-06 22:38:58.426328 +0900 JST m=+20.521475744
// 2 {39 1000000 burst} 2021-10-06 22:38:58.931505 +0900 JST m=+21.026664420
// 1 {40 1000000 burst} 2021-10-06 22:38:59.434309 +0900 JST m=+21.529480396
// 0 {41 1000000 normal} 2021-10-06 22:38:59.934581 +0900 JST m=+22.029763622
// 2 {42 1000000 normal} 2021-10-06 22:39:00.439749 +0900 JST m=+22.534943315
// 1 {43 1000000 normal} 2021-10-06 22:39:01.93656 +0900 JST m=+24.031788581
// 1 {46 1000000 burst} 2021-10-06 22:39:02.930257 +0900 JST m=+25.025508501
// 2 {45 1000000 burst} 2021-10-06 22:39:02.930306 +0900 JST m=+25.025556999
// 0 {44 1000000 burst} 2021-10-06 22:39:02.93028 +0900 JST m=+25.025531144
// 0 {47 1000000 burst} 2021-10-06 22:39:03.434201 +0900 JST m=+25.529463943
// 2 {48 1000000 burst} 2021-10-06 22:39:03.937669 +0900 JST m=+26.032943311
// 1 {49 1000000 burst} 2021-10-06 22:39:04.439099 +0900 JST m=+26.534385148
// 0 {50 1000000 normal} 2021-10-06 22:39:04.941177 +0900 JST m=+27.036474110
// 2 {51 1000000 normal} 2021-10-06 22:39:05.939519 +0900 JST m=+28.034839284
// limiter closed
// stop main
//  
// Compilation finished at Wed Oct  6 22:39:07
