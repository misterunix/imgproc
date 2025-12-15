package imgproc

import (
	"image"
)

// Helper function to calculate the error diffusion center around the pixel
func errorDiffustion(img *image.RGBA, errorTerm, errorMul float64, x, y int) {
	SetPixel(img, x, y, GetPixel(img, x, y)+int(errorMul*errorTerm))
}
