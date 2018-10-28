package draw

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"time"

	"github.com/CaterpillarSan/DrawImg/model"
)

type Icon struct {
	FrameColor color.Color
	FrameType  int
	ImageUrl   string
	rect       image.Rectangle
	img        *image.RGBA
	rectWidth  uint
}

// アイコンをセット
func NewIconList(cards []model.Card) []*Icon {

	var icons []*Icon

	// iconを配置する場所をランダムに決定
	boad := NewBoad()
	iconNum := len(cards)

	rand.Seed(time.Now().UnixNano())
	var rnum int

	for i := 0; i < iconNum; i++ {
		icon := &Icon{}

		// 画像
		rnum = rand.Intn(len(cards))
		icon.ImageUrl = cards[rnum].ImageUrl.String
		icon.FrameColor = getFrameColor(cards[rnum].EmoID)
		cards = append(cards[:rnum], cards[rnum+1:]...)

		// 設置場所
		rect, size, err := boad.GetRandomRect()
		if err != nil {
			// これ以上アイコンが置けない
			return icons
		}
		icon.rectWidth = uint(2 * size * WIDTH)
		icon.rect = *rect

		icons = append(icons, icon)
	}

	return icons
}

// emo IDからフレームの色を決定する
func getFrameColor(emoID int) color.Color {
	if emoID > len(EmoColor) {
		return GRAY
	}
	return EmoColor[emoID-1]
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
	if err := icon.NewPicture(); err != nil {
		return err
	}
	iconImg := icon.img.SubImage(icon.img.Rect)
	draw.Draw(distImg, icon.rect, iconImg, image.Pt(0, 0), draw.Over)

	return nil
}
