package myGame

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"golang.org/x/image/colornames"
)

var tempPosition pixel.Vec

var (
	goTo1Button = pixel.Rect{}
	goTo2Button = pixel.Rect{}
	goTo3Button = pixel.Rect{}
	goTo4Button = pixel.Rect{}
	goTo5Button = pixel.Rect{}
	goTo6Button = pixel.Rect{}
)

var (
	gotoButtonSlice = []pixel.Rect{}
)

func InitGoTo(win *pixelgl.Window, Txt *text.Text, bottleText string) {

	xOffSet := 100.0
	yOffSet := myPos.TopLefPos(win, Txt).Y - 100
	txtPos := pixel.V(0, 0)

	// imgPath := "assets/Screen.png"
	// imgFile, err := os.Open(imgPath)
	// if err != nil {
	// 	panic(err)
	// }
	// defer imgFile.Close()

	// img, _, err := image.Decode(imgFile)
	// if err != nil {
	// 	panic(err)
	// }

	// picData := pixel.PictureDataFromImage(img)

	// // 2. `pixel.NewSprite`を使って `*pixel.Sprite`オブジェクトを作ります。
	// sprite := pixel.NewSprite(picData, picData.Bounds())

	// // 3. テキストと同様に、`Draw()`メソッドを使って画像をウィンドウに描画します。
	// imgXCenter := win.Bounds().Center().X
	// imgYPos := win.Bounds().Center().Y // Adjust this value to change the vertical position of the image
	// sprite.Draw(win, pixel.IM.Moved(pixel.V(imgXCenter, imgYPos)))

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, bottleText)
	tempPosition = myPos.BotCenPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)

	//gotoSlice := []string{"1. Dungeon", "2. Town", "3. Equipment", "4. Job", "5. Save", "6. EXIT"}
	gotoSlice := []string{"1. ダンジョン", "2. 町", "3. ジョブ", "4. 鍛冶屋"}

	for _, gotoName := range gotoSlice {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, gotoName)
		yOffSet -= Txt.LineHeight + 40
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
		gotoButtonSlice = append(gotoButtonSlice, Txt.Bounds().Moved(txtPos))
	}
}

func GoToClickEvent(win *pixelgl.Window, mousePos pixel.Vec) myState.GameState {
	//TODO ページを作成したら追加
	//TODO: 全部この形式にする　やばいバグ
	if myState.CurrentGS == myState.GoToScreen && (win.JustPressed(pixelgl.Key1)) {
		myState.CurrentGS = myState.StageSelect
		log.Println("GoToScreen->Dungeon")
	} else if myState.CurrentGS == myState.GoToScreen && (win.JustPressed(pixelgl.Key2)) {
		myState.CurrentGS = myState.TownScreen
		log.Println("GoToScreen->Town")
	} else if myState.CurrentGS == myState.GoToScreen && (win.JustPressed(pixelgl.Key3)) {
		myState.CurrentGS = myState.JobSelect
		log.Println("GoToScreen->JobSelect")
	} else if myState.CurrentGS == myState.GoToScreen && (win.JustPressed(pixelgl.Key4)) {
		myState.CurrentGS = myState.BlackSmithScreen
		log.Println("GoToScreen->BlackSmithScreen")
	} else if myState.CurrentGS == myState.GoToScreen && (win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentGS = myState.StartScreen
		log.Println("GoToScreen->StartScreen")
	}
	return myState.CurrentGS
}
