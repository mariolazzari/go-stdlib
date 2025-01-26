package main

import (
	"fmt"
	"regexp"
)

func main() {
	// compile regex
	re := regexp.MustCompile(`\d+`)

	// matching
	match := re.MatchString("123")
	fmt.Println("Match 123:", match)

	// replace
	res := re.ReplaceAllString("abc123def456", "X")
	fmt.Println(res)
}
