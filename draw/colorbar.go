package draw

import (
	"image"
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

	// TODO わからん
	draw.Draw(t.Img, img.Rect, img, image.Point{0, IMG_SIZE - COLBAR_HEIGHT}, draw.Over)

}
