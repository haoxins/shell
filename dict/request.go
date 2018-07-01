package main

import "github.com/PuerkitoBio/goquery"
import "strings"
import "net/url"
import "fmt"

func query(word string) {
	// for now, only Bing
	query, err := url.ParseQuery(word)
	if err != nil {
		fmt.Printf("invalid word: %s, error: %v", word, err)
	}

	uri := strings.Replace(config.dicts[0].url, "{word}", query.Encode(), 1)
	doc, err := goquery.NewDocument(uri)
	if err != nil {
		fmt.Printf("query error: %v", err)
	}
	// phonetic
	doc.Find(".hd_prUS").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
	// word
	doc.Find(".def").Each(func(i int, s *goquery.Selection) {
		pos := s.Parent().Find(".pos").Text()
		def := s.Find("a, span").Text()

		if pos != "" {
			fmt.Printf("%s - %s \n", pos, def)
		} else {
			fmt.Println(def)
		}
	})

	// sentence
	doc.Find(".p1-11").Each(func(i int, s *goquery.Selection) {
		sen := s.Find("a, span").Text()
		fmt.Println(sen)
	})
}
