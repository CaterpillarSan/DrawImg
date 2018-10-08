package draw

func (t *Thumbnail) FillRect() {
	// TODO 背景用にRGPAセットをいくつか用意しておき, コピーする感じかなぁ...
	// メモリコピーと, for文でsetどっちが早いんだろう...
	// というか, 背景を画像で何種類か作っておいてimage/drawが正解そう...
	rect := t.Img.Rect

	boder := rect.Max.Y / 3

	for h := rect.Min.Y; h < rect.Max.Y; h++ {
		for v := rect.Min.X; v < rect.Max.X; v++ {
			if h/boder == 1 {
				t.Img.Set(v, h, LIGHT_GREEN)
			} else {
				t.Img.Set(v, h, SHIBU_GREEN)
			}
		}
	}
}
