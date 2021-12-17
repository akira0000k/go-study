package main
import "golang.org/x/tour/tree"
import (
	"fmt"
	//"time"
	"sync"
)
/*
   subject : A Tour of Go Exercise: Equivalent Binary Trees.  Test Same Function.
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
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
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
			} else {
				return //false
			}
		}
	}()
	
	var wg2 sync.WaitGroup
	wg2.Add(2)
	go func() {
		defer wg2.Done()
		Walk(t1, ch1)
	}()
	go func() {
		defer wg2.Done()
		Walk(t2, ch2)
	}()

	wg2.Wait()
	close(ch1)
	close(ch2)
	wg.Wait()
	
	return result
}

func main() {
	k1, k2, k3, k4 := 1, 2, 5, 5
	for i:=0; i<2; i++ {
		tt1 := tree.New(k1)
		tt2 := tree.New(k2)
		fmt.Println("tree1=", k1, tt1)
		fmt.Println("tree2=", k2, tt2)
		result := Same(tt1, tt2)
		fmt.Println(result)
		k1, k2 = k3, k4
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/14tree/" -*-
// Compilation started at Thu Oct 21 00:13:45
//  
// go run tree02.go
// tree1= 1 ((((1 (2)) 3 (4)) 5 ((6) 7 ((8) 9))) 10)
// tree2= 2 ((((2) 4 (6)) 8 (10 (12))) 14 ((16) 18 (20)))
// false
// tree1= 5 ((((((5) 10) 15 (20)) 25 (30)) 35) 40 ((45) 50))
// tree2= 5 ((5 ((((10) 15 (20)) 25) 30)) 35 ((40) 45 (50)))
// true
//  
// Compilation finished at Thu Oct 21 00:13:45
