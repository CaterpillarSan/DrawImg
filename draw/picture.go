package draw

import (
	"bytes"
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

	var pictures []*Picture

	rand.Seed(time.Now().UnixNano())

	// for i := 0; i < 4; i++ {
	pic := &Picture{}

	// 画像
	rnum := rand.Intn(len(cards))
	pic.ImageUrl = cards[rnum].ImageUrl.String
	pic.FrameColor = BAR_COLOR[cards[rnum].EmoID]
	cards = append(cards[:rnum], cards[rnum+1:]...)

	pictures = append(pictures, pic)
	// }

	return pictures
}

// pic型の情報をもとに, image型を作成
func (pic *Picture) ToPicture() error {
	// イメージを取り込み,image.Imageに変換
	img, err := pic.decodeImage()
	if err != nil {
		return err
	}

	// いい感じにリサイズ
	pic.img = resizeSquare(*img, IMG_SIZE)
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
	if t.Icons == nil {
		return nil
	}

	// TODO 設置場所を決める, どこで決めるべきか...
	// Iconの枚数によってIcon.rectが決まる
	for _, icon := range t.Icons {
		if err := icon.drawPicImage(t.Img); err != nil {
			return err
		}

	}
	return nil
}

func (pic *Picture) drawPicImage(distImg *image.RGBA) error {

	// アイコン画像生成
	if err := pic.ToPicture(); err != nil {
		return err
	}
	draw.Draw(distImg, distImg.Rect, *pic.img, image.ZP, draw.Src)

	return nil
}
