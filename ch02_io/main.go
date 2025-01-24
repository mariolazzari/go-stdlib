package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// open file
	file, err := os.Open("example.txt")
	checkError(err)
	defer file.Close()

	// read file
	content, err := io.ReadAll(file)
	checkError(err)
	fmt.Println("File content:", string(content))

	// create file
	file, err = os.Create("start.txt")
	checkError(err)
	file.WriteString("Hello, World!")
	defer file.Close()

	// copy file
	sourceFile, err := os.Open("start.txt")
	checkError(err)
	defer sourceFile.Close()

	destFile, err := os.Create("end.txt")
	checkError(err)
	defer destFile.Close()

	// cope content
	writtenByes, err := io.Copy(destFile, sourceFile)
	checkError(err)
	fmt.Printf("Copied %d bytes\n", writtenByes)

	var num int
	fmt.Println("Enter a number:")
	fmt.Scanf("%d", &num)
	fmt.Println("You entered:", num)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
