package main
import (
	"bufio"
	"fmt"
	"net/http"
)
/*
   subject : サンプルで学ぶ Go 言語　HTTP Clients　　https://www.spinute.org/go-by-example/
*/

//The Go standard library comes with excellent support for HTTP clients and servers in the net/http package.
//In this example we’ll use it to issue simple HTTP requests.

func main() {

	//Issue an HTTP GET request to a server.
	//http.Get is a convenient shortcut around creating an http.Client object and calling its Get method;
	//it uses the http.DefaultClient object which has useful default settings.

	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//Print the HTTP response status.
	fmt.Println("Response status:", resp.Status)

	//Print the first 5 lines of the response body.
	scanner := bufio.NewScanner(resp.Body)
	//for i := 0; scanner.Scan() && i < 5; i++ {
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
// -*- mode: compilation; default-directory: "~/go/src/practice/12HTTP/" -*-
// Compilation started at Tue Oct 19 22:17:17
//  
// go run http01.go
// Response status: 200 OK
// <!DOCTYPE html>
// <html>
//   <head>
//     <meta charset="utf-8">
//     <title>Go by Example</title>
//     <link rel=stylesheet href="site.css">
//   </head>
//   <body>
//     <div id="intro">
//       <h2><a href="./">Go by Example</a></h2>
//       <p>
//         <a href="http://golang.org">Go</a> is an
//         open source programming language designed for
//         building simple, fast, and reliable software.
//         Please read the official
//         <a href="https://golang.org/doc/tutorial/getting-started">documentation</a>
//         to learn a bit about Go code, tools packages,
//         and modules.
//       </p>
//  
//       <p>
//         <em>Go by Example</em> is a hands-on introduction
//         to Go using annotated example programs. Check out
//         the <a href="hello-world">first example</a> or
//         browse the full list below.
//       </p>
//  
//       <ul>
//       
//         <li><a href="hello-world">Hello World</a></li>
//       
//         <li><a href="values">Values</a></li>
//       
//         <li><a href="variables">Variables</a></li>
//       
//         <li><a href="constants">Constants</a></li>
//       
//         <li><a href="for">For</a></li>
//       
//         <li><a href="if-else">If/Else</a></li>
//       
//         <li><a href="switch">Switch</a></li>
//       
//         <li><a href="arrays">Arrays</a></li>
//       
//         <li><a href="slices">Slices</a></li>
//       
//         <li><a href="maps">Maps</a></li>
//       
//         <li><a href="range">Range</a></li>
//       
//         <li><a href="functions">Functions</a></li>
//       
//         <li><a href="multiple-return-values">Multiple Return Values</a></li>
//       
//         <li><a href="variadic-functions">Variadic Functions</a></li>
//       
//         <li><a href="closures">Closures</a></li>
//       
//         <li><a href="recursion">Recursion</a></li>
//       
//         <li><a href="pointers">Pointers</a></li>
//       
//         <li><a href="structs">Structs</a></li>
//       
//         <li><a href="methods">Methods</a></li>
//       
//         <li><a href="interfaces">Interfaces</a></li>
//       
//         <li><a href="errors">Errors</a></li>
//       
//         <li><a href="goroutines">Goroutines</a></li>
//       
//         <li><a href="channels">Channels</a></li>
//       
//         <li><a href="channel-buffering">Channel Buffering</a></li>
//       
//         <li><a href="channel-synchronization">Channel Synchronization</a></li>
//       
//         <li><a href="channel-directions">Channel Directions</a></li>
//       
//         <li><a href="select">Select</a></li>
//       
//         <li><a href="timeouts">Timeouts</a></li>
//       
//         <li><a href="non-blocking-channel-operations">Non-Blocking Channel Operations</a></li>
//       
//         <li><a href="closing-channels">Closing Channels</a></li>
//       
//         <li><a href="range-over-channels">Range over Channels</a></li>
//       
//         <li><a href="timers">Timers</a></li>
//       
//         <li><a href="tickers">Tickers</a></li>
//       
//         <li><a href="worker-pools">Worker Pools</a></li>
//       
//         <li><a href="waitgroups">WaitGroups</a></li>
//       
//         <li><a href="rate-limiting">Rate Limiting</a></li>
//       
//         <li><a href="atomic-counters">Atomic Counters</a></li>
//       
//         <li><a href="mutexes">Mutexes</a></li>
//       
//         <li><a href="stateful-goroutines">Stateful Goroutines</a></li>
//       
//         <li><a href="sorting">Sorting</a></li>
//       
//         <li><a href="sorting-by-functions">Sorting by Functions</a></li>
//       
//         <li><a href="panic">Panic</a></li>
//       
//         <li><a href="defer">Defer</a></li>
//       
//         <li><a href="recover">Recover</a></li>
//       
//         <li><a href="collection-functions">Collection Functions</a></li>
//       
//         <li><a href="string-functions">String Functions</a></li>
//       
//         <li><a href="string-formatting">String Formatting</a></li>
//       
//         <li><a href="regular-expressions">Regular Expressions</a></li>
//       
//         <li><a href="json">JSON</a></li>
//       
//         <li><a href="xml">XML</a></li>
//       
//         <li><a href="time">Time</a></li>
//       
//         <li><a href="epoch">Epoch</a></li>
//       
//         <li><a href="time-formatting-parsing">Time Formatting / Parsing</a></li>
//       
//         <li><a href="random-numbers">Random Numbers</a></li>
//       
//         <li><a href="number-parsing">Number Parsing</a></li>
//       
//         <li><a href="url-parsing">URL Parsing</a></li>
//       
//         <li><a href="sha1-hashes">SHA1 Hashes</a></li>
//       
//         <li><a href="base64-encoding">Base64 Encoding</a></li>
//       
//         <li><a href="reading-files">Reading Files</a></li>
//       
//         <li><a href="writing-files">Writing Files</a></li>
//       
//         <li><a href="line-filters">Line Filters</a></li>
//       
//         <li><a href="file-paths">File Paths</a></li>
//       
//         <li><a href="directories">Directories</a></li>
//       
//         <li><a href="temporary-files-and-directories">Temporary Files and Directories</a></li>
//       
//         <li><a href="testing">Testing</a></li>
//       
//         <li><a href="command-line-arguments">Command-Line Arguments</a></li>
//       
//         <li><a href="command-line-flags">Command-Line Flags</a></li>
//       
//         <li><a href="command-line-subcommands">Command-Line Subcommands</a></li>
//       
//         <li><a href="environment-variables">Environment Variables</a></li>
//       
//         <li><a href="http-clients">HTTP Clients</a></li>
//       
//         <li><a href="http-servers">HTTP Servers</a></li>
//       
//         <li><a href="context">Context</a></li>
//       
//         <li><a href="spawning-processes">Spawning Processes</a></li>
//       
//         <li><a href="execing-processes">Exec'ing Processes</a></li>
//       
//         <li><a href="signals">Signals</a></li>
//       
//         <li><a href="exit">Exit</a></li>
//       
//       </ul>
//  
//     <p class="footer">
//       by <a href="https://markmcgranaghan.com">Mark McGranaghan</a> and <a href="https://eli.thegreenplace.net">Eli Bendersky</a> | <a href="https://github.com/mmcgrana/gobyexample">source</a> | <a href="https://github.com/mmcgrana/gobyexample#license">license</a>
//     </p>
//  
//     </div>
//   </body>
// </html>
//  
// Compilation finished at Tue Oct 19 22:17:18
