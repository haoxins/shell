package main

import "github.com/PuerkitoBio/goquery"
import "strings"
import "fmt"
import "os"

func main() {

	if len(os.Args) > 1 {
		id := os.Args[1]
		doc, err := goquery.NewDocument("https://news.cnblogs.com/n/" + id)
		if err != nil {
			fmt.Printf("query error: %v", err)
		}

		doc.Find("#news_content #news_body p").Each(func(i int, s *goquery.Selection) {
			p := s.Text()

			fmt.Printf("%s \n", p)
		})

		return
	}

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
