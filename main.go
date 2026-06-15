package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("use: go run . [string]")
		os.Exit(1)
	}

	banner := "standard.txt"

	if len(os.Args) == 3 {
		banner = os.Args[2]
	}

	if !strings.HasSuffix(banner, ".txt") {
		banner += ".txt"
	}
	input := os.Args[1]
	if input == "" {
		return
	}

	if input == "\\n" {
		fmt.Println()
		return
	}
	data, err := LoadBanner(banner)
	if err != nil {
		fmt.Println("Error loading file", err)
		os.Exit(1)
	}
	text := strings.Split(input, "\\n")

	for _, texts := range text {
		if texts == "" {
			fmt.Println()
		} else {
			PrintArt(texts, data)
		}
	}
}
