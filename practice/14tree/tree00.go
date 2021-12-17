package main
import "golang.org/x/tour/tree"
import (
	"fmt"
)
/*
   subject : A Tour of Go Exercise: Equivalent Binary Trees.  study Tree struct.
*/
//type Tree struct {
// 	Left  *Tree
// 	Value int
// 	Right *Tree
//}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return true
}

func main() {
	tt := tree.New(1)
 
	fmt.Println(tt)
	fmt.Printf("%T\n", tt)
	for t:=tt;; {
		fmt.Println(t.Left, t.Value, t.Right)
		t = t.Left
		if t == nil { break}
	}
	for t:=tt;; {
		fmt.Println(t.Left, t.Value, t.Right)
		t = t.Right
		if t == nil { break}
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/14tree/" -*-
// Compilation started at Wed Oct 20 23:35:27
//  
// go run tree00.go
// ((((1 (2)) 3 (4)) 5 ((6) 7 ((8) 9))) 10)
// *tree.Tree
// (((1 (2)) 3 (4)) 5 ((6) 7 ((8) 9))) 10 ()
// ((1 (2)) 3 (4)) 5 ((6) 7 ((8) 9))
// (1 (2)) 3 (4)
// () 1 (2)
// (((1 (2)) 3 (4)) 5 ((6) 7 ((8) 9))) 10 ()
//  
// Compilation finished at Wed Oct 20 23:35:28
