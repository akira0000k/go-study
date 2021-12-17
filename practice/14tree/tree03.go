package main
import "golang.org/x/tour/tree"
import (
	"fmt"
	//"time"
	"sync"
)
/*
   subject : A Tour of Go Exercise: Equivalent Binary Trees.  length differ cause deadlock
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
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
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

func addTree(t *tree.Tree, val int) {
	if t.Right != nil {
		addTree(t.Right, val)
	} else {
		tt := new(tree.Tree)
		tt.Value = val
		t.Right = tt
		return
	}
}
	
func main() {
	tt1 := tree.New(1)
	tt2 := tree.New(1)
	addTree(tt1, 11)
	addTree(tt1, 12)
	addTree(tt1, 13)

	fmt.Println("tree1=", tt1)
	fmt.Println("tree2=", tt2)
	result := Same(tt1, tt2)
	fmt.Println(result)
}
// -*- mode: compilation; default-directory: "~/go/src/practice/14tree/" -*-
// Compilation started at Thu Oct 21 16:02:16
//  
// go run tree03.go
// tree1= ((((1 (2)) 3 (4)) 5 ((6) 7 ((8) 9))) 10 (11 (12 (13))))
// tree2= ((((1) 2 (3)) 4 (5 (6))) 7 ((8) 9 (10)))
// fatal error: all goroutines are asleep - deadlock!
//  
// goroutine 1 [semacquire]:
// sync.runtime_Semacquire(0xc000014228)
//  	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/sema.go:56 +0x45
// sync.(*WaitGroup).Wait(0xc000014220)
//  	/usr/local/Cellar/go/1.16.6/libexec/src/sync/waitgroup.go:130 +0x65
// main.Same(0xc00000c030, 0xc00000c120, 0xc00005cf38)
//  	/Users/Akira/go/src/practice/14tree/tree03.go:67 +0x193
// main.main()
//  	/Users/Akira/go/src/practice/14tree/tree03.go:95 +0x1bd
//  
// goroutine 6 [chan receive]:
// main.Same.func1(0xc000014210, 0xc0000141fa, 0xc000048070, 0xc0000480e0)
//  	/Users/Akira/go/src/practice/14tree/tree03.go:44 +0x94
// created by main.Same
//  	/Users/Akira/go/src/practice/14tree/tree03.go:37 +0xed
//  
// goroutine 7 [chan send]:
// main.Walk(0xc00000c240, 0xc000048070)
//  	/Users/Akira/go/src/practice/14tree/tree03.go:23 +0x45
// main.Walk(0xc00000c228, 0xc000048070)
//  	/Users/Akira/go/src/practice/14tree/tree03.go:25 +0x70
// main.Walk(0xc00000c210, 0xc000048070)
//  	/Users/Akira/go/src/practice/14tree/tree03.go:25 +0x70
// main.Walk(0xc00000c030, 0xc000048070)
//  	/Users/Akira/go/src/practice/14tree/tree03.go:25 +0x70
// main.Same.func2(0xc000014220, 0xc00000c030, 0xc000048070)
//  	/Users/Akira/go/src/practice/14tree/tree03.go:60 +0x5d
// created by main.Same
//  	/Users/Akira/go/src/practice/14tree/tree03.go:58 +0x14f
// exit status 2
//  
// Compilation exited abnormally with code 1 at Thu Oct 21 16:02:16
