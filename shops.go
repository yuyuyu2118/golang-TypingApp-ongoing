package main

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

type WeaponState int

const (
	weapon1 WeaponState = iota
	weapon2
	weapon3
)

var (
	buttonSlice = []pixel.Rect{}
)
var currentweaponState WeaponState

func initWeapon(win *pixelgl.Window, Txt *text.Text, windowHeightSize int) {
	//初期位置
	yOffSet := win.Bounds().H() / 4
	txtPos := pixel.V(0, 0)
	Txt.Color = colornames.White
	//csv読み込み
	weaponSlice := []string{"1. Wooden Stick", "2. Fruit Knife", "3. WoodenSword", "BackSpace. EXIT"}
	descWeapon1 := []string{"WoodenStick", "Attack Power: 1", "", "Description:", " A simple weapon made from a branch.", " Its attack power is very low.", "", "Unique Abilities: None"}
	descWeapon2 := []string{"Fruit Knife", "Attack Power: 3", "", "Description:", " A small knife used to cut vegetables", "and fruits.", " It is not very strong as a weapon.", "", "Unique Abilities: None"}
	descWeapon3 := []string{"Wooden Sword", "Attack Power: 5", "", "Description:", " A wooden sword used for", "practice and training.", "It is not suitable for actual combat,", "but it is light and easy to handle.", "", "Unique Abilities: None"}

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
	yOffSet = win.Bounds().H() / 4
	var tempPosition pixel.Matrix
	if win.Pressed(pixelgl.Key1) {
		currentweaponState = weapon1
	} else if win.Pressed(pixelgl.Key2) {
		currentweaponState = weapon2
	} else if win.Pressed(pixelgl.Key3) {
		currentweaponState = weapon3
	}
	switch currentweaponState {
	case weapon1:
		for _, value := range descWeapon1 {
			Txt.Clear()
			fmt.Fprintln(Txt, value)
			xOffSet = win.Bounds().W() / 3
			yOffSet -= Txt.TabWidth + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition = pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	case weapon2:
		for _, value := range descWeapon2 {
			Txt.Clear()
			fmt.Fprintln(Txt, value)
			xOffSet = win.Bounds().W() / 3
			yOffSet -= Txt.TabWidth + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition = pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	case weapon3:
		for _, value := range descWeapon3 {
			Txt.Clear()
			fmt.Fprintln(Txt, value)
			xOffSet = win.Bounds().W() / 3
			yOffSet -= Txt.TabWidth + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition = pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	}
}

func weaponClickEvent(win *pixelgl.Window, mousePos pixel.Vec, currentGameState GameState) GameState {
	//TODO ページを作成したら追加
	if buttonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1) {
		currentweaponState = weapon1
		log.Println("WeaponShop->weapon1")
	} else if buttonSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.Key2) {
		currentweaponState = weapon2
		log.Println("WeaponShop->weapon2")
	} else if buttonSlice[2].Contains(mousePos) || win.JustPressed(pixelgl.Key3) {
		currentweaponState = weapon3
		log.Println("WeaponShop->weapon3")
	} else if buttonSlice[3].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace) {
		currentweaponState = weapon1
		currentGameState = TownScreen
		log.Println("WeaponShop->TownScreen")
	}
	return currentGameState
}
