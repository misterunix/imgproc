package imgproc

import (
	"image"
	"reflect"
	"testing"
)

func TestBWLuminance(t *testing.T) {
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
			if got := BWLuminance(tt.args.imgIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BWLuminance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBWLuminanceGamma(t *testing.T) {
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
			if got := BWLuminanceGamma(tt.args.imgIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BWLuminanceGamma() = %v, want %v", got, tt.want)
			}
		})
	}
}
