package main

import (
	"fmt"
	"math/cmplx"
)
/*
   subject : complex64/128 type  (no complex type)
 */
func usecomplex() {
	//var c complex //./practice21.go:10:8: use of builtin complex not in function call
	//var c complex128 = 1+2i
	//var f float //./practice21.go:12:8: undefined: float
	var f float32
	
	var c complex64
	c = 1+2i
	fmt.Println(c, f)
	fmt.Printf("%T  %v\n", c, c)
}

func main() {
	fmt.Println("start")
	const npow=30
	c := -1.0+0i
	q := c
	d := 2
	for i:=0; i<npow; i++ {
		q = cmplx.Sqrt(q)
		d *= 2
	}
	fmt.Println("q=", q)
	fmt.Println("d=", d)
	for i, j, v :=0, 0, 1+0i; i<2*d+1; i++ {
		if j==0 {
			fmt.Println(i, j, v)
		}
		v *= q
		if j++; j == d {
			j = 0
		}
	}

}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sat Oct 30 15:12:56
//  
// go run practice21.go
// start
// q= (0.9999999999999999+2.925836158534322e-09i)
// d= 2147483648
// 0 0 (1+0i)
// 2147483648 0 (0.9999997287132282-1.0347541797053884e-13i)
// 4294967296 0 (0.9999994574255068-3.1504826000482843e-13i)
//  
// Compilation finished at Sat Oct 30 15:13:10
