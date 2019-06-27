// +build js,wasm

package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"reflect"
	"syscall/js"
	"time"
	"unsafe"

	"github.com/po3rin/img2circle"
)

type GifAnimeCreator struct {
	inBuf                  []uint8
	outBuf                 bytes.Buffer
	onImgLoadCb, initMemCb js.Func
	sourceImg              image.Image

	console js.Value
	done    chan struct{}
}

func New() *GifAnimeCreator {
	return &GifAnimeCreator{
		console: js.Global().Get("console"),
		done:    make(chan struct{}),
	}
}

func (g *GifAnimeCreator) Start() {
	// Setup functions
	g.setupInitMemCb()
	js.Global().Set("initMem", g.initMemCb)

	g.setupOnImgLoadCb()
	js.Global().Set("loadImage", g.onImgLoadCb)

	<-g.done
	g.log("Shutting down app")
	g.onImgLoadCb.Release()
}

func (g *GifAnimeCreator) ConvertImage(argStartFlag string, argEndFlag string, argLoopFlag string) {
	// sourceImg is already decoded

	// Set image
	cropper, err := img2circle.NewCropper(img2circle.Params{Src: g.sourceImg})
	if err != nil {
		log.Fatal(err)
	}
	result := cropper.CropCircle()

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, result); err != nil {
		fmt.Println("error:png\n", err)
		return
	}

	g.outBuf = *buf
}

// updateImage writes the image to a byte buffer and then converts it to base64.
// Then it sets the value to the src attribute of the target image.
func (g *GifAnimeCreator) updateImage(start time.Time) {
	g.console.Call("log", "updateImage:", start.String())
	g.ConvertImage("left", "right", "false")
	out := g.outBuf.Bytes()
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&out))
	ptr := uintptr(unsafe.Pointer(hdr.Data))
	// set pointer and length to JS function
	// 画像がセットされたポインターと画像の長さをJSの関数に渡す
	js.Global().Call("displayImage", ptr, len(out))
	g.console.Call("log", "time taken:", time.Now().Sub(start).String())
	g.outBuf.Reset()
}

// utility function to log a msg to the UI from inside a callback
func (s *GifAnimeCreator) log(msg string) {
	js.Global().Get("document").
		Call("getElementById", "status").
		Set("innerText", msg)
}
