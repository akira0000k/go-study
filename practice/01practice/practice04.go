package main

import "fmt"

//var i int;
/*
   subject : for range array map
*/
func main() {
	m := map[string]int{"orange": 100, "apple": 50, "banana": 30}
	for k, v := range m {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
	return
	main1()
	main2()
}
func main1() {
	fmt.Println("Start practice")
	//a := [...]int{3, 5, 11, 20}
	a := []int{3, 5, 11, 20}
	a = append(a, 30)
	a = append(a, 33)
	//for i=0; i<10; i++ {
	for _, v := range a {
		fmt.Printf("i=%d\n", v)
	}

	m := map[string]int{ "key1":3, "key2":50 }
	for key, val := range m {
		fmt.Printf("key=%s val=%d\n", key, val)
	}

	
	fmt.Println("End practice\n")
}


func main2() {
	fmt.Println("Start practice")

	//a := make([]int, 0, 1024)
	var a []int
	for i:=0; i<10; i++ {
		a = append(a, 2*i)
	}
	var i, v int
 	for i, v = range a {
		fmt.Printf("a[%d]=%d\n", i, v)
 	}
	fmt.Printf("a[%d]=%d\n", i, v)
	i, v = 3, 3
 	for i, v := range a {
		fmt.Printf("a[%d]=%d\n", i, v)
 	}
	fmt.Printf("a[%d]=%d\n", i, v)

	fmt.Println("End practice")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sat Oct 30 15:58:47
//  
// go run practice04.go
// Start practice
// i=3
// i=5
// i=11
// i=20
// i=30
// i=33
// key=key1 val=3
// key=key2 val=50
// End practice
//  
// Start practice
// a[0]=0
// a[1]=2
// a[2]=4
// a[3]=6
// a[4]=8
// a[5]=10
// a[6]=12
// a[7]=14
// a[8]=16
// a[9]=18
// a[9]=18
// a[0]=0
// a[1]=2
// a[2]=4
// a[3]=6
// a[4]=8
// a[5]=10
// a[6]=12
// a[7]=14
// a[8]=16
// a[9]=18
// a[3]=3
// End practice
//  
// Compilation finished at Sat Oct 30 15:58:47
