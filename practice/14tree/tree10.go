package main
import "golang.org/x/tour/tree"
import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)
/*
   subject : A Tour of Go Exercise: Equivalent Binary Trees.  check create and compare time
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
	
func main2() {
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
// Compilation started at Fri Oct 22 16:14:12
//  
// go run tree10.go
// main start 2021-10-22 16:14:13.128179 +0900 JST m=+0.000152807
// compare start 2021-10-22 16:14:27.417382 +0900 JST m=+14.288927326
// false
// compare end 2021-10-22 16:14:31.000738 +0900 JST m=+17.872174786
//  
// Compilation finished at Fri Oct 22 16:14:31

func main() {
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
// Compilation started at Fri Oct 22 16:16:26
//  
// go run tree10.go
// main start 2021-10-22 16:16:27.250038 +0900 JST m=+0.000126085
// compare start 2021-10-22 16:16:39.845038 +0900 JST m=+12.594748367
// false
// compare end 2021-10-22 16:16:43.428431 +0900 JST m=+16.178033602
//  
// Compilation finished at Fri Oct 22 16:16:43
