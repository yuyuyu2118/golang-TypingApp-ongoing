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
	"golang.org/x/image/colornames"
)

type DungeonState int

const (
	dungeonNil DungeonState = iota
	dungeon1
	dungeon2
	dungeon3
	dungeon4
	dungeon5
	dungeon6
	dungeon7
	dungeon8
	dungeon9
	dungeon10
)

var keyToDungeon = map[pixelgl.Button]DungeonState{
	pixelgl.Key1: dungeon1,
	pixelgl.Key2: dungeon2,
	pixelgl.Key3: dungeon3,
	pixelgl.Key4: dungeon4,
	pixelgl.Key5: dungeon5,
	pixelgl.Key6: dungeon6,
	pixelgl.Key7: dungeon7,
	pixelgl.Key8: dungeon8,
	pixelgl.Key9: dungeon9,
	pixelgl.Key0: dungeon10,
}

var stageSlice = []string{"1. ???", "2. ???", "3. ???", "4. ???", "5. ???", "6. ???", "7. ???", "8. ???", "9. ???", "0. ???"}
var stageNum = []string{"dungeon0", "dungeon1", "dungeon2", "dungeon3", "dungeon4", "dungeon5", "dungeon6", "dungeon7", "dungeon8", "dungeon9"}
var stageName = []string{"スライム", "バード", "プラント", "ゴブリン", "ゾンビ", "フェアリー", "スカル", "ウィザード", "ソルジャー", "ドラゴン"}

var StageNum int

// TODO: これと同じボタンは消してもOK
var (
	stage1Button = pixel.Rect{}
)

var (
	dungeonButtonSlice = []pixel.Rect{}
)
var currentdungeonState DungeonState

func InitStage(win *pixelgl.Window, Txt *text.Text) {
	xOffSet := myPos.TopLefPos(win, Txt).X + 400
	yOffSet := myPos.TopLefPos(win, Txt).Y
	txtPos := pixel.V(0, 0)

	for i, v := range stageName {
		if event.UnlockNewDungeonEventInstance.Dungeons[i] {
			stageSlice[i] = strconv.Itoa(i+1) + ". " + v
		}
	}
	if event.UnlockNewDungeonEventInstance.Dungeons[9] {
		stageSlice[9] = "0. " + stageName[9]
	}

	for _, dungeonName := range stageSlice {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, dungeonName)
		yOffSet -= Txt.LineHeight + 25
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
		dungeonButtonSlice = append(dungeonButtonSlice, Txt.Bounds().Moved(txtPos))
	}

	for i := 0; i < len(keyToDungeon)-1; i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if win.Pressed(key) && event.UnlockNewDungeonEventInstance.Dungeons[i] {
			currentdungeonState = DungeonState(i + 1)
			break
		}
	}
	if win.Pressed(pixelgl.Key0) && event.UnlockNewDungeonEventInstance.Dungeons[9] {
		currentdungeonState = dungeon10
	}
	//TODO: ダンジョン説明を追加する
	// if currentdungeonState >= dungeon1 && currentdungeonState <= dungeon10 {
	// 	DescriptionWeapon(win, descWeapon, int(currentdungeonState)-1)
	// }
}

func StageClickEvent(win *pixelgl.Window, mousePos pixel.Vec) myState.GameState {

	for i := 0; i < len(keyToDungeon)-1; i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if (dungeonButtonSlice[i].Contains(mousePos) || win.Pressed(key)) && event.UnlockNewDungeonEventInstance.Dungeons[i] && myState.CurrentGS == myState.StageSelect {
			currentdungeonState = DungeonState(i + 1)
			myState.CurrentGS = myState.PlayingScreen
			StageNum = i
			log.Println("ダンジョンセレクト", i+1)
			break
		}
	}

	if (dungeonButtonSlice[9].Contains(mousePos) || win.JustPressed(pixelgl.Key0)) && event.UnlockNewDungeonEventInstance.Dungeons[9] && myState.CurrentGS == myState.StageSelect {
		currentdungeonState = DungeonState(9)
		myState.CurrentGS = myState.PlayingScreen
		StageNum = 9
		log.Println("ダンジョンセレクト", 9)
	} else if myState.CurrentGS == myState.StageSelect && (win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentGS = myState.GoToScreen
		log.Println("StageScreen -> GoToScreen")
	}

	log.Println("PlayStage is", StageNum)
	return myState.CurrentGS
}
