//go:build wasm

package main

import (
	"fmt"
	"syscall/js"
)

func setImage(this js.Value, args []js.Value) interface{} {
	var img []byte
	js.CopyBytesToGo(img, args[0])
	width := args[1].Int()
	height := args[2].Int()
	return true
}

func main() {
	export := make(map[string]interface{})
	export["setImage"] = js.FuncOf(setImage)
	js.Global().Set("convolve", export)

	<-make(chan bool)
}
