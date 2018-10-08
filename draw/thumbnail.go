package draw

import (
	"fmt"
	"image"
)

type Thumbnail struct {
	Img   *image.RGBA
	Icons []*Icon
	Title *Text
}

func NewThumbnail(urls []string) *Thumbnail {
	t := &Thumbnail{}
	x := 0
	y := 0
	width := IMG_SIZE
	height := IMG_SIZE

	t.Img = image.NewRGBA(image.Rect(x, y, width, height))
	icons, err := CreateIcons(urls)
	if err != nil {
		panic(err)
	}

	t.Icons = icons

	t.FillRect()
	if err = t.PutIcons(); err != nil {
		fmt.Println(err)
	}
	t.SetTitle()
	return t
}
