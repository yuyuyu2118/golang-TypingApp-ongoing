package myGame

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

var (
	townButtonSlice = []pixel.Rect{}
)

func InitTown(win *pixelgl.Window, Txt *text.Text) {
	xOffSet := myPos.TopLefPos(win, Txt).X + 300
	yOffSet := myPos.TopLefPos(win, Txt).Y - 50
	txtPos := pixel.V(0, 0)

	//gotoSlice := []string{"1. Dungeon", "2. Town", "3. Equipment", "4. Job", "5. Save", "6. EXIT"}
	townSlice := []string{"1. 武器店", "2. 防具店", "3. アクセサリー店", "4. 鍛冶屋", "5. 装備", "BackSpace. 戻る"}

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

func TownClickEvent(win *pixelgl.Window, mousePos pixel.Vec) GameState {
	//TODO ページを作成したら追加
	if townButtonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1) {
		CurrentGS = WeaponShop
		log.Println("Town->WeaponShop")
	} else if townButtonSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.Key2) {
		CurrentGS = ArmorShop
		log.Println("Town->ArmorShop")
	} else if townButtonSlice[2].Contains(mousePos) || win.JustPressed(pixelgl.Key3) {
		CurrentGS = AccessoryShop
		log.Println("Town->AccessoryShop")
	} else if townButtonSlice[3].Contains(mousePos) || win.JustPressed(pixelgl.Key4) {
		CurrentGS = BlackSmith
		log.Println("Town->BlackSmith")
	} else if townButtonSlice[4].Contains(mousePos) || win.JustPressed(pixelgl.Key5) {
		CurrentGS = EquipmentScreen
		log.Println("Town->EquipmentScreen")
	} else if townButtonSlice[5].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace) {
		CurrentGS = GoToScreen
		log.Println("Town->GoToScreen")
	}
	return CurrentGS
}
