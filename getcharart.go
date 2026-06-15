package main

func GetCharArt(r rune, banner map[rune][]string) []string {
	if r < 32 || r > 126 {
		return []string{}
	}

	return banner[r]
}
