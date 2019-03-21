package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func keysearch(url, keyword string) {
	var client http.Client
	client.Timeout = 30 * time.Second

	res, err := client.Get(url)
	handle(err)

	body, err := ioutil.ReadAll(res.Body)
	handle(err)

	if strings.Contains(string(body), keyword) {
		fmt.Println("Match found for: " + keyword + " in url: " + url)
	} else {
		fmt.Println("Not found!")
	}
}
