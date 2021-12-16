package main
import (
	"bytes"
	"fmt"
	"regexp"
)
/*
   subject : Regular Expression
*/
func main() {

	match, _ := regexp.MatchString(`p([a-z]+)ch`, "peach")
	fmt.Println(match)//true

	r, _ := regexp.Compile(`p([a-z]+)ch`)
	fmt.Println(r)                                       //p([a-z]+)ch
	fmt.Printf("%v %T\n", r, r)                          //p([a-z]+)ch *regexp.Regexp

	fmt.Println("MatchString               ", r.MatchString("peach"))//true

	fmt.Println("FindString                ", r.FindString("peach punch"))//peach

	fmt.Println("FindStringIndex           ", r.FindStringIndex("peach punch"))//[0 5]

	fmt.Println("FindStringSubmatch        ", r.FindStringSubmatch("peach punch"))//[peach ea]

	fmt.Println("FindStringSubmatchIndex   ", r.FindStringSubmatchIndex("peach punch"))//[0 5 1 3]

	fmt.Println("FindAllString -1          ", r.FindAllString("peach punch pinch", -1))//[peach punch pinch]

	fmt.Println("FindAllStringSubmatchIndex", r.FindAllStringSubmatchIndex("peach punch pinch", -1))//[[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	fmt.Println("FindAllString 2           ", r.FindAllString("peach punch pinch", 2))//[peach punch]

	fmt.Println("Match []byte              ", r.Match([]byte("peach")))//true

	r = regexp.MustCompile(`p([a-z]+)ch`)
	fmt.Println(r)                                       //p([a-z]+)ch
	fmt.Printf("%v %T\n", r, r)                          //p([a-z]+)ch *regexp.Regexp

	fmt.Println("ReplaceAllString          ", r.ReplaceAllString("a peach", "<fruit>"))//a <fruit>

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println("ReplaceAllFunc            ", string(out))                             //a PEACH
}
// -*- mode: compilation; default-directory: "~/go/src/practice/08misc/" -*-
// Compilation started at Sat Oct  9 19:30:40
//  
// go run misc08.go
// true
// p([a-z]+)ch
// p([a-z]+)ch *regexp.Regexp
// MatchString                true
// FindString                 peach
// FindStringIndex            [0 5]
// FindStringSubmatch         [peach ea]
// FindStringSubmatchIndex    [0 5 1 3]
// FindAllString -1           [peach punch pinch]
// FindAllStringSubmatchIndex [[0 5 1 3] [6 11 7 9] [12 17 13 15]]
// FindAllString 2            [peach punch]
// Match []byte               true
// p([a-z]+)ch
// p([a-z]+)ch *regexp.Regexp
// ReplaceAllString           a <fruit>
// ReplaceAllFunc             a PEACH
//  
// Compilation finished at Sat Oct  9 19:30:41
