package main
import "golang.org/x/tour/tree"
import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)
/*
   subject : A Tour of Go Exercise: Equivalent Binary Trees.  stop compare flag
*/
//type Tree struct {
// 	Left  *Tree
// 	Value int
// 	Right *Tree
//}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, stp *bool){
	if *stp {
		return
	}
	if t.Left != nil {
		Walk(t.Left, ch, stp)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch, stp)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)//, 10)
	ch2 := make(chan int)//, 10)
	var result bool
	var stop bool
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
			stop = true
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
		Walk(t1, ch1, &stop)
	}()
	go func() {
		defer wg.Done()
		defer close(ch2)
		Walk(t2, ch2, &stop)
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

func newRandomTree(from, to int) *tree.Tree {
	n := to - from
	tree := addTree(nil, rand.Intn(n) + from)
	for i:=0; i<n; i++ {
		addTree(tree, rand.Intn(n) + from)
	}
	return tree
}
func newTree(from, to int) *tree.Tree {
	n := to - from + 1
	var tree *tree.Tree
	islice := make([]int, n)
	j := from
	for i:=0; i<n; i++ {
		islice[i] = j
		//islice = append(islice, j)
		j+=1
	}
	leng := n
	for i:=0; i<n; i++ {
		idx := rand.Intn(leng)
		num := islice[idx]
		tree = addTree(tree, num)
		islice[idx] = islice[leng - 1]
		leng -= 1
		//fmt.Println(tree)
	}
	return tree
}
	
func main() {
	fmt.Println("main start", time.Now())
	tt1 := newTree(1, 2)
	tt2 := newTree(1, 10000000)
	//fmt.Println("tree1=", tt1)
	//fmt.Println("tree2=", tt2)
	fmt.Println("compare start", time.Now())
	fmt.Println(Same(tt1, tt2))
	fmt.Println("compare end", time.Now())
}
// -*- mode: compilation; default-directory: "~/go/src/practice/14tree/" -*-
// Compilation started at Fri Oct 22 16:37:29
//  
// go run tree11.go
// main start 2021-10-22 16:37:30.27843 +0900 JST m=+0.000122961
// compare start 2021-10-22 16:37:44.48976 +0900 JST m=+14.211336302
// false
// compare end 2021-10-22 16:37:44.489878 +0900 JST m=+14.211453888
//  
// Compilation finished at Fri Oct 22 16:37:44

func main2() {
	fmt.Println("main start", time.Now())
	tt1 := newTree(1, 2)
	tt2 := newRandomTree(1, 10000000)
	//fmt.Println("tree1=", tt1)
	//fmt.Println("tree2=", tt2)
	fmt.Println("compare start", time.Now())
	fmt.Println(Same(tt1, tt2))
	fmt.Println("compare end", time.Now())
}
// -*- mode: compilation; default-directory: "~/go/src/practice/14tree/" -*-
// Compilation started at Fri Oct 22 16:30:11
//  
// go run tree11.go
// main start 2021-10-22 16:30:12.510376 +0900 JST m=+0.000126560
// compare start 2021-10-22 16:30:24.966376 +0900 JST m=+12.456025435
// false
// compare end 2021-10-22 16:30:24.966483 +0900 JST m=+12.456132017
//  
// Compilation finished at Fri Oct 22 16:30:25
