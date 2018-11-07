package main

import (
	"image/jpeg"
	"os"

	"github.com/CaterpillarSan/DrawImg/draw"
	"github.com/CaterpillarSan/DrawImg/model"
	"gopkg.in/guregu/null.v3"
)

func main() {

	cards := []model.Card{
		// model.Card{ImageUrl: null.NewString("./pictures/0.png", true), EmoID: 2},
		// model.Card{ImageUrl: null.NewString("./pictures/1.png", true), EmoID: 2},
		// model.Card{ImageUrl: null.NewString("./pictures/2.png", true), EmoID: 3},
		// model.Card{ImageUrl: null.NewString("./pictures/3.png", true), EmoID: 4},
		// model.Card{ImageUrl: null.NewString("./pictures/4.png", true), EmoID: 5},
		// model.Card{ImageUrl: null.NewString("./pictures/5.png", true), EmoID: 1},
		// model.Card{ImageUrl: null.NewString("./pictures/6.png", true), EmoID: 2},
		// model.Card{ImageUrl: null.NewString("./pictures/7.png", true), EmoID: 3},
		// model.Card{ImageUrl: null.NewString("./pictures/8.png", true), EmoID: 4},
		// model.Card{ImageUrl: null.NewString("./pictures/9.png", true), EmoID: 5},
		// model.Card{ImageUrl: null.NewString("./pictures/10.jpg", true), EmoID: 1},
		// model.Card{ImageUrl: null.NewString("./pictures/11.png", true), EmoID: 2},
		model.Card{ImageUrl: null.NewString("./pictures/emo.jpg", true), EmoID: 3},
	}

	title := "We are Yakudo"
	// title := "a"
	// title := "12345678901234567890123456789012"
	var err error

	t := draw.NewThumbnail(title, cards)
	if err = t.FillRect(); err != nil {
		panic(err)
	}
	if err = t.PutIcons(); err != nil {
		panic(err)
	}
	if err = t.SetTitle(); err != nil {
		panic(err)
	}

	emos := []int{1, 2, 3, 4}
	t.DrawColorBar(emos)
	// 出力用ファイル作成(エラー処理は略)
	file, _ := os.Create("sample.jpg")
	defer file.Close()

	// JPEGで出力(100%品質)
	if err := jpeg.Encode(file, t.Img, &jpeg.Options{100}); err != nil {
		panic(err)
	}
}
