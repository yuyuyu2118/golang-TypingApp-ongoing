package myUtil

import (
	"image/color"
	"io/ioutil"
	"os"
	"unicode"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"github.com/yuyuyu2118/typingGo/myIo"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
)

var (
	StartTxt        *text.Text
	ScreenTxt       *text.Text
	DescriptionTxt  *text.Text
	BasicTxt        *text.Text
	HunterBulletTxt *text.Text
)

func InitTxtFontLoading() {
	//fontPath := "assets\\fonts\\NotoSans-Black.ttf"
	japanFontPath := "assets/fonts/mplus-1c-black.ttf"
	japanFontPathBold := "assets/fonts/mplus-1c-bold.ttf"
	BasicTxt = initializeAnyText(japanFontPath, 40, colornames.White)
	StartTxt = initAnyJapanText(japanFontPathBold, 70, colornames.White)
	ScreenTxt = initAnyJapanText(japanFontPath, 40, colornames.White)
	DescriptionTxt = initAnyJapanText(japanFontPath, 30, colornames.White)
	HunterBulletTxt = initAnyJapanText(japanFontPathBold, 60, colornames.White)
	// startTxt := initializeAnyText(fontPath, 80, colornames.White)
	// endTxt := initializeAnyText(fontPath, 60, colornames.White)
}

func initializeText(face font.Face, color color.Color) *text.Text {
	basicAtlas := text.NewAtlas(face, text.ASCII)
	basicTxt := text.New(pixel.V(50, 500), basicAtlas)
	basicTxt.Color = color
	return basicTxt
}

func initializeAnyText(fontPath string, size int, color color.Color) *text.Text {
	face, _ := LoadTTF(fontPath, float64(size))
	return initializeText(face, color)
}

func initText(face font.Face, color color.Color) *text.Text {
	Atlas := text.NewAtlas(face, text.ASCII)
	Txt := text.New(pixel.V(0, 0), Atlas)
	Txt.Color = color
	return Txt
}

func initAnyText(fontPath string, size int, color color.Color) *text.Text {
	face, _ := LoadTTF(fontPath, float64(size))
	return initText(face, color)
}

func initJapanText(face font.Face, color color.Color) *text.Text {
	//TODO: ここが重い
	var customKanjiRunes []rune
	runes := myIo.CsvToSlice1Line("assets/fonts/kanji.csv")
	for _, r := range runes {
		customKanjiRunes = append(customKanjiRunes, rune(r))
	}
	customKanji := CustomRangeTable(customKanjiRunes)

	Atlas := text.NewAtlas(face, text.ASCII, text.RangeTable(unicode.P),
		text.RangeTable(unicode.Hiragana), text.RangeTable(unicode.Katakana),
		text.RangeTable(customKanji), text.RangeTable(CustomRangeTable([]rune{'ー'})))
	Txt := text.New(pixel.V(0, 0), Atlas)
	return Txt
}

func initAnyJapanText(fontPath string, size int, color color.Color) *text.Text {
	face := LoadJapanFont(fontPath, float64(size))
	return initJapanText(face, color)
}

func LoadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}

func LoadJapanFont(fontPath string, size float64) font.Face {
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	tt, err := truetype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}

	face := truetype.NewFace(tt, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	return face
}

func CustomRangeTable(runes []rune) *unicode.RangeTable {
	ranges := make([]unicode.Range16, len(runes))
	for i, r := range runes {
		ranges[i] = unicode.Range16{
			Lo:     uint16(r),
			Hi:     uint16(r),
			Stride: 1,
		}
	}
	return &unicode.RangeTable{R16: ranges}
}

func DrawCenteredText(win *pixelgl.Window, atlas *text.Atlas, line1, line2 string) {
	japanFontPath := "assets/fonts/mplus-1c-black.ttf"
	tempTxt := initAnyJapanText(japanFontPath, 40, colornames.White)
	tempTxt.Color = pixel.RGB(1, 1, 1) // テキストの色を設定

	tempTxt.WriteString(line1)
	line1Width := tempTxt.BoundsOf(line1).W()
	tempTxt.Dot.X -= line1Width / 2

	winX, winY := win.Bounds().Max.X, win.Bounds().Max.Y
	tempTxt.Orig.X = (winX - line1Width) / 2
	tempTxt.Orig.Y = winY/2 + atlas.LineHeight()/2

	tempTxt.Draw(win, pixel.IM.Scaled(tempTxt.Orig, 1))

	tempTxt.Clear()
	tempTxt.WriteString(line2)
	line2Width := tempTxt.BoundsOf(line2).W()

	tempTxt.Dot.X -= line2Width / 2
	tempTxt.Orig.X = (winX - line2Width) / 2
	tempTxt.Orig.Y = winY/2 - atlas.LineHeight()/2

	tempTxt.Draw(win, pixel.IM.Scaled(tempTxt.Orig, 1))
}
