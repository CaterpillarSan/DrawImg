package draw

import (
	"image"
	"image/draw"

	"github.com/Iwark/text2img"
)

type Text struct {
	Sentence string
	length   int
}

func NewText(sen string) *Text {
	text := &Text{sen, len(sen)}
	return text
}

func (t *Thumbnail) SetTitle() error {
	// タイトルが短かった時の処理
	title := t.Title.Sentence
	// 日本語でも, あえてbyte数を数えて見る
	PADDING := " "
	for len(title) < 16 {
		title = PADDING + title + PADDING
	}

	d, err := text2img.NewDrawer(text2img.Params{
		Width:           IMG_SIZE,
		Height:          TITLE_HEIGHT,
		FontPath:        FONT_PATH,
		BackgroundColor: TITLE_BACKGROUND,
		TextColor:       TITLE_CHAR,
	})
	if err != nil {
		return err
	}

	textImg, err := d.Draw(title)
	if err != nil {
		return err
	}

	rect := image.Rect(0, IMG_SIZE-TITLE_HEIGHT-COLBAR_HEIGHT, IMG_SIZE, IMG_SIZE-COLBAR_HEIGHT)

	// サムネイル画像と合成
	draw.Draw(t.Img, rect, textImg, image.ZP, draw.Over)

	return nil

}
