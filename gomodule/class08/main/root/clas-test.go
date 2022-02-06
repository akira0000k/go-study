package main

import (
	"fmt"
	"go.mod/persons"
)

func main() {
	fmt.Println("main start")

	var me persons.Person = persons.NewPerson()
	
	fmt.Println(me)

	me.Name = "akira"
	fmt.Println(me.Ban())
	me.SetClassname("Retired")

	you := persons.NewPerson()
	you.Name = "hana"
	you.Temper = "sleepy"
	you.SetClassname("OnotherWorld")
	
	fmt.Println(me)
	fmt.Println(you)
	
	fmt.Println(me.Temper)
	fmt.Println(me.GetClassname())
	fmt.Println(you.Temper)
	fmt.Println(you.GetClassname())

	nb := persons.PConcat(me, you)
	fmt.Println(nb)
	
	pp := persons.NConcat(&me, &you)
	fmt.Println(pp)
	fmt.Println(*pp)
	fmt.Println(pp.Name, pp.Temper, pp.GetClassname())
	
	
	fmt.Println("main end")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/07class/" -*-
// Compilation started at Tue Oct  5 17:33:29
//  
// go run class04.go
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
// Compilation finished at Tue Oct  5 17:33:30
