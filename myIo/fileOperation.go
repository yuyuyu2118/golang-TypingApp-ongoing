package myIo

import (
	"fmt"
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

func LoadSpriteSheet(imagePaths []string) ([]*pixel.Sprite, error) {
	var sprites []*pixel.Sprite

	for _, path := range imagePaths {
		imgFile, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("failed to open image: %w", err)
		}
		defer imgFile.Close()

		img, _, err := image.Decode(imgFile)
		if err != nil {
			return nil, fmt.Errorf("failed to decode image: %w", err)
		}

		picData := pixel.PictureDataFromImage(img)
		sprite := pixel.NewSprite(picData, picData.Bounds())

		sprites = append(sprites, sprite)
	}

	return sprites, nil
}

func LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}
