# img2circle

<img src="https://img.shields.io/badge/go-v1.11-blue.svg"/> [![CircleCI](https://circleci.com/gh/po3rin/img2circle.svg?style=shield)](https://circleci.com/gh/po3rin/img2circle) <a href="https://codeclimate.com/github/po3rin/img2circle/maintainability"><img src="https://api.codeclimate.com/v1/badges/8c9276a15d62f99fccf0/maintainability" /></a> [![GolangCI](https://golangci.com/badges/github.com/po3rin/img2circle.svg)](https://golangci.com)

Package img2circle lets you generate an image cropped a circular image out of a rectangular.

<img src="src/cover.png">

## Installation

```bash
$ go get github.com/po3rin/img2circle/cmd/img2circle
```

## Try this on Web

you enabled to try here !! (developed by Go + Wasm)

https://po3rin.github.io/img2circle/web/

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
    img, _ := os.Open(*imgPath)
    defer img.Close()
    src, _, _ := image.Decode(img)

    // use img2circle packege.
    c, _ := img2circle.NewCropper(img2circle.Params{Src: src})
    result := c.CropCircle()

    file, _ := os.Create("cropped.png")
    defer file.Close()
    _ = png.Encode(file, result)
}
```
