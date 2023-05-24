package main

import (
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	event "github.com/yuyuyu2118/typingGo/Event"
	"github.com/yuyuyu2118/typingGo/battle"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
	"golang.org/x/image/colornames"
)

const (
	winHSize = 1440
)

var startTime time.Time
var Ticker *time.Ticker
var saveContent string

// TODO: Utilに入れる
var language bool

func run() {
	rand.Seed(time.Now().UnixNano())
	win, _ := initializeWindow()
	myPos.SetCfg(winHSize)

	fontPath := "assets\\fonts\\NotoSans-Black.ttf"
	japanFontPath := "assets/fonts/PixelMplus12-Regular.ttf"
	basicTxt := initializeAnyText(fontPath, 40, colornames.White)
	screenTxt := initAnyJapanText(japanFontPath, 40, colornames.White)
	descriptionTxt := initAnyJapanText(japanFontPath, 30, colornames.White)
	startTxt := initializeAnyText(fontPath, 80, colornames.White)
	endTxt := initializeAnyText(fontPath, 60, colornames.White)

	myGame.SaveFileCheck("player\\save\\save.csv")
	loadContent := myGame.SaveFileLoad("player\\save\\save.csv")
	//playerStatusインスタンスを生成
	player := player.NewPlayerStatus(loadContent[1])
	event.CreateWeaponPurchaseEvent(loadContent[2])

	enemyInfo := enemy.CreateEnemyInstance()
	var enemySettings []enemy.EnemyStatus
	for _, enemy := range *enemyInfo {
		enemySettings = append(enemySettings, enemy)
	}
	//loadMonsterAnimation
	var enemySprites [][]*pixel.Sprite
	var enemyWaitSprites []*pixel.Sprite
	enemyNameSlice := []string{"Slime", "Bird", "Plant", "Goblin", "Zombie", "Fairy", "Skull", "Wizard", "Solidier"}
	enemyPath := "assets\\monster\\"
	var enemyPathBar []string
	// enemyPathBar := "assets\\monster\\Skull\\SkullA_Attack0.png"
	var enemyPaths []string
	for i, eNameSlice := range enemyNameSlice {
		enemyPaths = append(enemyPaths, enemyPath+eNameSlice)
		enemyPathBar = append(enemyPathBar, enemyPath+eNameSlice+"\\"+eNameSlice+"A_Wait0.png")
		enemyWaitSprites = enemy.SetEnemyAnimation(enemyPaths[i], eNameSlice+"A_Wait")
		enemySprites = append(enemySprites, enemyWaitSprites)
	}

	frame := 0
	last := time.Now()

	stage := myGame.NewStageInf(0)

	for !win.Closed() {
		switch myGame.CurrentGS {
		case myGame.StartScreen:
			myGame.InitStartScreen(win, startTxt)
			if win.JustPressed(pixelgl.KeyEnter) {
				myGame.CurrentGS = myGame.GoToScreen
				log.Println("Press:Enter -> GameState:jobSelect")
			}
			//testMode
			if win.JustPressed(pixelgl.KeyT) {
				myGame.CurrentGS = myGame.TestState
				log.Println("TestMode")
			}
			if win.JustPressed(pixelgl.KeyEscape) {
				win.Destroy()
				os.Exit(0)
			}
		case myGame.GoToScreen:
			//TODO: Saveの削除
			initScreenInformation(win, screenTxt, player)
			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.GoToClickEvent(win, win.MousePosition())
			}
		case myGame.StageSelect:
			initScreenInformation(win, screenTxt, player)
			//TODO: Key入力受付
			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.Key7) || win.JustPressed(pixelgl.Key8) || win.JustPressed(pixelgl.Key9) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.StageClickEvent(win, win.MousePosition(), stage)
			}
		case myGame.TownScreen:
			initScreenInformation(win, screenTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.TownClickEvent(win, win.MousePosition())
			}
		case myGame.WeaponShop:
			initScreenInformation(win, descriptionTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.WeaponClickEvent(win, win.MousePosition())
			}
		case myGame.ArmorShop:
			initScreenInformation(win, descriptionTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.ArmorClickEvent(win, win.MousePosition())
			}
		case myGame.AccessoryShop:
			initScreenInformation(win, descriptionTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.AccessoryClickEvent(win, win.MousePosition())
			}
		case myGame.EquipmentScreen:
			initScreenInformation(win, screenTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.EquipmentClickEvent(win, win.MousePosition())
			}
		case myGame.JobSelect:
			initScreenInformation(win, screenTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.JobClickEvent(win, win.MousePosition(), player)
				myGame.SaveGame("player\\save\\save.csv", 1, player)
			}

		case myGame.PlayingScreen:
			initScreenInformation(win, basicTxt, player)
			log.Println(stage.StageNum)

			dt := time.Since(last).Seconds()
			if dt >= 0.2 { // アニメーション速度を調整 (ここでは0.2秒ごとに更新)
				frame = (frame + 1) % len(enemySprites[stage.StageNum])
				last = time.Now()
			}
			enemy.SetEnemySprite(win, &enemySettings[stage.StageNum], enemyPathBar[stage.StageNum], enemySprites[stage.StageNum], frame)
			enemy.SetEnemySpriteText(win, screenTxt, &enemySettings[stage.StageNum])
			//TODO 手持ちアイテムバー、攻撃力や防御力の表示UI追加
			player.SetPlayerBattleInf(win, basicTxt)

			elapsed := time.Since(startTime)
			battle.InitBattleTextV2(win, basicTxt, elapsed)
			myGame.CurrentGS = battle.BattleTypingV2(win, player, &enemySettings[stage.StageNum], elapsed)
			if myGame.CurrentGS == myGame.BattleEnemyScreen {
				startTime = time.Now()
			}
		case myGame.BattleEnemyScreen:
			initScreenInformation(win, basicTxt, player)

			dt := time.Since(last).Seconds()
			if dt >= 0.2 { // アニメーション速度を調整 (ここでは0.2秒ごとに更新)
				frame = (frame + 1) % len(enemySprites[stage.StageNum])
				last = time.Now()
			}
			enemy.SetEnemySprite(win, &enemySettings[stage.StageNum], enemyPathBar[stage.StageNum], enemySprites[stage.StageNum], frame)
			enemy.SetEnemySpriteText(win, screenTxt, &enemySettings[stage.StageNum])
			//TODO 手持ちアイテムバー、攻撃力や防御力の表示UI追加
			player.SetPlayerBattleInf(win, basicTxt)

			elapsed := time.Since(startTime)
			battle.InitBattleTextV2(win, basicTxt, elapsed)
			myGame.CurrentGS = battle.BattleTypingV2(win, player, &enemySettings[stage.StageNum], elapsed)
			if myGame.CurrentGS == myGame.PlayingScreen {
				startTime = time.Now()
			}
		case myGame.EndScreen:
			myGame.InitEndScreen(win, endTxt)
			myGame.CurrentGS = battle.BattleEndScreen(win, endTxt, player, &enemySettings[stage.StageNum])
			//TODO
			if !myUtil.GetSaveReset() {
				//saveContent = "NoName,30,30,3,1,50,0,2," + strconv.Itoa(player.Gold) + "," + player.Job + "," + strconv.Itoa(player.AP) + ",Japanese,"
				myGame.SaveGame("player\\save\\save.csv", 1, player)
				myUtil.SetSaveReset(true)
			}
		case myGame.TestState:
			myGame.TestMode(win, basicTxt)
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
