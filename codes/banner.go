package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	bannerMap := make(map[rune][]string)
	charCode := ' '

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error reading file")
	}
	if len(file) == 0 {
		return nil, errors.New("error empty file")
	}

	data := strings.ReplaceAll(string(file), "\r\n", "\n")

	
	rawFiles := strings.Split(data, "\n\n")

	for _, raw := range rawFiles {
		lines := strings.Split(raw, "\n")

		if len(lines) < 8 {
			continue
		}
		bannerMap[charCode] = lines[:8]
		charCode++
	}

	return bannerMap, nil
}
