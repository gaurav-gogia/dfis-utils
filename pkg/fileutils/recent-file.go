package fileutils

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func Recent(root string, max int) error {
	var files []filenode

	err := filepath.Walk(root, func(filepath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		files = append(files, filenode{fpath: filepath, fname: info.Name(), ftime: info.ModTime()})

		return nil
	})
	if err != nil {
		return err
	}

	sort.SliceStable(files, func(i, j int) bool {
		return files[i].ftime.UnixNano() > files[j].ftime.UnixNano()
	})

	fmt.Println()
	for i, file := range files {
		fmt.Printf("File %d: %s, %v\n", i+1, file.fname, file.ftime)

		max--
		if max == 0 {
			break
		}
	}

	return nil
}
