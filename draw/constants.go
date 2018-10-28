package draw

import "image/color"

const IMG_SIZE = 500

const MAX_ICON_NUM = 10

const OUT_OF_BOAD = 4 // = num of type
const BOAD_SIZE = OUT_OF_BOAD*2 + 11

const WIDTH = IMG_SIZE / 10

const (
	DIAMOND = iota
	CIRCLE
)

const FONT_PATH = "./fonts/07LogoTypeGothic-Condense.ttf"
const FONT_SIZE = IMG_SIZE / 12

// ****************************** //
// 色シリーズ

var DARK_GREEN = color.RGBA{17, 75, 95, 200}     // #114B5F
var LIGHT_GREEN = color.RGBA{214, 253, 153, 255} // #D6F599
var SHIBU_GREEN = color.RGBA{168, 194, 86, 255}  // #A8C256
var MUSYOKU_TOUMEI = color.RGBA{0, 0, 0, 0}      // #A8C256

// Emo ID関連
// $emo1-color: #dedede
// $emo2-color: #f9c7db
// $emo3-color: #c3ffd1
// $emo4-color: #cdf9ff
// $emo5-color: #ffb3b3
var GRAY = color.RGBA{222, 222, 222, 100}       // #dedede
var PINK = color.RGBA{249, 119, 219, 200}       // #f9c7db
var LIME = color.RGBA{195, 255, 209, 200}       // #c3ffd1
var LIGHT_CYAN = color.RGBA{205, 249, 255, 200} // #cdf9ff
var LIGHT_PINK = color.RGBA{255, 179, 179, 200} // #ffb3b3

var EmoColor = []color.Color{GRAY, PINK, LIME, LIGHT_CYAN, LIGHT_PINK}
