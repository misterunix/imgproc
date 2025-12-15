package imgproc

import (
	"image"
	"reflect"
	"testing"
)

func TestBWBT2020(t *testing.T) {
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
			if got := BWBT2020(tt.args.imgIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BWBT2020() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBWLuminanceGammaBT2020(t *testing.T) {
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
			if got := BWLuminanceGammaBT2020(tt.args.imgIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BWLuminanceGammaBT2020() = %v, want %v", got, tt.want)
			}
		})
	}
}
