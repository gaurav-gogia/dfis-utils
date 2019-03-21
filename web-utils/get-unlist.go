package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func getunlist(url, wordfile string) {
	var thread int
	done := make(chan bool)

	file, err := os.Open(wordfile)
	defer file.Close()
	handle(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		go findurl(url, scanner.Text(), done)
		thread++

		if thread > 50 {
			<-done
			thread--
		}
	}

	for thread > 0 {
		<-done
		thread--
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Could not read file")
	}
}

func findurl(base, filepath string, done chan bool) {
	target, err := url.Parse(base)
	handle(err)

	target.Path = filepath

	res, err := http.Head(target.String())
	handle(err)

	if res.StatusCode == 200 {
		if err := down("./data/"+filepath, target.String()); err == nil {
			save("./foundfile", []string{filepath})
		}
	}
}

func down(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
