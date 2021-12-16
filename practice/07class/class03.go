package main
import (
	"fmt"
)
/*
   subject: class test. func concat
*/

type Object struct {
	classname string
}
func (o *Object) SetClassname(name string) {
	o.classname = name
}
func newObject() Object {
	//var obj = Object{"Object"}
	var obj Object
	obj.SetClassname("Object")
	return obj
}

type Human struct {
	temper string
	Object
}
func (hu *Human) Ban() string {
	hu.temper = "angry"
	return "nandayou"
}
func newHuman() Human {
	//var hu Human = Human{ "calm", newObject() }
	var hu Human
	hu.temper = "easy"
	hu.Object = newObject()
	hu.SetClassname("Human")
	return hu
}
		
type Person struct {
	name string
	Human
}
func newPerson() Person {
	//var ps Person = Person{ "noname", newHuman() }
	var ps Person // = Person{}
	ps.name = "Anon"
	ps.Human = newHuman()
	ps.SetClassname("Person")
	return ps
}

func pConcat(p1, p2 Person) Person {
	np := newPerson()
	np.name = p1.name + p2.name
	np.temper = p1.temper + p2.temper
	np.SetClassname(p1.classname)
	return np
}
func nConcat(p1, p2 *Person) *Person {
	np := new(Person)
	np.name = p1.name + p2.name
	np.temper = p1.temper + p2.temper
	np.SetClassname(p1.classname)
	return np
	//return &np
}

func main() {
	fmt.Println("main start")

	var me Person = newPerson()
	
	fmt.Println(me)

	me.name = "akira"
	fmt.Println(me.Ban())
	me.SetClassname("Retired")

	you := newPerson()
	you.name = "hana"
	you.temper = "sleepy"
	you.classname = "OnotherWorld"
	
	fmt.Println(me)
	fmt.Println(you)
	
	fmt.Println(me.temper)
	fmt.Println(me.classname)
	fmt.Println(you.temper)
	fmt.Println(you.classname)

	nb := pConcat(me, you)
	fmt.Println(nb)
	
	pp := nConcat(&me, &you)
	fmt.Println(pp)
	fmt.Println(*pp)
	fmt.Println(pp.name, pp.temper, pp.classname)
	
	
	fmt.Println("main end")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/07class/" -*-
// Compilation started at Tue Oct  5 10:26:11
//  
// go run class03.go
// main start
// {Anon {easy {Person}}}
// nandayou
// {akira {angry {Retired}}}
// {hana {sleepy {OnotherWorld}}}
// angry
// Retired
// sleepy
// OnotherWorld
// {akirahana {angrysleepy {Retired}}}
// &{akirahana {angrysleepy {Retired}}}
// {akirahana {angrysleepy {Retired}}}
// akirahana angrysleepy Retired
// main end
//  
// Compilation finished at Tue Oct  5 10:26:11
