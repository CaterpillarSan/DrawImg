package draw

import (
	"image"
	"image/color"
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
		BackgroundColor: MUSYOKU_TOUMEI,
		TextColor:       (ColorMap[t.BKColor][0]).(color.RGBA),
	})
	if err != nil {
		return err
	}

	// TODO タイトルが短かった時の処理をちゃんとする
	textImg, err := d.Draw(t.Title.Sentence)
	if err != nil {
		return err
	}

	var rect image.Rectangle
	switch t.Title.placement {
	case "header":
		rect = image.Rect(0, -WIDTH, IMG_SIZE, WIDTH*3)
	case "footer":
		rect = image.Rect(0, IMG_SIZE-WIDTH*3, IMG_SIZE, IMG_SIZE)
	}

	// サムネイル画像と合成
	draw.Draw(t.Img, rect, textImg, image.Pt(0, 0), draw.Over)

	return nil

}
