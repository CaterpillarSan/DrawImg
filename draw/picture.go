package draw

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/CaterpillarSan/DrawImg/model"
)

type Picture struct {
	FrameColor color.Color
	ImageUrl   string
	rect       image.Rectangle
	img        *image.Image
}

// cardsから, サムネに使用する写真データをランダムに抜き出す
func NewPicList(cards []model.Card) []*Picture {

	if len(cards) == 0 {
		return []*Picture{}
	}

	var pictures []*Picture

	rand.Seed(time.Now().UnixNano())

	// 画像を4枚まで選出
	for i := 0; i < 4; i++ {
		pic := &Picture{}

		// 画像
		rnum := rand.Intn(len(cards))
		pic.ImageUrl = cards[rnum].ImageUrl.String
		cards = append(cards[:rnum], cards[rnum+1:]...)

		pictures = append(pictures, pic)
		if len(cards) == 0 {
			break
		}
	}

	return pictures
}

// pic型の情報をもとに, image型を作成
func (pic *Picture) ToPicture(rect image.Rectangle) error {
	// イメージを取り込み,image.Imageに変換
	img, err := pic.decodeImage()
	if err != nil {
		return err
	}

	// いい感じにリサイズ
	pic.img = resizeSquare(*img, rect)
	return nil
}

func (pic *Picture) decodeImage() (*image.Image, error) {
	u, err := url.Parse(pic.ImageUrl)
	if err != nil {
		return nil, err
	}
	var originImg image.Image

	if u.Scheme == "https" && u.Host == "s3-ap-northeast-1.amazonaws.com" {
		// Get image
		originImg, err = getImageFromUrl(pic.ImageUrl)
	} else {
		// デバッグ用, ローカルのファイルを取り出す
		originImg, err = getImageFromLocal(pic.ImageUrl)
		// return nil, errors.New("Unauthorized URL") // TODO : エラーちゃんとする
	}
	if err != nil {
		return nil, err
	}

	return &originImg, nil
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

// 写真の描画
func (t *Thumbnail) DrawPictures() error {
	if t.Pics == nil {
		return nil
	}

	picNum := len(t.Pics)
	if picNum == 0 {
		// picが一枚もなかった
		return nil
	}

	// Iconの枚数によってIcon.rectが決まる
	for i, icon := range t.Pics {
		if err := icon.drawPicImage(t.Img, picNum, i); err != nil {
			return err
		}

	}
	return nil
}

func (pic *Picture) drawPicImage(distImg *image.RGBA, picNum int, picID int) error {
	var rect image.Rectangle
	if picID == 0 {
		// 1枚目 IMG_SIZE * IMG_SIZE か IMG_SIZE * 0.625
		if picNum == 1 {
			// 画像が全部で1枚
			rect = image.Rect(0, 0, IMG_SIZE, IMG_SIZE)
		} else {
			// 画像が複数枚ある時の, 最初の一枚
			rect = image.Rect(0, 0, LEFT_PIC_WIDTH, IMG_SIZE)
		}
	} else {
		if picNum == 1 {
			// 0割を避けるエラー処理
			return errors.New("pictureの数が合わない")
		}
		// 2枚目以降
		width := IMG_SIZE / (picNum - 1)
		rect = image.Rect(LEFT_PIC_WIDTH, (picID-1)*width, IMG_SIZE, picID*width)
	}

	// アイコン画像生成
	if err := pic.ToPicture(rect); err != nil {
		return err
	}
	draw.Draw(distImg, rect, *pic.img, image.ZP, draw.Src)

	return nil
}
