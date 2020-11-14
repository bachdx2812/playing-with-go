package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("#hnmain table.itemlist tbody tr.athing td.title a.storylink", func(e *colly.HTMLElement) {
		fmt.Println("Title: ", e.Text)
		fmt.Println("URL: ", e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://news.ycombinator.com/")
}
