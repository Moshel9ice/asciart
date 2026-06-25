package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("usage: go run . [A]")
		os.Exit(1)
	}

	userInput := os.Args[1]
	

	Funcall, _ := LoadBanner("standard.txt") 
	

	// input := "A"
	lines := RenderLines(userInput, Funcall)
	for _, line := range lines {
		fmt.Println(line)
	}

	
	

}
