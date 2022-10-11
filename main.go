package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Chapter struct {
	Ch   string `json:"chapter"`
	Page string `json:"page"`
}

func main() {
	c := colly.NewCollector()

	story := make([]Chapter, 0, 361)

	c.OnHTML(".chapter", func(e *colly.HTMLElement) {

		chapter := Chapter{
			Ch:   e.ChildText("h2"),
			Page: e.ChildText("p"),
		}

		story = append(story, chapter)

	})

	c.Visit("https://www.gutenberg.org/files/2600/2600-h/2600-h.htm")

	for i := range story {
		if story[i].Page == "" {
			fmt.Println(story[i].Ch)
		}
	}
}
