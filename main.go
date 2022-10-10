package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Chapter struct {
	Chapters string `json:"chapter"`
	Page     string `json:"page"`
}

func main() {
	c := colly.NewCollector()

	c.OnHTML(".chapter", func(e *colly.HTMLElement) {

		Chapter := Chapter{
			Chapters: e.ChildText("h2"),
			Page:     e.ChildText("p"),
		}

		fmt.Println(Chapter)
	})

	c.Visit("https://www.gutenberg.org/files/2600/2600-h/2600-h.htm")
}
