package fileutils

import (
	"crypto/rand"
	"dfis-utils/pkg/helper"
	"encoding/base32"
	"io/fs"
	"os"
	"path/filepath"
)

func Del(passes int, bufferSize int64, fpath string) error {
	info, err := os.Stat(fpath)
	if err != nil {
		return err
	}

	for i := 0; i < passes; i++ {
		err = writerandom(info.Size(), bufferSize, fpath)
		if err != nil {
			return err
		}
	}

	newPath := getNewPath(fpath)
	err = os.Rename(fpath, newPath)
	if err != nil {
		return err
	}

	return os.Remove(newPath)
}

func getNewPath(oldPath string) string {
	oldDir := filepath.Dir(oldPath)
	oldFileName := filepath.Base(oldPath)
	oldNameLen := int64(len(oldFileName))
	randomName := getrandom(oldNameLen)
	name := base32.StdEncoding.EncodeToString(randomName)
	return filepath.Join(oldDir, name)
}

func getrandom(size int64) []byte {
	data := make([]byte, size)
	rand.Read(data)
	return data
}

func writerandom(size, bufferSize int64, fpath string) error {
	fd, err := os.OpenFile(fpath, os.O_WRONLY, fs.ModePerm)
	if err != nil {
		return err
	}
	defer fd.Close()

	done := make(chan bool)
	var active int
	for i := int64(0); i <= size; i += int64(512) {
		active++

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
