package draw

import (
	"image"
)

type Thumbnail struct {
	Img   *image.RGBA
	Icons []*Icon
	Title *Text
}

func NewThumbnail() *Thumbnail {
	t := &Thumbnail{}
	x := 0
	y := 0
	width := IMG_SIZE
	height := IMG_SIZE

	t.Img = image.NewRGBA(image.Rect(x, y, width, height))
	icons, err := CreateIcons(4)
	if err != nil {
		panic(err)
	}

	t.Icons = icons

	t.FillRect()
	t.PutIcons()
	return t
}
