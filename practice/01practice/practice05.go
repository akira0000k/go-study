package main

import "fmt"

type StrSlice []string

/*
   subject : Each for slice   golangの日記　Go言語(golang)のループ for, for..range, foreach, while
 */ 

func (a StrSlice) Each(f func(i int, v string)) {
	for i, v := range a {
		f(i, v)
	}
}

type IntSlice []int

func (a IntSlice) Each(f func(i int, v int)) {
	for i, v := range a {
		f(i, v)
	}
}

func main() {
	fmt.Println("Start practice")

	StrSlice{"foo", "bar", "baz"}.
		Each(func(i int, v string) {
			fmt.Println(i, v)
		})

	IntSlice{10, 20, 30}.Each(func(i int, v int) {
		fmt.Println(i, v)
	})

	fmt.Println("End practice")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sat Oct 30 15:54:18
//  
// go run practice05.go
// Start practice
// 0 foo
// 1 bar
// 2 baz
// 0 10
// 1 20
// 2 30
// End practice
//  
// Compilation finished at Sat Oct 30 15:54:19
