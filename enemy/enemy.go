package enemy

import (
	"fmt"
	"log"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myIo"
	"github.com/yuyuyu2118/typingGo/myPos"
	"golang.org/x/image/colornames"
)

type EnemyStatus struct {
	Name       string
	MaxHP      float64
	HP         float64
	OP         float64
	DP         float64
	Gold       int
	Attack     bool
	AttackTick float64
	DropAP     int
	EnemySize  float64
	DropItems  []string
}

var EnemyPath = "assets\\monster\\"
var EnemyNameSlice = []string{"Slime", "Bird", "Plant", "Goblin", "Zombie", "Fairy", "Skull", "Wizard", "Solidier"}

func CreateEnemyInstance() *[]EnemyStatus {
	temp := myIo.CsvToSliceAll("enemy/enemySettings/enemy.csv")
	var instance []EnemyStatus

	for _, value := range temp {
		var DropItems []string
		Name := value[0]
		MaxHP, _ := strconv.ParseFloat(value[1], 64)
		HP, _ := strconv.ParseFloat(value[2], 64)
		OP, _ := strconv.ParseFloat(value[3], 64)
		DP, _ := strconv.ParseFloat(value[4], 64)
		Gold, _ := strconv.Atoi((value[5]))
		Attack := false
		AttackTick, _ := strconv.ParseFloat(value[7], 64)
		DropAP, _ := strconv.Atoi((value[8]))
		EnemySize, _ := strconv.ParseFloat(value[9], 64)
		DropItems = append(DropItems, value[10])
		DropItems = append(DropItems, value[11])
		DropItems = append(DropItems, value[12])

		tempInstance := EnemyStatus{
			Name:       Name,
			MaxHP:      MaxHP,
			HP:         HP,
			OP:         OP,
			DP:         DP,
			Gold:       Gold,
			Attack:     Attack,
			AttackTick: AttackTick,
			DropAP:     DropAP,
			EnemySize:  EnemySize,
			DropItems:  DropItems,
		}
		instance = append(instance, tempInstance)
	}
	log.Println(instance)
	return &instance
}

var EnemySettings []EnemyStatus
var EnemyPathBar []string
var EnemySprites [][]*pixel.Sprite

func CreateEnemySettings() {
	//csvファイルからEnemyの情報を1行ずつ取り出して、enemySettingsスライスに格納
	enemyInfo := CreateEnemyInstance()
	for _, enemy := range *enemyInfo {
		EnemySettings = append(EnemySettings, enemy)
	}

	//Animationに使うスプライトをスライスのスライスに格納、enemyPathBarには体力表示用のスライスを格納

	for _, eNameSlice := range EnemyNameSlice {
		EnemyPathBar = append(EnemyPathBar, EnemyPath+eNameSlice+"\\"+eNameSlice+"A_Wait0.png")
		EnemySprites = append(EnemySprites, SetEnemyAnimation(EnemyPath+eNameSlice, eNameSlice+"A_Wait"))
	}
}

func SetEnemyHPBar(win *pixelgl.Window, scaledSize pixel.Vec, HP float64, MaxHP float64, pos pixel.Vec) {
	rectWidth := scaledSize.X * (HP / MaxHP)
	var rect pixel.Rect
	if HP > 0 {
		rect = pixel.R(
			win.Bounds().Center().X-(rectWidth/2),
			win.Bounds().Center().Y-50,
			win.Bounds().Center().X+(rectWidth/2),
			win.Bounds().Center().Y,
		)
	} else {
		rect = pixel.R(
			win.Bounds().Center().X,
			win.Bounds().Center().Y-50,
			win.Bounds().Center().X,
			win.Bounds().Center().Y,
		)
	}
	imd := imdraw.New(nil)
	imd.Color = colornames.Green
	imd.Push(rect.Min, rect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func SetEnemyHPBarOut(win *pixelgl.Window, scaledSize pixel.Vec, pos pixel.Vec) {
	rect := pixel.R(
		win.Bounds().Center().X-pos.X/2,
		win.Bounds().Center().Y-50,
		win.Bounds().Center().X+pos.X/2,
		win.Bounds().Center().Y,
	)
	imd := imdraw.New(nil)
	imd.Color = colornames.Red
	imd.Push(rect.Min, rect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func SetEnemyPic(win *pixelgl.Window, enemyInf *EnemyStatus, path string, scaleFactor float64) {
	pic, _ := myIo.OpenDecodePictureData(path)
	picMonster := pixel.NewSprite(pic, pic.Bounds())

	scaledSize := pic.Bounds().Size().Scaled(scaleFactor)
	transMat := pixel.IM.Moved(win.Bounds().Center().Add(pixel.V(0, 35))).Scaled(win.Bounds().Center(), scaleFactor)
	picMonster.Draw(win, transMat)

	barPosition := pixel.V(
		picMonster.Picture().Bounds().W()*scaleFactor,
		picMonster.Picture().Bounds().H()*scaleFactor,
	)

	pic.Bounds()

	SetEnemyHPBarOut(win, scaledSize, barPosition)
	SetEnemyHPBar(win, scaledSize, enemyInf.HP, enemyInf.MaxHP, barPosition)
}

func SetEnemyText(win *pixelgl.Window, Txt *text.Text, enemy *EnemyStatus) {
	// cp := constantProvider{}
	// WinHSize := cp.GetConstant()
	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "EnemyHP : ", enemy.HP)
	myPos.DrawPos(win, Txt, myPos.TopCenterPos(win, Txt))
}
