package imgproc

import (
	"image"
	"image/color"
)

// dither an image using only two states. Produces ugly images.
// value <= 127 = 0, values > 127 = 255
func DitherSimple(imgIn *image.RGBA) *image.RGBA {

	//imgOut := image.NewRGBA(imgIn.Bounds())

	ww := imgIn.Bounds().Dx()
	hh := imgIn.Bounds().Dy()

	var gray uint8
	for y := range hh {
		for x := range ww {
			r, _, _, _ := imgIn.At(x, y).RGBA()
			rv := uint8(r >> 8)

			if rv > 127 {
				gray = 255
			} else {
				gray = 0
			}

			imgIn.Set(x, y, color.Gray{gray})

		}
	}

	return imgIn

}
