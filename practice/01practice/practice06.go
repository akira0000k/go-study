package main

import "fmt"
/*
   subject : サンプルで学ぶ Go 言語：Range
*/
func main() {
	main1()
	main2()
}
func main1() {
	fmt.Println("start main1")
	nums := []int{2, 3, 4}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	for i := range nums {
		fmt.Println("index:", i)
	}
	for _, num := range nums {
		fmt.Println("value:", num)
	}

	kvs := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "cabage",
	}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	
	for k := range kvs {
		fmt.Println("key:", k,
		)
	}
	
	for i, c := range "gone with the wind." {
		//fmt.Println(i, c)
		fmt.Printf("i=%d  c=%c (%d)\n", i, c, c)
	}
}

func main2() {
	fmt.Println("start main2")
	var kvs map[string]string
	kvs = map[string]string {
		"a": "apple",
		"b": "banana",
		"c": "cacao",
	}
	var k, v string
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	fmt.Printf("%s -> %s\n", k, v)
	for k, _ = range kvs {
		fmt.Printf("%s -> %s\n", k, kvs[k])
	}
	fmt.Printf("%s -> %s\n", k, v)


	for k, v :=
		range map[string]int{
			"key":3,
			"kcia":99,
			"kpop":22,
		} {
		fmt.Printf("%s -> %d\n", k, v)
	}
	fmt.Printf("%s -> %s\n", k, v)

	//var m map[string]int
	//m["kkk"]=321 //panic: assignment to entry in nil map

	m := map[string]int{}
	m["kkk"]=321 //ok
	m["cia"]=567
	m["kgb"]=987
	for k, v := range m {
		fmt.Println(k, v)
	}

}
// -*- mode: compilation; default-directory: "~/go/src/practice/01practice/" -*-
// Compilation started at Sat Oct 30 15:39:56
//  
// go run practice06.go
// start main1
// sum: 9
// index: 1
// index: 0
// index: 1
// index: 2
// value: 2
// value: 3
// value: 4
// a -> apple
// b -> banana
// c -> cabage
// key: b
// key: c
// key: a
// i=0  c=g (103)
// i=1  c=o (111)
// i=2  c=n (110)
// i=3  c=e (101)
// i=4  c=  (32)
// i=5  c=w (119)
// i=6  c=i (105)
// i=7  c=t (116)
// i=8  c=h (104)
// i=9  c=  (32)
// i=10  c=t (116)
// i=11  c=h (104)
// i=12  c=e (101)
// i=13  c=  (32)
// i=14  c=w (119)
// i=15  c=i (105)
// i=16  c=n (110)
// i=17  c=d (100)
// i=18  c=. (46)

// start main2
// a -> apple
// b -> banana
// c -> cacao
//  -> 
// a -> apple
// b -> banana
// c -> cacao
// c -> 
// key -> 3
// kcia -> 99
// kpop -> 22
// c -> 
// kkk 321
// cia 567
// kgb 987
//  
// Compilation finished at Sat Oct 30 15:39:57
