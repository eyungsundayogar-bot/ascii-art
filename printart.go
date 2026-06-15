package main

import "fmt"

func PrintArt(text string, banner map[rune][]string) {
	for row := range 8 {
		for _, char := range text {
			charArt := GetCharArt(char, banner)
			fmt.Print(charArt[row])
		}
		fmt.Println()
	}
}
