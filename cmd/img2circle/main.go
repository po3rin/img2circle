package main

import (
	"flag"
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
	c, err := img2circle.NewCroper(img2circle.Params{
		ImgPath: *imgPath,
	})
	if err != nil {
		log.Fatal(err)
	}
	result := c.CropCircle()
	file, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_ = png.Encode(file, result)
	if err != nil {
		log.Fatal(err)
	}
}
