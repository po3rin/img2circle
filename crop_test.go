package img2circle_test

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"testing"

	"github.com/po3rin/img2circle"
)

func TestCropCircle(t *testing.T) {
	tests := []struct {
		path string
	}{
		{
			path: "testdata/gopher.jpeg",
		},
		{
			path: "testdata/aws.png",
		},
	}

	for _, tt := range tests {
		img, err := os.Open(tt.path)
		if err != nil {
			log.Fatal(err)
		}
		defer img.Close()
		src, _, err := image.Decode(img)
		if err != nil {
			log.Fatal(err)
		}

		cropper, err := img2circle.NewCropper(img2circle.Params{Src: src})
		if err != nil {
			t.Fatalf("not expected error: %v", err.Error())
		}
		_ = cropper.CropCircle()
	}
}
