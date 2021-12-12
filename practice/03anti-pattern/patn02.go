package main

import (
	"fmt"
	//"math/rand"
	//"sync"
	//"sync/atomic"
	"time"
)

/*
 subject: goroutineに終了eventを送るケース。
*/
type Job int

func worker(jobChan chan Job, done chan struct{}) {
	fmt.Println("A start")
	for {
		select {
		case j := <-jobChan:
			fmt.Println("worker receive job", j)
		case <-done:
			fmt.Println("worker done")
			return
		}
	}
}

func B(jobs []Job) {
	fmt.Println("B start")

	jobChan := make(chan Job)
	done := make(chan struct{})

	go worker(jobChan, done)

	for _, j := range jobs {
		fmt.Println("jobChan<-", j)
		jobChan <-j
		time.Sleep(time.Second / 2)
		if j == 3 {
			fmt.Println("done@3 <-struct{}{}"); done <-struct{}{}
			break
		}

	}
	time.Sleep(time.Second * 2)

	//fmt.Println("done <-struct{}{}"); done <-struct{}{}

	time.Sleep(time.Second * 2)
	fmt.Println("B end")
}

func main() {
	fmt.Println("start")

	jobs := []Job{5, 4, 3, 2, 1, 0}
	B(jobs)

	fmt.Println("end")
}
// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/" -*-
// Compilation started at Sat Sep 25 19:32:27
//  
// go run patn02.go
// start
// B start
// jobChan<- 5
// A start
// worker receive job 5
// jobChan<- 4
// worker receive job 4
// jobChan<- 3
// worker receive job 3
// done@3 <-struct{}{}
// worker done
// B end
// end
//  
// Compilation finished at Sat Sep 25 19:32:34
