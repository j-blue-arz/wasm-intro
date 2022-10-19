//go:build wasm

package main

import (
	"syscall/js"
)

func convolve(this js.Value, args []js.Value) interface{} {
	inputBuffer := make([]byte, args[0].Get("length").Int())
	js.CopyBytesToGo(inputBuffer, args[0])
	width := args[1].Int()
	height := args[2].Int()

	image := canvasImage{inputBuffer, 4, width, height}

	resultImage := convolveImage(image)

	size := resultImage.width * resultImage.height

	result := js.Global().Get("Uint8ClampedArray").New(size * 4)
	js.CopyBytesToJS(result, resultImage.buffer)

	return result
}

func main() {
	js.Global().Set("convolve", js.FuncOf(convolve))

	<-make(chan bool)
}
