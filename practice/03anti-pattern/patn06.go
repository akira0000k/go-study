package main

import (
	"fmt"
	//"math/rand"
	"sync"
	//"sync/atomic"
	"time"
)

/*
 subject: channel buffer を使った　worker起動数制限の方法
*/


type Job struct {
	id int
}
func (j *Job) Run() {
	fmt.Println("start", j.id, "....")
	time.Sleep(time.Second)
	fmt.Println(".....end", j.id)
}



func A(jobs []Job) {
	sem := make(chan struct{}, 10)
	var wgrp sync.WaitGroup
	
	for _, j := range jobs {
		sem <- struct{}{}
		wgrp.Add(1)
		go func(j Job) {
			defer wgrp.Done()
			j.Run()
			<-sem
		}(j)
	}
	wgrp.Wait()
}

func main() {
	fmt.Println("start main")

	var jobs []Job

	for i:=1; i<=30; i++ {
		j := Job{ i }
		jobs = append(jobs, j)
	}

	A(jobs)

	fmt.Println("end main")
}
// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/03anti-pattern/" -*-
// Compilation started at Sun Sep 26 23:35:04
//  
// go run patn06-2.go
// start main
// start 10 ....
// start 8 ....
// start 1 ....
// start 5 ....
// start 9 ....
// start 3 ....
// start 4 ....
// start 6 ....
// start 2 ....
// start 7 ....
// .....end 7
// .....end 4
// .....end 9
// .....end 1
// .....end 2
// .....end 6
// .....end 8
// .....end 5
// .....end 3
// .....end 10
// start 11 ....
// start 12 ....
// start 20 ....
// start 13 ....
// start 14 ....
// start 15 ....
// start 19 ....
// start 17 ....
// start 16 ....
// start 18 ....
// .....end 16
// start 21 ....
// .....end 11
// .....end 20
// start 22 ....
// .....end 13
// start 24 ....
// .....end 12
// .....end 18
// start 26 ....
// start 23 ....
// .....end 15
// start 27 ....
// start 25 ....
// .....end 14
// start 28 ....
// .....end 19
// start 29 ....
// .....end 17
// start 30 ....
// .....end 30
// .....end 24
// .....end 22
// .....end 26
// .....end 23
// .....end 29
// .....end 28
// .....end 25
// .....end 21
// .....end 27
// end main
//  
// Compilation finished at Sun Sep 26 23:35:07
