package fileutils

import (
	"crypto/rand"
	"os"

	"golang.org/x/sys/unix"
)

func Del(fpath string) error {
	info, err := os.Stat(fpath)
	if err != nil {
		panic(err)
	}

	if err := writerandom(info.Size(), fpath); err != nil {
		return err
	}

	return os.Remove(fpath)
}

func getrandom(size int64) []byte {
	data := make([]byte, size)
	rand.Read(data)
	return data
}

func writerandom(size int64, fpath string) error {
	fd, err := unix.Open(fpath, unix.O_WRONLY, 0777)
	defer unix.Close(fd)
	if err != nil {
		return err
	}

	for i := int64(0); i <= size; i += 500 {
		var data []byte

		if size-i <= 500 {
			data = getrandom(size - i)
		} else {
			data = getrandom(500)
		}

		if _, err := unix.Write(fd, data); err != nil {
			return err
		}
	}

	return nil
}
