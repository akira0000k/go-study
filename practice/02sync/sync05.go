package main
import (
	"fmt"
	"time"
)
/*
   subject: Range over Channels       サンプルで学ぶ Go 言語
*/

func main() {
	_ = time.Second

	fmt.Println("start")
	queue := make(chan string, 2)
	queue<-"one"
	queue<-"two"
	close(queue)

	i := 0
	for item :=range queue{
		fmt.Println(i, item)
		//if i++; i>=2 {
		// 	break
		//}
	}
	//close(queue)

	fmt.Println("END")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Fri Oct 29 18:33:13
//  
// go run sync05.go
// start
// 0 one
// 0 two
// END
//  
// Compilation finished at Fri Oct 29 18:33:14
