package main

import (
	"fmt"
	"log"
	"os"
)

func info() {
	info, err := os.Stat(FILEPATH)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("\nName: ", info.Name())
	fmt.Println("Size: ", info.Size())
	fmt.Println("Mode: ", info.Mode())
	fmt.Println("Modified Time: ", info.ModTime())
	fmt.Println("Is Directory: ", info.IsDir())
}
