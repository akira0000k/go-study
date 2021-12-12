package main
import (
	"fmt"
	"time"
	"sync"
)
/*
   subject: WaitGroups. param wg should be adress
*/

func worker(id int, wg sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	//_ = time.Second
	var wg sync.WaitGroup

	for i:=1; i<=5; i++ {
		wg.Add(1)
		go worker(i, wg)
	}

	fmt.Println("Waiting.....")
	wg.Wait()
	fmt.Println(".....all worker done")
}
//  -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/" -*-
//  Compilation started at Fri Sep 17 21:31:37
//   
//  go run sync09-2.go
//  Waiting.....
//  Worker 5 starting
//  Worker 2 starting
//  Worker 3 starting
//  Worker 1 starting
//  Worker 4 starting
//  Worker 4 done
//  Worker 1 done
//  Worker 3 done
//  Worker 2 done
//  panic: sync: negative WaitGroup counter
//   
//  goroutine 8 [running]:
//  sync.(*WaitGroup).Add(0xc000100000, 0xffffffffffffffff)
//   	/usr/local/Cellar/go/1.16.6/libexec/src/sync/waitgroup.go:74 +0x147
//  sync.(*WaitGroup).Done(0xc000100000)
//   	/usr/local/Cellar/go/1.16.6/libexec/src/sync/waitgroup.go:99 +0x34
//  main.worker(0x3, 0x0, 0xc000000003)
//   	/Users/Akira/Desktop/work/go/practice/sync09-2.go:18 +0x198
//  created by main.main
//   	/Users/Akira/Desktop/work/go/practice/sync09-2.go:26 +0xaf
//  exit status 2
//   
//  Compilation exited abnormally with code 1 at Fri Sep 17 21:31:39

	
