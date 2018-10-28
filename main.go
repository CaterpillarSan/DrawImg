package main

import (
	"image/jpeg"
	"os"

	"github.com/CaterpillarSan/DrawImg/draw"
)

func main() {

	urls := []string{
		"./pictures/0.png",
		"./pictures/1.png",
		"./pictures/2.png",
		"./pictures/3.png",
		"./pictures/4.png",
		"./pictures/5.png",
		"./pictures/6.png",
		"./pictures/7.png",
		"./pictures/8.png",
		"./pictures/9.png",
		"./pictures/10.jpg",
		"./pictures/11.png",
	}

	title := "We are Yakudo!"
	var err error

	t := draw.NewThumbnail(title, urls)
	t.FillRect()
	if err = t.PutIcons(); err != nil {
		panic(err)
	}
	if err = t.SetTitle(); err != nil {
		panic(err)
	}

	// 出力用ファイル作成(エラー処理は略)
	file, _ := os.Create("sample.jpg")
	defer file.Close()

	// JPEGで出力(100%品質)
	if err := jpeg.Encode(file, t.Img, &jpeg.Options{100}); err != nil {
		panic(err)
	}
}
