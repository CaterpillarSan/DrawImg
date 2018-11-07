package draw

import "image/color"

const IMG_SIZE = 512

const OUT_OF_BOAD = 4 // = num of type
const ACTIVE_FIEAD = 11
const BOAD_SIZE = OUT_OF_BOAD*2 + ACTIVE_FIEAD

const WIDTH = IMG_SIZE / ACTIVE_FIEAD
const COLBAR_HEIGHT = IMG_SIZE / 25

const (
	DIAMOND = 0
	CIRCLE  = 1
)

const FONT_PATH = "./metadata/fonts/07LogoTypeGothic-Condense.ttf"
const BACKGROUND_PATH = "./metadata/backgrounds/"

// ****************************** //
// 色シリーズ

var MUSYOKU_TOUMEI = color.RGBA{0, 0, 0, 0}

// Emo ID関連
// $emo1-color: #dedede
// $emo2-color: #f9c7db
// $emo3-color: #c3ffd1
// $emo4-color: #cdf9ff
// $emo5-color: #ffb3b3
// https://coolors.co/dedede-f9c7db-c3ffd1-cdf9ff-ffb3b3
var GRAY = color.RGBA{222, 123, 222, 150}       // #dedede
var PINK = color.RGBA{249, 119, 219, 150}       // #f9c7db
var LIME = color.RGBA{195, 255, 209, 150}       // #c3ffd1
var LIGHT_CYAN = color.RGBA{205, 249, 255, 150} // #cdf9ff
var LIGHT_PINK = color.RGBA{255, 179, 179, 150} // #ffb3b3

var BAR_COLOR = []color.Color{
	GRAY,
	PINK,
	LIME,
	LIGHT_CYAN,
	LIGHT_PINK,
}

var TITLE_BACKGROUND = color.RGBA{0, 0, 0, 100}
var TITLE_CHAR = color.RGBA{255, 255, 255, 255}

// blue
// https://coolors.co/2f0601-554a41-26c485-a3e7fc-32908f
var BLACK_BEAN = color.RGBA{47, 6, 1, 255}               //#2F0601
var DARK_PUCE = color.RGBA{85, 58, 65, 255}              //#554A41
var MOUNTAIN_MEADOW = color.RGBA{38, 196, 133, 255}      //#26C485
var FRESH_AIR = color.RGBA{163, 231, 252, 255}           //#A3E7FC
var ILLUMINATING_EMERALD = color.RGBA{50, 144, 143, 255} //#32908F

// green
// https://coolors.co/0a3200-fffbbd-379634-74f2ce-e78f8e
var DARK_GREEN = color.RGBA{10, 50, 0, 255}           //#0A3200
var VERY_PALE_YELLOW = color.RGBA{255, 251, 189, 255} //#FFFBBD
var WAGENINGEN_GREEN = color.RGBA{55, 150, 52, 255}   //#379634
var AQUAMARINE = color.RGBA{116, 242, 206, 255}       //#74F2CE
var RUDDY_PINK = color.RGBA{231, 143, 142, 255}       //#E78F8E

// pink
// https://coolors.co/32213a-dd9ac2-d5d5d5-fffc36-4a1942
var DARK_PURPLE = color.RGBA{50, 33, 58, 255}       //#32213A
var KOBI = color.RGBA{221, 154, 194, 255}           //#DD9AC2
var LIGHT_GRAY = color.RGBA{213, 213, 213, 255}     //#D5D5D5
var ELECTRIC_YELLOW = color.RGBA{255, 252, 54, 255} //#FFFC36
var RUSSIAN_VIOKET = color.RGBA{74, 25, 66, 255}    //#4A1942

// white
// https://coolors.co/011936-ffb0ff-acf46f-86dcff-ff86b3
var MAASTRICHT_BLUE = color.RGBA{1, 25, 54, 255}       //#011936
var ELECTRIC_LAVENDER = color.RGBA{255, 176, 255, 255} //#FFB0FF
var INCHWORM = color.RGBA{172, 244, 111, 255}          //#ACF46F
var PALE_CYAN = color.RGBA{134, 220, 255, 255}         //#86DCFF
var TICKLE_ME_PINK = color.RGBA{255, 134, 179, 255}    //#FF86B3

//yellow
// https://coolors.co/432818-ffe6a7-b8b42d-bdd9bf-6f1d1b
var BISTRE = color.RGBA{67, 40, 24, 255}         //#432818
var NAVAO_WHITE = color.RGBA{255, 230, 167, 255} //#FFE6A7
var OLD_GOLD = color.RGBA{184, 180, 45, 255}     //#B8B42D
var SILVER = color.RGBA{189, 217, 191, 255}      //#BDD9BF
var PRUNE = color.RGBA{111, 29, 27, 255}         //#6F1D1B

// TODO Mapからキー一覧生成すべき?
var BKColors = []string{"blue", "green", "pink", "white", "yellow"}
var ColorMap = map[string][]color.Color{
	"blue": []color.Color{
		BLACK_BEAN,
		DARK_PUCE,
		MOUNTAIN_MEADOW,
		FRESH_AIR,
		ILLUMINATING_EMERALD,
	},
	"green": []color.Color{
		DARK_GREEN,
		VERY_PALE_YELLOW,
		WAGENINGEN_GREEN,
		AQUAMARINE,
		RUDDY_PINK,
	},
	"pink": []color.Color{
		DARK_PURPLE,
		KOBI,
		LIGHT_GRAY,
		ELECTRIC_YELLOW,
		RUSSIAN_VIOKET,
	},
	"white": []color.Color{
		MAASTRICHT_BLUE,
		ELECTRIC_LAVENDER,
		INCHWORM,
		PALE_CYAN,
		TICKLE_ME_PINK,
	},
	"yellow": []color.Color{
		BISTRE,
		NAVAO_WHITE,
		OLD_GOLD,
		SILVER,
		PRUNE,
	},
}
