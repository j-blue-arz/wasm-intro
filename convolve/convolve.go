package main

type RGBAImage struct {
	red    []byte
	green  []byte
	blue   []byte
	alpha  []byte
	width  int
	height int
}

func MakeRGBAImage(red, green, blue, alpha []byte, width, height int) *RGBAImage {
	return &RGBAImage{red, green, blue, alpha, width, height}
}

type grayImage struct {
	buffer []byte
	width  int
	height int
}

// kernel is expected to be normalized
type Kernel3 [9]float32

func (image grayImage) get(row, col int) byte {
	return image.buffer[row*image.width+col]
}

func (image grayImage) set(row, col int, value byte) {
	image.buffer[row*image.width+col] = value
}

func (k Kernel3) get(row, col int) float32 {
	return k[row*3+col]
}

// The input and output image represent a row-wise image.
// A pixel in row x and column y
// will be located at x*width + y with regard to the image width
//
// The returned image has its size reduced by 2 in both directions.
func convolveGray(img grayImage, k Kernel3) grayImage {
	width := img.width - 2
	height := img.height - 2
	result := grayImage{make([]byte, width*height), width, height}
	for row := 1; row < img.height-1; row++ {
		for col := 1; col < img.width-1; col++ {
			value := convolveGrayPixel(img, k, row, col)
			result.set(row-1, col-1, byte(value))
		}
	}
	return result
}

func convolveGrayPixel(img grayImage, kern Kernel3, row, col int) float32 {
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
func ConvolveRGBA(img RGBAImage, k Kernel3) RGBAImage {
	red := convolveGray(grayImage{img.red, img.width, img.height}, k)
	green := convolveGray(grayImage{img.green, img.width, img.height}, k)
	blue := convolveGray(grayImage{img.blue, img.width, img.height}, k)
	alpha := convolveGray(grayImage{img.alpha, img.width, img.height}, k)

	return RGBAImage{red.buffer, green.buffer, blue.buffer, alpha.buffer, img.width, img.height}
}
