package webutils

import (
	"fmt"
	"net/http"
)

func Heads(url string) error {
	res, err := http.Head(url)
	if err != nil {
		return err
	}

	for key, val := range res.Header {
		fmt.Printf("%s: %s\n", key, val)
	}

	return nil
}
