package imgproc

import (
	"image"
	"reflect"
	"testing"
)

func TestBWBT709(t *testing.T) {
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
			if got := BWBT709(tt.args.imgIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BWBT709() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBWLuminanceGammaBT709(t *testing.T) {
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
			if got := BWLuminanceGammaBT709(tt.args.imgIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BWLuminanceGammaBT709() = %v, want %v", got, tt.want)
			}
		})
	}
}
