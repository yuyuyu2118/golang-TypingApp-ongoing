package myGame

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
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

var (
	townButtonSlice = []pixel.Rect{}
)

func InitTown(win *pixelgl.Window, Txt *text.Text) {
	townMessageBox := myPos.NewMessageBox(win, myUtil.MessageTxt, colornames.White, colornames.White, 5, 0, 0, 1, 0.4)
	townMessageBox.DrawMessageBox()
	townMessageBox.DrawMessageTxt("どのお店へ行きますか？キーボードに対応する数字を入力してください。\n1. 武器店\n2. 防具店\n3. アクセサリー店\n4. 装備\n\nBackSpaceキーでタイトルに戻る")
}

func TownClickEvent(win *pixelgl.Window, mousePos pixel.Vec) myState.GameState {
	//TODO ページを作成したら追加
	if myState.CurrentGS == myState.TownScreen && (win.JustPressed(pixelgl.Key1)) {
		myState.CurrentGS = myState.WeaponShop
		log.Println("Town->WeaponShop")
	} else if myState.CurrentGS == myState.TownScreen && (win.JustPressed(pixelgl.Key2)) {
		myState.CurrentGS = myState.ArmorShop
		log.Println("Town->ArmorShop")
	} else if myState.CurrentGS == myState.TownScreen && (win.JustPressed(pixelgl.Key3)) {
		myState.CurrentGS = myState.AccessoryShop
		log.Println("Town->AccessoryShop")
		/*} else if myState.CurrentGS == myState.TownScreen && (townButtonSlice[3].Contains(mousePos) || win.JustPressed(pixelgl.Key4)) {
		myState.CurrentGS = myState.BlackSmith
		log.Println("Town->BlackSmith")*/
	} else if myState.CurrentGS == myState.TownScreen && (win.JustPressed(pixelgl.Key4)) {
		myState.CurrentGS = myState.EquipmentScreen
		log.Println("Town->EquipmentScreen")
	} else if myState.CurrentGS == myState.TownScreen && (win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentGS = myState.GoToScreen
		log.Println("Town->GoToScreen")
	}
	return myState.CurrentGS
}
