package main

import "fmt"
import "github.com/tlorens/go-ibgetkey"
import "time"
import "sync"
/*
   subject : 逆引きGolang goroutine の実行を確実に終了できる?
*/
func main() {
	fmt.Println("killtest start")
	kill := make(chan bool)
	finished := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(1)
	go killableGoroutine(kill, finished, &wg)

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
	wg.Wait()
	fmt.Println("main end")
}

func killableGoroutine(kill, finished chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	
	fmt.Println("Started goroutine. Push \".\" to kill me.")
	for i := 0; i < 50; i++ {
		select {
		case <-kill:
			fmt.Println()
			fmt.Println("sub Killed")
			return
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
	fmt.Println()
	fmt.Println("Finished..push any key to abort.")
	//finished <- true
	close(finished)
	return
}
// Akira@MBP 02sync % go build gyaku01-2.go

// Akira@MBP 02sync % ./gyaku01-2
// killtest start
// Started goroutine. Push "." to kill me.
// ...............
// sub Killed
// main end
// Akira@MBP 02sync % ./gyaku01-2
// killtest start
// Started goroutine. Push "." to kill me.
// ..................................................
// Finished..push any key to abort.
// main end
// Akira@MBP 02sync % ./gyaku01-2
// killtest start
// Started goroutine. Push "." to kill me.
// ..................................................
// Finished..push any key to abort.
//  
//  
//  
//  
//  
// ^[[D^C
// Akira@MBP 02sync %
