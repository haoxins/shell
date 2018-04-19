package main

import "github.com/parnurzeal/gorequest"

const host = "https://api.iextrading.com/1.0"

var request = gorequest.New()

var cList = []string{"nvda", "baba", "jd"}
