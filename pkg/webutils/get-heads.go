package webutils

import (
	"fmt"
	"net/http"
)

func Heads(url string) {
	res, err := http.Head(url)
	handle(err)

	for key, val := range res.Header {
		fmt.Printf("%s: %s\n", key, val)
	}
}
