package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) > 1 {
		id := os.Args[1]
		doc, err := goquery.NewDocument("https://news.ycombinator.com/item?id=" + id)
		if err != nil {
			fmt.Printf("query error: %v", err)
		}

		doc.Find(".comment-tree .athing.comtr").Each(func(i int, s *goquery.Selection) {
			text := s.Find(".default .comment .c00").Text()
			fmt.Printf("%s\n", text)
		})

		return
	}

	doc, err := goquery.NewDocument("https://news.ycombinator.com/")
	if err != nil {
		fmt.Printf("query error: %v", err)
	}

	doc.Find("tr.athing").Each(func(i int, s *goquery.Selection) {
		id, _ := s.Attr("id")
		title := s.Find("td.title a").Text()
		href, _ := s.Find("td.title a").Attr("href")

		fmt.Printf("%s  %s \n\n%s \n\n\n", title, id, href)
	})
}
