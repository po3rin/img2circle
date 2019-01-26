package img2circle_test

import (
	_ "image/jpeg"
	_ "image/png"
	"testing"

	"github.com/po3rin/img2circle"
)

func TestCropCircle(t *testing.T) {
	tests := []struct {
		params img2circle.Params
	}{
		{
			params: img2circle.Params{
				ImgPath: "testdata/gopher.jpeg",
			},
		},
		{
			params: img2circle.Params{
				ImgPath: "testdata/aws.png",
			},
		},
	}

	for _, tt := range tests {
		croper, err := img2circle.NewCroper(tt.params)
		if err != nil {
			t.Fatalf("not expected error: %v", err.Error())
		}
		_ = croper.CropCircle()
	}
}
