package main

import "github.com/PuerkitoBio/goquery"
import "strings"
import "fmt"

func main() {
	doc, err := goquery.NewDocument("https://news.cnblogs.com")
	if err != nil {
		fmt.Printf("query error: %v", err)
	}

	doc.Find(".news_block .content .news_entry").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		id, _ := s.Find("a").Attr("href")

		id = strings.Replace(id, "/", "", -1)
		id = strings.Replace(id, "n", "", 1)

		fmt.Printf("%s %s \n\n", title, id)
	})
}
