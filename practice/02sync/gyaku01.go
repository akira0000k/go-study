package main

import "fmt"
import "github.com/tlorens/go-ibgetkey"
import "time"

func main() {
	kill := make(chan bool)
	finished := make(chan bool)
	go killableGoroutine(kill, finished)
	targetkey := "."
	t := int(targetkey[0])
loop:
	for {
		input := keyboard.ReadKey()
		select {
		case <-finished:
			break loop
		default:
			if input == t {
				kill <- true
				break loop
			}
		}
	}
}

func killableGoroutine(kill, finished chan bool) {
	fmt.Println("Started goroutine. Push \".\" to kill me.")
	for i := 0; i < 50; i++ {
		select {
		case <-kill:
			fmt.Println()
			fmt.Println("Killed")
			finished <- true
			return
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
	fmt.Println()
	fmt.Println("Finished..push any key to abort.")
	finished <- true
	return
}
// Akira@MBP 02sync % go build gyaku00.go
// Akira@MBP 02sync % ./gyaku00
// Started goroutine. Push "." to kill me.
// ..................................................
// Finished..push any key to abort.
// Akira@MBP 02sync % ./gyaku00
// Started goroutine. Push "." to kill me.
// ...............
// Killed
// Akira@MBP 02sync % ./gyaku00
// Started goroutine. Push "." to kill me.
// ..................................................
// Finished..push any key to abort.
//  
//  
//  
//  
//  
// ^C
// Akira@MBP 02sync %
