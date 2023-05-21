package myIo

import (
	"image"
	"os"

	"github.com/faiface/pixel"
	"github.com/pkg/errors"
)

// in:画像パス out:pixelで使えるpictureData
func OpenDecodePictureData(sheetPath string) (sheet pixel.Picture, err error) {
	sheetFile, err := os.Open(sheetPath)
	if err != nil {
		err = errors.Wrap(err, "This error is os.Open error")
		return nil, err
	}
	defer sheetFile.Close()

	sheetImg, _, err := image.Decode(sheetFile)
	if err != nil {
		err = errors.Wrap(err, "This error is image.Decode error")
		return nil, err
	}
	sheet = pixel.PictureDataFromImage(sheetImg)

	return sheet, nil
}

// //in:pictureDataとframe幅 out:2D画像スライス
// func makePixelRectSlice2D(sheet pixel.Picture, frameWidth float64) []pixel.Rect {
// 	var frames []pixel.Rect
// 	for x := 0.0; x+frameWidth <= sheet.Bounds().Max.X; x += frameWidth {
// 		frames = append(frames, pixel.R(
// 			x,
// 			0,
// 			x+frameWidth,       //この関数の第三引数 つまりフレーム毎の幅
// 			sheet.Bounds().H(), //固定長のsheetの高さを取得 つまりフレーム毎の高さ
// 		))
// 	}
// 	return frames
// }

// //in:csvのパスと2D画像スライス out:アニメーション情報
// func csvOpenReadCreateAnimation(descPath string, frames []pixel.Rect) (anims map[string][]pixel.Rect, err error) {
// 	descFile, err := os.Open(descPath)
// 	anims = make(map[string][]pixel.Rect)

// 	if err != nil {
// 		err = errors.Wrap(err, "This error is os.Open error")
// 		return nil, err
// 	}
// 	defer descFile.Close()

// 	desc := csv.NewReader(descFile)
// 	for {
// 		anim, err := desc.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			err = errors.Wrap(err, "This error is csvRead error")
// 			return nil, err
// 		}
// 		name := anim[0]
// 		start, _ := strconv.Atoi(anim[1])
// 		end, _ := strconv.Atoi(anim[2])
// 		anims[name] = frames[start : end+1]
// 	}
// 	return anims, nil
// }

// // in:画像パス csvパス フレーム幅 out:アニメーションの画像とデータ
// func loadAnimationSheet(sheetPath, descPath string, frameWidth float64) (sheet pixel.Picture, anims map[string][]pixel.Rect) {

// 	sheet, err := openDecodePictureData(sheetPath)
// 	checkErrorPanic(err)

// 	frames := makePixelRectSlice2D(sheet, frameWidth)

// 	anims, err = csvOpenReadCreateAnimation(descPath, frames)
// 	checkErrorPanic(err)

// 	return sheet, anims
// }
