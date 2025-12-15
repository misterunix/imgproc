package imgproc

import (
	"image"
	"image/color"
	"math"
)

// converts an image to black and white using the BT.2020 standard
func BWBT2020(imgIn *image.RGBA) *image.RGBA {
	imgOut := image.NewRGBA(imgIn.Bounds())

	ww := imgIn.Bounds().Dx()
	hh := imgIn.Bounds().Dy()

	for y := range hh {
		for x := range ww {
			r, g, b, _ := imgIn.At(x, y).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8

			// Luma coefficients

			gg := (float32(r) * 0.2627) +
				(float32(g) * 0.6780) +
				(float32(b) * 0.0593)

			gray := uint8(gg)

			imgOut.Set(x, y, color.Gray{gray})
		}
	}

	return imgOut

}

// converts an image to black and white using the Luminance method with
// gamma correction and BT.2020
func BWLuminanceGammaBT2020(imgIn *image.RGBA) *image.RGBA {
	imgOut := image.NewRGBA(imgIn.Bounds())

	ww := imgIn.Bounds().Dx()
	hh := imgIn.Bounds().Dy()

	for y := range hh {
		for x := range ww {
			r, g, b, _ := imgIn.At(x, y).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8

			rL := float64(r) / 255.0
			gL := float64(g) / 255.0
			bL := float64(b) / 255.0

			rG := math.Pow(float64(rL), 2.2)
			gG := math.Pow(float64(gL), 2.2)
			bG := math.Pow(float64(bL), 2.2)

			rgN := rG * 255
			ggN := gG * 255
			bgN := bG * 255
			// Luminance coefficients
			gg := (float32(rgN) * 0.2627) +
				(float32(ggN) * 0.6780) +
				(float32(bgN) * 0.0593)

			// Gamma correction
			// gg = (gg / 255.0)
			// gg = (gg * gg * gg)

			gray := uint8(gg)

			imgOut.Set(x, y, color.Gray{gray})
		}
	}

	return imgOut
}
