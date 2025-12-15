package imgproc

import (
	"image"
	"reflect"
	"testing"
)

func TestBWBT601(t *testing.T) {
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
			if got := BWBT601(tt.args.imgIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BWBT601() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBWLuminanceGammaBT601(t *testing.T) {
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
			if got := BWLuminanceGammaBT601(tt.args.imgIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BWLuminanceGammaBT601() = %v, want %v", got, tt.want)
			}
		})
	}
}
