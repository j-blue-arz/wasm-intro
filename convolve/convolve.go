package main

type grayImage struct {
	buffer []uint8
	width  int
	height int
}

// kernel is expected to be normalized
type kernel3 [9]float32

func (image grayImage) get(row, col int) uint8 {
	return image.buffer[row*image.width+col]
}

func (image grayImage) set(row, col int, value uint8) {
	image.buffer[row*image.width+col] = value
}

func (k kernel3) get(row, col int) float32 {
	return k[row*3+col]
}

// The input and output image represent a row-wise image.
// A pixel in row x and column y
// will be located at x*width + y with regard to the image width
//
// The returned image has its size reduced by 2 in both directions.
func convolveGray(img grayImage, k kernel3) (grayImage, error) {

	width := img.width - 2
	height := img.height - 2
	result := grayImage{make([]uint8, width*height), width, height}
	for row := 1; row < img.height-1; row++ {
		for col := 1; col < img.width-1; col++ {
			value := convolveGrayPixel(img, k, row, col)
			result.set(row-1, col-1, uint8(value))
		}
	}
	return result, nil
}

func convolveGrayPixel(img grayImage, kern kernel3, row, col int) float32 {
	var value float32
	for x, kx := col-1, 2; x <= col+1; x, kx = x+1, kx-1 {
		for y, ky := row-1, 2; y <= row+1; y, ky = y+1, ky-1 {
			value += float32(img.get(y, x)) * kern.get(ky, kx)
		}
	}
	return value
}

// The input and output image have the same format as JS's ImageData object:
// rgb image is provided by four separate channels for
// Each color component is represented by an integer between 0 and 255.
// Each component is assigned a consecutive index within the array
//
// The returned image has size width-1, height-1
func convolveRGBA() {

}