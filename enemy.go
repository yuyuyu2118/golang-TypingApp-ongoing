package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type EnemyStatus struct {
	enemyMaxHP float64
	enemyHP    float64
	enemyOP    float64
	enemyDP    float64
	enemyGold  int
	enemyName  string
}

func newEnemyStatus(MaxHP float64, HP float64, OP float64, DP float64, Gold int, Name string) *EnemyStatus {
	return &EnemyStatus{MaxHP, HP, OP, DP, Gold, Name}
}

func setEnemyHPBar(win *pixelgl.Window, scaledSize pixel.Vec, HP float64, MaxHP float64, pos pixel.Vec) {
	rectWidth := scaledSize.X * ((MaxHP - (MaxHP - HP)) * 0.01)
	rect := pixel.R(
		win.Bounds().Center().X-(rectWidth/2),
		win.Bounds().Center().Y-50,
		win.Bounds().Center().X+(rectWidth/2),
		win.Bounds().Center().Y,
	)
	imd := imdraw.New(nil)
	imd.Color = colornames.Green
	imd.Push(rect.Min, rect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func setEnemyHPBarOut(win *pixelgl.Window, scaledSize pixel.Vec, pos pixel.Vec) {
	rect := pixel.R(
		win.Bounds().Center().X-pos.X/2,
		win.Bounds().Center().Y-50,
		win.Bounds().Center().X+pos.X/2,
		win.Bounds().Center().Y,
	)
	imd := imdraw.New(nil)
	imd.Color = colornames.Red
	imd.Push(rect.Min, rect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func setEnemyPic(win *pixelgl.Window, enemyInf *EnemyStatus, path string, scaleFactor float64) {
	pic, _ := openDecodePictureData(path)
	picMonster := pixel.NewSprite(pic, pic.Bounds())

	scaledSize := pic.Bounds().Size().Scaled(scaleFactor)
	transMat := pixel.IM.Moved(win.Bounds().Center().Add(pixel.V(0, 35))).Scaled(win.Bounds().Center(), scaleFactor)
	picMonster.Draw(win, transMat)

	barPosition := pixel.V(
		picMonster.Picture().Bounds().W()*scaleFactor,
		picMonster.Picture().Bounds().H()*scaleFactor,
	)

	pic.Bounds()

	setEnemyHPBarOut(win, scaledSize, barPosition)
	setEnemyHPBar(win, scaledSize, enemyInf.enemyHP, enemyInf.enemyMaxHP, barPosition)
}
