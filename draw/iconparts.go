package draw

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"time"
)

type Icon struct {
	frame_color color.Color
	ImageUrl    string
	rect        image.Rectangle
	pic         *Picture
}

var x_width = IMG_SIZE / (MAX_ICON_NUM/2 + 1)
var y_width = IMG_SIZE / 6 //これは決め打ち, 間の帯のサイズ次第
var r = ICON_RADIUS

var iconPoints = [6]image.Rectangle{
	image.Rect(1*x_width-r, y_width-r, 1*x_width+r, y_width+r),
	image.Rect(2*x_width-r, y_width-r, 2*x_width+r, y_width+r),
	image.Rect(3*x_width-r, y_width-r, 3*x_width+r, y_width+r),
	image.Rect(1*x_width-r, 5*y_width-r, 1*x_width+r, 5*y_width+r),
	image.Rect(2*x_width-r, 5*y_width-r, 2*x_width+r, 5*y_width+r),
	image.Rect(3*x_width-r, 5*y_width-r, 3*x_width+r, 5*y_width+r),
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

	// アイコン画像生成
	pic, err := NewPicture(icon.ImageUrl)
	if err != nil {
		return err
	}
	icon.pic = pic

	draw.Draw(distImg, icon.rect, *pic.Img, image.Pt(0, 0), draw.Over)

	return nil
}

// アイコンをセット
// ライブラリに書く必要は別にない...
func CreateIcons(urls []string) ([]*Icon, error) {

	var icons []*Icon

	// iconを配置する場所をランダムに決定
	points := iconPoints[:]
	iconNum := len(urls)

	rand.Seed(time.Now().UnixNano())
	var rnum int
	for i := 0; i < iconNum && i < MAX_ICON_NUM; i++ {
		icon := &Icon{}
		// 設置場所
		rnum = rand.Intn(len(points))
		icon.rect = points[rnum]
		points = append(points[:rnum], points[rnum+1:]...)

		// 画像
		rnum = rand.Intn(len(urls))
		icon.ImageUrl = urls[rnum]
		urls = append(urls[:rnum], urls[rnum+1:]...)

		icons = append(icons, icon)
	}

	return icons, nil
}
