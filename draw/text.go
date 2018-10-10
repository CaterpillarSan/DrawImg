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
	d, err := text2img.NewDrawer(text2img.Params{
		Width:           IMG_SIZE,
		Height:          IMG_SIZE / 3,
		FontPath:        FONT_PATH,
		FontSize:        FONT_SIZE,
		BackgroundColor: MUSYOKU_TOUMEI,
		TextColor:       DARK_GREEN,
	})
	if err != nil {
		return err
	}
	textImg, err := d.Draw(t.Title.Sentence)
	if err != nil {
		return err
	}

	rect := image.Rect(0, IMG_SIZE/3, IMG_SIZE, IMG_SIZE*2/3)

	// サムネイル画像と合成
	draw.Draw(t.Img, rect, textImg, image.Pt(0, 0), draw.Over)

	// file, _ := os.Create("test.jpg")
	// err = jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
	return nil

}
