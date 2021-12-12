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
 subject: error戻しチャネルの詰まりを sync.Once を使って解消する
*/

var once sync.Once

func doSomething() error {
	fmt.Println("doSomething start")

	time.Sleep(time.Second)
	
	fmt.Println("doSomething err")
	return errors.New("something error")
}


func A(wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()
	fmt.Println("A start")

	if err := doSomething(); err != nil {
		//
		fmt.Println("something returns", err)
		once.Do(func() {
			fmt.Println("errChan <-err")
			errChan <-err
		})
	}
	fmt.Println("A end<----------")
}

func B() error {
	fmt.Println("B start")

	wg := new(sync.WaitGroup)
	errChan := make(chan error)
	done := make(chan struct{})

	wg.Add(2)
	go A(wg, errChan)
	go A(wg, errChan)

	go func() {
		fmt.Println("wg.Waiting......")
		wg.Wait()
		fmt.Println("......exit wg.Wait")
		close(done)
	}()

	// 到達順序は保証外(まあ大丈夫)
	select {
	case <-done:
		fmt.Println("case <-done:")
		return nil
	case err := <-errChan:
		fmt.Println("case err := <-errChan:")
		return err
	}
}

func main() {
	fmt.Println("start")

	err := B()
	fmt.Println("B() err=", err)
	

	fmt.Println("end")
}
// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/" -*-
// Compilation started at Sat Sep 25 21:02:06
//  
// go run patn04.go
// start
// B start
// wg.Waiting......
// A start
// doSomething start
// A start
// doSomething start
// doSomething err
// doSomething err
// something returns something error
// errChan <-err
// A end<----------
// something returns something error
// A end<----------
// ......exit wg.Wait
// case err := <-errChan:
// B() err= something error
// end
//  
// Compilation finished at Sat Sep 25 21:02:08
