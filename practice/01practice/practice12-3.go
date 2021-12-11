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
	_, e := func2(42)
	fmt.Println(e.Error())

	//fmt.Println(e.(type)) //./practice12-3.go:26:15: use of .(type) outside type switch
	//arger := *(*argError)(e) //./practice12-3.go:27:23: cannot convert e (type error) to type *argError: need type assertion
	
	// 型アサーションを使って自作したエラー型のインスタンスを作る
	ae, ok := e.(*argError)
 
	if
	//ae, ok := e.(*argError);
	ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sat Oct 30 15:36:04
//  
// go run practice12-3.go
// 42 - can't work with it
// 42
// can't work with it
//  
// Compilation finished at Sat Oct 30 15:36:05
