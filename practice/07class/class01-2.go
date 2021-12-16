package main
import (
	"fmt"
)
/*
   subject: class test. Has-A
*/

type Object struct {
	classname string
}

type Human struct {
	temper string
	obj Object
}

type Person struct {
	name string
	hu Human
}

func main() {
	fmt.Println("main start")

	var me Person
	me.name = "akira"
	me.hu.temper = "angry"
	me.hu.obj.classname = "Person"

	fmt.Println(me)

	fmt.Println(me.hu.temper)
	fmt.Println(me.hu.obj.classname)
	
	fmt.Println("main end")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/07class/" -*-
// Compilation started at Mon Oct  4 22:26:31
//  
// go run class01-2.go
// main start
// {akira {angry {Person}}}
// angry
// Person
// main end
//  
// Compilation finished at Mon Oct  4 22:26:32
