package main
import (
	"fmt"
)
/*
   subject: class test. Is-A
*/

type Object struct {
	classname string
}

type Human struct {
	temper string
	Object
}

type Person struct {
	name string
	Human
}

func main() {
	fmt.Println("main start")

	var me Person
	me.name = "akira"
	me.temper = "angry"
	me.classname = "Person"

	fmt.Println(me)

	fmt.Println(me.temper)
	fmt.Println(me.classname)
	
	fmt.Println("main end")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/07class/" -*-
// Compilation started at Mon Oct  4 22:15:57
//  
// go run class01.go
// main start
// {akira {angry {Person}}}
// angry
// Person
// main end
//  
// Compilation finished at Mon Oct  4 22:15:58
