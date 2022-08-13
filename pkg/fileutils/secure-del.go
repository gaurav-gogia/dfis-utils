package fileutils

import (
	"crypto/rand"
	"dfis-utils/pkg/helper"
	"encoding/base32"
	"io/fs"
	"os"
	"path/filepath"
)

func Del(passes int, fpath string) error {
	info, err := os.Stat(fpath)
	if err != nil {
		return err
	}

	for i := 0; i < passes; i++ {
		err = writerandom(info.Size(), fpath)
		if err != nil {
			return err
		}
	}

	name := base32.StdEncoding.EncodeToString(getrandom(20))
	oldDir := filepath.Dir(fpath)
	newPath := filepath.Join(oldDir, name)
	err = os.Rename(fpath, newPath)
	if err != nil {
		return err
	}

	return os.Remove(newPath)
}

func getrandom(size int64) []byte {
	data := make([]byte, size)
	rand.Read(data)
	return data
}

func writerandom(size int64, fpath string) error {
	fd, err := os.OpenFile(fpath, os.O_WRONLY, fs.ModePerm)
	if err != nil {
		return err
	}
	defer fd.Close()

	done := make(chan bool)
	var active int
	for i := int64(0); i <= size; i += int64(512) {
		active++

		bufferSize := int64(512)
		if size-i < 512 {
			bufferSize = size - i
		}

		go writeOffset(fd, i, bufferSize, done)
		if active > helper.MAXGOROUTINES {
			<-done
			active--
		}
	}

	for active > 0 {
		<-done
		active--
	}

	return nil
}

func writeOffset(fd *os.File, offset, bufferSize int64, done chan bool) {
	data := getrandom(bufferSize)
	fd.WriteAt(data, offset)
	done <- true
}
