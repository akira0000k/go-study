package main
import (
	"fmt"
	"time"
)
/*
   subject: Closing Channels  サンプルで学ぶGo言語
*/
var jobs chan int //for named func
var done chan bool
func monkey() {
	for {
		j, more := <-jobs
		if more {
			fmt.Println("received job", j)
		} else {
			fmt.Println("received all jobs")
			done<-true
			return
		}
	}
}

func main() {
	_ = time.Second

	fmt.Println("start")

	//jobs := make(chan int, 5) //for closure
	//done := make(chan bool)
	jobs = make(chan int, 5)
	done = make(chan bool)
	//time.Sleep(time.Second)

	//go monkey() //named func
	//go func() { //direct exec
	baby := func () { //var baby
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done<-true
				return
			}
		}
	}
	go baby() //var baby
	//}() //direct exec

	for j:=1; j<=3; j++ {
		jobs<-j
		fmt.Println("sent job", j)
		time.Sleep(time.Millisecond / 10)
	}
	fmt.Println("sent all jobs")
	close(jobs)

	fmt.Println("receive done=", <-done)
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Fri Oct 29 18:31:56
//  
// go run sync04.go
// start
// sent job 1
// received job 1
// sent job 2
// received job 2
// sent job 3
// received job 3
// sent all jobs
// received all jobs
// receive done= true
//  
// Compilation finished at Fri Oct 29 18:31:57
