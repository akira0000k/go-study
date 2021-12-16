package main
import (
	"fmt"
)

/*
   subject : read map dict recursively
*/ 
type any interface{}
type dict map[string]any

//var p0 dict= dict {
// 	"name": "Akira",
// 	"age": 70,
// 	"address": dict {
// 		"zip": "135-0062",
// 		"tel": "03-3534-7320",
// 		"city": "shinonome",
// 		},
// 	}
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
		"name": "Akira",
		"age": 70,
		"address": dict {
			"zip": "135-0062",
			"tel": "03-3534-7320",
			"city": "shinonome",
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
// -*- mode: compilation; default-directory: "~/go/src/practice/09json/" -*-
// Compilation started at Sat Oct  9 20:54:08
//  
// go run practice17.go
// start
// Akira 70 135-0062
// -------
// key= age val= 70
// key= address val= map[city:shinonome tel:03-3534-7320 zip:135-0062]
// key= name val= Akira
// -------
// key= name val= Akira
// key= age val= 70
// key= city val= shinonome
// key= zip val= 135-0062
// key= tel val= 03-3534-7320
//  
// Compilation finished at Sat Oct  9 20:54:09
