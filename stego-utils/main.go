package main

import "fmt"

func main() {
	fmt.Println("Generating image with random noise")
	noise("noise.png")

	fmt.Println("Creating a stengnographic archive")
	create("./noise.png", "./data/arch.zip", "steggy.png", "./data/t1.txt", "./data/t2.txt")

	fmt.Println("Detecting a stegnographic archive")
	detect("./steggy.png")

	fmt.Println("Done!")
}

func handerr(err error) {
	if err != nil {
		panic(err)
	}
}
