package main
import "golang.org/x/tour/tree"
import (
	"fmt"
	"time"
	"sync"
	"math/rand"
	"context"
)
/*
   subject : A Tour of Go Exercise: Equivalent Binary Trees.  use context
*/
//type Tree struct {
// 	Left  *Tree
// 	Value int
// 	Right *Tree
//}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(ctx context.Context, t *tree.Tree, ch chan int){
	select {
	case <-ctx.Done():
		return
	default:
	}
	if t.Left != nil {
		Walk(ctx, t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(ctx, t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(ctx context.Context, t1, t2 *tree.Tree) bool {
	ctx2, cancel := context.WithCancel(ctx)
	defer cancel()
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
			cancel()
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
		Walk(ctx2, t1, ch1)
	}()
	go func() {
		defer wg.Done()
		defer close(ch2)
		Walk(ctx2, t2, ch2)
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
	
func main1() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println("main start", time.Now())
	tt1 := newTree(1, 2)
	tt2 := newTree(1, 10000000)
	//fmt.Println("tree1=", tt1)
	//fmt.Println("tree2=", tt2)
	fmt.Println("compare start", time.Now())
	fmt.Println(Same(ctx, tt1, tt2))
	fmt.Println("compare end", time.Now())
}
// -*- mode: compilation; default-directory: "~/go/src/practice/14tree/" -*-
// Compilation started at Fri Oct 22 17:17:05
//  
// go run tree12.go
// main start 2021-10-22 17:17:05.607695 +0900 JST m=+0.000132593
// compare start 2021-10-22 17:17:19.905251 +0900 JST m=+14.298038803
// false
// compare end 2021-10-22 17:17:19.905414 +0900 JST m=+14.298201586
//  
// Compilation finished at Fri Oct 22 17:17:19

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println("main start", time.Now())
	tt1 := newTree(1, 2)
	tt2 := newRandomTree(1, 100)
	fmt.Println("tree1=", tt1)
	fmt.Println("tree2=", tt2)
	fmt.Println("compare start", time.Now())
	fmt.Println(Same(ctx, tt2, tt1))
	fmt.Println("compare end", time.Now())
}
// -*- mode: compilation; default-directory: "~/go/src/practice/14tree/" -*-
// Compilation started at Fri Oct 22 17:23:38
//  
// go run tree12.go
// main start 2021-10-22 17:23:38.976561 +0900 JST m=+0.000172836
// tree1= ((1) 2)
// tree2= ((((((1 (1 (2 (4)))) 6 (6)) 7) 8 ((((8 (8)) 10) 11 (11)) 12 (13))) 14 (16 (((16) 17 (17)) 18 (18 (19))))) 21 (((21) 22 ((((((22) 23 ((23) 24 (24))) 25 ((26) 27)) 29 ((29) 31 (32))) 33 (36 ((37 (37)) 38))) 39 (((40) 41 (42 (43))) 45 ((45) 47)))) 48 (((((48 (48)) 49) 52 (((53) 54 (55)) 56 (58 (59)))) 60 (((((60) 61 (61)) 62 (62 ((62) 64 ((65) 66)))) 68 ((72 (72)) 75 (75 (76 (((77) 78) 79 ((79 (81)) 82 ((82 (82 (82))) 83 (83 (83))))))))) 89 (89 (89 (((89) 90 (91 (93))) 94))))) 95 (95))))
// compare start 2021-10-22 17:23:38.97699 +0900 JST m=+0.000602510
// false
// compare end 2021-10-22 17:23:38.977081 +0900 JST m=+0.000693366
//  
// Compilation finished at Fri Oct 22 17:23:38
