package main
import (
	"fmt"
	"time"
)
/*
   subject: Timers
*/

func main() {
	//_ = time.Second

	fmt.Println("start")

	timer1 := time.NewTimer(time.Second)
	fmt.Println(time.Now(), "Timer 1 start")

	//conut one two
	fmt.Println(<-timer1.C, "Timer 1 fired")


	
	timer2 := time.NewTimer(2 * time.Second)
	fmt.Println(time.Now(), "Timer 2 start")
	go func() {
		//count one
		fmt.Println(<-timer2.C, "Timer 2 fired")
	}()
	time.Sleep(time.Second * 2 *999/1000) //not enough precision
	//time.Sleep(time.Second * 2 *99/100) //stop
	fmt.Println(time.Now(), "wake up")
	
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println(time.Now(), "Timer 2 stopped")
	} else {
		fmt.Println(time.Now(), "no Timer 2")
	}

	time.Sleep(2*time.Second)

	fmt.Println(time.Now(), "end")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Fri Oct 29 18:39:12
//  
// go run sync06.go
// start
// 2021-10-29 18:39:12.943994 +0900 JST m=+0.000185879 Timer 1 start
// 2021-10-29 18:39:13.949157 +0900 JST m=+1.005318165 Timer 1 fired
// 2021-10-29 18:39:13.949348 +0900 JST m=+1.005509272 Timer 2 start
// 2021-10-29 18:39:15.951215 +0900 JST m=+3.007316451 Timer 2 fired
// 2021-10-29 18:39:15.951264 +0900 JST m=+3.007365158 wake up
// 2021-10-29 18:39:15.951385 +0900 JST m=+3.007486375 no Timer 2
// 2021-10-29 18:39:17.95666 +0900 JST m=+5.012701466 end
//  
// Compilation finished at Fri Oct 29 18:39:17
