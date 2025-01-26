package main

import (
	"fmt"
	"strings"
)

const dogs = "dogs"
const cats = "cats"
const DOGS = "DOGS"

func main() {
	fmt.Println("Comparing dogs and cats", strings.Compare(dogs, cats))
	fmt.Println("Comparing cats and cats", strings.Compare(cats, cats))
	fmt.Println("Comparing cats and dogs", strings.Compare(cats, dogs))

	fmt.Println("Equal fold dogs and cats", strings.EqualFold(dogs, cats))
	fmt.Println("Equal fold cats and cats", strings.EqualFold(cats, cats))
	fmt.Println("Equal fold dogs and DOGS", strings.EqualFold(dogs, DOGS))

	fmt.Println("To upper dogs", strings.ToUpper(dogs))
	fmt.Println("To lower DOGS", strings.ToLower(DOGS))

	fmt.Println("Trim spaces", strings.TrimSpace("  Hello, World!  "))
	fmt.Println("Trim hello", strings.Trim("Hello, World!", "Hello"))

}
