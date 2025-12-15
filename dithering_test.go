package imgproc

import (
	"image"
	"testing"
)

func Test_errorDiffustion(t *testing.T) {
	type args struct {
		img       *image.RGBA
		errorTerm float64
		errorMul  float64
		x         int
		y         int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorDiffustion(tt.args.img, tt.args.errorTerm, tt.args.errorMul, tt.args.x, tt.args.y)
		})
	}
}
