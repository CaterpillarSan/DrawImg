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
	ImageUrl   string
	rect       image.Rectangle
	img        *image.Image
}

// アイコンをセット
func NewIconList(cards []model.Card, boad Boad, bkColor string) []*Icon {

	var icons []*Icon

	rand.Seed(time.Now().UnixNano())

	// for i := 0; i < 4; i++ {
	icon := &Icon{}

	// 画像
	rnum := rand.Intn(len(cards))
	icon.ImageUrl = cards[rnum].ImageUrl.String
	icon.FrameColor = BAR_COLOR[cards[rnum].EmoID]
	cards = append(cards[:rnum], cards[rnum+1:]...)

	icons = append(icons, icon)
	// }

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
	if err := icon.NewPicture(); err != nil {
		return err
	}
	draw.Draw(distImg, distImg.Rect, *icon.img, image.Pt(0, 0), draw.Over)

	return nil
}
