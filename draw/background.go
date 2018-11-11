package draw

import "image/color"

// 背景を準備する
func (t *Thumbnail) FillRect() error {

	for i := 0; i < IMG_SIZE; i++ {
		for j := 0; j < IMG_SIZE; j++ {
			t.Img.Set(i, j, color.RGBA{255, 255, 255, 0})
		}
	}
	return nil
}
