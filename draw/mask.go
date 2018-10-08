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

func NewPicture(imageUrl string) (*Picture, error) {
	pic := &Picture{}
	// イメージをimage.Imageに変換し, リサイズ
	img, err := decodeImage(imageUrl)
	if err != nil {
		return nil, err
	}

	// アイコンを丸く切り取る
	out := cutImage(img)

	// ふちをつける
	drawBounds(out, DARK_GREEN)

	pic.Img = out
	return pic, nil
}

func decodeImage(imageUrl string) (*image.Image, error) {
	f, err := os.Open(imageUrl)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	originImg, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	smallImg := resize.Thumbnail(ICON_RADIUS*2, ICON_RADIUS*2, originImg, resize.Lanczos3)

	return &smallImg, nil
}

// 円形に切り取る
func cutImage(in *image.Image) *image.RGBA {
	// 土台となる無地のimage
	out := image.NewRGBA(image.Rect(0, 0, 2*ICON_RADIUS, 2*ICON_RADIUS))
	// 円形以外の場合はここを変える
	mask := &circle{image.Pt(ICON_RADIUS, ICON_RADIUS), ICON_RADIUS}
	// 画像切り取り
	draw.DrawMask(out, out.Bounds(), *in, image.ZP, mask, image.ZP, draw.Over)
	return out
}

// 枠を書く
func drawBounds(img *image.RGBA, col color.Color) {
	r := ICON_RADIUS
	bold := ICON_RADIUS / 7
	// TODO radianの刻みもIMAGE_SIZEによって変えるべき
	for rad := 0.0; rad < 2.0*float64(ICON_RADIUS); rad += 0.01 {
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
