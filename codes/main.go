package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) <2 || len(os.Args) > 3 {
		fmt.Println("usage: go run . [A]")
		os.Exit(1)
	}

	banner := "standard.txt"

	if len(os.Args) == 3{
		banner = os.Args[2]
		if !strings.HasSuffix(banner, ".txt"){
			banner+=".txt"
		}

	}
	userInput := os.Args[1]

	Funcall, err:= LoadBanner(banner)
	if err!=nil{
		fmt.Println("no banner file")
		return
	}

	data := GenerateArt(userInput, Funcall)
	fmt.Print(data)

}
