package main

import (
	"errors"
	"fmt"
)

func func1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{3, 7, 42, 44, 800} {
		//fmt.Println(func1(i))
		
		if
		r, e := func1(i);
		e != nil {
			//fmt.Println(e.Error())
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sat Oct 30 15:35:02
//  
// go run practice12-1.go
// f1 worked: 6
// f1 worked: 10
// f1 failed: can't work with 42
// f1 worked: 47
// f1 worked: 803
//  
// Compilation finished at Sat Oct 30 15:35:02
