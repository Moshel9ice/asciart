package main

import (
	"strings"
)


func RenderLines(input string, banner map[rune][]string) []string {
	var result []string

	for row := 0; row < 8; row++ {

		var char strings.Builder

		for _, ch := range input {

			char.WriteString(banner[ch][row])
		}
		result = append(result, char.String())
	}
	return result

}
