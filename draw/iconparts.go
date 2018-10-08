package draw

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"
	"time"
)

type Icon struct {
	frame_color color.Color
	image_url   string
	C           Circle
}

type Circle struct {
	r int
	p image.Point
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

func (icon *Icon) DrawIconImage(img *image.RGBA) error {
	icon.C.DrawBounds(img, DARK_GREEN)
	return nil
}

// 枠を書く
func (c *Circle) DrawBounds(img *image.RGBA, col color.Color) {
	for rad := 0.0; rad < 2.0*float64(c.r); rad += 0.1 {
		for i := 0; i < 5; i++ {
			x := int(float64(c.p.X) + float64(c.r+i)*math.Cos(rad))
			y := int(float64(c.p.Y) + float64(c.r+i)*math.Sin(rad))
			img.Set(x, y, col)

		}
	}
}

// アイコンをセット
// ライブラリに書く必要は別にない...
func CreateIcons(num int) ([]*Icon, error) {
	if num > MAX_ICON_NUM {
		return nil, errors.New("cannot create icons over 6")
	}

	var icons []*Icon

	// iconを配置する場所をランダムに決定
	var points = icon_points[:]

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		icon := &Icon{}
		rnum := rand.Intn(len(points))
		icon.C = Circle{r: IMG_SIZE / 10, p: points[rnum]}
		fmt.Println(icon.C)
		points = unset(points, rnum)
		icons = append(icons, icon)
	}

	return icons, nil
}

// TODO 末尾削除もできる? GCちゃんと効く?
func unset(arr []image.Point, i int) []image.Point {
	return append(arr[:i], arr[i+1:]...)
}
