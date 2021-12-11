package main

import (
	"fmt"
)
/*
   subject : サンプルで学ぶGo言語: Variadic function
 */

func sum(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
	return total
}


func main() {
	fmt.Println("start")

	fmt.Println(sum(1,2,3,4,5))
	fmt.Println(sum(1,2,3,4,5,6,7,8,9,10))
	fmt.Println(sum(1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20))
	sl := []int{1,2,3}
	fmt.Println(sum(sl...))

	a := [...]int{1,2,3,4,5,6}
	//var ar []int = a[:]
	ar := a[:]  //convert array to slice
	fmt.Println(sum(ar...))
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Thu Oct 28 21:25:02
//  
// go run practice24.go
// start
// [1 2 3 4 5] 15
// 15
// [1 2 3 4 5 6 7 8 9 10] 55
// 55
// [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20] 210
// 210
// [1 2 3] 6
// 6
// [1 2 3 4 5 6] 21
// 21
//  
// Compilation finished at Thu Oct 28 21:25:03
