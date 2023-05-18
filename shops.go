package main

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

var (
	woodenKatanaButton = pixel.Rect{}
	ironKatanaButton   = pixel.Rect{}
	goldKatanaButton   = pixel.Rect{}
	buttonSlice        = []pixel.Rect{}
)

func initWeapon(win *pixelgl.Window, Txt *text.Text, windowHeightSize int) {
	//初期位置
	yOffSet := win.Bounds().H() / 4
	txtPos := pixel.V(0, 0)
	Txt.Color = colornames.White
	//csv読み込み
	weaponSlice := []string{"1. WoodenKatana", "2. IronKatana", "3. goldKatanaButton", "4", "5", "6", "7", "8", "9", "10"}
	descWeapon1 := []string{"WoodenKatana", "OP : 3", "この剣はサンプルの剣です\n説明はサンプルです。", "強化回数", "未強化", "強化上昇値", "特殊能力", "この剣の特殊能力は？です。\nサンプルです。"}

	for _, weaponName := range weaponSlice {
		Txt.Clear()
		fmt.Fprintln(Txt, weaponName)
		yOffSet -= Txt.LineHeight + 20
		txtPos = pixel.V(0, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
		buttonSlice = append(buttonSlice, Txt.Bounds().Moved(txtPos))
	}

	xOffSet := win.Bounds().W() / 3
	var tempPosition pixel.Matrix
	if win.Pressed(pixelgl.Key1) {
		Txt.Clear()
		fmt.Fprintln(Txt, descWeapon1[0])
		xOffSet = win.Bounds().W() / 3
		yOffSet = win.Bounds().H() / 4
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition = pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)

		Txt.Clear()
		fmt.Fprintln(Txt, descWeapon1[1])
		xOffSet = win.Bounds().W() / 3
		yOffSet -= Txt.TabWidth + 50
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition = pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)

		Txt.Clear()
		fmt.Fprintln(Txt, descWeapon1[2])
		xOffSet = win.Bounds().W() / 3
		yOffSet -= Txt.TabWidth + 100
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition = pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
	}
}

func weaponClickEvent(win *pixelgl.Window, mousePos pixel.Vec, currentGameState GameState) GameState {
	//TODO ページを作成したら追加
	if buttonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1) {
		currentGameState = WeaponShop
		log.Println("Town->WeaponShop")
	} else if buttonSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.Key2) {
		currentGameState = ArmorShop
		log.Println("Town->ArmorShop")
	} else if buttonSlice[2].Contains(mousePos) || win.JustPressed(pixelgl.Key3) {
		currentGameState = AccessoryShop
		log.Println("Town->AccessoryShop")
	}
	return currentGameState
}
