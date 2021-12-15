package main
import(
	"fmt"
	"context"
	"time"
	"sync"
)
/*
   subject: GolangのContextを頑張って理解していく
   それが僕には楽しかったんです。

*/
var wg sync.WaitGroup

func main() {
	// 空のContextを生成
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	_ = cancel
	wg.Add(1)
	go parentRoutine(ctx)

	//time.Sleep(3 * time.Second)
	//cancel()
	
	wg.Wait()//time.Sleep(20 * time.Second)

	fmt.Println("main func finish")
}

func parentRoutine(ctx context.Context) {
	defer wg.Done()
	
	subCtx, cancel := context.WithCancel(ctx)
	subsubCtx, cancel2 := context.WithCancel(subCtx)
	//subsubCtx, _ := context.WithCancel(subCtx)

	wg.Add(1)
	go childRoutine(subCtx, "sub context")
	wg.Add(1)
	go childRoutine(subsubCtx, "sub sub context")

	time.Sleep(5 * time.Second)

	fmt.Println("cancel2()..."); cancel2(); fmt.Println("...done cancel2()")

	// time.Sleep(5 * time.Second)
	//  
	fmt.Println("cancel()......"); cancel(); fmt.Println("......done cancel()")
	//  
	// tCtx, _ := context.WithTimeout(ctx, time.Second * 5)
	// go childRoutine(tCtx, "context with timeout")

	fmt.Println("parent routine finish")
}

func childRoutine(ctx context.Context, prefix string) {
	defer wg.Done()
	for i:=0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("routine %s cancelled......\n", prefix)
		        time.Sleep(time.Second * 3)
			fmt.Printf("....... %s return.\n", prefix)
			return

		case <-time.After(1 * time.Second):
			fmt.Printf("routine %s has value %d \n", prefix, i)
		}
	}
}
// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/06context/" -*-
// Compilation started at Wed Sep 29 20:46:15
//  
// go run cont01.go
// routine sub context has value 0 
// routine sub sub context has value 0 
// routine sub sub context has value 1 
// routine sub context has value 1 
// routine sub context has value 2 
// routine sub sub context has value 2 
// routine sub context has value 3 
// routine sub sub context has value 3 
// cancel2()...
// ...done cancel2()
// cancel()......
// ......done cancel()
// parent routine finish
// routine sub sub context cancelled......
// routine sub context cancelled......
// ....... sub sub context return.
// ....... sub context return.
// main func finish
//  
// Compilation finished at Wed Sep 29 20:46:24
