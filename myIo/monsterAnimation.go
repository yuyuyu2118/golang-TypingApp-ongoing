package myIo

import (
	"fmt"
	"image"
	"os"

	"github.com/faiface/pixel"
)

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
