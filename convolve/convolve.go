package main

type color int

const (
	red   color = 0
	green color = 1
	blue  color = 2
)

// all color and alpha
const numChannels = 4

var colorChannels = []color{red, green, blue}

// Representation of a ImageData object as retrieved from a JS canvas
// the buffer has size width*height*4.
// the values for r, g, b, and a are interleaved int the buffer.
// That is, the blue value of the pixel at row 20 and column 100 is at 20*width+100+2
type canvasImage struct {
	buffer []byte
	width  int
	height int
}

// kernel is expected to be normalized
type kernel3 [9]float32

func (image canvasImage) index(row, col int, c color) int {
	return (row*image.width+col)*numChannels + int(c)
}

func (image canvasImage) get(row, col int, c color) byte {
	return image.buffer[image.index(row, col, c)]
}

func (image canvasImage) set(row, col int, c color, value byte) {
	image.buffer[image.index(row, col, c)] = value
}

func (k kernel3) get(row, col int) float32 {
	return k[row*3+col]
}

// The returned image has its size reduced by 2 in both directions.
func convolveImage(img canvasImage, k kernel3) canvasImage {
	width := img.width - 2
	height := img.height - 2
	result := canvasImage{make([]byte, width*height*4), width, height}
	for row := 1; row < img.height-1; row++ {
		for col := 1; col < img.width-1; col++ {
			for _, color := range colorChannels {
				value := convolvePixel(img, k, row, col, color)
				result.set(row-1, col-1, color, byte(value))
			}
		}
	}
	return result
}

func convolvePixel(img canvasImage, kern kernel3, row, col int, c color) float32 {
	var value float32
	for x, kx := col-1, 2; x <= col+1; x, kx = x+1, kx-1 {
		for y, ky := row-1, 2; y <= row+1; y, ky = y+1, ky-1 {
			value += float32(img.get(y, x, c)) * kern.get(ky, kx)
		}
	}
	return value
}
