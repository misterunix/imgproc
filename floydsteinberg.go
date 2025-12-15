package imgproc

import (
	"image"
)

// False Floyd-Steinberg dithering implementation
func DitherFalseFloydSteinberg(imgIn *image.RGBA) *image.RGBA {

	// because of the way look ahead dithering works, it needs to
	// modify te same image as it is reading from

	ww := imgIn.Bounds().Dx()
	hh := imgIn.Bounds().Dy()

	for y := range hh {
		for x := range ww {

			// keep everything in float as long as possible

			var diffErr float64

			pixel := GetPixel(imgIn, x, y)
			if pixel >= 127 {
				diffErr = float64(DiffGray(255, pixel))
			} else {
				diffErr = float64(DiffGray(0, pixel))
			}

			diffErr = (float64(diffErr) / 8.0)

			// I think this is cleaner, if not slower
			errorDiffustion(imgIn, diffErr, 3.0, x+1, y)
			errorDiffustion(imgIn, diffErr, 2.0, x-1, y+1)
			errorDiffustion(imgIn, diffErr, 3.0, x, y+1)

			if pixel >= 127 {
				SetPixel(imgIn, x, y, 255)
			} else {
				SetPixel(imgIn, x, y, 0)
			}

		}
	}

	return imgIn

}

//
//
//

// dither using the Floyd-Steinberg implementation
func DitherFloydSteinberg(imgIn *image.RGBA) *image.RGBA {

	//iOut := image.NewRGBA(imgIn.Bounds())

	// because of the way look ahead dithering works, it needs to
	// modify te same image as it is reading from

	ww := imgIn.Bounds().Dx()
	hh := imgIn.Bounds().Dy()

	for y := range hh {
		for x := range ww {

			// keep everything in float as long as possible

			var diffErr float64

			pixel := GetPixel(imgIn, x, y)
			if pixel >= 127 {
				diffErr = float64(DiffGray(255, pixel))
			} else {
				diffErr = float64(DiffGray(0, pixel))
			}

			diffErr = (float64(diffErr) / 16.0)

			// this is cleaner, if not slower
			errorDiffustion(imgIn, diffErr, 7.0, x+1, y)
			errorDiffustion(imgIn, diffErr, 3.0, x-1, y+1)
			errorDiffustion(imgIn, diffErr, 5.0, x, y+1)
			errorDiffustion(imgIn, diffErr, 1.0, x+1, y+1)

			if pixel >= 127 {
				SetPixel(imgIn, x, y, 255)
			} else {
				SetPixel(imgIn, x, y, 0)
			}

		}
	}

	return imgIn

}
