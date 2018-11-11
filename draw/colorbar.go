package draw

import (
	"image"
	"image/color"
	"image/draw"
)

func (t *Thumbnail) DrawColorBar(emos []int) {

	colWidth := IMG_SIZE / len(emos)

	img := image.NewRGBA(image.Rect(0, IMG_SIZE-COLBAR_HEIGHT, IMG_SIZE, IMG_SIZE))

	for i := 0; i < IMG_SIZE; i++ {
		for j := 0; j <= COLBAR_HEIGHT; j++ {
			img.Set(i, IMG_SIZE-j, BAR_COLOR[i/colWidth])
		}
	}

	mask := &skeleton{100, t.Img.Bounds()}
	draw.DrawMask(t.Img, img.Bounds(), img, image.Point{0, IMG_SIZE - COLBAR_HEIGHT}, mask, image.ZP, draw.Over)
}

// 透かし
type skeleton struct {
	alpha uint8
	rect  image.Rectangle
}

func (s *skeleton) ColorModel() color.Model {
	return color.AlphaModel
}

func (s *skeleton) Bounds() image.Rectangle {
	return s.rect
}

func (s *skeleton) At(x, y int) color.Color {
	return color.Alpha{s.alpha}
}
