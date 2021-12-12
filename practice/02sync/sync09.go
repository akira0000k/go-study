package main
import (
	"fmt"
	"time"
	"sync"
)
/*
   subject: WaitGroups
*/

func worker(id int, wg *sync.WaitGroup) {
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
		go worker(i, &wg)
	}
	fmt.Println("Waiting.....")
	wg.Wait()
	fmt.Println(".....all worker done")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Fri Oct 29 18:56:50
//  
// go run sync09.go 
// Waiting.....
// Worker 5 starting
// Worker 3 starting
// Worker 2 starting
// Worker 4 starting
// Worker 1 starting
// Worker 2 done
// Worker 3 done
// Worker 1 done
// Worker 5 done
// Worker 4 done
// .....all worker done
//  
// Compilation finished at Fri Oct 29 18:56:51
