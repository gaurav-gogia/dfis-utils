package main

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
)

func create(original, stegno, ziparch, data1, data2 string) {
	prep(data1, data2, ziparch)

	org, err := os.Open(original)
	handerr(err)
	defer org.Close()

	zipped, err := os.Open(ziparch)
	handerr(err)
	defer zipped.Close()

	mut, err := os.Create(stegno)
	handerr(err)
	defer mut.Close()

	_, err = io.CopyBuffer(mut, org, nil)
	handerr(err)

	_, err = io.CopyBuffer(mut, zipped, nil)
	handerr(err)
}

func prep(data1, data2, ziparch string) {
	// read files
	one, err := ioutil.ReadFile(data1)
	handerr(err)

	two, err := ioutil.ReadFile(data2)
	handerr(err)

	// get zip file creator
	out, err := os.Create(ziparch)
	handerr(err)
	defer out.Close()

	w := zip.NewWriter(out)
	defer w.Close()

	// add files
	f, err := w.Create("./data/t1.txt")
	handerr(err)
	_, err = f.Write(one)
	handerr(err)

	f, err = w.Create("./data/t2.txt")
	handerr(err)
	_, err = f.Write(two)
	handerr(err)
}
