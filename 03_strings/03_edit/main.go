package main

import (
	"fmt"
	"strings"
)

const sentence = "Mario Lazzari senior full stack developer"
const spaces = "    Mario   Lazzari   senior  full    stack    developer       "

const name = "Mario"
const surname = "Lazzari"

func main() {
	fmt.Println("Sentence", sentence)

	// split
	tokens := strings.Split(sentence, " ")
	fmt.Println("Split sentence by space", tokens)

	// fields
	fields := strings.Fields(sentence)
	fmt.Println("Fields sentece", fields)
	fmt.Println("Too meny spaces", spaces)
	fields = strings.Fields(spaces)
	fmt.Println("Too meny spaces fields", fields)

	// join
	toJojn := []string{name, surname}
	fullName := strings.Join(toJojn, " ")
	fmt.Println("Join name and surname", fullName)

	// replace
	links := strings.Replace("link link link", "k", "ked", 1)
	fmt.Println("Replace k with ked:", links)

	colors := "green green green"
	colors = strings.ReplaceAll(colors, "green", "blue")
	fmt.Println("Replace all green with blue:", colors)

	// string builder
	var sb strings.Builder
	sb.WriteString("Mario")
	sb.WriteString(" ")
	sb.WriteString("Lazzari")
	fmt.Println("String builder:", sb.String())

	// repeat
	fmt.Println("Repeat 3 times Mario:", strings.Repeat("Mario", 3))
}
