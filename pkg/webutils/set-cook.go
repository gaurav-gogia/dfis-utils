package webutils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Setcook(url, key, value, method string) error {
	var client http.Client

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set(key, value)
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", data)
	return nil
}
