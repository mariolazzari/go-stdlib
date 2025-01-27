package main

import (
	"fmt"
	"net/url"
)

func main() {
	// url parsing
	u, err := url.Parse("https://mariolazzari.it:3333/path?key=value#fragment")
	if err != nil {
		panic(err)
	}
	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("Host:", u.Host)
	fmt.Println("Path:", u.Path)
	fmt.Println("Raw query:", u.RawQuery)
	fmt.Println("Fragment:", u.Fragment)

	// query params
	q := u.Query()
	v := q.Get("key")
	fmt.Println("Query value:", v)

	// set params
	q.Set("x", "y")
	u.RawQuery = q.Encode()
	fmt.Println(u.RawQuery)

	// encoding
	encoded := url.QueryEscape("Testing, one two three")
	fmt.Println(encoded)

	// decoding
	decoded, _ := url.QueryUnescape(encoded)
	fmt.Println(decoded)

	// URL building
	u = &url.URL{Scheme: "https", Host: "mariolazzari.it", Path: "/path", RawQuery: "key=value"}
	fmt.Println(u.String())
}
