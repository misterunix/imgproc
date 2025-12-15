package imgproc

import (
	"image"
	"reflect"
	"testing"
)

func TestConvertToRGBA(t *testing.T) {
	type args struct {
		img image.Image
	}
	tests := []struct {
		name string
		args args
		want *image.RGBA
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToRGBA(tt.args.img); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToRGBA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadImageFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    *image.RGBA
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadImageFromFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadImageFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadImageFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveImageToFile(t *testing.T) {
	type args struct {
		filename string
		img      *image.RGBA
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveImageToFile(tt.args.filename, tt.args.img); (err != nil) != tt.wantErr {
				t.Errorf("SaveImageToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDiffGray(t *testing.T) {
	type args struct {
		from int
		to   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiffGray(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("DiffGray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPixel(t *testing.T) {
	type args struct {
		imgIn *image.RGBA
		x     int
		y     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPixel(tt.args.imgIn, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("GetPixel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetPixel(t *testing.T) {
	type args struct {
		imgIn *image.RGBA
		x     int
		y     int
		gray  int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetPixel(tt.args.imgIn, tt.args.x, tt.args.y, tt.args.gray)
		})
	}
}

func TestCalcRMSE(t *testing.T) {
	type args struct {
		img1 *image.RGBA
		img2 *image.RGBA
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalcRMSE(tt.args.img1, tt.args.img2)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalcRMSE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalcRMSE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeImageToPNG(t *testing.T) {
	type args struct {
		img *image.RGBA
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeImageToPNG(tt.args.img)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeImageToPNG() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeImageToPNG() = %v, want %v", got, tt.want)
			}
		})
	}
}
