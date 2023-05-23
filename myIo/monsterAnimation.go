package myIo

import (
	"fmt"
	"image"
	_ "image/png"
	"io/ioutil"
	"os"
	"strings"

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

func GroupSlimeImages(directory string) [][]string {
	var (
		waitImages   []string
		attackImages []string
		slimeImages  [][]string
	)

	// ディレクトリ内のファイル一覧を取得
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return slimeImages
	}

	// ファイル名をスライスに分類
	for _, file := range files {
		fileName := file.Name()
		if strings.HasPrefix(fileName, "SlimeA_Wait") {
			waitImages = append(waitImages, fileName)
		} else if strings.HasPrefix(fileName, "SlimeA_Attack") {
			attackImages = append(attackImages, fileName)
		}
	}

	slimeImages = append(slimeImages, waitImages)
	slimeImages = append(slimeImages, attackImages)

	return slimeImages
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
