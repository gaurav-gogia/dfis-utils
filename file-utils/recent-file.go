package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func recent(root string) {
	var files []filenode

	filepath.Walk(root, func(filepath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		files = append(files, filenode{fpath: filepath, fname: info.Name(), ftime: info.ModTime()})

		return nil
	})

	sort.SliceStable(files, func(i, j int) bool {
		return files[i].ftime.UnixNano() > files[j].ftime.UnixNano()
	})

	fmt.Println()
	if len(files) <= 10 {
		for i := 0; i < len(files); i++ {
			fmt.Printf("File %d: %s, %v\n", i, files[i].fname, files[i].ftime)
		}
		return
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("File %d: %s, %v\n", i, files[i].fname, files[i].ftime)
	}
}
