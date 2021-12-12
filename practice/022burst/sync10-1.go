package main
import (
	"fmt"
	"time"
	//"sync"
)
/*
   subject: Rate Limiting
*/

func main() {
	fmt.Println("start")
	
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
