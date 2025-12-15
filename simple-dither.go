package imgproc

import (
	"image"
	"image/color"
	"image/draw"
)

// dither an image using only two states. Produces ugly images.
// value <= 127 = 0, values > 127 = 255
//
// imgIn is passed as a pointer for speed, but is not modified.
func DitherSimple(imgIn *image.RGBA) *image.RGBA {

	// use draw to create, copy the input image to a new image
	// and then modify that image in place
	imgMod := image.NewRGBA(imgIn.Bounds())
	draw.Draw(imgMod, imgMod.Bounds(), imgIn, imgIn.Bounds().Min, draw.Src)

	//imgOut := image.NewRGBA(imgIn.Bounds())

	//	imgMod := image.NewRGBA(imgIn.Bounds())

	ww := imgMod.Bounds().Dx()
	hh := imgMod.Bounds().Dy()

	var gray uint8
	for y := range hh {
		for x := range ww {
			r, _, _, _ := imgMod.At(x, y).RGBA()
			rv := uint8(r >> 8)

			if rv > 127 {
				gray = 255
			} else {
				gray = 0
			}

			imgMod.Set(x, y, color.Gray{gray})

		}
	}

	return imgMod

}
