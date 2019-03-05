package main

import (
	"image"
	"image/png"
	"log"
	"os"
	"strconv"

	L "./resize"
)

type imageSize struct {
	height int
	width  int
}

func saveImage(image image.Image, name string) {
	// open file to save
	dstFile, err := os.Create("./asset/" + name + ".png")
	if err != nil {
		log.Fatal(err)
	}
	// encode as .png to the file
	err = png.Encode(dstFile, image)
	// close the file
	dstFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	images := []imageSize{{height: 29, width: 29}, {height: 40, width: 40}}

	existingImageFile, err := os.Open("test.png")
	if err != nil {
		log.Fatal(err)
	}
	img, err := png.Decode(existingImageFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, image := range images {
		dstImg := L.Resize(img, image.height, image.width, L.Lanczos)
		saveImage(dstImg, "icon_"+strconv.Itoa(image.height))
	}

}
