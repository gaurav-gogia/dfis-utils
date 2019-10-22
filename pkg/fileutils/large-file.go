package fileutils

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type filenode struct {
	fpath string
	fname string
	ftime time.Time
	fsize int64
}

func Large(root string, max int) error {
	var files []filenode

	err := filepath.Walk(root, func(filepath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		files = append(files, filenode{fpath: filepath, fname: info.Name(), fsize: info.Size()})

		return nil
	})
	if err != nil {
		return err
	}

	sort.SliceStable(files, func(i, j int) bool {
		return files[i].fsize > files[j].fsize
	})

	fmt.Println()
	for i, file := range files {
		fmt.Printf("File %d: %s, %db\n", i+1, file.fname, file.fsize)
		max--

		if max == 0 {
			break
		}
	}

	return nil
}
