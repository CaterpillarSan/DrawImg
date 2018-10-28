package draw

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"time"
)

type Icon struct {
	frame_color color.Color
	ImageUrl    string
	rect        image.Rectangle
	pic         *Picture
	rectWidth   uint
}

// アイコン設置可能な位置の列挙

// アイコンをセット
func NewIconList(urls []string) []*Icon {

	var icons []*Icon

	// iconを配置する場所をランダムに決定
	boad := NewBoad()
	iconNum := len(urls)

	rand.Seed(time.Now().UnixNano())

	var rnum int
	for i := 0; i < iconNum; i++ {
		icon := &Icon{}
		// 設置場所
		rect, size, err := boad.GetRandomRect()
		if err != nil {
			// これ以上アイコンが置けない
			return icons
		}
		icon.rectWidth = uint(2 * size * boad.width)
		icon.rect = *rect

		// 画像
		rnum = rand.Intn(len(urls))
		icon.ImageUrl = urls[rnum]
		urls = append(urls[:rnum], urls[rnum+1:]...)

		icons = append(icons, icon)
	}

	return icons
}

// アイコン(=写真)を描画
func (t *Thumbnail) PutIcons() error {
	if t.Icons == nil {
		return nil
	}
	for _, icon := range t.Icons {
		if err := icon.drawIconImage(t.Img); err != nil {
			return err
		}

	}
	return nil
}

func (icon *Icon) drawIconImage(distImg *image.RGBA) error {

	// アイコン画像生成
	pic, err := icon.NewPicture()
	if err != nil {
		return err
	}
	icon.pic = pic
	iconImg := pic.Img.SubImage(pic.Img.Rect)
	draw.Draw(distImg, icon.rect, iconImg, image.Pt(0, 0), draw.Over)

	return nil
}
