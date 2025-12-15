package imgproc

import (
	"image"
	"image/color"
)

// converts an image to black and white using the simple average method
// (r+g+b)/3
//
// imgIn is passed as a pointer for speed, but is not modified.
func BWSimple(imgIn *image.RGBA) *image.RGBA {
	imgOut := image.NewRGBA(imgIn.Bounds())

	ww := imgIn.Bounds().Dx()
	hh := imgIn.Bounds().Dy()

	for y := range hh {
		for x := range ww {
			r, g, b, _ := imgIn.At(x, y).RGBA()
			rf := float64(uint8(r >> 8))
			bf := float64(uint8(b >> 8))
			gf := float64(uint8(g >> 8))
			grayf := (rf + gf + bf) / 3.0 // float64 simple average

			gray := uint8(grayf)

			imgOut.Set(x, y, color.Gray{gray})
		}
	}
	return imgOut

}
