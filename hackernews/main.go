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

		getInfoOfAnArticle(article{title, url, author})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://news.ycombinator.com/")
}

type article struct {
	title  string
	url    string
	author string
}

func getInfoOfAnArticle(article article) {
	fmt.Println("Title: ", article.title)
	fmt.Println("Url: ", article.url)
	fmt.Println("Author: ", article.author)
	fmt.Println("-----------------")
}
