# DrawImg
Goで画像を生成し, S3に上がっている写真と任意の文字列を合成し, S3にあげたい.

## しんちょく
![しんちょく](https://github.com/CaterpillarSan/DrawImg/blob/master/sample.jpg "現在のmasterブランチの生成物")

## 構成
- main.go : メイン. ファイル出力もここで.
- draw package : 一連の画像生成周りをここに実装します
- sample.jpg : `run main.go` で生成される画像サンプル

## TODO
 - [ ] 土台, 背景生成
 - [ ] Circle, Rectangleを画像内の任意の位置に配置
 - [ ] ローカルの画像を取り込み, 合成
 - [ ] 取り込んだ画像のリサイズ?
 - [ ] 取り込んだ画像に枠をつけたり...
 - [ ] S3からGETした画像を取り込めるように
 - [ ] 生成した画像をローカルではなくS3に投げられるように
 - [ ] 背景に自由度を持たせる (設計方針要相談)
 - [ ] フィルターなど, 画像生成のバリエーションを増やす(nice to have)
 
 
## 考えること
- 背景を複数種類用意するとして, 以下のどちらか, またはより用方法があるか
  - 色の組み合わせセットだけ用意し, 画像生成時に色を塗る.
  ```
   for h := rect.Min.Y; h < rect.Max.Y; h++ {
		for v := rect.Min.X; v < rect.Max.X; v++ {
			if h/boder == 1 {
				t.Img.Set(v, h, c1)
			} else {
				t.Img.Set(v, h, c2)
			}
		}
	}
  ```
  - 色を塗った状態のimageオブジェクトを複数種類用意しておき, 複製して使う
- for文とメモリコピー走るのどっちが重いか...
