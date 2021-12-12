package main

import "fmt"
import "github.com/tlorens/go-ibgetkey"
import "time"
/*
   subject : 逆引きGolang goroutine の実行を終了させられない事も。
*/
func main() {
	fmt.Println("main start")
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
				<-finished
				break loop
			}
		}
	}
	fmt.Println("main end")
}

func killableGoroutine(kill, finished chan bool) {
	fmt.Println("Started goroutine. Push \".\" to kill me.")
	for i := 0; i < 50; i++ {
		select {
		case <-kill:
			fmt.Println()
			fmt.Println("sub Killed")
			finished <- true
			return
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
	fmt.Println()
	fmt.Println("Finished..push any key to abort.")
	finished <- true // 相手がいないと投げられないで固まる
	return
}
// Akira@MBP 02sync % ./gyaku01
// main start
// Started goroutine. Push "." to kill me.
// .....................
// sub Killed
// main end
// Akira@MBP 02sync % ./gyaku01
// main start
// Started goroutine. Push "." to kill me.
// ..................................................
// Finished..push any key to abort.
// main end
// Akira@MBP 02sync % ./gyaku01
// main start
// Started goroutine. Push "." to kill me.
// ..................................................
// Finished..push any key to abort.
// main end
// Akira@MBP 02sync % ./gyaku01
// main start
// Started goroutine. Push "." to kill me.
// ..................................................
// Finished..push any key to abort.
//  
//  
//  
//  02sync % ./gyaku01
// main start
// Started goroutine. Push "." to kill me.
// ..................................................
// Finished..push any key to abort.
//  
//  
// ^C
// Akira@MBP 02sync %
