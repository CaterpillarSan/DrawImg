package draw

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"

	"github.com/nfnt/resize"
)

func (icon *Icon) NewPicture() error {
	// イメージをimage.Imageに変換し, リサイズ
	img, err := icon.decodeImage()
	if err != nil {
		return err
	}

	icon.img = img
	return nil
}

func (icon *Icon) decodeImage() (*image.Image, error) {
	u, err := url.Parse(icon.ImageUrl)
	if err != nil {
		return nil, err
	}
	var originImg image.Image

	if u.Scheme == "https" && u.Host == "s3-ap-northeast-1.amazonaws.com" {
		// Get image
		originImg, err = getImageFromUrl(icon.ImageUrl)
	} else {
		// デバッグ用, ローカルのファイルを取り出す
		originImg, err = getImageFromLocal(icon.ImageUrl)
		// return nil, errors.New("Unauthorized URL") // TODO : エラーちゃんとする
	}
	if err != nil {
		return nil, err
	}

	smallImg := resizeSquare(originImg, IMG_SIZE)
	return &smallImg, nil
}

func getImageFromUrl(imgUrl string) (image.Image, error) {
	resp, err := http.Get(imgUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	return img, nil
}

func getImageFromLocal(imgUrl string) (image.Image, error) {
	f, err := os.Open(imgUrl)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	originImg, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return originImg, nil
}

// 画像を正方形に整形
func resizeSquare(img image.Image, width uint) image.Image {
	xx := img.Bounds().Dx()
	yy := img.Bounds().Dy()
	len := int(math.Min(float64(xx), float64(yy)))
	var point image.Point
	if xx > yy {
		point = image.Pt((xx-yy)/2, 0)
	} else {
		point = image.Pt(0, (yy-xx)/2)
	}
	// 土台となる無地のimage
	out := image.NewRGBA(image.Rect(0, 0, len, len))
	draw.Draw(out, out.Bounds(), img, point, draw.Src)
	smallImg := resize.Thumbnail(width, width, out, resize.Lanczos3)
	return smallImg
}

// 画像を切り取る
func cutImage(in *image.Image, frameType int, col color.Color, r int) *image.RGBA {
	// 土台となる無地のimage
	out := image.NewRGBA(image.Rect(0, 0, 2*r, 2*r))
	var mask image.Image // 円形以外の場合はここを変える
	switch frameType {
	case DIAMOND:
		mask = &diamond{image.Pt(r, r), r}
	case CIRCLE:
		mask = &circle{image.Pt(r, r), r}
	default:
		mask = &diamond{image.Pt(r, r), r}
	}
	// 画像切り取り
	draw.DrawMask(out, out.Bounds(), *in, image.ZP, mask, image.ZP, draw.Over)
	// 枠
	drawBounds(out, frameType, col, r)
	return out
}

// 枠を書く
func drawBounds(out *image.RGBA, frameType int, col color.Color, r int) {
	switch frameType {
	case DIAMOND:
		drawDiamondBounds(out, frameType, col, r)
	case CIRCLE:
		drawCircleBounds(out, frameType, col, r)
	default:
		drawDiamondBounds(out, frameType, col, r)
	}

}

func drawDiamondBounds(img *image.RGBA, frameType int, col color.Color, r int) {
	bold := r / 7
	for i := 0; i < r; i++ {
		for j := 0; j < bold && i+j <= r; j++ {
			img.Set(i+j, r-i, col)
			img.Set(i+j, r+i, col)
			img.Set(2*r-i-j, r+i, col)
			img.Set(2*r-i-j, r-i, col)
		}
	}
}

func drawCircleBounds(img *image.RGBA, frameType int, col color.Color, r int) {
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

// ************************************************************** //
// 正方形作ったけどいらなかった...

type square struct {
	width  int
	height int
}

func (s *square) ColorModel() color.Model {
	return color.AlphaModel
}

func (s *square) Bounds() image.Rectangle {
	return image.Rect(0, 0, s.width, s.height)
}

func (s *square) At(x, y int) color.Color {
	if s.width > s.height {
		d := (s.width - s.height) / 2
		if x < d || d < x {
			return color.Alpha{0}
		} else {
			return color.Alpha{255}
		}
	} else {
		d := (s.height - s.width) / 2
		if y < d || d < x {
			return color.Alpha{0}
		} else {
			return color.Alpha{255}
		}
	}
}

// *************************************************************** //
// 画像マスク用
// interface image.Image

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

type diamond struct {
	p image.Point
	r int
}

func (d *diamond) ColorModel() color.Model {
	return color.AlphaModel
}

func (d *diamond) Bounds() image.Rectangle {
	return image.Rect(d.p.X-d.r, d.p.Y-d.r, d.p.X+d.r, d.p.Y+d.r)
}

func (d *diamond) At(x, y int) color.Color {
	if y > x-d.r && y < x+d.r && y > -x+d.r && y < -x+3*d.r {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

// *************************************************************** //
