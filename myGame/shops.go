package myGame

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

type ArmorState int
type AccessoryState int

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
	armorPath     = "assets/shop/armor.csv"
	descArmor     = CsvToSlice(armorPath)
	accessoryPath = "assets/shop/accessory.csv"
	descAccessory = CsvToSlice(accessoryPath)
)

var (
	buttonSlice  = []pixel.Rect{}
	buySellSlice = []pixel.Rect{}
)

var currentarmorState ArmorState
var currentaccessoryState AccessoryState

var tempInt []int

func InitArmor(win *pixelgl.Window, Txt *text.Text) {
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

func ArmorClickEvent(win *pixelgl.Window, mousePos pixel.Vec) GameState {
	//TODO ページを作成したら追加
	if CurrentGS == ArmorShop && (buttonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1)) {
		currentarmorState = armor1
		log.Println("ArmorShop->armor1")
	} else if CurrentGS == ArmorShop && (buttonSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.Key2)) {
		currentarmorState = armor2
		log.Println("ArmorShop->armor2")
	} else if CurrentGS == ArmorShop && (buttonSlice[2].Contains(mousePos) || win.JustPressed(pixelgl.Key3)) {
		currentarmorState = armor3
		log.Println("ArmorShop->armor3")
	} else if CurrentGS == ArmorShop && (buttonSlice[3].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace)) {
		currentarmorState = armor1
		CurrentGS = TownScreen
		log.Println("ArmorShop->TownScreen")
	}
	return CurrentGS
}

func InitAccessory(win *pixelgl.Window, Txt *text.Text) {
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

func AccessoryClickEvent(win *pixelgl.Window, mousePos pixel.Vec) GameState {
	//TODO ページを作成したら追加
	if CurrentGS == AccessoryShop && (buttonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1)) {
		currentaccessoryState = accessory1
		log.Println("AccessoryShop->accessory1")
	} else if CurrentGS == AccessoryShop && (buttonSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace)) {
		currentaccessoryState = accessory1
		CurrentGS = TownScreen
		log.Println("AccessoryShop->TownScreen")
	}
	return CurrentGS
}

func CsvToSlice(path string) [][]string {
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
