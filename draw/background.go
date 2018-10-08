package draw

import "image/color"

func (t Thumbnail) FillRect() {
	// TODO 背景用にRGPAセットをいくつか用意しておき, コピーする感じかなぁ...
	// メモリコピーと, for文でsetどっちが早いんだろう...
	rect := t.Img.Rect

	boder := rect.Max.Y / 3

	c1 := color.RGBA{214, 253, 153, 50} // #D6F599
	c2 := color.RGBA{168, 194, 86, 50}  // #A8C256

	for h := rect.Min.Y; h < rect.Max.Y; h++ {
		for v := rect.Min.X; v < rect.Max.X; v++ {
			if h/boder == 1 {
				t.Img.Set(v, h, c1)
			} else {
				t.Img.Set(v, h, c2)
			}
		}
	}
}
