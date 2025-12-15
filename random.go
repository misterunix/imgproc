package imgproc

import (
	"image"

	"math/rand/v2"
)

// dither an image using random probability. Produces ugly images.
func DitherRandom(imgIn *image.RGBA) *image.RGBA {

	ww := imgIn.Bounds().Dx()
	hh := imgIn.Bounds().Dy()

	var gray uint8
	for y := range hh {
		for x := range ww {
			r, _, _, _ := imgIn.At(x, y).RGBA()
			rn := uint8(rand.IntN(255))
			rv := uint8(r >> 8)

			if rv > rn {
				gray = 255
			} else {
				gray = 0
			}
			SetPixel(imgIn, x, y, int(gray))
			//imgIn.Set(x, y, color.Gray{gray})
		}
	}

	return imgIn

}
