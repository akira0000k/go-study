package main

import "fmt"

/*
   subject : サンプルで学ぶ Go 言語：Structs
 */ 
type person struct {
	name string
	age  int
}
type Person struct {
	name string
	age  int
	call int
}

func newPerson(name string) *Person {

	p := Person{name: name}
	p.age = 42
	p.call = 0
	return &p
}

func (p *Person) getAge() int {
	p.call++
	return p.age
}
func (q Person) getAgeage() int {
	q.call++         //nonsense
	return q.age
}


func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	fmt.Println("------------")

	//fmt.Println(newPerson("Jon"))
	np := newPerson("Jon")
	fmt.Println(np)
	
	fmt.Println((*np).name, np.age)
	fmt.Println(np.name, (*np).age)
	
	fmt.Println(np.getAgeage(), np.getAgeage())
	fmt.Println(np)                              //&{Jon 42 0}
	fmt.Println(np.getAge(), np.getAge())
	fmt.Println(np)                              //&{Jon 42 2}
	fmt.Println(np.getAgeage(), np.getAgeage())
	fmt.Println(np)                              //&{Jon 42 2}
	fmt.Println(np.getAge(), np.getAge())
	fmt.Println(np)                              //&{Jon 42 4}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sat Oct 30 00:23:08
//  
// go run practice08.go
// {Bob 20}
// {Alice 30}
// {Fred 0}
// &{Ann 40}
// Sean
// 50
// 51
// ------------
// &{Jon 42 0}
// Jon 42
// Jon 42
// 42 42
// &{Jon 42 0}
// 42 42
// &{Jon 42 2}
// 42 42
// &{Jon 42 2}
// 42 42
// &{Jon 42 4}
//  
// Compilation finished at Sat Oct 30 00:23:09
