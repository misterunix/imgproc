package imgproc

import (
	"image"
	"reflect"
	"testing"
)

func TestDitherFalseFloydSteinberg(t *testing.T) {
	type args struct {
		imgIn *image.RGBA
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
			if got := DitherFalseFloydSteinberg(tt.args.imgIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DitherFalseFloydSteinberg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDitherFloydSteinberg(t *testing.T) {
	type args struct {
		imgIn *image.RGBA
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
			if got := DitherFloydSteinberg(tt.args.imgIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DitherFloydSteinberg() = %v, want %v", got, tt.want)
			}
		})
	}
}
