package draw

import (
	"image/draw"
	"math/rand"
	"strconv"
	"time"
)

// 背景を準備する
func (t *Thumbnail) FillRect() error {

	rand.Seed(time.Now().UnixNano())

	backgroundUrl := BACKGROUND_PATH + t.BKColor + "/" + strconv.Itoa(rand.Intn(4)) + ".png"
	backgroundImg, err := getImageFromLocal(backgroundUrl)
	if err != nil {
		return err
	}
	backgroundImg = resizeSquare(backgroundImg, uint(IMG_SIZE))
	// image.Image -> image.RGBAの型変換がわからず...w
	b := backgroundImg.Bounds()
	draw.Draw(t.Img, t.Img.Bounds(), backgroundImg, b.Min, draw.Src)
	return nil
}
