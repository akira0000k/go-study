package main
import (
	"fmt"
)
/*
   subject: class test. object maker newSomeclass()
*/

type Object struct {
	classname string
}
func newObject() Object {
	obj := Object{"Object"}
	return obj
}

type Human struct {
	temper string
	Object
}
func newHuman() Human {
	var hu = Human{ "calm", Object{"Human"} }
	return hu
}
		
type Person struct {
	name string
	Human
}
func newPerson() Person {
	var ps Person = Person{ "noname", Human{ "calm", Object{"Person"}}}
	return ps
}

func main() {
	fmt.Println("main start")

	var me Person = newPerson()
	
	fmt.Println(me)

	me.name = "akira"
	me.temper = "angry"

	you := newPerson()
	you.name = "hana"
	you.temper = "sleepy"

	fmt.Println(me)
	fmt.Println(you)
	
	fmt.Println(me.temper)
	fmt.Println(me.classname)
	fmt.Println(you.temper)
	fmt.Println(you.classname)
	
	fmt.Println("main end")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/07class/" -*-
// Compilation started at Mon Oct  4 23:29:48
//  
// go run class02.go
// main start
// {noname {calm {Person}}}
// {akira {angry {Person}}}
// {hana {sleepy {Person}}}
// angry
// Person
// sleepy
// Person
// main end
//  
// Compilation finished at Mon Oct  4 23:29:49
