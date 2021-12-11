package main

import (
	"fmt"
)


type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func func2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{1, 5, 7, 42, 98, 156} {
		if
		r, e := func2(i);
		e != nil {
			fmt.Println("f2 failed:", e)  //e.Error())
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sat Oct 30 15:35:42
//  
// go run practice12-2.go
// f2 worked: 4
// f2 worked: 8
// f2 worked: 10
// f2 failed: 42 - can't work with it
// f2 worked: 101
// f2 worked: 159
//  
// Compilation finished at Sat Oct 30 15:35:43
