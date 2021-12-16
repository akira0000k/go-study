//Writing a basic HTTP server is easy using the net/http package.
package main
import (
	"fmt"
	"net/http"
)
/*
   subject : サンプルで学ぶ Go 言語：HTTP Servers    https://www.spinute.org/go-by-example/http-servers.html   
*/
//A fundamental concept in net/http servers is handlers.
//A handler is an object implementing the http.Handler interface.
//A common way to write a handler is by using the http.HandlerFunc adapter on functions with the appropriate signature.


func hello(w http.ResponseWriter, req *http.Request) {

	//Functions serving as handlers take a http.ResponseWriter and a http.Request as arguments.
	//The response writer is used to fill in the HTTP response.
	//Here our simple response is just “hello\n”.
	fmt.Fprintf(w, "hello\n")
}



func headers(w http.ResponseWriter, req *http.Request) {

	//This handler does something a little more sophisticated by reading all the HTTP request headers and echoing them into the response body.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}



func main() {

	//We register our handlers on server routes using the http.HandleFunc convenience function.
	//It sets up the default router in the net/http package and takes a function as an argument.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	//Finally, we call the ListenAndServe with the port and a handler.
	//nil tells it to use the default router we’ve just set up.
	http.ListenAndServe(":8090", nil)
}
//  
// Akira@MBP practice % ls
// total 25288
// drwxr-xr-x  31 Akira  staff      992 10 19 22:27 ./
// drwxr-xr-x   5 Akira  staff      160 10  1 20:07 ../
// drwxr-xr-x  33 Akira  staff     1056 10 16 19:38 01practice/
// drwxr-xr-x   8 Akira  staff      256 10  7 15:06 022burst/
// drwxr-xr-x  46 Akira  staff     1472 10  1 18:52 02sync/
// drwxr-xr-x  16 Akira  staff      512  9 29 18:28 03anti-pattern/
// drwxr-xr-x   4 Akira  staff      128 10  8 21:19 04tests/
// drwxr-xr-x   6 Akira  staff      192  9 26 21:50 05benchmark/
// drwxr-xr-x   6 Akira  staff      192 10  1 15:50 06context/
// drwxr-xr-x  14 Akira  staff      448 10  5 17:37 07class/
// drwxr-xr-x  21 Akira  staff      672 10 14 18:35 08misc/
// drwxr-xr-x   7 Akira  staff      224 10 11 10:16 09json/
// drwxr-xr-x  10 Akira  staff      320 10 16 19:33 10process/
// drwxr-xr-x  26 Akira  staff      832 10 19 21:03 11fileIO/
// drwxr-xr-x   5 Akira  staff      160 10 19 22:24 12HTTP/
// -rwxr-xr-x   1 Akira  staff  2087920 10 16 19:51 file01*
// -rw-r--r--   1 Akira  staff     2344 10 16 19:46 file01.go
// -rwxr-xr-x   1 Akira  staff  2324976 10 19 19:23 file05-2*
// -rw-r--r--   1 Akira  staff     2854 10 19 19:24 file05-2.go
// -rw-r--r--   1 Akira  staff     2790 10 19 19:16 file05-2.go~
// -rw-r--r--   1 Akira  staff       82 10  1 19:01 go.mod
// -rw-r--r--   1 Akira  staff      175 10  1 19:01 go.sum
// -rw-r--r--   1 Akira  staff     1442 10 19 22:24 http02.go
// -rwxr-xr-x   1 Akira  staff  2053328 10 16 18:07 process05*
// -rw-r--r--   1 Akira  staff     1476 10 16 18:00 process05.go
// -rwxr-xr-x   1 Akira  staff  2032032 10 16 18:09 process06*
// -rw-r--r--   1 Akira  staff      246 10 16 18:09 process06.go
// -rwxr-xr-x   1 Akira  staff  2197856 10 16 19:22 process07*
// -rw-r--r--   1 Akira  staff     1637 10 16 19:22 process07.go
// -rwxr-xr-x   1 Akira  staff  2197856 10 16 19:29 process08*
// -rw-r--r--   1 Akira  staff     1293 10 16 19:29 process08.go
// Akira@MBP practice % go build http02.go
// Akira@MBP practice % ./http02 &
// [1] 81172
// Akira@MBP practice % curl localhost:8090/hello
// hello
// Akira@MBP practice % 

// Akira@MBP practice % ps -ef | grep http
//     0   109     1   0 22 921  ??         0:36.22 /Library/PrivilegedHelperTools/licenseDaemon.app/Contents/MacOS/licenseDaemon --backurl https://activation.paceap.com/InitiateActivation
//   502 81172 79465   0 10:28PM ttys000    0:00.01 ./http02
//   502 81196 79465   0 10:31PM ttys000    0:00.00 grep http
// Akira@MBP practice % curl localhost:8090/hello/
// 404 page not found
// Akira@MBP practice % curl localhost:8090/hello
// hello
// Akira@MBP practice % curl localhost:8090/hello
// hello
// Akira@MBP practice % curl localhost:8090/hello
// hello
// Akira@MBP practice % killall http02
// [1]  + terminated  ./http02
// Akira@MBP practice % curl localhost:8090/hello
// curl: (7) Failed to connect to localhost port 8090: Connection refused
// Akira@MBP practice %
