package main

import (
	"fmt"
	"strings"
)

func main() {
	urls := []string{"example.org", "example.com"}
	tld := ".org"
	tlds := countTlds(urls, tld)
	fmt.Println(".org =", tlds)

	urls = []string{"example.net", "example.org", "test.net"}
	tld = ".net"
	tlds = countTlds(urls, tld)
	fmt.Println(".net =", tlds)
}

func countTlds(urls []string, tld string) int {
	tlds := 0

	for _, url := range urls {
		if strings.HasSuffix(url, tld) {
			tlds++
		}
	}

	return tlds
}
