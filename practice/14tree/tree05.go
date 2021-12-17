package main
import "golang.org/x/tour/tree"
import (
	"fmt"
	//"time"
	"sync"
)
/*
   subject : A Tour of Go Exercise: Equivalent Binary Trees.  complete addTree function(**).
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

func addTree(pt **tree.Tree, val int) {
	t := *pt
	if t == nil {
		tt := new(tree.Tree)
		tt.Value = val
		*pt = tt
	} else {
		if val < t.Value {
			addTree(&t.Left, val)
		} else {
			addTree(&t.Right, val)
		}
	}
}
	
func main() {
	tt1 := tree.New(1)
	tt2 := tree.New(2)
	var result bool
	
	fmt.Println("tree1=", tt1)
	fmt.Println("tree2=", tt2)
	result = Same(tt1, tt2)
	fmt.Println(result)

	tt1 = tree.New(1)
	tt2 = tree.New(1)

	fmt.Println("tree1=", tt1)
	fmt.Println("tree2=", tt2)
	result = Same(tt1, tt2)
	fmt.Println(result)

	addTree(&tt1, 13)
	addTree(&tt1, 12)
	addTree(&tt1, 11)

	fmt.Println("tree1=", tt1)
	fmt.Println("tree2=", tt2)
	result = Same(tt1, tt2)
	fmt.Println(result)

	addTree(&tt2, 12)
	addTree(&tt2, 11)
	addTree(&tt2, 13)
	addTree(&tt2, 16)
	addTree(&tt2, 14)
	addTree(&tt2, 15)

	fmt.Println("tree1=", tt1)
	fmt.Println("tree2=", tt2)
	result = Same(tt1, tt2)
	fmt.Println(result)

	var ttn *tree.Tree
	addTree(&ttn, 9)
	addTree(&ttn, 6)
	addTree(&ttn, 6)
	addTree(&ttn, 6)
	addTree(&ttn, 13)
	addTree(&ttn, 2)
	addTree(&ttn, 8)
	
	fmt.Println("tree1=", tt1)
	fmt.Println("treen=", ttn)
	result = Same(tt1, ttn)
	fmt.Println(result)
}
// -*- mode: compilation; default-directory: "~/go/src/practice/14tree/" -*-
// Compilation started at Thu Oct 21 19:44:18
//  
// go run tree05.go
// tree1= ((((1 (2)) 3 (4)) 5 ((6) 7 ((8) 9))) 10)
// tree2= ((((2) 4 (6)) 8 (10 (12))) 14 ((16) 18 (20)))
// false
// tree1= ((((((1) 2) 3 (4)) 5 (6)) 7) 8 ((9) 10))
// tree2= ((1 ((((2) 3 (4)) 5) 6)) 7 ((8) 9 (10)))
// true
// tree1= ((((((1) 2) 3 (4)) 5 (6)) 7) 8 ((9) 10 (((11) 12) 13)))
// tree2= ((1 ((((2) 3 (4)) 5) 6)) 7 ((8) 9 (10)))
// false
// tree1= ((((((1) 2) 3 (4)) 5 (6)) 7) 8 ((9) 10 (((11) 12) 13)))
// tree2= ((1 ((((2) 3 (4)) 5) 6)) 7 ((8) 9 (10 ((11) 12 (13 ((14 (15)) 16))))))
// false
// tree1= ((((((1) 2) 3 (4)) 5 (6)) 7) 8 ((9) 10 (((11) 12) 13)))
// treen= (((2) 6 (6 (6 (8)))) 9 (13))
// false
//  
// Compilation finished at Thu Oct 21 19:44:19
