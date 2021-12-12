package main

import "fmt"
import "os"
import "runtime/pprof"
import "time"
/*
   subject : 逆引き golang 実行中goroutineのリスト runtime/pprof.Lookup
*/
func main() {
	go goroutine1()
	go goroutine2()

	for i := 0; i<3; i++ {
		fmt.Printf("\n%d  %v\n", 1, time.Now())
		pprof.Lookup("goroutine").WriteTo(os.Stdout, 1) // 2はデバッグレベル。goroutineだけリストする、の意味。
		fmt.Printf("\n%d  %v\n", 2, time.Now())
		pprof.Lookup("goroutine").WriteTo(os.Stdout, 2) // 2はデバッグレベル。goroutineだけリストする、の意味。
		time.Sleep(1 * time.Second)                     // goroutineの起動のオーバヘッド待ち
	}
}

func goroutine1() {
	time.Sleep(1 * time.Second)
	fmt.Println("Goroutine1 finished")
}

func goroutine2() {
	time.Sleep(2 * time.Second)
	fmt.Println("Goroutine2 finished")
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Sun Oct 24 23:23:48
//  
// go run gyaku03.go
//  
// 1  2021-10-24 23:23:49.233193 +0900 JST m=+0.000175282
// goroutine profile: total 3
// 1 @ 0x1037405 0x1064432 0x10cd2ae 0x1067101
// #	0x1064431	time.Sleep+0xd1		/usr/local/Cellar/go/1.16.6/libexec/src/runtime/time.go:193
// #	0x10cd2ad	main.goroutine1+0x2d	/Users/Akira/go/src/practice/02sync/gyaku03.go:24
//  
// 1 @ 0x1037405 0x1064432 0x10cd34e 0x1067101
// #	0x1064431	time.Sleep+0xd1		/usr/local/Cellar/go/1.16.6/libexec/src/runtime/time.go:193
// #	0x10cd34d	main.goroutine2+0x2d	/Users/Akira/go/src/practice/02sync/gyaku03.go:29
//  
// 1 @ 0x106223d 0x10c180e 0x10c15e5 0x10be172 0x10cd10d 0x1036fd6 0x1067101
// #	0x106223c	runtime/pprof.runtime_goroutineProfileWithLabels+0x5c	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/mprof.go:716
// #	0x10c180d	runtime/pprof.writeRuntimeProfile+0xcd			/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:724
// #	0x10c15e4	runtime/pprof.writeGoroutine+0xa4			/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:684
// #	0x10be171	runtime/pprof.(*Profile).WriteTo+0x3f1			/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:331
// #	0x10cd10c	main.main+0x18c						/Users/Akira/go/src/practice/02sync/gyaku03.go:16
// #	0x1036fd5	runtime.main+0x255					/usr/local/Cellar/go/1.16.6/libexec/src/runtime/proc.go:225
//  
//  
// 2  2021-10-24 23:23:49.234031 +0900 JST m=+0.001013248
// goroutine 1 [running]:
// runtime/pprof.writeGoroutineStacks(0x111fb88, 0xc000134008, 0x0, 0x10aef31)
//  	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:693 +0x9f
// runtime/pprof.writeGoroutine(0x111fb88, 0xc000134008, 0x2, 0x1103338, 0xc000157e50)
//  	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:682 +0x45
// runtime/pprof.(*Profile).WriteTo(0x11a7280, 0x111fb88, 0xc000134008, 0x2, 0xc000157f38, 0x2)
//  	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:331 +0x3f2
// main.main()
//  	/Users/Akira/go/src/practice/02sync/gyaku03.go:18 +0x2a5
//  
// goroutine 18 [sleep]:
// time.Sleep(0x3b9aca00)
//  	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/time.go:193 +0xd2
// main.goroutine1()
//  	/Users/Akira/go/src/practice/02sync/gyaku03.go:24 +0x2e
// created by main.main
//  	/Users/Akira/go/src/practice/02sync/gyaku03.go:11 +0x47
//  
// goroutine 19 [sleep]:
// time.Sleep(0x77359400)
//  	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/time.go:193 +0xd2
// main.goroutine2()
//  	/Users/Akira/go/src/practice/02sync/gyaku03.go:29 +0x2e
// created by main.main
//  	/Users/Akira/go/src/practice/02sync/gyaku03.go:12 +0x5f
// Goroutine1 finished
//  
// 1  2021-10-24 23:23:50.243984 +0900 JST m=+1.010958782
// goroutine profile: total 2
// 1 @ 0x1037405 0x1064432 0x10cd34e 0x1067101
// #	0x1064431	time.Sleep+0xd1		/usr/local/Cellar/go/1.16.6/libexec/src/runtime/time.go:193
// #	0x10cd34d	main.goroutine2+0x2d	/Users/Akira/go/src/practice/02sync/gyaku03.go:29
//  
// 1 @ 0x106223d 0x10c180e 0x10c15e5 0x10be172 0x10cd10d 0x1036fd6 0x1067101
// #	0x106223c	runtime/pprof.runtime_goroutineProfileWithLabels+0x5c	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/mprof.go:716
// #	0x10c180d	runtime/pprof.writeRuntimeProfile+0xcd			/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:724
// #	0x10c15e4	runtime/pprof.writeGoroutine+0xa4			/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:684
// #	0x10be171	runtime/pprof.(*Profile).WriteTo+0x3f1			/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:331
// #	0x10cd10c	main.main+0x18c						/Users/Akira/go/src/practice/02sync/gyaku03.go:16
// #	0x1036fd5	runtime.main+0x255					/usr/local/Cellar/go/1.16.6/libexec/src/runtime/proc.go:225
//  
//  
// 2  2021-10-24 23:23:50.271794 +0900 JST m=+1.038768319
// goroutine 1 [running]:
// runtime/pprof.writeGoroutineStacks(0x111fb88, 0xc000134008, 0x0, 0x10aef31)
//  	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:693 +0x9f
// runtime/pprof.writeGoroutine(0x111fb88, 0xc000134008, 0x2, 0x1103338, 0xc000157e50)
//  	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:682 +0x45
// runtime/pprof.(*Profile).WriteTo(0x11a7280, 0x111fb88, 0xc000134008, 0x2, 0xc000157f38, 0x2)
//  	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:331 +0x3f2
// main.main()
//  	/Users/Akira/go/src/practice/02sync/gyaku03.go:18 +0x2a5
//  
// goroutine 19 [sleep]:
// time.Sleep(0x77359400)
//  	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/time.go:193 +0xd2
// main.goroutine2()
//  	/Users/Akira/go/src/practice/02sync/gyaku03.go:29 +0x2e
// created by main.main
//  	/Users/Akira/go/src/practice/02sync/gyaku03.go:12 +0x5f
// Goroutine2 finished
//  
// 1  2021-10-24 23:23:51.274421 +0900 JST m=+2.041388559
// goroutine profile: total 1
// 1 @ 0x106223d 0x10c180e 0x10c15e5 0x10be172 0x10cd10d 0x1036fd6 0x1067101
// #	0x106223c	runtime/pprof.runtime_goroutineProfileWithLabels+0x5c	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/mprof.go:716
// #	0x10c180d	runtime/pprof.writeRuntimeProfile+0xcd			/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:724
// #	0x10c15e4	runtime/pprof.writeGoroutine+0xa4			/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:684
// #	0x10be171	runtime/pprof.(*Profile).WriteTo+0x3f1			/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:331
// #	0x10cd10c	main.main+0x18c						/Users/Akira/go/src/practice/02sync/gyaku03.go:16
// #	0x1036fd5	runtime.main+0x255					/usr/local/Cellar/go/1.16.6/libexec/src/runtime/proc.go:225
//  
//  
// 2  2021-10-24 23:23:51.274981 +0900 JST m=+2.041948417
// goroutine 1 [running]:
// runtime/pprof.writeGoroutineStacks(0x111fb88, 0xc000134008, 0x0, 0x10aef31)
//  	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:693 +0x9f
// runtime/pprof.writeGoroutine(0x111fb88, 0xc000134008, 0x2, 0x1103338, 0xc000157e50)
//  	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:682 +0x45
// runtime/pprof.(*Profile).WriteTo(0x11a7280, 0x111fb88, 0xc000134008, 0x2, 0xc000157f38, 0x2)
//  	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/pprof/pprof.go:331 +0x3f2
// main.main()
//  	/Users/Akira/go/src/practice/02sync/gyaku03.go:18 +0x2a5
//  
// Compilation finished at Sun Oct 24 23:23:52
