package main

import (
	"fmt"
	"strconv"
	"sync"
	"io/ioutil"
	"os"
)
/*
   subject : 逆引きgolang goroutine間の競合
*/
func main() {
	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	write(0)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go countupGoroutine(wg, i, m)
	}
	wg.Wait()
	fmt.Println("result:", read())
	fmt.Println("end main")
}

func countupGoroutine(wg *sync.WaitGroup, id int, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	defer m.Unlock()
	
	counter := read() + 1
	write(counter)
	fmt.Printf("wk%d wrote %d\n", id, counter)
}

func write(i int) {
	s := strconv.Itoa(i)
	ioutil.WriteFile("count.txt", []byte(s), os.ModePerm)
}

func read() int {
	t, _ := ioutil.ReadFile("count.txt")
	i, _ := strconv.Atoi(string(t))
	return i
}
// -*- mode: compilation; default-directory: "~/go/src/practice/02sync/" -*-
// Compilation started at Sun Oct 24 23:54:26
//  
// go run gyaku05.go
// wk0 wrote 1
// wk9 wrote 2
// wk5 wrote 3
// wk6 wrote 4
// wk7 wrote 5
// wk8 wrote 6
// wk2 wrote 7
// wk3 wrote 8
// wk4 wrote 9
// wk1 wrote 10
// result: 10
// end main
//  
// Compilation finished at Sun Oct 24 23:54:27
