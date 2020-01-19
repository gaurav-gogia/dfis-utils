package webutils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func Getcoms(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	re := regexp.MustCompile("<!--(.|\n)*?-->")
	matches := re.FindAllString(string(body), -1)

	if matches == nil {
		return errors.New("No HTML comments found")
	}

	for _, match := range matches {
		fmt.Println(match)
	}

	return nil
}
