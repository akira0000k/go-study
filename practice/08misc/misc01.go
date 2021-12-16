package main

import (
	"fmt"
	"sort"
)
/*
   subject: sorting
*/
func main() {
	strs := []string{"ccc", "a", "b", "cc", "c", "aa", "bbb"}

	fmt.Println("Strings:", strs)
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4, 17, 247, -543}
	fmt.Println("Ints:   ", ints)
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	sort.Ints(ints)

	fmt.Println("Ints:   ", ints)
	s = sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
// -*- mode: compilation; default-directory: "~/go/src/practice/08misc/" -*-
// Compilation started at Wed Oct  6 23:20:01
//  
// go run misc01.go
// Strings: [ccc a b cc c aa bbb]
// Strings: [a aa b bbb c cc ccc]
// Ints:    [7 2 4 17 247 -543]
// Sorted:  false
// Ints:    [-543 2 4 7 17 247]
// Sorted:  true
//  
// Compilation finished at Wed Oct  6 23:20:01
