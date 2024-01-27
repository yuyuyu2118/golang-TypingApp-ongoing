package myGame

import (
	"log"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	event "github.com/yuyuyu2118/typingGo/Event"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
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
var stageName = []string{"スライム", "バード", "プラント", "コボルト", "ゾンビ", "フェアリー", "スカル", "ウィザード", "ソルジャー", "ドラゴン"}

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
	for i, v := range stageName {
		if event.UnlockNewDungeonEventInstance.Dungeons[i] {
			stageSlice[i] = strconv.Itoa(i+1) + ". " + v
		}
	}
	// 10番目のダンジョンの表示を修正
	if event.UnlockNewDungeonEventInstance.Dungeons[9] {
		stageSlice[9] = "0. " + stageName[9]
	}

	stageMessageBox := myPos.NewMessageBox(win, myUtil.MessageTxt, colornames.White, colornames.White, 5, 0, 0, 1, 0.4)
	stageMessageBox.DrawMessageBox()
	var dungeonOptions string
	for i, dungeonName := range stageSlice {
		// アンロックされていないダンジョンは "1. ???" のように表示
		if event.UnlockNewDungeonEventInstance.Dungeons[i] {
			dungeonOptions += strconv.Itoa(i+1) + ". " + stageName[i] + "\n"
		} else {
			dungeonOptions += dungeonName + "\n" // 既に "1. ???" の形式である
		}
	}
	stageMessageBox.DrawMessageTxt("どのモンスターと戦いますか？キーボードに対応する数字を入力してください。\n" + dungeonOptions + "\nBackSpaceキーでタイトルに戻る")

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
}

func StageClickEvent(win *pixelgl.Window, mousePos pixel.Vec) myState.GameState {

	for i := 0; i < len(keyToDungeon)-1; i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if (win.Pressed(key)) && event.UnlockNewDungeonEventInstance.Dungeons[i] && myState.CurrentGS == myState.StageSelect {
			currentdungeonState = DungeonState(i + 1)
			myState.CurrentGS = myState.PlayingScreen
			StageNum = i
			log.Println("ダンジョンセレクト", i+1)
			break
		}
	}

	if (win.JustPressed(pixelgl.Key0)) && event.UnlockNewDungeonEventInstance.Dungeons[9] && myState.CurrentGS == myState.StageSelect {
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
