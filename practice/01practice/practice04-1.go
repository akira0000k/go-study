package main

import (
	"fmt"
	"sort"
)

/*
   subject : for range array map and sort
*/
func main() {
	main1()
	main2()
}
func main1() {
	fmt.Println("Start ange map")
	m := map[string]int{"orange": 100, "xpple": 50, "banana": 30}
	for k, v := range m {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
	fmt.Println("End map practice\n")
}


func main2() {
	fmt.Println("Start sort")
	m := map[string]int{"orange":100,"xpple":50,"banana":30}
	var keys []string
	for key:=range m{
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _,key := range keys{
		fmt.Printf("key:%s, value:%d\n", key, m[key])
	}
	fmt.Println("End sort practice")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Thu Dec  2 17:02:24
//  
// go run practice04-1.go
// Start ange map
// key: xpple, value: 50
// key: banana, value: 30
// key: orange, value: 100
// End map practice
//  
// Start sort
// key:banana, value:30
// key:orange, value:100
// key:xpple, value:50
// End sort practice
//  
// Compilation finished at Thu Dec  2 17:02:24
