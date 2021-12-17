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
   subject : A Tour of Go Exercise: Equivalent Binary Trees.  cancel in main
*/
//type Tree struct {
// 	Left  *Tree
// 	Value int
// 	Right *Tree
//}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(ctx context.Context, t *tree.Tree, ch chan int) bool {
	select {
	case <-ctx.Done():
		return false
	default:
	}
	if t.Left != nil {
		if !Walk(ctx, t.Left, ch) {
			return false
		}
	}
	ch <- t.Value
	if t.Right != nil {
		if !Walk(ctx, t.Right, ch) {
			return false
		}
	}
	return true
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(ctx context.Context, t1, t2 *tree.Tree) (bool, bool) {
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
	var complete1, complete2 bool
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(ch1)
		complete1 = Walk(ctx2, t1, ch1)
	}()
	go func() {
		defer wg.Done()
		defer close(ch2)
		complete2 = Walk(ctx2, t2, ch2)
	}()

	wg.Wait()
	
	return result, complete1 && complete2
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

func newRandomTree(ctx context.Context, from, to int) *tree.Tree {
	n := to - from
	tree := addTree(nil, rand.Intn(n) + from)
	for i:=0; i<n; i++ {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		addTree(tree, rand.Intn(n) + from)
	}
	return tree
}
func newTree(ctx context.Context, from, to int) *tree.Tree {
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
		select {
		case <-ctx.Done():
			return nil
		default:
		}
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println("main start", time.Now())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("creating tree1...", time.Now())
		tt1 := newTree(ctx, 1, 1000000)
		if tt1 == nil {
			return
		}
		fmt.Println("creating tree2...", time.Now())
		tt2 := newTree(ctx, 1, 1000000)
		if tt2 == nil {
			return
		}
		//fmt.Println("tree1=", tt1)
		//fmt.Println("tree2=", tt2)
		fmt.Println("compare start", time.Now())
		result, ok := Same(ctx, tt2, tt1)
		fmt.Println(result)
		if ok {
			fmt.Println("compare end", time.Now())
		} else {
			fmt.Println("compare cancelled", time.Now())
		}
	}()
	go func() {
		//time.Sleep(time.Millisecond *  300)
		//time.Sleep(time.Millisecond *  1000)
		time.Sleep(time.Millisecond * 1500)
		//time.Sleep(time.Millisecond * 2500)
		fmt.Println("timeout...cancel", time.Now())
		cancel()
	}()
	wg.Wait()
	fmt.Println("main end", time.Now())
}
// -*- mode: compilation; default-directory: "~/go/src/practice/14tree/" -*-
// Compilation started at Fri Oct 22 19:22:32
//  
// go run tree13.go
// main start 2021-10-22 19:22:33.2288 +0900 JST m=+0.000155860
// creating tree1... 2021-10-22 19:22:33.228975 +0900 JST m=+0.000330912
// creating tree2... 2021-10-22 19:22:33.953997 +0900 JST m=+0.725354442
// compare start 2021-10-22 19:22:34.698702 +0900 JST m=+1.470062478
// timeout...cancel 2021-10-22 19:22:34.728974 +0900 JST m=+1.500333926
// false
// compare end 2021-10-22 19:22:34.729027 +0900 JST m=+1.500387231
// main end 2021-10-22 19:22:34.729033 +0900 JST m=+1.500392870
//  
// Compilation finished at Fri Oct 22 19:22:34
