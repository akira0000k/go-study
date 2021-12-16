package main

import (
	"fmt"
	"os"
)
/*
   subject: using defer for closing file
*/
func main() {
	f := createFile("./defer.txt")
	defer closeFile(f)

	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating...")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing...")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing...")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/08misc/" -*-
// Compilation started at Thu Oct  7 00:19:57
//  
// go run misc04.go
// creating...
// writing...
// closing...
//  
// Compilation finished at Thu Oct  7 00:19:58
