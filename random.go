package imgproc

import (
	"image"
	"image/draw"

	"math/rand/v2"
)

// dither an image using random probability. Produces ugly images.
//
// imgIn is passed as a pointer for speed, but is not modified.
func DitherRandom(imgIn *image.RGBA) *image.RGBA {

	// use draw to create, copy the input image to a new image
	// and then modify that image in place
	imgMod := image.NewRGBA(imgIn.Bounds())
	draw.Draw(imgMod, imgMod.Bounds(), imgIn, imgIn.Bounds().Min, draw.Src)

	//imgMod := image.NewRGBA(imgIn.Bounds())

	ww := imgMod.Bounds().Dx()
	hh := imgMod.Bounds().Dy()

	var gray uint8
	for y := range hh {
		for x := range ww {
			r, _, _, _ := imgMod.At(x, y).RGBA()
			rn := uint8(rand.IntN(255))
			rv := uint8(r >> 8)

			if rv > rn {
				gray = 255
			} else {
				gray = 0
			}
			SetPixel(imgMod, x, y, int(gray))
			//imgIn.Set(x, y, color.Gray{gray})
		}
	}

	return imgMod

}
