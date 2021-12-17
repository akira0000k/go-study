package main
import "golang.org/x/tour/tree"
import (
	"fmt"
	//"time"
	"sync"
)
/*
   subject : A Tour of Go Exercise: Equivalent Binary Trees.  Test Walk Function.
*/
//type Tree struct {
// 	Left  *Tree
// 	Value int
// 	Right *Tree
//}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return true
}

func main() {
	tt := tree.New(1)
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch {
			fmt.Println(val)
		}
	}()
	
	Walk(tt, ch)

	close(ch)
	wg.Wait()
}
// -*- mode: compilation; default-directory: "~/go/src/practice/14tree/" -*-
// Compilation started at Wed Oct 20 23:35:59
//  
// go run tree01.go
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10
//  
// Compilation finished at Wed Oct 20 23:35:59
