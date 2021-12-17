package main

import (
	"fmt"
)

/*
   Variadic Params
*/
func main() {
	fmt.Println("start main")

	fmt.Println(rtn2(5, 10))
	//fmt.Println("rtn2: ", rtn2(5, 10))//./learn03.go:14:28: multiple-value rtn2() in single-value context
	fmt.Println("rtn2: ", fmt.Sprint(rtn2(6,12)))
	
	prnp(3, 4, 5)

	fmt.Println("end main")
}


func rtn2(a, b int) (c, d int) {
	c = b
	d = a
	return
}

func prnp(intp ...int) {
	for i, v := range intp {
		fmt.Printf("param[%d] = %d\n", i, v)
	}
	return
}
// -*- mode: compilation; default-directory: "~/go/src/practice/15LearnXinYmin/" -*-
// Compilation started at Fri Dec 17 17:27:48
//  
// go run learn03.go
// start main
// 10 5
// rtn2:  12 6
// param[0] = 3
// param[1] = 4
// param[2] = 5
// end main
//  
// Compilation finished at Fri Dec 17 17:27:49
