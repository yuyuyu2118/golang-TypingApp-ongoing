package myGame

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
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
	xOffSet := myPos.TopLefPos(win, Txt).X + 400
	yOffSet := myPos.TopLefPos(win, Txt).Y - 50
	txtPos := pixel.V(0, 0)

	//gotoSlice := []string{"1. Dungeon", "2. Town", "3. Equipment", "4. Job", "5. Save", "6. EXIT"}
	townSlice := []string{"1. 武器店", "2. 防具店", "3. アクセサリー店", "4. 鍛冶屋", "5. 装備"}

	for _, townName := range townSlice {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, townName)
		yOffSet -= Txt.LineHeight + 25
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
		townButtonSlice = append(townButtonSlice, Txt.Bounds().Moved(txtPos))
	}
}

func TownClickEvent(win *pixelgl.Window, mousePos pixel.Vec) myState.GameState {
	//TODO ページを作成したら追加
	if myState.CurrentGS == myState.TownScreen && (townButtonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1)) {
		myState.CurrentGS = myState.WeaponShop
		log.Println("Town->WeaponShop")
	} else if myState.CurrentGS == myState.TownScreen && (townButtonSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.Key2)) {
		myState.CurrentGS = myState.ArmorShop
		log.Println("Town->ArmorShop")
	} else if myState.CurrentGS == myState.TownScreen && (townButtonSlice[2].Contains(mousePos) || win.JustPressed(pixelgl.Key3)) {
		myState.CurrentGS = myState.AccessoryShop
		log.Println("Town->AccessoryShop")
	} else if myState.CurrentGS == myState.TownScreen && (townButtonSlice[3].Contains(mousePos) || win.JustPressed(pixelgl.Key4)) {
		myState.CurrentGS = myState.BlackSmith
		log.Println("Town->BlackSmith")
	} else if myState.CurrentGS == myState.TownScreen && (townButtonSlice[4].Contains(mousePos) || win.JustPressed(pixelgl.Key5)) {
		myState.CurrentGS = myState.EquipmentScreen
		log.Println("Town->EquipmentScreen")
	} else if myState.CurrentGS == myState.TownScreen && (win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentGS = myState.GoToScreen
		log.Println("Town->GoToScreen")
	}
	return myState.CurrentGS
}
