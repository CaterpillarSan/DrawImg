package draw

import (
	"image"
)

type Thumbnail struct {
	Img *image.RGBA
}

func NewThumbnail() *Thumbnail {
	t := &Thumbnail{}
	x := 0
	y := 0
	width := 300
	height := 300

	t.Img = image.NewRGBA(image.Rect(x, y, width, height))

	t.FillRect()

	return t
}
