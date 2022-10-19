package main

import (
	"math"
)

type color int

// Representation of a ImageData object as retrieved from a JS canvas
// the buffer has size width*height*4.
// the values for r, g, b, and a are interleaved int the buffer.
// That is, the blue value of the pixel at row 20 and column 100 is at 20*width+100+2
type canvasImage struct {
	buffer      []byte
	numChannels int
	width       int
	height      int
}

func makeGrayImage(width, height int) canvasImage {
	return canvasImage{make([]byte, width*height), 1, width, height}
}

func makeRGBAImage(width, height int) canvasImage {
	return canvasImage{make([]byte, width*height*4), 4, width, height}
}

func (image canvasImage) index(row, col int, c color) int {
	return (row*image.width+col)*image.numChannels + int(c)
}

func (image canvasImage) get(row, col int, c color) byte {
	return image.buffer[image.index(row, col, c)]
}

func (image canvasImage) set(row, col int, c color, value byte) {
	image.buffer[image.index(row, col, c)] = value
}

type kernel [9]float64

func (k kernel) get(row, col int) float64 {
	return k[row*3+col]
}

var kernel_x = kernel{
	1.0 / 8.0, 0, -1.0 / 8.0,
	2.0 / 8.0, 0, -2.0 / 8.0,
	1.0 / 8.0, 0, -1.0 / 8.0,
}

var kernel_y = kernel{
	1.0 / 8.0, 2.0 / 8.0, 1.0 / 8.0,
	0, 0, 0,
	-1.0 / 8.0, -2.0 / 8.0, -1.0 / 8.0,
}

func toGray(red, green, blue byte) byte {
	return (red + green + blue) / 3
}

// The returned image has its size reduced by 2 in both directions.
func convolveImage(img canvasImage) canvasImage {
	grayImage := makeGrayImage(img.width, img.height)
	for pixel := 0; pixel < len(grayImage.buffer); pixel++ {
		colors := img.buffer[pixel*4 : pixel*4+3]
		red, green, blue := colors[0], colors[1], colors[2]
		grayImage.buffer[pixel] = toGray(red, green, blue)
	}
	width := img.width - 2
	height := img.height - 2
	convolved := makeGrayImage(width, height)
	min, max := byte(255), byte(0)
	for row := 1; row < img.height-1; row++ {
		for col := 1; col < img.width-1; col++ {
			value_x := convolvePixel(img, kernel_x, row, col)
			value_y := convolvePixel(img, kernel_y, row, col)
			value := byte(math.Sqrt(value_x*value_x + value_y*value_y))
			if min > value {
				min = value
			}
			if max < value {
				max = value
			}
			convolved.set(row-1, col-1, 0, byte(value))
		}
	}

	result := makeRGBAImage(width, height)
	for pixel, value := range convolved.buffer {
		outValue := byte(float64(value-min) / float64(max-min) * 255)
		result.buffer[pixel*4] = outValue
		result.buffer[pixel*4+1] = outValue
		result.buffer[pixel*4+2] = outValue
		result.buffer[pixel*4+3] = 255
	}
	return result
}

func convolvePixel(img canvasImage, kernel kernel, row, col int) float64 {
	var value float64
	for x, kx := col-1, 2; x <= col+1; x, kx = x+1, kx-1 {
		for y, ky := row-1, 2; y <= row+1; y, ky = y+1, ky-1 {
			value += float64(img.get(y, x, 0)) * kernel.get(ky, kx)
		}
	}
	return value
}
