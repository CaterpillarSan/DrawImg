package draw

import (
	"errors"
	"image"
	"math"
	"math/rand"
	"time"
)

// 使わない

type Boad struct {
	points     [BOAD_SIZE][BOAD_SIZE]bool
	titlePoint string
}

/*
ここにbit配列の説明を
*/

func NewBoad() Boad {
	b := Boad{}
	// bitmap的な
	points := [BOAD_SIZE][BOAD_SIZE]bool{}
	// タイトルを表示する位置を決める
	rand.Seed(time.Now().UnixNano())
	switch rand.Intn(2) {
	case 0:
		b.titlePoint = "header"
	default:
		b.titlePoint = "footer"
	}
	// boad初期化
	// 外周はfalse
	for i := 0; i < BOAD_SIZE; i++ {
		for j := 0; j < BOAD_SIZE; j++ {
			// 0 ~ 3, 15 ~ 18
			if i < OUT_OF_BOAD || i >= BOAD_SIZE-OUT_OF_BOAD ||
				j < OUT_OF_BOAD || j >= BOAD_SIZE-OUT_OF_BOAD {
				points[i][j] = false
			} else if b.titlePoint == "header" && j < OUT_OF_BOAD+2 {
				points[i][j] = false

			} else if b.titlePoint == "footer" && j > BOAD_SIZE-OUT_OF_BOAD-3 {
				points[i][j] = false
			} else {
				points[i][j] = true
			}
		}
	}

	b.points = points
	return b
}

func (b *Boad) atRect(x, y, size int) image.Rectangle {
	// size 1, 2, 3, 4 (* 2)
	if size < 1 || size > 4 {
		size = 2
	}
	x0 := (x - OUT_OF_BOAD - size) * WIDTH
	y0 := (y - OUT_OF_BOAD - size) * WIDTH
	x1 := (x - OUT_OF_BOAD + size) * WIDTH
	y1 := (y - OUT_OF_BOAD + size) * WIDTH

	return image.Rect(x0, y0, x1, y1)
}

func (b *Boad) fillBoad(x, y, size int) {
	for i := -size; i <= size; i++ {
		num := size - int(math.Abs(float64(i)))
		for j := -num; j <= num; j++ {
			b.points[x+i][y+j] = false
		}
	}
}

func (b *Boad) GetRandomRect() (*image.Rectangle, int, error) {

	rand.Seed(time.Now().UnixNano())

	// サイズ大きい順に置ければ置く
	// size:
	// - 4 -> 6回
	// - 3 -> 12回
	// - 2 -> 18回
	// - 1 -> 24回
	// 挑戦する
	for size := 4; size > 0; size-- {
		for i := 0; i < (5-size)*6; i++ {
			rx := rand.Intn(BOAD_SIZE-OUT_OF_BOAD*2) + OUT_OF_BOAD
			ry := rand.Intn(BOAD_SIZE-OUT_OF_BOAD*2) + OUT_OF_BOAD
			if b.isAbleToPut(rx, ry, size) {
				b.fillBoad(rx, ry, size)
				rect := b.atRect(rx, ry, size)
				return &rect, size, nil
			}
		}
	}

	return nil, 1, errors.New("Cannot put icon anymore.")
}

// Boadにおけるか
// はみ出るものも弾いている
// 今後どうしようか...
func (b *Boad) isAbleToPut(x, y, size int) bool {
	for i := -size; i <= size; i++ {
		num := size - int(math.Abs(float64(i)))
		for j := -num; j <= num; j++ {
			if b.points[x+i][y+j] == false {
				return false
			}
		}
	}
	return true
}
