package enemy

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myIo"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myUtil"
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

func SetEnemySprite(win *pixelgl.Window, frame int) {
	sprites := EnemySprites[myGame.StageNum]
	if myGame.StageNum == 0 {
		sprites[frame].Draw(win, pixel.IM.Moved(win.Bounds().Center().Add(pixel.V(0, 20))).Scaled(win.Bounds().Center(), EnemySettings[myGame.StageNum].EnemySize))
	} else if myGame.StageNum >= 1 && myGame.StageNum <= 2 {
		sprites[frame].Draw(win, pixel.IM.Moved(win.Bounds().Center().Add(pixel.V(0, 30))).Scaled(win.Bounds().Center(), EnemySettings[myGame.StageNum].EnemySize))
	} else if myGame.StageNum >= 3 && myGame.StageNum <= 4 {
		sprites[frame].Draw(win, pixel.IM.Moved(win.Bounds().Center().Add(pixel.V(0, 55))).Scaled(win.Bounds().Center(), EnemySettings[myGame.StageNum].EnemySize))
	} else if myGame.StageNum >= 5 && myGame.StageNum <= 6 {
		sprites[frame].Draw(win, pixel.IM.Moved(win.Bounds().Center().Add(pixel.V(0, 35))).Scaled(win.Bounds().Center(), EnemySettings[myGame.StageNum].EnemySize))
	} else if myGame.StageNum == 7 {
		sprites[frame].Draw(win, pixel.IM.Moved(win.Bounds().Center().Add(pixel.V(0, 50))).Scaled(win.Bounds().Center(), EnemySettings[myGame.StageNum].EnemySize))
	} else if myGame.StageNum == 8 {
		sprites[frame].Draw(win, pixel.IM.Moved(win.Bounds().Center().Add(pixel.V(0, 55))).Scaled(win.Bounds().Center(), EnemySettings[myGame.StageNum].EnemySize))
	} else if myGame.StageNum == 9 {
		sprites[frame].Draw(win, pixel.IM.Moved(win.Bounds().Center().Add(pixel.V(0, 55))).Scaled(win.Bounds().Center(), EnemySettings[myGame.StageNum].EnemySize))
	}

	pic, _ := myIo.OpenDecodePictureData(EnemyPathBar[myGame.StageNum])
	scaledSize := pic.Bounds().Size().Scaled(EnemySettings[myGame.StageNum].EnemySize)
	barPosition := pixel.V(
		sprites[0].Picture().Bounds().W()*EnemySettings[myGame.StageNum].EnemySize,
		sprites[0].Picture().Bounds().H()*EnemySettings[myGame.StageNum].EnemySize,
	)

	SetEnemyHPBarOut(win, scaledSize, barPosition)
	SetEnemyHPBar(win, scaledSize, EnemySettings[myGame.StageNum].HP, EnemySettings[myGame.StageNum].MaxHP, barPosition)
}

func SetEnemySpriteText(win *pixelgl.Window, Txt *text.Text, enemy *EnemyStatus) {
	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "HP: ", strconv.FormatFloat(enemy.HP, 'f', 1, 64))
	myPos.DrawPos(win, Txt, myPos.TopCenPos(win, Txt).Add(pixel.V(0, 0)))
}

func StartEnemyAnimation(win *pixelgl.Window, last *time.Time, frame *int) {
	dt := time.Since(*last).Seconds()
	if dt >= 0.2 { // アニメーション速度を調整 (ここでは0.2秒ごとに更新)
		*frame = (*frame + 1) % len(EnemySprites[myGame.StageNum])
		(*last) = time.Now()
	}
	SetEnemySprite(win, *frame)
	SetEnemySpriteText(win, myUtil.ScreenTxt, &EnemySettings[myGame.StageNum])
}
