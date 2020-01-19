package fileutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	filetype "github.com/h2non/filetype"
)

// File types
type FileType uint

const (
	FileUnknown FileType = iota
	FileImage
	FileArchive
	FileAudio
	FileVideo
)

// ParseFileType parses ftype and returns the corresponding
// FileType
func ParseFileType(ftype string) (FileType, error) {
	switch ftype {
	case "image":
		return FileImage, nil
	case "archive":
		return FileArchive, nil
	case "audio":
		return FileAudio, nil
	case "video":
		return FileVideo, nil
	default:
		return FileUnknown, fmt.Errorf("unknown file type '%s'", ftype)
	}
}

func Extract(root, dst string, ftype FileType) error {
	var count int64

	return filepath.Walk(root, func(filepath string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			buf, err := ioutil.ReadFile(filepath)
			if err != nil {
				return err
			}

			switch ftype {
			case FileImage:
				copyimage(&buf, &count, dst, info.Name())
			case FileArchive:
				copyarchive(&buf, &count, dst, info.Name())
			case FileAudio:
				copyaudio(&buf, &count, dst, info.Name())
			case FileVideo:
				copyvideo(&buf, &count, dst, info.Name())
			default:
				return fmt.Errorf("wrong file type")
			}
		}
		return nil
	})
}

func copyimage(buf *[]byte, count *int64, dst, name string) {
	if filetype.IsImage(*buf) {
		*count++
		ioutil.WriteFile(dst+name, *buf, 0644)
		fmt.Printf("\rImage File Found: %s, Count: %v", name, *count)
	}
}
func copyvideo(buf *[]byte, count *int64, dst, name string) {
	if filetype.IsVideo(*buf) {
		*count++
		ioutil.WriteFile(dst+name, *buf, 0644)
		fmt.Printf("\rVideo File Found: %s, Count: %v", name, *count)
	}
}
func copyaudio(buf *[]byte, count *int64, dst, name string) {
	if filetype.IsAudio(*buf) {
		*count++
		ioutil.WriteFile(dst+name, *buf, 0644)
		fmt.Printf("\rAudio File Found: %s, Count: %v", name, *count)
	}
}
func copyarchive(buf *[]byte, count *int64, dst, name string) {
	if filetype.IsArchive(*buf) {
		*count++
		ioutil.WriteFile(dst+name, *buf, 0644)
		fmt.Printf("\rArchive File Found: %s, Count: %v", name, *count)
	}
}
