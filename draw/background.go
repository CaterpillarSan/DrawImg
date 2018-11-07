package draw

import "image/color"

// 背景を準備する
func (t *Thumbnail) FillRect() error {

	// rand.Seed(time.Now().UnixNano())
	//
	// backgroundUrl := BACKGROUND_PATH + t.BKColor + "/" + strconv.Itoa(rand.Intn(4)) + ".png"
	// backgroundImg, err := getImageFromLocal(backgroundUrl)
	// if err != nil {
	// 	return err
	// }
	// backgroundImg = resizeSquare(backgroundImg, uint(IMG_SIZE))
	// // image.Image -> image.RGBAの型変換がわからず...w
	// b := backgroundImg.Bounds()
	// draw.Draw(t.Img, t.Img.Bounds(), backgroundImg, b.Min, draw.Src)
	for i := 0; i < IMG_SIZE; i++ {
		for j := 0; j < IMG_SIZE; j++ {
			t.Img.Set(i, j, color.RGBA{255, 255, 255, 0})
		}
	}
	return nil
}
