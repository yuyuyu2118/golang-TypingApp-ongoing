package myGame

import (
	"fmt"
	"log"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	event "github.com/yuyuyu2118/typingGo/Event"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/player"
	"golang.org/x/image/colornames"
)

type WeaponState int

const (
	weaponNil WeaponState = iota
	weapon1
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

var weaponSlice = []string{"1. ???", "2. ???", "3. ???", "4. ???", "5. ???", "6. ???", "7. ???", "8. ???", "9. ???", "0. ???", "BackSpace. EXIT"}
var weaponNum = []string{"weapon0", "weapon1", "weapon2", "weapon3", "weapon4", "weapon5", "weapon6", "weapon7", "weapon8", "weapon9"}
var weaponName = []string{"木の棒", "果物ナイフ", "木刀", "ドレインソード", "スタンハンマー", "鉄の剣", "隼の剣", "勇者の剣", "名刀村正", "死神の大鎌"}

var (
	weaponPath = "assets/shop/weapon.csv"
	descWeapon = CsvToSlice(weaponPath)
)
var currentweaponState WeaponState

func InitWeapon(win *pixelgl.Window, Txt *text.Text, topText string) {
	xOffSet := 100.0
	yOffSet := myPos.TopLefPos(win, Txt).Y - 100
	txtPos := pixel.V(0, 0)
	Txt.Color = colornames.White

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, topText)
	tempPosition = myPos.TopCenPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)

	if event.WeaponPurchaseEventInstance.Weapon1 {
		weaponSlice[0] = "1. 木の棒"
	}
	if event.WeaponPurchaseEventInstance.Weapon2 {
		weaponSlice[1] = "2. 果物ナイフ"
	}
	if event.WeaponPurchaseEventInstance.Weapon3 {
		weaponSlice[2] = "3. 木刀"
	}
	if event.WeaponPurchaseEventInstance.Weapon4 {
		weaponSlice[3] = "4. ドレインソード"
	}
	if event.WeaponPurchaseEventInstance.Weapon5 {
		weaponSlice[4] = "5. スタンハンマー"
	}
	if event.WeaponPurchaseEventInstance.Weapon6 {
		weaponSlice[5] = "6. 鉄の剣"
	}
	if event.WeaponPurchaseEventInstance.Weapon7 {
		weaponSlice[6] = "7. 隼の剣"
	}
	if event.WeaponPurchaseEventInstance.Weapon8 {
		weaponSlice[7] = "8. 勇者の剣"
	}
	if event.WeaponPurchaseEventInstance.Weapon9 {
		weaponSlice[8] = "9. 名刀村正"
	}
	if event.WeaponPurchaseEventInstance.Weapon0 {
		weaponSlice[9] = "0. 死神の大鎌"
	}
	//weaponSlice := []string{"1. 木の棒", "2. 果物ナイフ", "3. 木刀", "4. ドレインソード", "5. スタンハンマー", "6. 鉄の剣", "7. 隼の剣", "8. 勇者の剣", "9. 名刀村正", "0. 死神の大鎌", "BackSpace. EXIT"}

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

	if win.Pressed(pixelgl.Key1) && event.WeaponPurchaseEventInstance.Weapon1 {
		currentweaponState = weapon1
	} else if win.Pressed(pixelgl.Key2) && event.WeaponPurchaseEventInstance.Weapon2 {
		currentweaponState = weapon2
	} else if win.Pressed(pixelgl.Key3) && event.WeaponPurchaseEventInstance.Weapon3 {
		currentweaponState = weapon3
	} else if win.Pressed(pixelgl.Key4) && event.WeaponPurchaseEventInstance.Weapon4 {
		currentweaponState = weapon4
	} else if win.Pressed(pixelgl.Key5) && event.WeaponPurchaseEventInstance.Weapon5 {
		currentweaponState = weapon5
	} else if win.Pressed(pixelgl.Key6) && event.WeaponPurchaseEventInstance.Weapon6 {
		currentweaponState = weapon6
	} else if win.Pressed(pixelgl.Key7) && event.WeaponPurchaseEventInstance.Weapon7 {
		currentweaponState = weapon7
	} else if win.Pressed(pixelgl.Key8) && event.WeaponPurchaseEventInstance.Weapon8 {
		currentweaponState = weapon8
	} else if win.Pressed(pixelgl.Key9) && event.WeaponPurchaseEventInstance.Weapon9 {
		currentweaponState = weapon9
	} else if win.Pressed(pixelgl.Key0) && event.WeaponPurchaseEventInstance.Weapon0 {
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
	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "B. 買う")
	yOffSet -= Txt.LineHeight + 20
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition := pixel.IM.Moved(txtPos)
	Txt.Draw(win, tempPosition)
	buySellSlice = append(buySellSlice, Txt.Bounds().Moved(txtPos))
	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "S. 売る")
	xOffSet += Txt.TabWidth + 100
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	Txt.Draw(win, tempPosition)
	buySellSlice = append(buySellSlice, Txt.Bounds().Moved(txtPos))
}

func WeaponClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *player.PlayerStatus) myState.GameState {
	var tempWeapon string
	//TODO ページを作成したら追加
	if (buttonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1)) && event.WeaponPurchaseEventInstance.Weapon1 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon1
		log.Println("WeaponShop->weapon1")
	} else if (buttonSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.Key2)) && event.WeaponPurchaseEventInstance.Weapon2 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon2
		log.Println("WeaponShop->weapon2")
	} else if (buttonSlice[2].Contains(mousePos) || win.JustPressed(pixelgl.Key3)) && event.WeaponPurchaseEventInstance.Weapon3 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon3
		log.Println("WeaponShop->weapon3")
	} else if (buttonSlice[3].Contains(mousePos) || win.JustPressed(pixelgl.Key4)) && event.WeaponPurchaseEventInstance.Weapon4 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon4
		log.Println("WeaponShop->weapon4")
	} else if (buttonSlice[4].Contains(mousePos) || win.JustPressed(pixelgl.Key5)) && event.WeaponPurchaseEventInstance.Weapon5 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon5
		log.Println("WeaponShop->weapon5")
	} else if (buttonSlice[5].Contains(mousePos) || win.JustPressed(pixelgl.Key6)) && event.WeaponPurchaseEventInstance.Weapon6 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon6
		log.Println("WeaponShop->weapon6")
	} else if (buttonSlice[6].Contains(mousePos) || win.JustPressed(pixelgl.Key7)) && event.WeaponPurchaseEventInstance.Weapon7 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon7
		log.Println("WeaponShop->weapon7")
	} else if (buttonSlice[7].Contains(mousePos) || win.JustPressed(pixelgl.Key8)) && event.WeaponPurchaseEventInstance.Weapon8 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon8
		log.Println("WeaponShop->weapon8")
	} else if (buttonSlice[8].Contains(mousePos) || win.JustPressed(pixelgl.Key9)) && event.WeaponPurchaseEventInstance.Weapon9 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon9
		log.Println("WeaponShop->weapon9")
	} else if (buttonSlice[9].Contains(mousePos) || win.JustPressed(pixelgl.Key0)) && event.WeaponPurchaseEventInstance.Weapon0 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon0
		log.Println("WeaponShop->weapon0")
	} else if buttonSlice[10].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace) && myState.CurrentGS == myState.WeaponShop {
		myState.CurrentGS = myState.TownScreen
		log.Println("WeaponShop->TownScreen")
	}
	if (buySellSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.KeyB)) && player.Gold >= 100 {
		if currentweaponState == weapon1 {
			//TODO: ファイルからweaponのインスタンス作成、weaponインスタンスからGoldを参照
			player.Gold -= 100
			tempWeapon = "weapon1"
		} else if currentweaponState == weapon2 {
			player.Gold -= 300
			tempWeapon = "weapon2"
		} else if currentweaponState == weapon3 {
			player.Gold -= 500
			tempWeapon = "weapon3"
		} else if currentweaponState == weapon4 {
			player.Gold -= 1500
			tempWeapon = "weapon4"
		} else if currentweaponState == weapon5 {
			player.Gold -= 3000
			tempWeapon = "weapon5"
		} else if currentweaponState == weapon6 {
			player.Gold -= 2000
			tempWeapon = "weapon6"
		} else if currentweaponState == weapon7 {
			player.Gold -= 4000
			tempWeapon = "weapon7"
		} else if currentweaponState == weapon8 {
			player.Gold -= 6000
			tempWeapon = "weapon8"
		} else if currentweaponState == weapon9 {
			player.Gold -= 8000
			tempWeapon = "weapon9"
		} else if currentweaponState == weapon0 {
			player.Gold -= 10000
			tempWeapon = "weapon0"
		}
		SaveWeaponPurchaseEvent(SaveFilePath, 3, tempWeapon, player)
		SaveGame(SaveFilePath, 1, player)
	} else if buySellSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.KeyS) {
		log.Println("Sell")
		for _, value := range player.PossessedWeapon {
			tempPossessedWeapon, _ := strconv.Atoi(value)
			tempInt = append(tempInt, tempPossessedWeapon)
		}
		if currentweaponState == weapon1 && tempInt[0] >= 1 {
			//TODO: ファイルからweaponのインスタンス作成、weaponインスタンスからGoldを参照
			player.Gold += 50
			tempWeapon = "weapon1"
		} else if currentweaponState == weapon2 && tempInt[1] >= 1 {
			player.Gold += 150
			tempWeapon = "weapon2"
		} else if currentweaponState == weapon3 && tempInt[2] >= 1 {
			player.Gold += 250
			tempWeapon = "weapon3"
		} else if currentweaponState == weapon4 && tempInt[3] >= 1 {
			player.Gold += 750
			tempWeapon = "weapon4"
		} else if currentweaponState == weapon5 && tempInt[4] >= 1 {
			player.Gold += 1500
			tempWeapon = "weapon5"
		} else if currentweaponState == weapon6 && tempInt[5] >= 1 {
			player.Gold += 1000
			tempWeapon = "weapon6"
		} else if currentweaponState == weapon7 && tempInt[6] >= 1 {
			player.Gold += 2000
			tempWeapon = "weapon7"
		} else if currentweaponState == weapon8 && tempInt[7] >= 1 {
			player.Gold += 3000
			tempWeapon = "weapon8"
		} else if currentweaponState == weapon9 && tempInt[8] >= 1 {
			player.Gold += 4000
			tempWeapon = "weapon9"
		} else if currentweaponState == weapon0 && tempInt[9] >= 1 {
			player.Gold += 5000
			tempWeapon = "weapon0"
		}
		SaveWeaponSellEvent(SaveFilePath, 3, tempWeapon, player)
		SaveGame(SaveFilePath, 1, player)
		tempInt = tempInt[:0]
	}
	return myState.CurrentGS
}
