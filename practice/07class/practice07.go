package main

import "fmt"
import "strconv"

/*
   subject : struct function1　　とほほのGo言語入門
 */ 
type Person struct {
	name string
	age int
}

func (p *Person) SetPerson(name string, age int) {
	p.name = name
	p.age = age
}
 
func (p *Person) GetPerson() (string, int) {
	page := strconv.Itoa(p.age)
	return p.name + "(" + page + ")", p.age
}

func main() {
	var p1 Person
	p1.SetPerson("Yamada", 26)
	name, age := p1.GetPerson()
	fmt.Println(name, age)

	fmt.Println(p1)
	//fmt.Println(Person{"akira"})   //./practice7.go:28:21: too few values in Person{...}
	fmt.Println(Person{name:"akira"})
	fmt.Println(Person{age: 27})
	fmt.Println(Person{"akira", 75})
	fmt.Println(Person{age:100, name:"saboru"})
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sat Oct 30 00:27:02
//  
// go run practice07.go
// Yamada(26) 26
// {Yamada 26}
// {akira 0}
// { 27}
// {akira 75}
// {saboru 100}
//  
// Compilation finished at Sat Oct 30 00:27:03
