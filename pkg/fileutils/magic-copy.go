package fileutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	filetype "gopkg.in/h2non/filetype.v0"
)

func Extract(root, dst string, in int) {
	var count int64

	filepath.Walk(root, func(filepath string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			buf, err := ioutil.ReadFile(filepath)

			switch in {
			case IMAGE:
				copyimage(&buf, &count, dst, info.Name())
			case ARCHIVE:
				copyarchive(&buf, &count, dst, info.Name())
			case AUDIO:
				copyaudio(&buf, &count, dst, info.Name())
			case VIDEO:
				copyvideo(&buf, &count, dst, info.Name())
			default:
				fmt.Println("Wrong Choice")
				return nil
			}

			return err
		}
		return nil
	})
}

func copyimage(buf *[]byte, count *int64, dst, name string) {
	if filetype.IsImage(*buf) {
		*count++
		ioutil.WriteFile(dst+name, *buf, 0644)
		fmt.Printf("\rImage File Found: %s, Count: %v", name, count)
	}
}
func copyvideo(buf *[]byte, count *int64, dst, name string) {
	if filetype.IsVideo(*buf) {
		*count++
		ioutil.WriteFile(dst+name, *buf, 0644)
		fmt.Printf("\rImage File Found: %s, Count: %v", name, count)
	}
}
func copyaudio(buf *[]byte, count *int64, dst, name string) {
	if filetype.IsAudio(*buf) {
		*count++
		ioutil.WriteFile(dst+name, *buf, 0644)
		fmt.Printf("\rImage File Found: %s, Count: %v", name, count)
	}
}
func copyarchive(buf *[]byte, count *int64, dst, name string) {
	if filetype.IsArchive(*buf) {
		*count++
		ioutil.WriteFile(dst+name, *buf, 0644)
		fmt.Printf("\rImage File Found: %s, Count: %v", name, count)
	}
}
