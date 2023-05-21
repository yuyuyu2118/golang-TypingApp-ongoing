package main

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myPos"
	"golang.org/x/image/colornames"
)

var (
	equip1Button = pixel.Rect{}
	// equip2Button = pixel.Rect{}
	// equip3Button = pixel.Rect{}
)

func initEquipment(win *pixelgl.Window, Txt *text.Text) {

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "Where are you going?")
	tempPosition = myPos.TopCenterPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "1. GoToScreen")
	tempPosition = myPos.CenterLeftPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
	equip1Button = Txt.Bounds().Moved(tempPosition)
}

func equipmentClickEvent(win *pixelgl.Window, mousePos pixel.Vec) myGame.GameState {
	//TODO ページを作成したら追加
	if equip1Button.Contains(mousePos) || win.JustPressed(pixelgl.Key1) {
		myGame.CurrentGS = myGame.GoToScreen
		log.Println("equipment->GoToScreen")
	}
	return myGame.CurrentGS
}

// func initPlayerEquipment(win *pixelgl.Window, Txt *text.Text, player *PlayerStatus) {
// 	Txt.Clear()
// 	Txt.Color = colornames.White
// 	fmt.Fprintln(Txt, "Weapon: ", "\nArmor: ", "\nAccessory: ")
// 	xOffSet := 0.0
// 	yOffSet := win.Bounds().H()/3 - Txt.LineHeight*3
// 	txtPos := pixel.V(xOffSet, yOffSet)
// 	tempPosition := pixel.IM.Moved(txtPos)
// 	Txt.Draw(win, tempPosition)
// }
