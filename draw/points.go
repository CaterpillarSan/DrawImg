package draw

import (
	"errors"
	"image"
	"math"
	"math/rand"
	"time"
)

type Boad struct {
	points [BOAD_SIZE][BOAD_SIZE]bool
}

/*
ここにbit配列の説明を
*/

func NewBoad() Boad {
	// bitmap的な
	points := [BOAD_SIZE][BOAD_SIZE]bool{}
	// boad初期化
	// 外周はfalse
	for i := 0; i < BOAD_SIZE; i++ {
		for j := 0; j < BOAD_SIZE; j++ {
			// 0 ~ 3, 15 ~ 18
			if i < OUT_OF_BOAD || i >= BOAD_SIZE-OUT_OF_BOAD ||
				j < OUT_OF_BOAD || j >= BOAD_SIZE-OUT_OF_BOAD {
				points[i][j] = false
			} else {
				points[i][j] = true
			}
		}
	}

	return Boad{points}
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

	// 30回でいいやろ~~
	for i := 0; i < 20; i++ {
		rx := rand.Intn(BOAD_SIZE-OUT_OF_BOAD*2) + OUT_OF_BOAD
		ry := rand.Intn(BOAD_SIZE-OUT_OF_BOAD*2) + OUT_OF_BOAD
		rnum := rand.Intn(10)
		// iconsizeのランダムな決定
		var rsize int
		switch {
		case rnum == 0:
			rsize = 1
		case rnum < 4:
			rsize = 2
		case rnum < 8:
			rsize = 3
		default:
			rsize = 4
		}
		if b.isAbleToPut(rx, ry, rsize) {
			b.fillBoad(rx, ry, rsize)
			rect := b.atRect(rx, ry, rsize)
			return &rect, rsize, nil
		}
	}

	// size =2 が置けるかな
	for i := 0; i < 10; i++ {
		rx := rand.Intn(BOAD_SIZE-OUT_OF_BOAD*2) + OUT_OF_BOAD
		ry := rand.Intn(BOAD_SIZE-OUT_OF_BOAD*2) + OUT_OF_BOAD
		rsize := 2
		if b.isAbleToPut(rx, ry, rsize) {
			b.fillBoad(rx, ry, rsize)
			rect := b.atRect(rx, ry, rsize)
			return &rect, rsize, nil
		}
	}
	// size = 1が置けるかな
	for i := 0; i < 10; i++ {
		rx := rand.Intn(BOAD_SIZE-OUT_OF_BOAD*2) + OUT_OF_BOAD
		ry := rand.Intn(BOAD_SIZE-OUT_OF_BOAD*2) + OUT_OF_BOAD
		rsize := 1
		if b.isAbleToPut(rx, ry, rsize) {
			b.fillBoad(rx, ry, rsize)
			rect := b.atRect(rx, ry, rsize)
			return &rect, rsize, nil
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
