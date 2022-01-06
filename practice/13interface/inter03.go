package main

import "fmt"

func main() {
	fmt.Println("start")

	type intf interface {
		hello()
	}

	var t intf
	var s sta

	t = &s
	t.hello()
	t.hello()
	t.hello()

	var b stb
	t = &b
	t.hello()
	t.hello()
	t.hello()
	b.vai()
	b.vai()
	t.(*stb).vai()
	t.(*stb).vai()
	t.(*stb).vai()
	t.(*stb).vai()

	type kenta interface {
		vai()
	}
	var k kenta
	k = &b
	k.vai()
	k.vai()
	k.vai()
	k.(*stb).hello()

	var v stv
	k = &v
	k.vai()
	k.vai()
	k.(*stv).jiro()
	k.(*stv).jiro()

}

//----hello no.1-----
type sta struct {
	seq int
}
func (t *sta) hello() {
	t.seq++
	fmt.Println("hello", t.seq)
}

//----hello no.2-----vai no1.------
type stb struct {
	seq float64
}
func (t *stb) hello() {
	t.seq += 1.0
	fmt.Printf("hello %.1f\n", t.seq)
}
func (t *stb) vai() {
	t.seq += 1.1111111
	fmt.Printf("vai   %.6f\n", t.seq)
}

//----vai no.2-----
type stv struct {
	seq float64
}
func (t *stv) jiro() {
	t.seq += 3.0
	fmt.Printf("jiro  %.1f\n", t.seq)
}
func (t *stv) vai() {
	t.seq += 1.1111111
	fmt.Printf("vai   %.6f\n", t.seq)
}
// -*- mode: compilation; default-directory: "~/go/src/go-study/practice/13interface/" -*-
// Compilation started at Tue Jan  4 22:43:28
//  
// go run inter03.go
// start
// hello 1
// hello 2
// hello 3
// hello 1.0
// hello 2.0
// hello 3.0
// vai   4.111111
// vai   5.222222
// vai   6.333333
// vai   7.444444
// vai   8.555556
// vai   9.666667
// vai   10.777778
// vai   11.888889
// vai   13.000000
// hello 14.0
// vai   1.111111
// vai   2.222222
// jiro  5.2
// jiro  8.2
//  
// Compilation finished at Tue Jan  4 22:43:29
