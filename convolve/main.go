//go:build wasm

package main

import (
	"fmt"
	"syscall/js"
)

func convolve(this js.Value, args []js.Value) interface{} {
	inputBuffer := make([]byte, args[0].Get("length").Int())
	js.CopyBytesToGo(inputBuffer, args[0])
	fmt.Println("len(inputBuffer): ", len(inputBuffer))
	width := args[1].Int()
	height := args[2].Int()

	image := canvasImage{inputBuffer, width, height}
	/* kernel := kernel3{
		1.0 / 8.0, 0, -1.0 / 8.0,
		2.0 / 8.0, 0, -2.0 / 8.0,
		1.0 / 8.0, 0, -1.0 / 8.0,
	} */

	kernel := kernel3{
		0.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 0.0, 0.0,
	}

	resultImage := convolveImage(image, kernel)

	size := resultImage.width * resultImage.height

	result := js.Global().Get("Uint8ClampedArray").New(size * 4)
	fmt.Println("size * 4: ", size*4)
	fmt.Println("len(outputBuffer): ", len(resultImage.buffer))
	js.CopyBytesToJS(result, resultImage.buffer)

	return result
}

func main() {
	js.Global().Set("convolve", js.FuncOf(convolve))

	<-make(chan bool)
}
