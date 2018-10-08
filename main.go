package main

import (
	"image/jpeg"
	"os"

	"github.com/CaterpillarSan/DrawImg/draw"
)

func main() {

	t := draw.NewThumbnail()
	// RectからRGBAを作る(ゼロ値なので黒なはず)

	// 出力用ファイル作成(エラー処理は略)
	file, _ := os.Create("sample.jpg")
	defer file.Close()

	// JPEGで出力(100%品質)
	if err := jpeg.Encode(file, t.Img, &jpeg.Options{100}); err != nil {
		panic(err)
	}
}
