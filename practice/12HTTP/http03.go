//URLs provide a uniform way to locate resources. Here’s how to parse URLs in Go.
package main
import (
	"fmt"
	"net"
	"net/url"
)
/*
   subject : サンプルで学ぶ Go 言語：URL Parsing     https://www.spinute.org/go-by-example/url-parsing.html
*/
func main() {

	//We’ll parse this example URL, which includes a scheme, authentication info, host, port, path, query params, and query fragment.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	//Parse the URL and ensure there are no errors.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	//Accessing the scheme is straightforward.
	fmt.Println(u.Scheme)// postgres

	//User contains all authentication info; call Username and Password on this for individual values.
	fmt.Println(u.User)// user:pass
	fmt.Println(u.User.Username())// user
	p, _ := u.User.Password()
	fmt.Println(p)// pass

	//The Host contains both the hostname and the port, if present. Use SplitHostPort to extract them.
	fmt.Println(u.Host)// host.com:5432
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)// host.com:5432
	fmt.Println(port)// 5432

	//Here we extract the path and the fragment after the #.
	fmt.Println(u.Path)// /path
	fmt.Println(u.Fragment)// f

	//To get query params in a string of k=v format, use RawQuery.
	//You can also parse query params into a map.
	//The parsed query param maps are from strings to slices of strings,
	//so index into [0] if you only want the first value.
	fmt.Println(u.RawQuery)// k=v
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)// map[k:[v]]
	fmt.Println(m["k"][0])// v
}
// -*- mode: compilation; default-directory: "~/go/src/practice/12HTTP/" -*-
// Compilation started at Tue Oct 19 23:02:13
//  
// go run http03.go
// postgres
// user:pass
// user
// pass
// host.com:5432
// host.com
// 5432
// /path
// f
// k=v
// map[k:[v]]
// v
//  
// Compilation finished at Tue Oct 19 23:02:14
