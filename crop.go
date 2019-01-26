package img2circle

import (
	"image"
	"image/color"
	"image/draw"
)

// Cropper for crop circle.
type Cropper interface {
	CropCircle() *image.RGBA
	setDst()
	setSrc(src image.Image) error
}

type cropper struct {
	src       image.Image
	dst       *image.RGBA
	srcWidth  int
	srcHeight int
	radius    int
}

// Params is parameters for NewDrawer functio
type Params struct {
	Src image.Image
	// PosX    int
	// PosY    int
}

// NewCropper init cropper from Params
func NewCropper(params Params) (Cropper, error) {
	d := &cropper{}
	err := d.setSrc(params.Src)
	if err != nil {
		return d, err
	}
	d.setDst()
	return d, nil
}

func (c *cropper) setSrc(src image.Image) error {
	b := src.Bounds()
	srcWidth := b.Max.X
	srcHeight := b.Max.Y

	var radius int
	if srcWidth <= srcHeight {
		radius = srcWidth / 2
	} else {
		radius = srcHeight / 2
	}

	c.src = src
	c.srcWidth = srcWidth
	c.srcHeight = srcHeight
	c.radius = radius

	return nil
}

func (c *cropper) setDst() {
	rect := image.Rect(0, 0, c.srcWidth, c.srcHeight)
	dst := image.NewRGBA(rect)
	fillRect(dst, color.RGBA{0, 0, 0, 0})
	c.dst = dst
}

// CropCircle crop a circle image out of image.
func (c *cropper) CropCircle() *image.RGBA {
	circle := &circle{p: image.Point{c.srcWidth / 2, c.srcHeight / 2}, r: c.radius}
	dst := c.dst

	draw.DrawMask(dst, dst.Bounds(), c.src, image.ZP, circle, image.ZP, draw.Over)
	return dst
}

func fillRect(img *image.RGBA, col color.Color) {
	rect := img.Rect
	for h := rect.Min.Y; h < rect.Max.Y; h++ {
		for v := rect.Min.X; v < rect.Max.X; v++ {
			img.Set(v, h, col)
		}
	}
}
