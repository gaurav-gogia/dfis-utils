package webutils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Setcook(url string) {
	var client http.Client

	req, err := http.NewRequest("GET", url, nil)
	handle(err)

	req.Header.Set("Cookie", "session_id=something")
	res, err := client.Do(req)
	handle(err)

	data, err := ioutil.ReadAll(res.Body)
	handle(err)
	fmt.Printf("%s\n", data)
}
