package imgproc

import (
	"bytes"
	"errors"

	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// convert image.Image to image.RGBA.
func ConvertToRGBA(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	rgbaImg := image.NewRGBA(bounds)
	draw.Draw(rgbaImg, bounds, img, bounds.Min, draw.Src)
	return rgbaImg
}

//
//
//

// load image from file returning an image.RGBA
func LoadImageFromFile(filename string) (*image.RGBA, error) {
	var tImg image.Image  // temporary image
	var outImg image.RGBA // image.RGBA to return

	file, err := os.Open(filename)
	if err != nil {
		fullError := "Error loading image file: '" +
			filename + "'\n" + err.Error()
		return &outImg, errors.New(fullError)
	}
	defer file.Close()

	tImg, _, err = image.Decode(file)
	if err != nil {
		fullError := "Error decoding image file: '" +
			filename + "'\n" + err.Error()
		return &outImg, errors.New(fullError)
	}

	outImg = *ConvertToRGBA(tImg)

	return &outImg, nil
}

//
//
//

// save image to file
func SaveImageToFile(filename string, img *image.RGBA) error {
	file, err := os.Create(filename)
	if err != nil {
		fullError := "Error creating image file: '" +
			filename + "'\n" + err.Error()
		return errors.New(fullError)
	}
	defer file.Close() // make sure the file closes

	if strings.HasSuffix(strings.ToLower(filename), ".jpg") ||
		strings.HasSuffix(strings.ToLower(filename), ".jpeg") {
		err = jpeg.Encode(file, img, nil)
		if err != nil {
			fullError := "Error encoding image file: '" +
				filename + "'\n" + err.Error()
			return errors.New(fullError)
		}
	}
	if strings.HasSuffix(strings.ToLower(filename), ".png") {
		err = png.Encode(file, img)
		if err != nil {
			fullError := "Error encoding image file: '" +
				filename + "'\n" + err.Error()
			return errors.New(fullError)
		}
	}
	return nil
}

//
//
//

// helper function to calculate the difference in gray values
func DiffGray(from, to int) int {
	return to - from
}

// Helper function to get the pixel value. If x or y are out of bounds,
// return 0. The returned value is the gray value of the pixel.
func GetPixel(imgIn *image.RGBA, x, y int) int {

	if x < 0 || y < 0 || x >= imgIn.Bounds().Dx() ||
		y >= imgIn.Bounds().Dy() {
		return 0
	}
	rv, _, _, _ := imgIn.At(x, y).RGBA()
	r := uint8(rv >> 8)

	return int(r)
}

// Helper function to set the pixel value. If the value < 0 or > 255 then
// limit the value. If x or y are out of bounds, do nothing
func SetPixel(imgIn *image.RGBA, x, y, gray int) {
	if x < 0 || y < 0 || x >= imgIn.Bounds().Dx() ||
		y >= imgIn.Bounds().Dy() {
		return
	}
	if gray < 0 {
		gray = 0
	}
	if gray > 255 {
		gray = 255
	}
	imgIn.Set(x, y, color.Gray{uint8(gray)})
}

//
//
//

// Calculate the Root Mean Square Error (RMSE) between two images.
func CalcRMSE(img1, img2 *image.RGBA) (float64, error) {
	if img1.Bounds() != img2.Bounds() {
		return 0, errors.New("CalcRMSE: images are not the same size") // images must be the same size
	}

	var sum float64
	count := 0

	for y := img1.Bounds().Min.Y; y < img1.Bounds().Max.Y; y++ {
		for x := img1.Bounds().Min.X; x < img1.Bounds().Max.X; x++ {
			pixel1 := GetPixel(img1, x, y)
			pixel2 := GetPixel(img2, x, y)
			diff := DiffGray(pixel1, pixel2)
			sum += float64(diff * diff)
			count++
		}
	}

	if count == 0 {
		return 0, nil // avoid division by zero
	}

	return sum / float64(count), nil
}

//
//
//

func EncodeImageToPNG(img *image.RGBA) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
