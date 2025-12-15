package imgproc

import (
	"image"
	"reflect"
	"testing"
)

func TestDitherUnknown1(t *testing.T) {
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
			if got := DitherUnknown1(tt.args.imgIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DitherUnknown1() = %v, want %v", got, tt.want)
			}
		})
	}
}
