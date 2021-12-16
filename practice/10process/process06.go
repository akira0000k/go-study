
package main
import(
	"fmt"
	"os"
)
/*
   subject : Command-Line Arguments
*/
func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
// Akira@MBP practice % go build process06.go
//  
// Akira@MBP practice % ./process06 
// panic: runtime error: index out of range [3] with length 1
//  
// goroutine 1 [running]:
// main.main()
//  	/Users/Akira/go/src/practice/process06.go:13 +0x20d
//  
// Akira@MBP practice % ./process06 1 2 3
// [./process06 1 2 3]
// [1 2 3]
// 3
// Akira@MBP practice %
