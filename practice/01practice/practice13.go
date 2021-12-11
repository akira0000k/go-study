package main

import (
	"fmt"
)

/*
   subject : array initialize
 */ 
func main1() {
	fmt.Println("Start")


	var ia [5]int
	ia[0] = 3
	ia[1] = 23
	ia[2] = 113
	fmt.Println(ia)

	var ja [5]int = [5]int{1,2,3,4}
	ja[0] = 3
	ja[1] = 23
	ja[2] = 113
	fmt.Println(ja)
	
	a := [5]int{1,2,3,4,5}
	a[0] = 3
	a[1] = 23
	a[2] = 113
	fmt.Println(a)

	b := [...]int{1,2}//,3,4,5,6}
	b[0] = 3
	b[1] = 23
	//b[2] = 113 //./practice13.go:32:3: invalid array index 2 (out of bounds for 2-element array)
	fmt.Println(b)

}

func main2() {
	fmt.Println("Start")

	var sa [3]string
	sa[1] = "BB"
	fmt.Println(sa)

	var sb [3]string = [3]string{"aa", "bb"}//, "cc"}
	sb[1] = "BB"
	fmt.Println(sb)

	sc := [3]string{"aa", "bb", "cc"}
	sc[1] = "BB"
	fmt.Println(sc)
}

func main3() {
	fmt.Println("Start")

	var i int
	i = 33
	fmt.Println(i)

	r := 34.567
	//var ii int = int(34.567) //./practice13.go:61:18: constant 34.567 truncated to integer
	//var ii int = r  //./practice13.go:62:6: cannot use r (type float64) as type int in assignment
	var ii int = int(r)
	fmt.Println(ii)
}

func main() {
	main1()
	main2()
	main3()
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Fri Oct 29 23:52:43
//  
// go run practice13.go
// Start
// [3 23 113 0 0]
// [3 23 113 4 0]
// [3 23 113 4 5]
// [3 23]
// Start
// [ BB ]
// [aa BB ]
// [aa BB cc]
// Start
// 33
// 34
//  
// Compilation finished at Fri Oct 29 23:52:44
