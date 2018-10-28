package draw

import (
	"image"

	"github.com/CaterpillarSan/DrawImg/model"
)

type Thumbnail struct {
	Img   *image.RGBA
	Icons []*Icon
	Title *Text
}

func NewThumbnail(title string, cards []model.Card) *Thumbnail {
	t := &Thumbnail{}
	x := 0
	y := 0
	width := IMG_SIZE
	height := IMG_SIZE

	// イメージの土台
	t.Img = image.NewRGBA(image.Rect(x, y, width, height))

	// アイコン画像一覧
	t.Icons = NewIconList(cards)

	// タイトル
	t.Title = NewText(title)

	return t
}
