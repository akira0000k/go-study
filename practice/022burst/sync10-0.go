package main
import (
	"fmt"
	"time"
	//"sync"
)
/*
   subject: Rate Limiting. limiter
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

// 	fmt.Println("===============================")
// 
// 	nBurst := 10
// 	burstyLimiter := make(chan time.Time, nBurst)
// 
// 	for i:=0; i<nBurst; i++ {
// 		burstyLimiter<-time.Now()
// 	}
// 
// 	go func() {
// 		for t := range time.Tick(200*time.Millisecond) {
// 			burstyLimiter <- t
// 		}
// 	}()
// 
// 	fmt.Println("-------------------------------")
// 
// 	nReq := 20
// 	burstyRequests := make(chan int, nReq)
// 	for i:=1; i<=nReq; i++{
// 		burstyRequests <- i
// 	}
// 	close(burstyRequests)
// 
// 	
// 	for req := range burstyRequests {
// 		<-burstyLimiter
// 		fmt.Println("request", req, time.Now())
// 	}
}
