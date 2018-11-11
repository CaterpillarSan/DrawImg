package draw

import (
	"image"
	"image/color"
	"image/draw"
)

func (t *Thumbnail) DrawColorBar() {

	colWidth := IMG_SIZE / len(t.EmoIdList)

	img := image.NewRGBA(image.Rect(0, IMG_SIZE-COLBAR_HEIGHT, IMG_SIZE, IMG_SIZE))

	for emo := 0; emo < len(t.EmoIdList); emo++ {
		for i := colWidth * emo; i < colWidth*(emo+1); i++ {
			for j := 0; j <= COLBAR_HEIGHT; j++ {
				img.Set(i, IMG_SIZE-j, BAR_COLOR[t.EmoIdList[emo]-1])
			}
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
