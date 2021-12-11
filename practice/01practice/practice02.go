package main

import "fmt"
import "time"

/*
   subject : string,  go routine
 */ 
func main() {
	main1()
	main2()
}
func main1() {
	fmt.Println("Start main1")

	var st, str, strn string="st", "str", "strn"

	fmt.Println("st=", st)
	fmt.Println("str=", str)
	fmt.Println("strn=", strn)
	strg := st + str + strn
	fmt.Println("strg=", strg)

	cct := ""
	for i:=0; i<3; i++ {
		cct += "\t"
	}
	fmt.Println(cct)
	
	fmt.Println("End main1\n")
}

func funcA() {
	for i:=0; i<30; i++ {
		fmt.Printf("A%d\n", i)
		time.Sleep(1*time.Millisecond)
	}
}
func main2() {
	fmt.Println("Start practice")

	main1()
	
	go funcA()
	//time.Sleep(1*time.Millisecond)
	for i:=0; i<10; i++ {
		fmt.Printf("M%d\n", i)
		time.Sleep(2*time.Millisecond)
	}
	fmt.Println("\nEnd practice")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sat Oct 30 16:04:16
//  
// go run practice02.go
// Start main1
// st= st
// str= str
// strn= strn
// strg= ststrstrn
//  			
// End main1
//  
// Start practice
// Start main1
// st= st
// str= str
// strn= strn
// strg= ststrstrn
//  			
// End main1
//  
// M0
// A0
// A1
// M1
// A2
// A3
// M2
// A4
// A5
// M3
// A6
// A7
// M4
// A8
// A9
// M5
// A10
// M6
// A11
// A12
// M7
// A13
// A14
// M8
// A15
// M9
// A16
// A17
// A18
//  
// End practice
//  
// Compilation finished at Sat Oct 30 16:04:17
