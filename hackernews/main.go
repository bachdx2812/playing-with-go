package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("table.itemlist > tbody > tr.athing td:nth-of-type(3)", func(e *colly.HTMLElement) {
		title := e.ChildText("a:nth-of-type(1)")
		url := e.ChildAttr("a:nth-of-type(1)", "href")
		author := e.ChildText("span")

		fmt.Println("Title: ", title)
		fmt.Println("Url: ", url)
		fmt.Println("Author: ", author)
		fmt.Println("-----------------")
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://news.ycombinator.com/")
}
