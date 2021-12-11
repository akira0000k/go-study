package main

/*
   subject : サンプルで学ぶ Go 言語：Errors
 */ 
import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
	ext  int
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s. info:%d", e.arg, e.prob, e.ext)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it", 33}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
		fmt.Println(ae.ext)
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sat Oct 30 00:08:01
//  
// go run practice12.go
// f1 worked: 10
// f1 failed: can't work with 42
// f2 worked: 10
// f2 failed: 42 - can't work with it. info:33
// 42
// can't work with it
// 33
//  
// Compilation finished at Sat Oct 30 00:08:03
