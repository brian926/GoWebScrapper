package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("https://www.gutenberg.org"),
	)

	c.OnHTML("div.chapter", func(e *colly.HTMLElement) {

		fmt.Println(e.Text)
	})

	c.Visit("https://www.gutenberg.org/files/2600/2600-h/2600-h.htm")
}
