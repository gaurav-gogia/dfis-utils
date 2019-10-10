package fileutils

import (
	"fmt"
	"os"
)

func Info(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	fmt.Println("\nName: ", info.Name())
	fmt.Println("Size: ", info.Size())
	fmt.Println("Mode: ", info.Mode())
	fmt.Println("Modified Time: ", info.ModTime())
	fmt.Println("Is Directory: ", info.IsDir())
	return nil
}
