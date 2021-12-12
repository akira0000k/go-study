package main

import "fmt"
import "github.com/tlorens/go-ibgetkey"
import "time"
/*
   subject : 逆引きgolang goroutineを停止する?
*/
func main() {
	com := make(chan string)
	finished := make(chan bool)
	go stoppableGoroutine(com, finished)
	targetkey := "."
	t := int(targetkey[0])
	running := true
loop:
	for {
		input := keyboard.ReadKey()
		select {
		case <-finished:
			break loop
		default:
			if input == t {
				if running == true {
					fmt.Print("P")
					com <- "stop"
					fmt.Print("p")
					running = false
				} else {
					fmt.Print("T")
					com <- "start"
					fmt.Print("t")
					running = true
				}
			}
		}
	}
}

func stoppableGoroutine(command chan string, finished chan bool) {
	fmt.Println("Started goroutine. Push \".\" to stop/start me.")
	running := true
	i := 0
	for i < 50 {
		select {
		case com := <-command:
			if com == "stop" {
				running = false
			} else {
				running = true
			}
		default:
		}
		if running == true {
			fmt.Print(".")
			//time.Sleep(100 * time.Millisecond)
			i++
		}
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
	fmt.Println("Finished..push any key to abort.")
	finished <- true
	return
}
// Akira@MBP 02sync % ./gyaku02
// Started goroutine. Push "." to stop/start me.
// ..................................................
// Finished..push any key to abort.
// Akira@MBP 02sync % ./gyaku02
// Started goroutine. Push "." to stop/start me.
// .................................................PpT.t
// Finished..push any key to abort.
// Akira@MBP 02sync % ./gyaku02
// Started goroutine. Push "." to stop/start me.
// ..................................................
// Finished..push any key to abort.
// Akira@MBP 02sync % ./gyaku02
// Started goroutine. Push "." to stop/start me.
// ..................................................P
// Finished..push any key to abort.
//  
//  
//  
//  
//  
// ^C
// Akira@MBP 02sync %
