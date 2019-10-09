package fileutils

import (
	"crypto/rand"
	"os"

	"golang.org/x/sys/unix"
)

func Del(fpath string) {
	info, err := os.Stat(fpath)
	if err != nil {
		panic(err)
	}

	writerandom(info.Size(), fpath)
	os.Remove(fpath)
}

func getrandom(size int64) []byte {
	data := make([]byte, size)
	rand.Read(data)
	return data
}

func writerandom(size int64, fpath string) {
	fd, err := unix.Open(fpath, unix.O_WRONLY, 0777)
	defer unix.Close(fd)
	if err != nil {
		panic(err)
	}

	for i := int64(0); i <= size; i += BUFSIZE {
		var data []byte

		if size-i <= BUFSIZE {
			data = getrandom(size - i)
		} else {
			data = getrandom(BUFSIZE)
		}

		unix.Write(fd, data)
	}
}
