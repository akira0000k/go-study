package main

import "fmt"

/*
   subject : multi return, multi assignment
 */ 

//func funckey(name string, age int) (string, bool) {
// 	return name, true
//}

//./practice10.go:9:6: funckey redeclared in this block	previous declaration at ./practice10.go:6:45
func funckey(name string) (string, bool) {
	return name, true
}

func main1() {
	fmt.Println("start")
	fmt.Println(funckey("stuvw"))
	//fmt.Printf("%s %d\n", funckey("stuvw"))  //./practice10.go:17:31: multiple-value funckey() in single-value context
	name, ok := funckey("stuvw")
	fmt.Printf("%s %t\n", name, ok)
	fmt.Println("end")
}

func funckt(name string) (bool, int, bool, string) {
	return true, 31,
		false, name
}
func main2() {
	fmt.Println("start")

	ok, age,
		ng, name := funckt("akira")
	fmt.Println(ok,
		age,
		ng,
		name,
	)
	
	ok2, age, ng, name := funckt("akira") //少なくとも一つ新変数が必要だが、ほかは代入になる
	fmt.Println(ok2, age, ng, name)

	ok, age2, ng, name := false, 55, true, "takashi"
	fmt.Println(ok, age2, ng, name)
	
	fmt.Println(funckt("hana"))
	//fmt.Println(funckt("hana"), 1, 2, 3) //./practice10.go:48:20: multiple-value funckt() in single-value context
	
	fmt.Println("end")
}

func main() {
	main1()
	main2()
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sat Oct 30 00:18:16
//  
// go run practice10.go
// start
// stuvw true
// stuvw true
// end
// start
// true 31 false akira
// true 31 false akira
// false 55 true takashi
// true 31 false hana
// end
//  
// Compilation finished at Sat Oct 30 00:18:17
