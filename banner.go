package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	if filename == "" {
		return nil, fmt.Errorf("filename is empty")
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error failed to read file %v", err)
	}
	character := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(character, "\n")

	banner := make(map[rune][]string)

	currentChar := rune(32)

	for i := 1; i < len(lines); i += 9 {
		if i+8 > len(lines) {
			break
		}
		charArt := lines[i : i+8]
		banner[currentChar] = charArt
		currentChar++
	}
	if len(banner) == 0 {
		return nil, fmt.Errorf("banner file is empty")
	}
	if len(banner) != 95 {
		return nil, fmt.Errorf("banner file is incomplete expected 95 got %d", len(banner))
	}
	return banner, nil
}
