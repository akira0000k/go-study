package main

import "fmt"
import "math"
/*
   subject : 逆引きgolang goroutine間の通信
*/
func main() {
	queue := make(chan int)//, 3) // 3はキューの深さ
	go sqrtGoroutine(queue)

	line := 0
loop:
	for {
		fmt.Scanln(&line)
		if line == -1 {
			break loop
		} else {
			queue <- line
		}
	}
	close(queue)
	fmt.Println("end main")
}

func sqrtGoroutine(queue chan int) {
	for n := range queue {
		if int(n) >= 0 {
			val := math.Sqrt(float64(n))
			fmt.Printf("Square(%d) = %f\n", int(n), val)
		} else {
			fmt.Println("?")
		}
	}
	fmt.Println("exit sub")
}
// Akira@MBP 02sync % go build gyaku04.go
// Akira@MBP 02sync % ./gyaku04
// 3
// Square(3) = 1.732051
// 4
// Square(4) = 2.000000
// 5
// Square(5) = 2.236068
// 6
// Square(6) = 2.449490
// -1
// exit sub
// end main
// Akira@MBP 02sync % ./gyaku04
// 3
// Square(3) = 1.732051
// -1
// end main
// Akira@MBP 02sync % ./gyaku04
// -1
// end main
// Akira@MBP 02sync %
