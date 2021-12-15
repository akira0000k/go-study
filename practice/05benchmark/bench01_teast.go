package main

import (
	"fmt"
	"testing"
)
/*
   subject: how to use benchmark. $ go test -bench . -benchmem
   go標準のbenchmark機能の使い方  より
*/
func BenchmarkAppend_AllocateEveryTime(b *testing.B) {
    base := []string{}
    b.ResetTimer()
    // Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
    for i := 0; i < b.N; i++ {
        // 都度append
        base = append(base, fmt.Sprintf("no%d", i))
    }
}

func BenchmarkAppend_AllocateOnce(b *testing.B) {
    //最初に長さを決める
    base := make([]string, b.N)
    b.ResetTimer()
    // Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
    for i := 0; i < b.N; i++ {
        base[i] = fmt.Sprintf("no%d", i)
    }
}
// Akira@MBP 05benchmark % go test
// testing: warning: no tests to run
// PASS
// ok  	practice/05benchmark	0.570s
//  
//  
// Akira@MBP 05benchmark % go test -bench .
// goos: darwin
// goarch: amd64
// pkg: practice/05benchmark
// cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
// BenchmarkAppend_AllocateEveryTime-4   	 5953842	       183.1 ns/op
// BenchmarkAppend_AllocateOnce-4        	10914158	       120.6 ns/op
// PASS
// ok  	practice/05benchmark	2.870s
//  
//  
// Akira@MBP 05benchmark % go test -bench . -benchmem
// goos: darwin
// goarch: amd64
// pkg: practice/05benchmark
// cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
// BenchmarkAppend_AllocateEveryTime-4   	 5955879	       184.5 ns/op	     111 B/op	       1 allocs/op
// BenchmarkAppend_AllocateOnce-4        	10811313	       120.2 ns/op	      23 B/op	       1 allocs/op
// PASS
// ok  	practice/05benchmark	2.864s
//  
//  
// Akira@MBP 05benchmark %
