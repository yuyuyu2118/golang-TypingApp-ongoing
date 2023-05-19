package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

type WeaponState int
type ArmorState int
type AccessoryState int

const (
	weapon1 WeaponState = iota
	weapon2
	weapon3
)
const (
	armor1 ArmorState = iota
	armor2
	armor3
)
const (
	accessory1 AccessoryState = iota
	accessory2
	accessory3
)

var (
	weaponPath    = "assets/shop/weapon.csv"
	descWeapon    = csvToSlice(weaponPath)
	armorPath     = "assets/shop/armor.csv"
	descArmor     = csvToSlice(armorPath)
	accessoryPath = "assets/shop/accessory.csv"
	descAccessory = csvToSlice(accessoryPath)
)

var (
	buttonSlice = []pixel.Rect{}
)
var currentweaponState WeaponState
var currentarmorState ArmorState
var currentaccessoryState AccessoryState

func initWeapon(win *pixelgl.Window, Txt *text.Text, windowHeightSize int) {
	//初期位置
	yOffSet := win.Bounds().H() / 4
	txtPos := pixel.V(0, 0)
	Txt.Color = colornames.White
	//csv読み込み

	weaponSlice := []string{"1. Wooden Stick", "2. Fruit Knife", "3. WoodenSword", "BackSpace. EXIT"}

	for _, weaponName := range weaponSlice {
		Txt.Clear()
		fmt.Fprintln(Txt, weaponName)
		yOffSet -= Txt.LineHeight + 20
		txtPos = pixel.V(0, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
		buttonSlice = append(buttonSlice, Txt.Bounds().Moved(txtPos))
	}

	xOffSet := win.Bounds().W()/3 - 200
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
		for _, value := range descWeapon[0] {
			Txt.Clear()
			fmt.Fprintln(Txt, value)
			yOffSet -= Txt.TabWidth + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition = pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	case weapon2:
		for _, value := range descWeapon[1] {
			Txt.Clear()
			fmt.Fprintln(Txt, value)
			yOffSet -= Txt.TabWidth + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition = pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	case weapon3:
		for _, value := range descWeapon[2] {
			Txt.Clear()
			fmt.Fprintln(Txt, value)
			yOffSet -= Txt.TabWidth + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition = pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	}
}

func weaponClickEvent(win *pixelgl.Window, mousePos pixel.Vec) GameState {
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

func initArmor(win *pixelgl.Window, Txt *text.Text, windowHeightSize int) {
	//初期位置
	yOffSet := win.Bounds().H() / 4
	txtPos := pixel.V(0, 0)
	Txt.Color = colornames.White
	//csv読み込み

	armorSlice := []string{"1. Pot Lid", "2. Leather Shield", "3. Silver Shield", "BackSpace. EXIT"}

	for _, armorName := range armorSlice {
		Txt.Clear()
		fmt.Fprintln(Txt, armorName)
		yOffSet -= Txt.LineHeight + 20
		txtPos = pixel.V(0, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
		buttonSlice = append(buttonSlice, Txt.Bounds().Moved(txtPos))
	}

	xOffSet := win.Bounds().W()/3 - 200
	yOffSet = win.Bounds().H() / 4
	var tempPosition pixel.Matrix
	if win.Pressed(pixelgl.Key1) {
		currentarmorState = armor1
	} else if win.Pressed(pixelgl.Key2) {
		currentarmorState = armor2
	} else if win.Pressed(pixelgl.Key3) {
		currentarmorState = armor3
	}

	switch currentarmorState {
	case armor1:
		for _, value := range descArmor[0] {
			Txt.Clear()
			fmt.Fprintln(Txt, value)
			yOffSet -= Txt.TabWidth + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition = pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	case armor2:
		for _, value := range descArmor[1] {
			Txt.Clear()
			fmt.Fprintln(Txt, value)
			yOffSet -= Txt.TabWidth + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition = pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	case armor3:
		for _, value := range descArmor[2] {
			Txt.Clear()
			fmt.Fprintln(Txt, value)
			yOffSet -= Txt.TabWidth + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition = pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	}
}

func armorClickEvent(win *pixelgl.Window, mousePos pixel.Vec) GameState {
	//TODO ページを作成したら追加
	if buttonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1) {
		currentarmorState = armor1
		log.Println("ArmorShop->armor1")
	} else if buttonSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.Key2) {
		currentarmorState = armor2
		log.Println("ArmorShop->armor2")
	} else if buttonSlice[2].Contains(mousePos) || win.JustPressed(pixelgl.Key3) {
		currentarmorState = armor3
		log.Println("ArmorShop->armor3")
	} else if buttonSlice[3].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace) {
		currentarmorState = armor1
		currentGameState = TownScreen
		log.Println("ArmorShop->TownScreen")
	}
	return currentGameState
}

func initAccessory(win *pixelgl.Window, Txt *text.Text, windowHeightSize int) {
	//初期位置
	yOffSet := win.Bounds().H() / 4
	txtPos := pixel.V(0, 0)
	Txt.Color = colornames.White
	//csv読み込み

	accessorySlice := []string{"1. Copper Bracelet", "BackSpace. EXIT"}

	for _, accessoryName := range accessorySlice {
		Txt.Clear()
		fmt.Fprintln(Txt, accessoryName)
		yOffSet -= Txt.LineHeight + 20
		txtPos = pixel.V(0, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
		buttonSlice = append(buttonSlice, Txt.Bounds().Moved(txtPos))
	}

	xOffSet := win.Bounds().W()/3 - 200
	yOffSet = win.Bounds().H() / 4
	var tempPosition pixel.Matrix
	if win.Pressed(pixelgl.Key1) {
		currentaccessoryState = accessory1
	} else if win.Pressed(pixelgl.Key2) {
		currentaccessoryState = accessory2
	} else if win.Pressed(pixelgl.Key3) {
		currentaccessoryState = accessory3
	}

	switch currentaccessoryState {
	case accessory1:
		for _, value := range descAccessory[0] {
			Txt.Clear()
			fmt.Fprintln(Txt, value)
			yOffSet -= Txt.TabWidth + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition = pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	}
}

func accessoryClickEvent(win *pixelgl.Window, mousePos pixel.Vec) GameState {
	//TODO ページを作成したら追加
	if buttonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1) {
		currentaccessoryState = accessory1
		log.Println("AccessoryShop->accessory1")
	} else if buttonSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace) {
		currentaccessoryState = accessory1
		currentGameState = TownScreen
		log.Println("AccessoryShop->TownScreen")
	}
	return currentGameState
}

func csvToSlice(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}
