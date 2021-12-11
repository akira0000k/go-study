package main

import "fmt"

/*
   subject : func戻り値のカッコ　サンプルで学ぶ Go 言語：Structs
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

func newPerson(name string) (*Person) {

	p := Person{name: name}
	p.age = 42
	p.call = 0
	return &p
}

func (p *Person) getAge() int { //戻り値にカッコはなくても良い
	p.call++
	p.age++
	return p.age
}
func (p *Person) getAge2() (int) { //あっても良い
	p.call++
	p.age++
	return p.age
}
func (p *Person) getAge3() (newage int) { //なくちゃだめ
	p.call++
	p.age++
	newage = p.age
	return
}
func (q Person) getAgeage() (int, int) { //なくちゃだめ
	q.call++         //nonsense
	q.age++
	return q.age, q.call
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
	
	fmt.Println(np.getAgeage())
	fmt.Println(np.getAgeage())
	fmt.Println(np)                              //&{Jon 42 0}
	fmt.Println(np.getAge(), np.getAge())
	fmt.Println(np)                              //&{Jon 42 2}
	fmt.Println(np.getAgeage())
	fmt.Println(np.getAgeage())
	fmt.Println(np)                              //&{Jon 42 2}
	fmt.Println(np.getAge2(), np.getAge())
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
