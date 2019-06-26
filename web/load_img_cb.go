// +build js,wasm

package main

import (
	"bytes"
	"image/png"
	"reflect"
	"syscall/js"
	"time"
	"unsafe"
)

func (g *GifAnimeCreator) setupOnImgLoadCb() {
	g.onImgLoadCb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		reader := bytes.NewReader(g.inBuf)
		var err error
		g.sourceImg, err = png.Decode(reader)
		if err != nil {
			g.log(err.Error())
			return nil
		}
		g.log("Ready for operations")
		start := time.Now()
		g.updateImage(start)
		return nil
	})
}

func (g *GifAnimeCreator) setupInitMemCb() {
	// The length of the image array buffer is passed.
	// Then the buf slice is initialized to that length.
	// And a pointer to that slice is passed back to the browser
	g.initMemCb = js.FuncOf(func(this js.Value, i []js.Value) interface{} {
		length := i[0].Int()
		g.console.Call("log", "length:", length)
		// make buffer by image length
		// 画像サイズ分のバッファを作成する
		g.inBuf = make([]uint8, length)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&g.inBuf))
		ptr := uintptr(unsafe.Pointer(hdr.Data))
		// pass pointer to JS by calling function
		// JS側の関数の引数にポインターをセットして渡す
		js.Global().Call("gotMem", ptr)
		return nil
	})
}
