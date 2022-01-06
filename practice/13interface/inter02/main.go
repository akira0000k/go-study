package main

import (
	"fmt"
	"go-study/practice/13interface/inter02/funky"
	"go-study/practice/13interface/inter02/monkey"
)

type carol interface {
	GetAmero(i int) string
	GetBmero(i int) string
}

func main() {
	fmt.Println("** start main **")

	var t carol
	//var f funky.Funkytype = funky.Funkytype{ "Ho Ho Ho" }
	var f = funky.Funkytype{ "Ho Ho Ho" }
	t = &f
	
	println(t.GetAmero(1)) //(t  Funkytype)
	println(t.GetBmero(2)) //(t *Funkytype) 一つでも*がつく関数があると、&f のようにpointerにしなければならない

	m := monkey.Monkeytype{ "boo boo boo" }
	t = m

	println(t.GetAmero(3)) //(t Monkeytype)
	println(t.GetBmero(4)) //(t Monkeytype)
	
	fmt.Println("** END main **")
}
// -*- mode: compilation; default-directory: "~/go/src/go-study/practice/13interface/inter02/" -*-
// Compilation started at Tue Jan  4 22:58:24
//  
// go run main.go
// ** start main **
// [1] 君は Funky Baby おどけてるよ Ho Ho Ho
// [2] 君がいなけりゃ Baby I'm blue Ho Ho Ho
// [3] 君は Monkey Baby いかれてるよ boo boo boo
// [4] 君がいなけりゃ Baby I'm so sad boo boo boo
// ** END main **
//  
// Compilation finished at Tue Jan  4 22:58:24
