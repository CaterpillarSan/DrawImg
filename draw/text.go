package draw

import (
	"image"
	"image/draw"

	"github.com/Iwark/text2img"
)

type Text struct {
	Sentence  string
	length    int
	placement string
}

func NewText(sen string, boad Boad) *Text {
	text := &Text{sen, len(sen), boad.titlePoint}
	return text
}

func (t *Thumbnail) SetTitle() error {
	d, err := text2img.NewDrawer(text2img.Params{
		Width:           IMG_SIZE,
		Height:          WIDTH * 3,
		FontPath:        FONT_PATH,
		BackgroundColor: TITLE_BACKGROUND,
		TextColor:       TITLE_CHAR,
	})
	if err != nil {
		return err
	}

	// TODO タイトルが短かった時の処理をちゃんとする
	textImg, err := d.Draw(t.Title.Sentence)
	if err != nil {
		return err
	}

	rect := image.Rect(0, IMG_SIZE-WIDTH*3, IMG_SIZE, IMG_SIZE)

	// サムネイル画像と合成
	draw.Draw(t.Img, rect, textImg, image.Pt(0, 0), draw.Over)

	return nil

}
