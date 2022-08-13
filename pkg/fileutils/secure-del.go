package fileutils

import (
	"crypto/rand"
	"io/fs"
	"os"
)

func Del(fpath string) error {
	_, err := os.Stat(fpath)
	if err != nil {
		panic(err)
	}

	if err := writerandom(fpath); err != nil {
		return err
	}

	return os.Remove(fpath)
}

func getrandom(size int64) []byte {
	data := make([]byte, size)
	rand.Read(data)
	return data
}

func writerandom(fpath string) error {
	fd, err := os.OpenFile(fpath, os.O_WRONLY, fs.ModePerm)
	if err != nil {
		return err
	}
	defer fd.Close()

	for i := int64(0); i <= 10; i++ {
		var data []byte

		data = getrandom(500)

		if _, err := fd.Write(data); err != nil {
			return err
		}
	}

	return nil
}
