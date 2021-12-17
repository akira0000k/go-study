package main
import (
	"fmt"
	"io/ioutil" // Implements some I/O utility functions.
	"net/http"  // Yes, a web server!
)

func main() {
	learnWebProgramming()
}


/*
   ******************************
   Web Programming
   ******************************
*/
// A single function from package http starts a web server.
func learnWebProgramming() {

	// First parameter of ListenAndServe is TCP address to listen to.
	// Second parameter is an interface, specifically http.Handler.
	go func() {
		err := http.ListenAndServe(":8080", pair{})
		fmt.Println(err) // don't ignore errors
	}()

	requestServer()
}

// Define pair as a struct with two fields, ints named x and y.
type pair struct {
	x, y int
}

// Make pair an http.Handler by implementing its only method, ServeHTTP.
func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Serve data with a method of http.ResponseWriter.
	w.Write([]byte("You learned Go in Y minutes!"))
}

func requestServer() {
	resp, err := http.Get("http://localhost:8080")
	fmt.Println(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("\nWebserver said: `%s`\n", string(body))
}
// -*- mode: compilation; default-directory: "~/go/src/practice/15LearnXinYmin/" -*-
// Compilation started at Thu Oct 28 00:32:02
//  
// go run learn04.go
// <nil>
//  
// Webserver said: `You learned Go in Y minutes!`
// Compilation finished at Thu Oct 28 00:32:04
