package main

import "github.com/parnurzeal/gorequest"

const host = "https://api.iextrading.com/1.0"
const zhHost = "https://hq.sinajs.cn"

var request = gorequest.New()

var cList = []string{"nvda", "baba", "jd"}

var zhList = []string{"sz000725"}
