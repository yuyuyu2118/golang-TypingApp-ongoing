package main

import (
	"image"
	"image/gif"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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
	win, _ := initializeWindow()
	rand.Seed(time.Now().UnixNano())
	myGame.SaveFileCheck("assets\\save\\save.csv")
	loadContent := myGame.SaveFileLoad("assets\\save\\save.csv")
	log.Println(loadContent[1])

	f, err := os.Open("assets/monster/Slime/スライム_アニメーション.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	g, err := gif.DecodeAll(f)
	if err != nil {
		panic(err)
	}

	// Create a palette for each image.
	palettes := make([]*image.Paletted, len(g.Image))
	for i, img := range g.Image {
		palettes[i] = img
	}

	frame := 0
	dt := float64(g.Delay[frame]) / 100.0

	fontPath := "assets\\fonts\\NotoSans-Black.ttf"
	japanFontPath := "assets/fonts/PixelMplus12-Regular.ttf"
	basicTxt := initializeAnyText(fontPath, 40, colornames.White)
	screenTxt := initAnyJapanText(japanFontPath, 40, colornames.White)
	descriptionTxt := initAnyJapanText(japanFontPath, 30, colornames.White)
	startTxt := initializeAnyText(fontPath, 80, colornames.White)
	endTxt := initializeAnyText(fontPath, 60, colornames.White)

	myPos.SetCfg(winHSize)
	//playerStatusインスタンスを生成
	stage := myGame.NewStageInf(0)
	playerLoadInfo := loadContent[1]
	player := player.NewPlayerStatus(playerLoadInfo)
	log.Println(player)
	enemyInfo := enemy.CreateEnemyInstance()
	enemyKnight := (*enemyInfo)[0]
	// for _, enemy := range *enemyInfo {
	// 	enemyKnight := enemy
	// }

	last := time.Now()
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
		case myGame.GoToScreen:
			// Update the current frame.
			dt -= time.Since(last).Seconds()
			last = time.Now()
			if dt <= 0 {
				frame = (frame + 1) % len(g.Image)
				dt = float64(g.Delay[frame]) / 100.0
			}

			// Draw the current frame.
			win.Clear(pixel.RGB(1, 1, 1))
			img := g.Image[frame]
			sprite := pixel.NewSprite(pixel.PictureDataFromImage(img), pixel.Rect{
				Min: pixel.Vec{X: float64(img.Bounds().Min.X), Y: float64(img.Bounds().Min.Y)},
				Max: pixel.Vec{X: float64(img.Bounds().Max.X), Y: float64(img.Bounds().Max.Y)},
			})
			sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
			//TODO: Saveの削除
			initScreenInformation(win, screenTxt, player)
			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.GoToClickEvent(win, win.MousePosition())
			}
		case myGame.StageSelect:
			initScreenInformation(win, screenTxt, player)
			//TODO: Key入力受付
			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.KeyBackspace) {
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
				saveContent = "NoName,30,30,3,1,50,0,2," + strconv.Itoa(player.Gold) + "," + player.Job + "," + strconv.Itoa(player.AP) + ",Japanese,"
				myGame.SaveGame("assets\\save\\save.csv", 1, saveContent)
			}
		// case myGame.SaveScreen:
		// 	initScreenInformation(win, basicTxt, player)

		// 	if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) {
		// 		myGame.CurrentGS = myGame.SaveClickEvent(win, win.MousePosition(), player)
		// 	}

		case myGame.PlayingScreen:
			initScreenInformation(win, basicTxt, player)
			enemy.SetEnemyPic(win, &enemyKnight, "assets\\monster\\monster1.png", enemyKnight.EnemySize)
			enemy.SetEnemyText(win, basicTxt, &enemyKnight)
			//TODO 手持ちアイテムバー、攻撃力や防御力の表示UI追加
			player.SetPlayerBattleInf(win, basicTxt)

			elapsed := time.Since(startTime)
			battle.InitBattleTextV2(win, basicTxt, elapsed)
			myGame.CurrentGS = battle.BattleTypingV2(win, player, &enemyKnight, elapsed)
			if myGame.CurrentGS == myGame.BattleEnemyScreen {
				startTime = time.Now()
			}
		case myGame.BattleEnemyScreen:
			initScreenInformation(win, basicTxt, player)
			enemy.SetEnemyPic(win, &enemyKnight, "assets\\monster\\monster1.png", enemyKnight.EnemySize)
			enemy.SetEnemyText(win, basicTxt, &enemyKnight)
			//TODO 手持ちアイテムバー、攻撃力や防御力の表示UI追加
			player.SetPlayerBattleInf(win, basicTxt)

			elapsed := time.Since(startTime)
			battle.InitBattleTextV2(win, basicTxt, elapsed)
			myGame.CurrentGS = battle.BattleTypingV2(win, player, &enemyKnight, elapsed)
			if myGame.CurrentGS == myGame.PlayingScreen {
				startTime = time.Now()
			}
		case myGame.EndScreen:
			myGame.InitEndScreen(win, endTxt)
			myGame.CurrentGS = battle.BattleEndScreen(win, endTxt, player, &enemyKnight)
			//TODO
			if !myUtil.GetSaveReset() {
				saveContent = "NoName,30,30,3,1,50,0,2," + strconv.Itoa(player.Gold) + "," + player.Job + "," + strconv.Itoa(player.AP) + ",Japanese,"
				myGame.SaveGame("assets\\save\\save.csv", 1, saveContent)
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
