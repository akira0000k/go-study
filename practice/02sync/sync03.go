package main
import (
	"fmt"
	"time"
)
/*
   subject: Non-Blocking Channel Operations   channel select using default
*/
func main() {
	_ = time.Second
	//messages := make(chan string)
	messages := make(chan string, 1)
	signals := make(chan bool, 1)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	signals<-true

	for i:=0; i<3; i++ {
		select {
		case msg := <-messages:
			fmt.Println("received message", msg)
		case sig := <-signals:
			fmt.Println("received signal", sig)
		//case to := <-time.After(time.Second):
		// 	fmt.Println("timeout", to) //never
		default:
			fmt.Println("no activity") //ok
		}
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Fri Oct 29 18:25:33
//  
// go run sync03.go
// no message received
// sent message hi
// received message hi
// received signal true
// no activity
//  
// Compilation finished at Fri Oct 29 18:25:34
