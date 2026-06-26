package main

import (
	"fmt"
	"os"
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {
	var result strings.Builder
	userInput := strings.Split(input, "\\n")

	for i, line := range userInput {
		if line == "" {
			if i < len(userInput)-1 {
				result.WriteString("\n")
			}
			continue
		}
		for _, ch := range input {
			if ch < ' ' || ch > '~' {
				fmt.Println("non ascii character")
				os.Exit(1)
			}
		}
		row := RenderLines(line, banner)
		for _, lines := range row {
			result.WriteString(lines)
			result.WriteString("\n")
		}

		if i < len(userInput)-1 {
			result.WriteString("\n")
		}

	}

	return result.String()

}
