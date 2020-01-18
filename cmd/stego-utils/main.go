package main

import (
	"fmt"

	"dfis-utils/pkg/stegoutils"
)

func main() {
	fmt.Println("Generating image with random noise")
	stegoutils.Noise("noise.png")

	fmt.Println("Creating a stengnographic archive")
	stegoutils.Create("./noise.png", "./data/arch.zip", "steggy.png", "./data/t1.txt", "./data/t2.txt")

	fmt.Println("Detecting a stegnographic archive")
	stegoutils.Detect("./steggy.png")
}
