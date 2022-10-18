package main

import "testing"

func TestConvolveGray(t *testing.T) {
	image := grayImage{
		buffer: []uint8{
			0, 0, 0, 0,
			0, 50, 50, 50,
			0, 50, 100, 100,
			0, 50, 100, 150,
		}, width: 4, height: 4,
	}

	k := kernel3{
		1.0 / 8.0, 0, -1.0 / 8.0,
		2.0 / 8.0, 0, -2.0 / 8.0,
		1.0 / 8.0, 0, -1.0 / 8.0,
	}

	result := convolveGray(image, k)

	if result.width != 2 {
		t.Fatalf("Expected 2x2 image, but width was %d", result.width)
	}
	if result.height != 2 {
		t.Fatalf("Expected 2x2 image, but height was %d", result.height)
	}
	if len(result.buffer) != 4 {
		t.Fatalf("Expected 2x2 image, but buffer size was was %d", len(result.buffer))
	}

	expected := grayImage{
		buffer: []uint8{
			25, 6,
			43, 25,
		}, width: 2, height: 2,
	}

	for row := 0; row < 2; row++ {
		for col := 0; col < 2; col++ {
			if result.get(row, col) != expected.get(row, col) {
				t.Fatalf("Pixel at (%d, %d) was expected as %d, but was %d",
					row, col, expected.get(row, col), result.get(row, col))
			}
		}
	}
}
