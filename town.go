package main

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"golang.org/x/image/colornames"
)

var (
	town1Button = pixel.Rect{}
	town2Button = pixel.Rect{}
	town3Button = pixel.Rect{}
	town4Button = pixel.Rect{}
	town5Button = pixel.Rect{}
	town6Button = pixel.Rect{}
)

func initTown(win *pixelgl.Window, Txt *text.Text) {

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "Where are you going?")
	tempPosition = myPos.TopCenterPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "1. WeaponShop")
	tempPosition = myPos.CenterLeftPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
	town1Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "2. ArmorShop")
	tempPosition = myPos.CenterPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
	town2Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "3. AccessoryShop")
	tempPosition = myPos.CenterRightPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
	town3Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "4. BlackSmith")
	tempPosition = myPos.BottleLeftPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
	town4Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "5. Equipment")
	tempPosition = myPos.BottleCenterPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
	town5Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "6. Exit")
	tempPosition = myPos.BottleRightPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
	town6Button = Txt.Bounds().Moved(tempPosition)
}

func townClickEvent(win *pixelgl.Window, mousePos pixel.Vec) GameState {
	//TODO ページを作成したら追加
	if town1Button.Contains(mousePos) || win.JustPressed(pixelgl.Key1) {
		currentGameState = WeaponShop
		log.Println("Town->WeaponShop")
	} else if town2Button.Contains(mousePos) || win.JustPressed(pixelgl.Key2) {
		currentGameState = ArmorShop
		log.Println("Town->ArmorShop")
	} else if town3Button.Contains(mousePos) || win.JustPressed(pixelgl.Key3) {
		currentGameState = AccessoryShop
		log.Println("Town->AccessoryShop")
	} else if town4Button.Contains(mousePos) || win.JustPressed(pixelgl.Key4) {
		currentGameState = BlackSmith
		log.Println("Town->BlackSmith")
	} else if town5Button.Contains(mousePos) || win.JustPressed(pixelgl.Key5) {
		currentGameState = EquipmentScreen
		log.Println("Town->EquipmentScreen")
	} else if town6Button.Contains(mousePos) || win.JustPressed(pixelgl.Key6) {
		currentGameState = GoToScreen
		log.Println("Town->GoToScreen")
	}
	return currentGameState
}
