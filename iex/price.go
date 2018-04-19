package main

import (
	"fmt"
)

func getPrice() {
	for _, v := range cList {
		_, body, _ := request.Get(host + "/stock/" + v + "/price").End()
		fmt.Println(v, ":", body)
	}
}
