package main

import (
	"fmt"
	"math"
)

/*
   subject : for Loop
 */ 
func main1() {
	fmt.Println("Start practice")
	var i int=99;
	fmt.Println("first i=", i)
	
	//for i:=0;; {
	//i:=0; for i<5 {
	//for i=0; i<10; i++ {
	for i:=0; i<10; i++ {
		fmt.Printf("i=%d\n", i)
		//if (i == 10) {break}
		//i++
		if i:=333; i==5 {
			fmt.Printf("i==5\n")
		}
	}
	fmt.Println("last i=", i)

	fmt.Println("End practice")
}
func main() {
	fmt.Println("Start practice")
loop:
	for i, j:=0,1 ; i<3; i++ {
		for i := 0; i<10; i,j = i+1, j*2 {
			//for i, j := 0, 1; i<10; i,j = i+1, j*2 {
			//for i,j:=0,1; i<10; i++ {
			fmt.Println(i+1, j)
			if j >= int(math.Pow10(5)) { break loop }
			//j*=2
		}
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sat Oct 30 16:01:55
//  
// go run practice03.go
// Start practice
// 1 1
// 2 2
// 3 4
// 4 8
// 5 16
// 6 32
// 7 64
// 8 128
// 9 256
// 10 512
// 1 1024
// 2 2048
// 3 4096
// 4 8192
// 5 16384
// 6 32768
// 7 65536
// 8 131072
//  
// Compilation finished at Sat Oct 30 16:01:56
