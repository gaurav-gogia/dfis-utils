package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func large(root string) {
	var files []filenode

	filepath.Walk(root, func(filepath string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, filenode{fpath: filepath, fname: info.Name(), fsize: info.Size()})

		return nil
	})

	sort.SliceStable(files, func(i, j int) bool {
		return files[i].fsize > files[j].fsize
	})

	fmt.Println()
	if len(files) <= 10 {
		for i := 0; i < len(files); i++ {
			fmt.Printf("File %d: %s, %d\n", i, files[i].fname, files[i].fsize)
		}
		return
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("File %d: %s, %d\n", i, files[i].fname, files[i].fsize)
	}
}
