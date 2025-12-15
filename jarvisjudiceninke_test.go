package imgproc

import (
	"image"
	"reflect"
	"testing"
)

func TestDitherJarvisJudiceNinke(t *testing.T) {
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
			if got := DitherJarvisJudiceNinke(tt.args.imgIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DitherJarvisJudiceNinke() = %v, want %v", got, tt.want)
			}
		})
	}
}
