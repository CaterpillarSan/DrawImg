package draw

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	_ "image/png"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/nfnt/resize"
)

type Icon struct {
	frame_color color.Color
	ImageUrl    string
	C           Circle
}

type Circle struct {
	Rad int
	Pt  image.Point
}

var x_width = IMG_SIZE / (MAX_ICON_NUM/2 + 1)
var y_width = IMG_SIZE / 6

var icon_points = [6]image.Point{
	image.Point{x_width * 1, y_width},
	image.Point{x_width * 2, y_width},
	image.Point{x_width * 3, y_width},
	image.Point{x_width * 1, IMG_SIZE - y_width},
	image.Point{x_width * 2, IMG_SIZE - y_width},
	image.Point{x_width * 3, IMG_SIZE - y_width},
}

// アイコン(=写真)を描画
func (t *Thumbnail) PutIcons() error {
	if t.Icons == nil {
		return errors.New("There is no icons.")
	}
	for _, icon := range t.Icons {
		if err := icon.DrawIconImage(t.Img); err != nil {
			return err
		}

	}
	return nil
}

func (icon *Icon) DrawIconImage(distImg *image.RGBA) error {
	// 貼り付ける画像をimage.Image型として取得
	iconImg, err := decodeImage(icon.ImageUrl)
	if err != nil {
		return err
	}
	// 合成
	centerPt := icon.C.Pt
	rad := icon.C.Rad
	place := image.Rect(
		centerPt.X-rad, // x0
		centerPt.Y-rad, // y0
		centerPt.X+rad, // x1
		centerPt.Y+rad, // y1
	)

	// 画像を丸く切り取るなど
	cutImage(iconImg, icon.C)

	// 画像貼り付け
	draw.Draw(distImg, place, *iconImg, image.Pt(0, 0), draw.Over)

	// 淵をつける
	icon.C.drawBounds(distImg, DARK_GREEN)
	return nil
}

// 画像を読み込み, リサイズ
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

// 画像を丸く切り取る
func cutImage(img *image.Image, circle Circle) {
	return
}

// 枠を書く
func (c *Circle) drawBounds(img *image.RGBA, col color.Color) {
	for rad := 0.0; rad < 2.0*float64(c.Rad); rad += 0.1 {
		for i := 0; i < 5; i++ {
			x := int(float64(c.Pt.X) + float64(c.Rad+i)*math.Cos(rad))
			y := int(float64(c.Pt.Y) + float64(c.Rad+i)*math.Sin(rad))
			img.Set(x, y, col)

		}
	}
}

// アイコンをセット
// ライブラリに書く必要は別にない...
func CreateIcons(urls []string) ([]*Icon, error) {

	var icons []*Icon

	// iconを配置する場所をランダムに決定
	points := icon_points[:]
	iconNum := len(urls)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < iconNum && i < MAX_ICON_NUM; i++ {
		icon := &Icon{}
		// 設置場所
		rnum := rand.Intn(len(points))
		icon.C = Circle{Rad: ICON_RADIUS, Pt: points[rnum]}
		points = append(points[:rnum], points[rnum+1:]...)

		// 画像
		rnum = rand.Intn(len(urls))
		icon.ImageUrl = urls[rnum]
		urls = append(urls[:rnum], urls[rnum+1:]...)

		icons = append(icons, icon)
	}

	return icons, nil
}
