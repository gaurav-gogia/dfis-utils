package main

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
)

func create() {
	prep()

	org, err := os.Open("./noise.png")
	handerr(err)
	defer org.Close()

	zipped, err := os.Open("./data/arch.zip")
	handerr(err)
	defer zipped.Close()

	mut, err := os.Create("steggy.png")
	handerr(err)
	defer mut.Close()

	_, err = io.CopyBuffer(mut, org, nil)
	handerr(err)

	_, err = io.CopyBuffer(mut, zipped, nil)
	handerr(err)
}

func prep() {
	// read files
	one, err := ioutil.ReadFile("./data/t1.txt")
	handerr(err)

	two, err := ioutil.ReadFile("./data/t2.txt")
	handerr(err)

	// get zip file creator
	out, err := os.Create("./data/arch.zip")
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
