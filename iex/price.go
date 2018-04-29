package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func getPrice() {
	for _, v := range cList {
		_, body, _ := request.Get(host + "/stock/" + v + "/price").End()
		fmt.Printf("%s: %s\n", v, body)
	}

	for _, v := range zhList {
		_, body, _ := request.Get(zhHost + "/list=" + v).End()
		reader := transform.NewReader(bytes.NewReader([]byte(body)), simplifiedchinese.GBK.NewDecoder())
		d, e := ioutil.ReadAll(reader)
		if e != nil {
			panic(e.Error())
		}

		s := strings.Replace(string(d), `var hq_str_sz000725="`, "", 1)
		s = strings.Replace(s, `";`, "", 1)
		a := strings.Split(s, ",")

		fmt.Printf("%s: %s\n", a[0], a[3])
	}
}
