package main
import (
	"fmt"
)

/*
   subject : read map dict recursively　　　とほほのGo言語入門
*/ 
type any interface{}
type dict map[string]any

var p0 dict= dict {
	"name": "Taro",
	"age": 20,
	"address": dict {
		"zip": "123-4567",
		"tel": "01-2345-6789",
		"city": "waocity",
		},
	}

func printdic(d dict) {
	for k, v := range d {
		switch v.(type) {
		case dict: printdic(v.(dict))
		default: fmt.Println("key=", k, "val=", v)
		}
	}
}

func main() {
	fmt.Println("start")
	//fmt.Println(p0)

	p1 := dict {
		"name": "Hanako",
		"age": 16,
		"address": dict {
			"zip": "123-4567",
			"tel": "01-2345-6789",
			"city": "waocity",
		},
	}
	//fmt.Println(p1)
	fmt.Println(p1["name"], p1["age"], p1["address"].(dict)["zip"])
	fmt.Println("-------")
	for k, v := range p1 {
		fmt.Println("key=", k, "val=", v)
	}
	fmt.Println("-------")
	printdic(p1)
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Fri Nov 26 21:06:50
//  
// go run practice17.go
// start
// Hanako 16 123-4567
// -------
// key= name val= Hanako
// key= age val= 16
// key= address val= map[city:waocity tel:01-2345-6789 zip:123-4567]
// -------
// key= zip val= 123-4567
// key= tel val= 01-2345-6789
// key= city val= waocity
// key= name val= Hanako
// key= age val= 16
//  
// Compilation finished at Fri Nov 26 21:06:51
