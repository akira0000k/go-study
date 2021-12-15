package main
import(
	"fmt"
	"time"
	"context"
	"sync"
	"math/rand"
)
/*
  subject: golangでcontextパッケージを使う   write ahead log
  wgは追加した
*/
func infiniteLoop(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for {
		fmt.Println("Help")

		select {
		case val, ok := <-ctx.Done():
			fmt.Println("val:", val, "ok:", ok)
			fmt.Println("Done received.", ctx.Err())
			time.Sleep(time.Second / 2)
			fmt.Println("func exit")
			return
		default:
			time.Sleep(time.Second / 10)
		}
	}
}

func main() {
	fmt.Println("start main")

	ctx := context.Background()
	ctx, cancel1 := context.WithTimeout(ctx, time.Second)
	_ = cancel1
	
	var wg sync.WaitGroup

	wg.Add(1)
	go infiniteLoop(ctx, &wg)
	
	rand.Seed(time.Now().UnixNano() / 1000000)
	irand := rand.Intn(2)
	fmt.Println("rand:", irand)
	if irand == 0 {
		time.Sleep(time.Second/2)
		cancel1() //Done received. context canceled
	} else {
		//Done received. context deadline exceeded
	}
	wg.Wait()
	fmt.Println("end main")
}
// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/06context/" -*-
// Compilation started at Wed Sep 29 16:45:15
//  
// go run cont02.go
// start main
// rand: 1
// Help
// Help
// Help
// Help
// Help
// Help
// Help
// Help
// Help
// Help
// Help
// val: {} ok: false
// Done received. context deadline exceeded
// func exit
// end main
//  
// Compilation finished at Wed Sep 29 16:45:17
//  
// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/06context/" -*-
// Compilation started at Wed Sep 29 16:46:22
//  
// go run cont02.go
// start main
// rand: 0
// Help
// Help
// Help
// Help
// Help
// Help
// val: {} ok: false
// Done received. context canceled
// func exit
// end main
//  
// Compilation finished at Wed Sep 29 16:46:24
