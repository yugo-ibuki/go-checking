package main

import (
	"fmt"
	"time"
)

type Author struct {
	FirstName string
	LastName  string
}

type Book struct {
	Title      string
	Author     Author
	Publisher  string
	ReleasedAt time.Time
	ISBN       string
}

func main() {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	book := Book{
		Title: "Book Title",
		Author: Author{
			FirstName: "FirstName",
			LastName:  "LastName",
		},
		Publisher:  "Publisher",
		ISBN:       "444444444",
		ReleasedAt: time.Date(2017, time.June, 14, 0, 0, 0, 0, jst),
	}
	fmt.Println(book.Title)
}
