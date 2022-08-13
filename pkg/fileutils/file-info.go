package fileutils

import (
	"dfis-utils/pkg/helper"
	"fmt"
	"os"
)

func Info(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	size := helper.HumanizeSize(info.Size())
	fmt.Println("\nName: ", info.Name())
	fmt.Println("Size: ", size)
	fmt.Println("Mode: ", info.Mode())
	fmt.Println("Modified Time: ", info.ModTime())
	fmt.Println("Is Directory: ", info.IsDir())
	return nil
}
