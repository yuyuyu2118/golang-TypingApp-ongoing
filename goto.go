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
	goTo1Button = pixel.Rect{}
	goTo2Button = pixel.Rect{}
	goTo3Button = pixel.Rect{}
	goTo4Button = pixel.Rect{}
	goTo5Button = pixel.Rect{}
	goTo6Button = pixel.Rect{}
)

func initGoTo(win *pixelgl.Window, Txt *text.Text) {

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "Where are you going?")
	tempPosition = myPos.TopCenterPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "1. dungeon")
	tempPosition = myPos.CenterLeftPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
	goTo1Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "2. Town")
	tempPosition = myPos.CenterPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
	goTo2Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "3. Equipment")
	tempPosition = myPos.CenterRightPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
	goTo3Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "4. Job")
	tempPosition = myPos.BottleLeftPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
	goTo4Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "5. Save")
	tempPosition = myPos.BottleCenterPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
	goTo5Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "6.Exit")
	tempPosition = myPos.BottleRightPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
	goTo6Button = Txt.Bounds().Moved(tempPosition)
}

func goToClickEvent(win *pixelgl.Window, mousePos pixel.Vec) GameState {
	//TODO ページを作成したら追加
	if goTo1Button.Contains(mousePos) || win.JustPressed(pixelgl.Key1) {
		currentGameState = StageSelect
		log.Println("GoToScreen->Dungeon")
	} else if goTo2Button.Contains(mousePos) || win.JustPressed(pixelgl.Key2) {
		currentGameState = TownScreen
		log.Println("GoToScreen->Town")
	} else if goTo3Button.Contains(mousePos) || win.JustPressed(pixelgl.Key3) {
		currentGameState = EquipmentScreen
		log.Println("GoToScreen->Equipment")
	} else if goTo4Button.Contains(mousePos) || win.JustPressed(pixelgl.Key4) {
		currentGameState = JobSelect
		log.Println("GoToScreen->JobSelect")
	} else if goTo5Button.Contains(mousePos) || win.JustPressed(pixelgl.Key5) {
		currentGameState = SaveScreen
		log.Println("GoToScreen->Save")
	} else if goTo6Button.Contains(mousePos) || win.JustPressed(pixelgl.Key6) {
		currentGameState = StartScreen
		log.Println("GoToScreen->StartScreen")
	}
	return currentGameState
}
