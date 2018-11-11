package draw

import (
	"image"

	"github.com/CaterpillarSan/DrawImg/model"
)

type Thumbnail struct {
	Img       *image.RGBA
	Icons     []*Picture
	Title     *Text
	BKColor   string
	EmoIdList []int
}

func NewThumbnail(title string, cards []model.Card) *Thumbnail {
	t := &Thumbnail{}

	// イメージの土台
	t.Img = image.NewRGBA(image.Rect(0, 0, IMG_SIZE, IMG_SIZE))

	// EmoIDリスト
	t.EmoIdList = getEmoIdList(cards)

	// アイコン画像一覧
	t.Icons = NewPicList(cards)

	// タイトル
	t.Title = NewText(title)

	return t
}

func getEmoIdList(cards []model.Card) []int {
	var list []int
	for _, v := range cards {
		list = append(list, v.EmoID)
	}
	return list
}
