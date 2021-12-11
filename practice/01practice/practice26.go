package main
import (
	"fmt"
)
/*
   subject : assignment outside function
*/
func main() {
	fmt.Println("start main")

	fmt.Println(v1, v2, v3)
}

var v3 = v1 + v2
var v2 = v1 + 1
var v1 = isquare(5)


func isquare(i int) int {
	fmt.Println("isquare", i)
	return i * i
}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Fri Nov 26 20:18:29
//  
// go run practice26.go
// isquare 5
// start main
// 25 26 51
//  
// Compilation finished at Fri Nov 26 20:18:30
