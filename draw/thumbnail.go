package draw

import (
	"image"
	"math/rand"

	"github.com/CaterpillarSan/DrawImg/model"
)

type Thumbnail struct {
	Img     *image.RGBA
	Icons   []*Icon
	Title   *Text
	BKColor string
}

func NewThumbnail(title string, cards []model.Card) *Thumbnail {
	t := &Thumbnail{}

	// イメージの土台
	t.Img = image.NewRGBA(image.Rect(0, 0, IMG_SIZE, IMG_SIZE))

	// 背景の色を決める
	t.BKColor = BKColors[rand.Intn(len(BKColors))]

	// ボード
	boad := NewBoad()

	// アイコン画像一覧
	t.Icons = NewIconList(cards, boad, t.BKColor)

	// タイトル
	t.Title = NewText(title, boad)

	return t
}
