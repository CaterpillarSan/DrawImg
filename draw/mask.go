package draw

import (
	"image"
	"image/color"
	"image/draw"
	"math"
	"os"

	"github.com/nfnt/resize"
)

type Picture struct {
	Img *image.RGBA
}

func (i *Icon) NewPicture() (*Picture, error) {
	pic := &Picture{}
	// イメージをimage.Imageに変換し, リサイズ
	img, err := i.decodeImage()
	if err != nil {
		return nil, err
	}

	r := int(i.rectWidth / 2)
	// アイコンを丸く切り取る
	out := cutImage(img, r)

	// ふちをつける
	drawBounds(out, DARK_GREEN, r)

	pic.Img = out
	return pic, nil
}

func (i *Icon) decodeImage() (*image.Image, error) {
	f, err := os.Open(i.ImageUrl)
	if err != nil {
		return nil, err
	})
	defer f.Close()
	originImg, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	smallImg := resize.Thumbnail(i.rectWidth, i.rectWidth, originImg, resize.Lanczos3)

	return &smallImg, nil
}

// 円形に切り取る
func cutImage(in *image.Image, r int) *image.RGBA {
	// 土台となる無地のimage
	out := image.NewRGBA(image.Rect(0, 0, r*2, r*2))
	// 円形以外の場合はここを変える
	mask := &circle{image.Pt(r, r), r}
	// 画像切り取り
	draw.DrawMask(out, out.Bounds(), *in, image.ZP, mask, image.ZP, draw.Over)
	return out
}

// 枠を書く
func drawBounds(img *image.RGBA, col color.Color, r int) {
	bold := r / 7
	// TODO radianの刻みもIMAGE_SIZEによって変えるべき
	for rad := 0.0; rad < 2.0*float64(r); rad += 0.01 {
		for i := 0; i < bold; i++ {
			x := int(float64(r) + float64(r-i)*math.Cos(rad))
			y := int(float64(r) + float64(r-i)*math.Sin(rad))
			img.Set(x, y, col)

		}
	}
}

// *************************************************************** //
// 画像マスク用

type circle struct {
	p image.Point
	r int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}
