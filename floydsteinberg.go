package imgproc

import (
	"image"
	"image/draw"
)

// False Floyd-Steinberg dithering implementation
//
// imgIn is passed as a pointer for speed, but is not modified.
func DitherFalseFloydSteinberg(imgIn *image.RGBA) *image.RGBA {

	// use draw to create, copy the input image to a new image
	// and then modify that image in place
	imgMod := image.NewRGBA(imgIn.Bounds())
	draw.Draw(imgMod, imgMod.Bounds(), imgIn, imgIn.Bounds().Min, draw.Src)

	//imgMod := image.NewRGBA(imgIn.Bounds())

	// because of the way look ahead dithering works, it needs to
	// modify te same image as it is reading from

	ww := imgMod.Bounds().Dx()
	hh := imgMod.Bounds().Dy()

	for y := range hh {
		for x := range ww {

			// keep everything in float as long as possible

			var diffErr float64

			pixel := GetPixel(imgMod, x, y)
			if pixel >= 127 {
				diffErr = float64(DiffGray(255, pixel))
			} else {
				diffErr = float64(DiffGray(0, pixel))
			}

			diffErr = (float64(diffErr) / 8.0)

			// I think this is cleaner, if not slower
			errorDiffustion(imgMod, diffErr, 3.0, x+1, y)
			errorDiffustion(imgMod, diffErr, 2.0, x-1, y+1)
			errorDiffustion(imgMod, diffErr, 3.0, x, y+1)

			if pixel >= 127 {
				SetPixel(imgMod, x, y, 255)
			} else {
				SetPixel(imgMod, x, y, 0)
			}

		}
	}

	return imgMod

}

//
//
//

// dither using the Floyd-Steinberg implementation
func DitherFloydSteinberg(imgIn *image.RGBA) *image.RGBA {

	imgMod := image.NewRGBA(imgIn.Bounds())
	draw.Draw(imgMod, imgMod.Bounds(), imgIn, imgIn.Bounds().Min, draw.Src)

	// because of the way look ahead dithering works, it needs to
	// modify te same image as it is reading from

	ww := imgMod.Bounds().Dx()
	hh := imgMod.Bounds().Dy()

	for y := range hh {
		for x := range ww {

			// keep everything in float as long as possible

			var diffErr float64

			pixel := GetPixel(imgMod, x, y)
			if pixel >= 127 {
				diffErr = float64(DiffGray(255, pixel))
			} else {
				diffErr = float64(DiffGray(0, pixel))
			}

			diffErr = (float64(diffErr) / 16.0)

			// this is cleaner, if not slower
			errorDiffustion(imgMod, diffErr, 7.0, x+1, y)
			errorDiffustion(imgMod, diffErr, 3.0, x-1, y+1)
			errorDiffustion(imgMod, diffErr, 5.0, x, y+1)
			errorDiffustion(imgMod, diffErr, 1.0, x+1, y+1)

			if pixel >= 127 {
				SetPixel(imgMod, x, y, 255)
			} else {
				SetPixel(imgMod, x, y, 0)
			}

		}
	}

	return imgMod

}
