package main

import "fmt"

func main() {
	fmt.Println("Generating image with random noise")
	//	noise()

	fmt.Println("Creating a stengnographic archive")
	//	create()

	fmt.Println("Detecting a stegnographic archive")
	detect()

	fmt.Println("Done!")
}

func handerr(err error) {
	if err != nil {
		panic(err)
	}
}
