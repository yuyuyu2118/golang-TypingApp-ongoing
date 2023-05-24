package enemy

import (
	"fmt"
	"path/filepath"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myIo"
	"github.com/yuyuyu2118/typingGo/myPos"
	"golang.org/x/image/colornames"
)

func SetEnemyAnimation(directory string, fileName string) []*pixel.Sprite {
	// ディレクトリ内にある画像ファイルを検索する
	matches, err := filepath.Glob(filepath.Join(directory, fileName+"*.png"))
	if err != nil {
		panic(err)
	}

	var sprites []*pixel.Sprite
	for _, path := range matches {
		picture, err := myIo.LoadPicture(path)
		if err != nil {
			panic(err)
		}
		sprite := pixel.NewSprite(picture, picture.Bounds())
		sprites = append(sprites, sprite)
	}
	return sprites
}

func SetEnemySprite(win *pixelgl.Window, enemyInf *EnemyStatus, path string, sprites []*pixel.Sprite, frame int) {
	sprites[frame].Draw(win, pixel.IM.Moved(win.Bounds().Center().Add(pixel.V(0, 25))).Scaled(win.Bounds().Center(), enemyInf.EnemySize))
	pic, _ := myIo.OpenDecodePictureData(path)
	scaledSize := pic.Bounds().Size().Scaled(enemyInf.EnemySize)
	barPosition := pixel.V(
		sprites[0].Picture().Bounds().W()*enemyInf.EnemySize,
		sprites[0].Picture().Bounds().H()*enemyInf.EnemySize,
	)

	SetEnemyHPBarOut(win, scaledSize, barPosition)
	SetEnemyHPBar(win, scaledSize, enemyInf.HP, enemyInf.MaxHP, barPosition)
}

func SetEnemySpriteText(win *pixelgl.Window, Txt *text.Text, enemy *EnemyStatus) {
	// cp := constantProvider{}
	// WinHSize := cp.GetConstant()
	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "EnemyHP : ", enemy.HP)
	myPos.DrawPos(win, Txt, myPos.TopCenPos(win, Txt))
}
