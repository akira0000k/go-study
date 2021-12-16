package main
import (
	"fmt"
	"os"
)
/*
   subject : Exit
*/
func main() {
	defer fmt.Println("****END***!")
	os.Exit(3)//deferは実行されない
}
// Akira@MBP 10process % go run process04.go
// ****END***!

// Akira@MBP 10process % go run process04.go
// # command-line-arguments
// ./process04.go:9:2: undefined: os

// Akira@MBP 10process % go run process04.go
// exit status 3
// Akira@MBP 10process %
