package main

import "fmt"


/*
   subject : return value and flag
 */ 
func isPerson(name string) (string, bool) {
	if name == "akira" {
		return name, true
	} else if name == "tiger" {
		return name, false
	} else {
		return "what", false
	}
}
func isFamily(name string) (string, string, bool) {
	family := "Honda"
	switch name {
	case "akira":
		fallthrough
	case "hana", "takashi":
		return family, name, true
	default:
		return "", name, false
		
	}
}

func main() {
	family, namae, ok := isFamily("akira")
	fmt.Println(family, namae, ok)
	family2, namae2, ok := isFamily("hana")
	fmt.Println(family2, namae2, ok)

	
	family, namae, _ = isFamily("hana")
	fmt.Println(family, namae)
	family, _, _ = isFamily("takashi")
	fmt.Println(family)
	family, namae, ok = isFamily("debeso")
	fmt.Println(family, namae, ok)


	fmt.Println(isFamily("akira"))
	fmt.Println(isFamily("hana"))
	fmt.Println(isFamily("takashi"))
	fmt.Println(isFamily("fire"))

	
	name, ok := isPerson("akira")
	if ok {
		fmt.Printf("%s is a person.\n", name)
	}
	name, ok = isPerson("tiger")
	if ok {
		fmt.Printf("%s is a person.\n", name)
	}
	//name, ok := isPerson("bobozu")    //./practice9.go:26:11: no new variables on left side of :=
	name2, ok := isPerson("bobozu")   //ok is ok
	if ok {
		fmt.Printf("%s is a person.\n", name2)
	}
	fmt.Println(isPerson("akira"))   //akira true
	fmt.Println(isPerson("tiger"))   //akira true
	fmt.Println(isPerson("baka"))   //akira true
	//fmt.Println("res", isPerson("akira"))    //./practice9.go:31:29: multiple-value isPerson() in single-value context
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sun Nov 28 20:36:31
//  
// go run practice09.go
// Honda akira true
// Honda hana true
// Honda hana
// Honda
//  debeso false
// Honda akira true
// Honda hana true
// Honda takashi true
//  fire false
// akira is a person.
// akira true
// tiger false
// what false
//  
// Compilation finished at Sun Nov 28 20:36:32
