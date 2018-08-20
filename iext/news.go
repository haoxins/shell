package main

import (
	"encoding/json"
	"fmt"
)

func getNews() {
	type News struct {
		Datetime string
		Headline string
		URL      string
	}
	for _, v := range cList {
		_, body, _ := request.Get(host + "/stock/" + v + "/news/last/6").End()
		var news []News
		_ = json.Unmarshal([]byte(body), &news)
		for _, v := range news {
			fmt.Println(v.Headline)
			fmt.Println(v.URL)
			fmt.Println()
		}
		fmt.Print("\n\n\n")
	}
}
