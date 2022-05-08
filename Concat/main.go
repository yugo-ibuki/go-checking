package main

import (
	"log"
	"strings"
)

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

	// 以下のstrings.Builderを使う書き方の方が良い
	var builder strings.Builder
	builder.Grow(100) //　最大100文字以下と仮定できる場合
	for i, word := range src {
		if i != 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(word) // ここで書き込み
	}
	log.Print(builder.String()) // 出力
}
