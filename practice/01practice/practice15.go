package main
import (
	"fmt"
)

/*
   subject : interface  とほほのGo言語入門
 */ 
type Printable interface {
	ToString() string
}
//func PrintOut(p interface {ToString() string}) {
func PrintOut(p Printable) {
	fmt.Println(p.ToString())
}

type Any interface{}
//func PrintAny(a interface{}) {
func PrintAny(a Any) {
	if
	//q, ok := a.(interface {ToString() string});
	q, ok := a.(Printable);
	ok {
		fmt.Println(q.ToString())
	} else {
		fmt.Println("Not Printable")
	}
}
func PrintAll(a interface{}) {
	fmt.Printf("%#v\n", a)
	switch a.(type) {
	case bool: fmt.Printf("%t\n", a)
	case int:  fmt.Printf("%d\n", a)
	case float64: fmt.Printf("%f\n", a)
	case string:  fmt.Printf("%s\n", a)
	case Printable: fmt.Println(a.(Printable).ToString())
	default: fmt.Println("not printable")
	}
}

type Person struct {
	name string
}
func (p Person) ToString() string {
	return p.name
}

type Book struct {
	titl string
}
func (b Book) ToString() string {
	return b.titl
}


func main() {
	fmt.Println("Start")
	a1 := Person{ name: "Person A1" }
	a2 := Book{ titl: "Book A2" }
	fmt.Println("-----PrintOut")
	PrintOut(a1)
	PrintOut(a2)
	//PrintOut(123) //cannot use 123 (type int) as type Printable in argument to PrintOut: int does not implement Printable (missing ToString method)
	//PrintOut(3.14) //cannot use 3.14 (type float64) as type Printable in argument to PrintOut: float64 does not implement Printable (missing ToString method)
	//PrintOut("PAI") //cannot use "PAI" (type string) as type Printable in argument to PrintOut: string does not implement Printable (missing ToString method)
	fmt.Println("-----PrintAny")
	PrintAny(a1)
	PrintAny(a2)
	PrintAny(123)  //Not Printable
	PrintAny(3.14) //Not Printable
	PrintAny("PAI")//Not Printable
	fmt.Println("-----PrintALL")
	PrintAll(a1)
	PrintAll(a2)
	PrintAll(123)
	PrintAll(3.14) //not printable
	PrintAll("PAI")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Fri Nov 26 20:53:37
//  
// go run practice15.go
// Start
// -----PrintOut
// Person A1
// Book A2
// -----PrintAny
// Person A1
// Book A2
// Not Printable
// Not Printable
// Not Printable
// -----PrintALL
// main.Person{name:"Person A1"}
// Person A1
// main.Book{titl:"Book A2"}
// Book A2
// 123
// 123
// 3.14
// 3.140000
// "PAI"
// PAI
//  
// Compilation finished at Fri Nov 26 20:53:38
