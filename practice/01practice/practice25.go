package main
import (
	"fmt"
	"encoding/xml"
)

/*
   subject : サンプルで学ぶGo言語: XML の　struct initialize
*/
type pplant Plant
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}
func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}


func main() {
	fmt.Println("main start")

	type sss struct {
		id int
		name string
	}
	var s sss
	fmt.Println(s)
	s = sss{3, "akira"}
	fmt.Println(s)
	s = sss{name:"hana", id:4}
	fmt.Println(s)
	fmt.Println("\n------")
	
	fmt.Printf("%v\n", s)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%#v\n", s)
	fmt.Println("\n------")

	var p Plant
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	var pp pplant
	fmt.Printf("%v\n", pp)
	fmt.Printf("%+v\n", pp)
	fmt.Printf("%#v\n", pp)
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sun Oct 10 21:40:20
//  
// go run practice25.go
// main start
// {0 }
// {3 akira}
// {4 hana}
//  
// ------
// {4 hana}
// {id:4 name:hana}
// main.sss{id:4, name:"hana"}
//  
// ------
// Plant id=0, name=, origin=[]
// Plant id=0, name=, origin=[]
// main.Plant{XMLName:xml.Name{Space:"", Local:""}, Id:0, Name:"", Origin:[]string(nil)}
// {{ } 0  []}
// {XMLName:{Space: Local:} Id:0 Name: Origin:[]}
// main.pplant{XMLName:xml.Name{Space:"", Local:""}, Id:0, Name:"", Origin:[]string(nil)}
//  
// Compilation finished at Sun Oct 10 21:40:20
