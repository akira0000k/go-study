package main
import (
	"context"
	"fmt"
	"time"
	"sync"
)
/*
   subject: Go の Context を学ぶ
   ポイントは、コンテキストオブジェクトは、デコレータの様な使い方をするところ。
   この例では、context.Background()で渡ってきたのは、空のContextが渡ってきているが、
   ここで、context.WithCancel(ctx) の様にラップしてキャンセル機能を追加している。

   子のキャンセルは親に伝搬しない。
*/
var wg sync.WaitGroup

func infiniteLoop(ctx context.Context) {
	defer wg.Done()
	//innerCtx, cancel := context.WithCancel(ctx)//内部コンテクスト作成キャンセル付き
	ctx, cancel := context.WithCancel(ctx)//内部コンテクスト作成キャンセル付き
	_ = cancel
	defer cancel()//returnでキャンセル。メモリーリーク防止? (根本に向かっては伝搬しないので)

	for i:=0; ;i++ {
		fmt.Println("Waiting for time out", i)
		//time.Sleep(time.Second / 5)//ここで待ったらDoneを受け取れないうちにmainが終わる可能性あり。

		if i==2 {
			return
		}
		if i==1000 {
			cancel()//処理途中で止めることもできる main end: context deadline exceeded (何故に?) 
		}
		fmt.Print(".")//cancelしてもここへは来る
		select {
		case <-ctx.Done():
			fmt.Println("Exit now!")
			//time.Sleep(time.Second / 5)//先にmainが終わっちゃう
			fmt.Println("message:", ctx.Value("message").(string))
			return
		case <-time.After(time.Second / 5):
		//default://ょりtime.Afterを使った方がいい
		}
	}
}

func main() {
	ctx := context.Background()//空のcontextを作成
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)//タイムアウト,キャンセル機能追加
	ctx = context.WithValue(ctx, "message", "hi")//メッセージ追加
	_ = cancel
	//defer cancel()//何故必要
	ech := make(chan bool, 1)//どんな時でもblockされないようにbufferが一つ必要
	
	wg.Add(1)
	go infiniteLoop(ctx)

	go func() {
		wg.Wait()
		ech<-true
	}()
	
	//cancel() //main end: context canceled
	//fmt.Print(",")//ここへは来る
	select {
	case <-ctx.Done():
		fmt.Println("main end:", ctx.Err())
	case <-ech:
		fmt.Println("main end: goroutine returned")
		cancel()
	}
	//time.Sleep(time.Second / 2)//なくてもいけそう
}
// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/06context/" -*-
// Compilation started at Wed Sep 29 20:14:58
//  
// go run cont03.go
// Waiting for time out 0
// .Waiting for time out 1
// .Waiting for time out 2
// .Exit now!
// message: hi
// main end: context deadline exceeded
//  
// Compilation finished at Wed Sep 29 20:14:59

// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/06context/" -*-
// Compilation started at Wed Sep 29 22:38:03
//  
// go run cont03.go
// Waiting for time out 0
// .Waiting for time out 1
// .Waiting for time out 2
// main end: goroutine returned
//  
// Compilation finished at Wed Sep 29 22:38:04

// -*- mode: compilation; default-directory: "~/Desktop/work/go/practice/06context/" -*-
// Compilation started at Wed Sep 29 22:39:02
//  
// go run cont03.go
// Waiting for time out 0
// .Waiting for time out 1
// .Waiting for time out 2
// .Waiting for time out 3
// .Waiting for time out 4
// .Waiting for time out 5
// .Waiting for time out 6
// .Waiting for time out 7
// .Waiting for time out 8
// .Waiting for time out 9
// .Waiting for time out 10
// .Waiting for time out 11
// .Waiting for time out 12
// .Waiting for time out 13
// .Waiting for time out 14
// .Waiting for time out 15
// .Waiting for time out 16
// .Waiting for time out 17
// .Waiting for time out 18
// .Waiting for time out 19
// .Waiting for time out 20
// .Waiting for time out 21
// .Waiting for time out 22
// .Waiting for time out 23
// .Waiting for time out 24
// .Exit now!
// message: hi
// main end: context deadline exceeded
//  
// Compilation finished at Wed Sep 29 22:39:08
