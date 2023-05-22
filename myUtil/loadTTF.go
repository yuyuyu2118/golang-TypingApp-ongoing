package myUtil

import (
	"io/ioutil"
	"os"
	"unicode"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

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
