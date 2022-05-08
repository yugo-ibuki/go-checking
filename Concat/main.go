package main

import "log"

func main() {
	// この書き方は遅くなる
	src := []string{"Back", "To", "The", "Future", "Part", "III"}
	var title string
	for i, word := range src {
		if i != 0 {
			title += " "
		}
		title += word
	}
	log.Print(title)
}
