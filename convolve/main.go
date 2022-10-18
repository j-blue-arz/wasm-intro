//go:build wasm

package main

import (
	"fmt"
	"syscall/js"
)

func convolve(this js.Value, args []js.Value) interface{} {
	inputBuffer := make([]byte, args[0].Get("length").Int())
	js.CopyBytesToGo(inputBuffer, args[0])
	width := args[1].Int()
	height := args[2].Int()
	size := width * height
	red := inputBuffer[0:size]
	green := inputBuffer[size : 2*size]
	blue := inputBuffer[2*size : 3*size]
	alpha := inputBuffer[3*size:]

	image := rgbaImage{red, green, blue, alpha, width, height}
	kernel := kernel3{
		1.0 / 8.0, 0, -1.0 / 8.0,
		2.0 / 8.0, 0, -2.0 / 8.0,
		1.0 / 8.0, 0, -1.0 / 8.0,
	}

	resultImage := convolveRGBA(image, kernel)

	var outputBuffer []byte
	outputBuffer = append(outputBuffer, resultImage.red...)
	outputBuffer = append(outputBuffer, resultImage.green...)
	outputBuffer = append(outputBuffer, resultImage.blue...)
	outputBuffer = append(outputBuffer, resultImage.alpha...)

	size = resultImage.width * resultImage.height

	result := js.Global().Get("Uint8ClampedArray").New(size * 4)
	fmt.Println(size * 4)
	fmt.Println(len(outputBuffer))
	js.CopyBytesToJS(result, outputBuffer)

	return result
}

func main() {
	js.Global().Set("convolve", js.FuncOf(convolve))

	<-make(chan bool)
}
