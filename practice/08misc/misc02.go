package main

import (
	"fmt"
	"sort"
)
/*
   subject: sorting with original compare method
*/
type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	if len(s[i]) == len(s[j]) {
		return s[i] < s[j]
	} else {
		return len(s[i]) < len(s[j])
	}
}

func main() {
	fruits := []string{"abacaxi", "peach", "kiwi", "delaware", "melon", "abocado", "dorian", "banana"}
	fmt.Println(fruits)

	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

	iary := []int{ 3, 2, 5, 0, -9, 183, 66, -99 }
	fmt.Println(iary)
	sort.Sort(ordinaryInt(iary))
	fmt.Println(iary)
}

type ordinaryInt []int

func (oi ordinaryInt) Len() int {
	return len(oi)
}
func (oi ordinaryInt) Swap(i, j int) {
	oi[i], oi[j] = oi[j], oi[i]
}
func (oi ordinaryInt) Less(i, j int) bool {
	return oi[i] < oi[j]
}
