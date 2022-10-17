//go:build wasm

package main

import (
	"syscall/js"
)

func convolve(this js.Value, args []js.Value) interface{} {
	var img []byte
	js.CopyBytesToGo(img, args[0])
	width := args[1].Int()
	height := args[2].Int()
	size := width * height
	red := img[0:size]
	green := img[size : 2*size]
	blue := img[2*size : 3*size]
	alpha := img[3*size:]

	rgbaImage := MakeRGBAImage(red, green, blue, alpha, width, height)
	kernel := Kernel3{
		1.0 / 8.0, 0, -1.0 / 8.0,
		2.0 / 8.0, 0, -2.0 / 8.0,
		1.0 / 8.0, 0, -1.0 / 8.0,
	}

	return ConvolveRGBA(*rgbaImage, kernel)
}

func main() {
	js.Global().Set("convolve", convolve)

	<-make(chan bool)
}
