package main
import "fmt"
/*
   subject : interface{} に入れた値も型は持っていて %v で正しく表示される。`abcd` をint32に入れたりは出来ないと言う事。
*/
func main() {
	var i interface{}
	i = 255
	fmt.Println(i)
	fmt.Printf("T:%T  v:%v d:%d  f:%f  s:%s x:%x\n", i, i, i, i, i, i)
	var ui uint
	ui, ok := i.(uint)
	fmt.Println("uint:", ui, ok)

	
	i = 345.678
	fmt.Println(i)
	fmt.Printf("T:%T  v:%v d:%d  f:%f  s:%s x:%x\n", i, i, i, i, i, i)
	var f3 float32
	f3, ok = i.(float32)
	fmt.Println("float32:", f3, ok)
	
	i = "0123456789ABCDabcd"
	fmt.Println(i)
	fmt.Printf("T:%T  v:%v d:%d  f:%f  s:%s x:%x\n", i, i, i, i, i, i)
}
// -*- mode: compilation; default-directory: "~/go/src/practice/13interface/" -*-
// Compilation started at Fri Oct 22 23:50:44
//  
// go run inter01-2.go
// 255
// T:int  v:255 d:255  f:%!f(int=255)  s:%!s(int=255) x:ff
// uint: 0 false
// 345.678
// T:float64  v:345.678 d:%!d(float64=345.678)  f:345.678000  s:%!s(float64=345.678) x:0x1.59ad916872b02p+08
// float32: 0 false
// 0123456789ABCDabcd
// T:string  v:0123456789ABCDabcd d:%!d(string=0123456789ABCDabcd)  f:%!f(string=0123456789ABCDabcd)  s:0123456789ABCDabcd x:303132333435363738394142434461626364
//  
// Compilation finished at Fri Oct 22 23:50:45
