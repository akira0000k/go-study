package main

import (
	//"fmt"
	"testing"
)
/*
   subject: slice and channel compare
   Hori Blog
   Go の channel 処理パターン集
   高速に送りまくる
*/
func BenchmarkWithSimpleSlice(b *testing.B) {
	src := make([]int, 10000)
	dst := make([]int, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// write
		for i := range src {
			dst[i] = src[i]
		}
		// read
		for i := 0; i < len(src); i++ {
			_ = dst[i]
		}
	}
}

func BenchmarkWithChannel(b *testing.B) {
	src := make([]int, 10000)
	ch := make(chan int, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// write
		for i := range src {
			ch <- src[i]
		}
		// read
		for i := 0; i < len(src); i++ {
			<-ch
		}
	}
}
// Akira@MBP 05benchmark % go test -bench .
// goos: darwin
// goarch: amd64
// pkg: practice/05benchmark
// cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
// BenchmarkWithSimpleSlice-4   	  134312	      8413 ns/op
// BenchmarkWithChannel-4       	    2486	    480272 ns/op
// PASS
// ok  	practice/05benchmark	2.565s
// Akira@MBP 05benchmark %
