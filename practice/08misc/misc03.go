package main

import "os"
/*
   subject: panic stop
*/
func sub() {
	_, err := os.Create("/tmp/akira/file")
	if err != nil {
		panic(err)
	}
}

func main() {
	//panic("a problem")
	sub()
}
// -*- mode: compilation; default-directory: "~/go/src/practice/08misc/" -*-
// Compilation started at Thu Oct  7 00:07:24
//  
// go run misc03.go
// panic: open /tmp/akira/file: no such file or directory
//  
// goroutine 1 [running]:
// main.sub()
//  	/Users/Akira/go/src/practice/08misc/misc03.go:8 +0x7a
// main.main()
//  	/Users/Akira/go/src/practice/08misc/misc03.go:14 +0x25
// exit status 2
//  
// Compilation exited abnormally with code 1 at Thu Oct  7 00:07:24
