package myGame

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"golang.org/x/image/colornames"
)

type WeaponState int
type ArmorState int
type AccessoryState int

const (
	weapon1 WeaponState = iota
	weapon2
	weapon3
	weapon4
	weapon5
	weapon6
	weapon7
	weapon8
	weapon9
	weapon0
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
	descWeapon    = CsvToSlice(weaponPath)
	armorPath     = "assets/shop/armor.csv"
	descArmor     = CsvToSlice(armorPath)
	accessoryPath = "assets/shop/accessory.csv"
	descAccessory = CsvToSlice(accessoryPath)
)

var (
	buttonSlice = []pixel.Rect{}
)
var currentweaponState WeaponState
var currentarmorState ArmorState
var currentaccessoryState AccessoryState

func InitWeapon(win *pixelgl.Window, Txt *text.Text, topText string) {
	xOffSet := 100.0
	yOffSet := myPos.TopLefPos(win, Txt).Y - 100
	txtPos := pixel.V(0, 0)
	Txt.Color = colornames.White
	//csv読み込み

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, topText)
	tempPosition = myPos.TopCenPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)

	weaponSlice := []string{"1. 木の棒", "2. 果物ナイフ", "3. 木刀", "4. ドレインソード", "5. スタンハンマー", "6. 鉄の剣", "7. 隼の剣", "8. 勇者の剣", "9. 名刀村正", "0. 死神の大鎌", "BackSpace. EXIT"}

	for _, weaponName := range weaponSlice {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, weaponName)
		yOffSet -= Txt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
		buttonSlice = append(buttonSlice, Txt.Bounds().Moved(txtPos))
	}

	if win.Pressed(pixelgl.Key1) {
		currentweaponState = weapon1
	} else if win.Pressed(pixelgl.Key2) {
		currentweaponState = weapon2
	} else if win.Pressed(pixelgl.Key3) {
		currentweaponState = weapon3
	} else if win.Pressed(pixelgl.Key4) {
		currentweaponState = weapon4
	} else if win.Pressed(pixelgl.Key5) {
		currentweaponState = weapon5
	} else if win.Pressed(pixelgl.Key6) {
		currentweaponState = weapon6
	} else if win.Pressed(pixelgl.Key7) {
		currentweaponState = weapon7
	} else if win.Pressed(pixelgl.Key8) {
		currentweaponState = weapon8
	} else if win.Pressed(pixelgl.Key9) {
		currentweaponState = weapon9
	} else if win.Pressed(pixelgl.Key0) {
		currentweaponState = weapon0
	}

	xOffSet = myPos.TopLefPos(win, Txt).X + 300
	yOffSet = myPos.TopLefPos(win, Txt).Y - 50
	switch currentweaponState {
	case weapon1:
		for _, value := range descWeapon[0] {
			Txt.Clear()
			Txt.Color = colornames.White
			fmt.Fprintln(Txt, value)
			yOffSet -= Txt.LineHeight + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition := pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	case weapon2:
		for _, value := range descWeapon[1] {
			Txt.Clear()
			Txt.Color = colornames.White
			fmt.Fprintln(Txt, value)
			yOffSet -= Txt.LineHeight + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition := pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	case weapon3:
		for _, value := range descWeapon[2] {
			Txt.Clear()
			Txt.Color = colornames.White
			fmt.Fprintln(Txt, value)
			yOffSet -= Txt.LineHeight + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition := pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	case weapon4:
		for _, value := range descWeapon[3] {
			Txt.Clear()
			Txt.Color = colornames.White
			fmt.Fprintln(Txt, value)
			yOffSet -= Txt.LineHeight + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition := pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	case weapon5:
		for _, value := range descWeapon[4] {
			Txt.Clear()
			Txt.Color = colornames.White
			fmt.Fprintln(Txt, value)
			yOffSet -= Txt.LineHeight + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition := pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	case weapon6:
		for _, value := range descWeapon[5] {
			Txt.Clear()
			Txt.Color = colornames.White
			fmt.Fprintln(Txt, value)
			yOffSet -= Txt.LineHeight + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition := pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	case weapon7:
		for _, value := range descWeapon[6] {
			Txt.Clear()
			Txt.Color = colornames.White
			fmt.Fprintln(Txt, value)
			yOffSet -= Txt.LineHeight + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition := pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	case weapon8:
		for _, value := range descWeapon[7] {
			Txt.Clear()
			Txt.Color = colornames.White
			fmt.Fprintln(Txt, value)
			yOffSet -= Txt.LineHeight + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition := pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	case weapon9:
		for _, value := range descWeapon[8] {
			Txt.Clear()
			Txt.Color = colornames.White
			fmt.Fprintln(Txt, value)
			yOffSet -= Txt.LineHeight + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition := pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	case weapon0:
		for _, value := range descWeapon[9] {
			Txt.Clear()
			Txt.Color = colornames.White
			fmt.Fprintln(Txt, value)
			yOffSet -= Txt.LineHeight + 20
			txtPos = pixel.V(xOffSet, yOffSet)
			tempPosition := pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	}
}

func WeaponClickEvent(win *pixelgl.Window, mousePos pixel.Vec) GameState {
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
	} else if buttonSlice[3].Contains(mousePos) || win.JustPressed(pixelgl.Key4) {
		currentweaponState = weapon4
		log.Println("WeaponShop->weapon4")
	} else if buttonSlice[4].Contains(mousePos) || win.JustPressed(pixelgl.Key5) {
		currentweaponState = weapon5
		log.Println("WeaponShop->weapon5")
	} else if buttonSlice[5].Contains(mousePos) || win.JustPressed(pixelgl.Key6) {
		currentweaponState = weapon6
		log.Println("WeaponShop->weapon6")
	} else if buttonSlice[6].Contains(mousePos) || win.JustPressed(pixelgl.Key7) {
		currentweaponState = weapon7
		log.Println("WeaponShop->weapon7")
	} else if buttonSlice[7].Contains(mousePos) || win.JustPressed(pixelgl.Key8) {
		currentweaponState = weapon8
		log.Println("WeaponShop->weapon8")
	} else if buttonSlice[8].Contains(mousePos) || win.JustPressed(pixelgl.Key9) {
		currentweaponState = weapon9
		log.Println("WeaponShop->weapon9")
	} else if buttonSlice[9].Contains(mousePos) || win.JustPressed(pixelgl.Key0) {
		currentweaponState = weapon0
		log.Println("WeaponShop->weapon0")
	} else if buttonSlice[10].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace) {
		currentweaponState = weapon1
		CurrentGS = TownScreen
		log.Println("WeaponShop->TownScreen")
	}
	return CurrentGS
}

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
	if buttonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1) {
		currentaccessoryState = accessory1
		log.Println("AccessoryShop->accessory1")
	} else if buttonSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace) {
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
