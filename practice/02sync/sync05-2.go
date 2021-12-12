package main
import (
	"fmt"
	"time"
)
/*
   subject: use more flag after closing channel
*/

func main() {
	_ = time.Second

	fmt.Println("start")
	queue := make(chan string, 2)
	queue<-"one"
	queue<-"two"
	close(queue)

	for i:=0; i<10; i++{
		item, more := <-queue
		if !more {
			break
		}
		fmt.Println(i, item)
	}
	//close(queue)

	fmt.Println("END")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Fri Oct 29 18:33:35
//  
// go run sync05-2.go
// start
// 0 one
// 1 two
// END
//  
// Compilation finished at Fri Oct 29 18:33:36
