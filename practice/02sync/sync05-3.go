package main
import (
	"fmt"
	"time"
)
/*
   subject: use select default and more flag receiving channel
*/

func main() {
	_ = time.Second

	fmt.Println("start")
	queue := make(chan string, 2)
	queue<-"one"
	queue<-"two"
	//close(queue)

forloop:
	for i:=0; i<10; i++{
		select {
		case item, more := <-queue:
			if !more {
				fmt.Println("no more item") //closed
				break forloop
			}
			fmt.Println(i, item)
		default:
			fmt.Println("default") //empty
			break forloop
		}
	}
	close(queue)

	fmt.Println("END")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Fri Oct 29 18:34:01
//  
// go run sync05-3.go
// start
// 0 one
// 1 two
// default
// END
//  
// Compilation finished at Fri Oct 29 18:34:02
