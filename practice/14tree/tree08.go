package main
import "golang.org/x/tour/tree"
import (
	"fmt"
	"sync"
	"math/rand"
)
/*
   subject : A Tour of Go Exercise: Equivalent Binary Trees.  making sorted tree
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
		islice = append(islice[0:idx], islice[idx+1:]...)
		leng = len(islice)
		//fmt.Println(tree)
	}
	return tree
}
	
func main() {
	fm := 10
	to := 20

	tt1 := newTree(fm, to)
	tt2 := newTree(fm, to)
	fmt.Println("tree1=", tt1)
	fmt.Println("tree2=", tt2)
	fmt.Println(Same(tt1, tt2))

	tt1 = newTree(fm, to)
	tt2 = newTree(fm, to+1)
	fmt.Println("tree1=", tt1)
	fmt.Println("tree2=", tt2)
	fmt.Println(Same(tt1, tt2))

	tt1 = newTree(fm, to+1)
	tt2 = newTree(fm, to)
	fmt.Println("tree1=", tt1)
	fmt.Println("tree2=", tt2)
	fmt.Println(Same(tt1, tt2))

	tt1 = newTree(fm, to)
	tt2 = newTree(fm+1, to)
	fmt.Println("tree1=", tt1)
	fmt.Println("tree2=", tt2)
	fmt.Println(Same(tt1, tt2))
}
// -*- mode: compilation; default-directory: "~/go/src/practice/14tree/" -*-
// Compilation started at Thu Oct 21 23:38:25
//  
// go run tree08.go
// tree1= ((10) 11 (((12) 13 ((14) 15 ((16) 17))) 18 ((19) 20)))
// tree2= (10 ((11 (12)) 13 (14 ((((15) 16) 17) 18 ((19) 20)))))
// true
// tree1= (((10 (((11) 12 (13)) 14 ((15 (16)) 17))) 18 (19)) 20)
// tree2= (((((10 ((((11) 12) 13) 14)) 15) 16 (17)) 18 ((19) 20)) 21)
// false
// tree1= (((10) 11 (12)) 13 (((14) 15) 16 ((((17 (18)) 19) 20) 21)))
// tree2= (((10 ((11) 12)) 13 ((14 (15)) 16 (17 (18 (19))))) 20)
// false
// tree1= (10 (((11) 12 ((13) 14 ((15) 16))) 17 (18 (19 (20)))))
// tree2= (11 (((12) 13 ((14) 15 (16 (17)))) 18 ((19) 20)))
// false
//  
// Compilation finished at Thu Oct 21 23:38:26
