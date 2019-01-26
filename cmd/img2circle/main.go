package main

import (
	"flag"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/po3rin/img2circle"
)

var imgPath = flag.String("f", "", "path to the image")
var output = flag.String("o", "cropped.png", "path to the output image")

func main() {
	flag.Parse()
	if *imgPath == "" {
		log.Fatal("path flag is required")
	}

	img, err := os.Open(*imgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()

	src, _, err := image.Decode(img)
	if err != nil {
		log.Fatal(err)
	}

	c, err := img2circle.NewCroper(img2circle.Params{Src: src})
	if err != nil {
		log.Fatal(err)
	}

	result := c.CropCircle()
	file, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = png.Encode(file, result)
	if err != nil {
		log.Fatal(err)
	}
}
