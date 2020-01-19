package webutils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Keysearch(url, keyword string) error {
	var client http.Client
	client.Timeout = 30 * time.Second

	res, err := client.Get(url)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if strings.Contains(string(body), keyword) {
		fmt.Println("Match found for: " + keyword + " in url: " + url)
	} else {
		fmt.Println("Not found!")
	}

	return nil
}
