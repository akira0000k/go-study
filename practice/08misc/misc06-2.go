package main

import (
	"fmt"
	s "strings"
)
/*
   subject : String Functions. rune index function
*/
var pf = fmt.Printf
var p = fmt.Println

func main() {

	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.Contains(arg1, arg2)) 
	}("Contains", "test", "es")

	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.Count(arg1, arg2)) 
	}("Count", "test", "t")

	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.HasPrefix(arg1, arg2)) 
	}("HasPrefix", "test", "te")

	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.HasSuffix(arg1, arg2)) 
	}("HasSuffix", "test", "st")

	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.Index(arg1, arg2)) 
	}("Index", "test", "t")
	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.Index(arg1, arg2)) 
	}("Index", "test", "e")
	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.Index(arg1, arg2)) 
	}("Index", "test", "s")
	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.Index(arg1, arg2)) 
	}("Index", "test", "t")
	func(name, arg1, arg2 string) {
		pf("%s(%s, %s) = %v\n", name, arg1, arg2, s.Index(arg1, arg2)) 
	}("Index", "test", "x")

	func(name string, arg1 []string, arg2 string) {
		pf("%s(%v, %s) = %v\n", name, arg1, arg2, s.Join(arg1, arg2))
	}("Join", []string{"a", "b", "c"}, "-")
	
	func(name, arg1 string, arg2 int) {
		pf("%s(%s, %v) = %v\n", name, arg1, arg2, s.Repeat(arg1, arg2)) 
	}("Repeat", "a", 6)

	sReplace("foobaaaboo", "o", "0", -2)
	sReplace("foobaaaboo", "o", "0", -1)
	sReplace("foobaaaboo", "o", "0", 0)
	sReplace("foobaaaboo", "o", "0", 1)
	sReplace("foobaaaboo", "o", "0", 2)
	sReplace("foobaaaboo", "o", "0", 3)
	sReplace("foobaaaboo", "o", "0", 4)
	sReplace("foobaaaboo", "o", "0", 5)
	
	sSplit("a-b-c-d-e", "-")
	sSplit("a-b*c-*-e", "*")
	
	sToLower("TEST")
	sToLower("TeSt")
	sToUpper("test")
	sToUpper("tEsT")
	p("\n")
	sLen("hello")
	sLen("hello-HELLO")
	sChar("hello"[0])
	sChar("hello"[1])
	sChar("hello"[2])
	sChar("hello"[3])
	sChar("hello"[4])
	sRune("雲呑麺", 0)
	sRune("雲呑麺", 1)
	sRune("雲呑麺", 2)
	sRune("wang", 0)
	sRune("wang", 1)
	sRune("wang", 2)
	sRune("wang", 3)
}

func sReplace(arg1, arg2, arg3 string, arg4 int) {
	pf("%s(%s, %s, %s, %d) = %v\n", "Replace", arg1, arg2, arg3, arg4, s.Replace(arg1, arg2, arg3, arg4))
}
func sSplit(arg1, arg2 string) {
	pf("%s(%s, %s) = %v\n", "Split", arg1, arg2, s.Split(arg1, arg2)) 
}
func sToLower(arg1 string) {
	pf("%s(%s) = %v\n", "ToLower", arg1, s.ToLower(arg1)) 
}
func sToUpper(arg1 string) {
	pf("%s(%s) = %v\n", "ToUpper", arg1, s.ToUpper(arg1)) 
}
func sLen(arg1 string) {
	pf("%s(%s) = %v\n", "len", arg1, len(arg1)) 
}

func sChar(arg1 byte) {
	pf("%c %d\n", arg1, arg1)
}
func sRune(arg1 string, arg2 int) {
	count := 0
	for i, r := range arg1 {
		if count == arg2 {
			pf("%d %d %T %c\n", count, i, r, r)
			return
		}
		count++
	}
}

// -*- mode: compilation; default-directory: "~/go/src/practice/08misc/" -*-
// Compilation started at Fri Oct  8 22:58:27
//  
// go run misc06-2.go
// Contains(test, es) = true
// Count(test, t) = 2
// HasPrefix(test, te) = true
// HasSuffix(test, st) = true
// Index(test, t) = 0
// Index(test, e) = 1
// Index(test, s) = 2
// Index(test, t) = 0
// Index(test, x) = -1
// Join([a b c], -) = a-b-c
// Repeat(a, 6) = aaaaaa
// Replace(foobaaaboo, o, 0, -2) = f00baaab00
// Replace(foobaaaboo, o, 0, -1) = f00baaab00
// Replace(foobaaaboo, o, 0, 0) = foobaaaboo
// Replace(foobaaaboo, o, 0, 1) = f0obaaaboo
// Replace(foobaaaboo, o, 0, 2) = f00baaaboo
// Replace(foobaaaboo, o, 0, 3) = f00baaab0o
// Replace(foobaaaboo, o, 0, 4) = f00baaab00
// Replace(foobaaaboo, o, 0, 5) = f00baaab00
// Split(a-b-c-d-e, -) = [a b c d e]
// Split(a-b*c-*-e, *) = [a-b c- -e]
// ToLower(TEST) = test
// ToLower(TeSt) = test
// ToUpper(test) = TEST
// ToUpper(tEsT) = TEST
//  
//  
// len(hello) = 5
// len(hello-HELLO) = 11
// h 104
// e 101
// l 108
// l 108
// o 111
// 0 0 int32 雲
// 1 3 int32 呑
// 2 6 int32 麺
// 0 0 int32 w
// 1 1 int32 a
// 2 2 int32 n
// 3 3 int32 g
//  
// Compilation finished at Fri Oct  8 22:58:28
