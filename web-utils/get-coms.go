package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func getcoms(url string) {
	res, err := http.Get(url)
	handle(err)

	body, err := ioutil.ReadAll(res.Body)
	handle(err)

	re := regexp.MustCompile("<!--(.|\n)*?-->")
	matches := re.FindAllString(string(body), -1)

	if matches == nil {
		fmt.Println("No HTML comments found")
		return
	}

	for _, match := range matches {
		fmt.Println(match)
	}
}
