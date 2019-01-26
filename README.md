# img2circle

<img src="https://img.shields.io/badge/go-v1.11-blue.svg"/> [![CircleCI](https://circleci.com/gh/po3rin/img2circle.svg?style=shield)](https://circleci.com/gh/po3rin/img2circle) <a href="https://codeclimate.com/github/po3rin/img2circle/maintainability"><img src="https://api.codeclimate.com/v1/badges/8c9276a15d62f99fccf0/maintainability" /></a> [![GolangCI](https://golangci.com/badges/github.com/po3rin/img2circle.svg)](https://golangci.com)

Package img2circle lets you generate an image cropped a circular image out of a rectangular.

<img src="src/cover.png">

## Installation

```
$ go get github.com/po3rinimg2circle/cmd/img2circle
```

## Usage

as CLI tool.

```bash
$ img2circle -f testdata/gopher.jpeg -o cropped.png
```

as Code.

```go
package main

import (
    _ "image/jpeg"
    "image/png"
    "os"

    "github.com/po3rin/img2circle"
)

func main(){
    c, _ := img2circle.NewCroper(img2circle.Params{
		ImgPath: "testdata/gopher.jpeg",
	})
    result := c.CropCircle()
    file, _ := os.Create(*output)
	defer file.Close()

	_ = png.Encode(file, result)
}
```
