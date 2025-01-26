package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Mario Lazzari contains Mario?", strings.Contains("Mario Lazzari", "Mario"))
	fmt.Println("Mario Lazzari index of Lazzari?", strings.Index("Mario Lazzari", "Lazzari"))
	fmt.Println("Mario Lazzari index of x?", strings.Index("Mario Lazzari", "x"))

	fmt.Println("Mario Lazzari has prefix Mario?", strings.HasPrefix("Mario Lazzari", "Mario"))
	fmt.Println("Mario Lazzari has suffix Lazzari?", strings.HasSuffix("Mario Lazzari", "Lazzari"))
	fmt.Println("Mario Lazzari has prefix Maria?", strings.HasPrefix("Mario Lazzari", "Maria"))

}
