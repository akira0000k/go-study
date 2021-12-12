package main
import (
	"fmt"
	"time"
	//"sync"
)
/*
   subject: Rate Limiting. limiter   https://www.spinute.org/go-by-example/rate-limiting.html
*/

func main() {
 	nRequ := 20
	requests := make(chan int, nRequ)

	//burst request
	for i:=1; i<=nRequ; i++ {
		requests<-i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("===============================")
 
	nBurst := 10
	burstyLimiter := make(chan time.Time, nBurst)
 
	for i:=0; i<nBurst; i++ {
		burstyLimiter<-time.Now()
	}
 
	go func() {
		for t := range time.Tick(200*time.Millisecond) {
			burstyLimiter <- t
		}
	}()
 
	fmt.Println("-------------------------------")
 
	nReq := 20
	burstyRequests := make(chan int, nReq)
	for i:=1; i<=nReq; i++{
		burstyRequests <- i
	}
	close(burstyRequests)
 
	
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/022burst/" -*-
// Compilation started at Wed Oct  6 18:57:45
//  
// go run sync10.go
// request 1 2021-10-06 18:57:47.073399 +0900 JST m=+0.205359189
// request 2 2021-10-06 18:57:47.273378 +0900 JST m=+0.405344378
// request 3 2021-10-06 18:57:47.469543 +0900 JST m=+0.601516384
// request 4 2021-10-06 18:57:47.668884 +0900 JST m=+0.800864335
// request 5 2021-10-06 18:57:47.871523 +0900 JST m=+1.003509855
// request 6 2021-10-06 18:57:48.073325 +0900 JST m=+1.205318778
// request 7 2021-10-06 18:57:48.268174 +0900 JST m=+1.400175088
// request 8 2021-10-06 18:57:48.471191 +0900 JST m=+1.603198845
// request 9 2021-10-06 18:57:48.670582 +0900 JST m=+1.802596252
// request 10 2021-10-06 18:57:48.871553 +0900 JST m=+2.003574202
// request 11 2021-10-06 18:57:49.073292 +0900 JST m=+2.205319733
// request 12 2021-10-06 18:57:49.269769 +0900 JST m=+2.401803538
// request 13 2021-10-06 18:57:49.471974 +0900 JST m=+2.604015244
// request 14 2021-10-06 18:57:49.673278 +0900 JST m=+2.805326740
// request 15 2021-10-06 18:57:49.871848 +0900 JST m=+3.003903429
// request 16 2021-10-06 18:57:50.073245 +0900 JST m=+3.205306758
// request 17 2021-10-06 18:57:50.268358 +0900 JST m=+3.400426778
// request 18 2021-10-06 18:57:50.473117 +0900 JST m=+3.605192819
// request 19 2021-10-06 18:57:50.670583 +0900 JST m=+3.802665324
// request 20 2021-10-06 18:57:50.872964 +0900 JST m=+4.005053615
// ===============================
// -------------------------------
// request 1 2021-10-06 18:57:50.873118 +0900 JST m=+4.005207623
// request 2 2021-10-06 18:57:50.87314 +0900 JST m=+4.005229611
// request 3 2021-10-06 18:57:50.873153 +0900 JST m=+4.005242617
// request 4 2021-10-06 18:57:50.873164 +0900 JST m=+4.005252944
// request 5 2021-10-06 18:57:50.873226 +0900 JST m=+4.005315668
// request 6 2021-10-06 18:57:50.873247 +0900 JST m=+4.005336793
// request 7 2021-10-06 18:57:50.873257 +0900 JST m=+4.005346278
// request 8 2021-10-06 18:57:50.873268 +0900 JST m=+4.005357538
// request 9 2021-10-06 18:57:50.873281 +0900 JST m=+4.005370665
// request 10 2021-10-06 18:57:50.873292 +0900 JST m=+4.005381305
// request 11 2021-10-06 18:57:51.073714 +0900 JST m=+4.205809842
// request 12 2021-10-06 18:57:51.273634 +0900 JST m=+4.405736871
// request 13 2021-10-06 18:57:51.473269 +0900 JST m=+4.605378804
// request 14 2021-10-06 18:57:51.673551 +0900 JST m=+4.805667870
// request 15 2021-10-06 18:57:51.874464 +0900 JST m=+5.006587152
// request 16 2021-10-06 18:57:52.073254 +0900 JST m=+5.205384020
// request 17 2021-10-06 18:57:52.273785 +0900 JST m=+5.405921626
// request 18 2021-10-06 18:57:52.473586 +0900 JST m=+5.605729751
// request 19 2021-10-06 18:57:52.673286 +0900 JST m=+5.805436692
// request 20 2021-10-06 18:57:52.873724 +0900 JST m=+6.005881692
//  
// Compilation finished at Wed Oct  6 18:57:52
