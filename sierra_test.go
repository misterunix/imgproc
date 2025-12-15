package imgproc

import (
	"image"
	"reflect"
	"testing"
)

func TestDitherSierra1(t *testing.T) {
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
			if got := DitherSierra1(tt.args.imgIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DitherSierra1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDitherSierra2(t *testing.T) {
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
			if got := DitherSierra2(tt.args.imgIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DitherSierra2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDitherSierra2_4a(t *testing.T) {
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
			if got := DitherSierra2_4a(tt.args.imgIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DitherSierra2_4a() = %v, want %v", got, tt.want)
			}
		})
	}
}
