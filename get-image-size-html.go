package main

import (
	"bufio"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"

	"github.com/atotto/clipboard"
)

func main() {
	width, height := getImageDimension(os.Args[1])
	widthStr := strconv.Itoa(width)
	heightStr := strconv.Itoa(height)
	fmt.Println(textPtr)
	fmt.Println("width=", widthStr)
	fmt.Println("height=", heightStr)
	fmt.Println(len(os.Args), os.Args[1])
	fmt.Print("Press 'Enter' to continue...")
	clipboard.WriteAll(
		"width=\"" + widthStr + "\"" +
			" height=\"" + heightStr + "\"")

	bufio.NewReader(os.Stdin).ReadBytes('\n')

}

func getImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
}
