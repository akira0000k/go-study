package main

import (
	"fmt"
	//"math/rand"
	"sync"
	//"sync/atomic"
	"time"
	"errors"
)

/*
 subject: range で job queueを読み出し、channel closeで抜ける。
 送信側は1つ and 受信側は1つ
 並列化するほどじゃないけど受け取り処理はさっさと済ませて遅延処理させたいシンプルなパターン。
 channel の capacity が溢れない程度に受信側が上手いこと処理してくれている（間に合っている）場合はコレで足りる。
*/
func main() {
	fmt.Println("start main")

	err := A()
	fmt.Println(err)

	wgrp.Wait()
	fmt.Println("end main")
}


func (j *Job) Run() {
	fmt.Println("start", j.id, "....")
	time.Sleep(time.Second)
	fmt.Println(".....end", j.id)
}


func worker(queue chan Job) {
	defer wgrp.Done()
	
	for j := range queue {
		j.Run()
	}
}

func someHandle() (Job, error) {
	jobid++
	if jobid > 10 {
		return Job{}, errors.New("End of jobs")
	}
	time.Sleep(time.Millisecond * 10)
	j := Job{ jobid }
	return j, nil
}

func A() error {
	queue := make(chan Job, 100)
	defer close(queue)
	var j Job
	var err error
	
	wgrp.Add(1)
	go worker(queue)

	for {
		j, err = someHandle()
		if err != nil {
			break
		}
		queue <-j
	}
	return err
}

	
// file scope の宣言は順不同。後ろに書いても良い。
type Job struct {
	id int
}
var wgrp sync.WaitGroup
var jobid int
// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/03anti-pattern/" -*-
// Compilation started at Mon Sep 27 18:26:14
//  
// go run patn08.go
// start main
// start 1 ....
// End of jobs
// .....end 1
// start 2 ....
// .....end 2
// start 3 ....
// .....end 3
// start 4 ....
// .....end 4
// start 5 ....
// .....end 5
// start 6 ....
// .....end 6
// start 7 ....
// .....end 7
// start 8 ....
// .....end 8
// start 9 ....
// .....end 9
// start 10 ....
// .....end 10
// end main
//  
// Compilation finished at Mon Sep 27 18:26:25
