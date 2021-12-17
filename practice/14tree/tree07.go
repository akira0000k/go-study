package main
import "golang.org/x/tour/tree"
import (
	"fmt"
	"sync"
	"math/rand"
)
/*
   subject : A Tour of Go Exercise: Equivalent Binary Trees.  making random tree
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
	ch1 := make(chan int)//, 10)
	ch2 := make(chan int)//, 10)
	var result bool
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		result = false
		for {
			var v1, v2 int
			var ok1, ok2 bool
			v1, ok1 = <-ch1
			v2, ok2 = <-ch2
			if ok1 && ok2 && v1==v2 {
				continue
			} else if !ok1 && !ok2 {
				result = true
				return //true
			}

			for ok1 || ok2 {
				if ok1 {
					_, ok1 = <-ch1
				}
				if ok2 {
					_, ok2 = <-ch2
				}
			}
			return //false
		}
	}()
	
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(ch1)
		Walk(t1, ch1)
	}()
	go func() {
		defer wg.Done()
		defer close(ch2)
		Walk(t2, ch2)
	}()

	wg.Wait()
	
	return result
}

func addTree(t *tree.Tree, val int) *tree.Tree {
	if t == nil {
		tt := new(tree.Tree)
		tt.Value = val
		return tt
	} else {
		if val < t.Value {
			t.Left = addTree(t.Left, val)
		} else {
			t.Right = addTree(t.Right, val)
		}
	}
	return t
}

func newTree(from, to int) *tree.Tree {
	n := to - from
	tree := addTree(nil, rand.Intn(n) + from)
	for i:=0; i<n; i++ {
		addTree(tree, rand.Intn(n) + from)
	}
	return tree
}
	
func main() {
	var result bool
	fm := 10
	to := 20

	imax := 100000
	for i:=0; i<imax; i++ {
		tt1 := newTree(fm, to)
		tt2 := newTree(fm, to)
		
		result = Same(tt1, tt2)
		if result {
			fmt.Println(i)
			fmt.Println("tree1=", tt1)
			fmt.Println("tree2=", tt2)
			fmt.Println(result)
			return
		}
	}
	fmt.Println("all false", imax)
}
// -*- mode: compilation; default-directory: "~/go/src/practice/14tree/" -*-
// Compilation started at Thu Oct 21 23:04:15
//  
// go run tree07.go
// 16028
// tree1= (((10 (12)) 13 (13 (14))) 15 ((15) 16 ((16) 17 (18))))
// tree2= (10 ((((12) 13 (13)) 14 ((15 (15)) 16 (16))) 17 (18)))
// true
//  
// Compilation finished at Thu Oct 21 23:04:16
