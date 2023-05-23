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
	EnemySize  float64
	DropAP     int
}

func CreateEnemyInstance() *[]EnemyStatus {
	temp := myIo.CsvToSliceAll("enemy/enemySettings/enemy.csv")
	var instance []EnemyStatus

	for _, value := range temp {
		log.Println(value[0], value[1])
		Name := value[0]
		MaxHP, _ := strconv.ParseFloat(value[1], 64)
		HP, _ := strconv.ParseFloat(value[2], 64)
		OP, _ := strconv.ParseFloat(value[3], 64)
		DP, _ := strconv.ParseFloat(value[4], 64)
		Gold, _ := strconv.Atoi((value[5]))
		Attack := false
		AttackTick, _ := strconv.ParseFloat(value[7], 64)
		DropAP, _ := strconv.Atoi((value[8]))

		tempInstance := EnemyStatus{
			Name:       Name,
			MaxHP:      MaxHP,
			HP:         HP,
			OP:         OP,
			DP:         DP,
			Gold:       Gold,
			Attack:     Attack,
			AttackTick: AttackTick,
			EnemySize:  4.0,
			DropAP:     DropAP,
		}
		instance = append(instance, tempInstance)
	}
	return &instance
}

func SetEnemyHPBar(win *pixelgl.Window, scaledSize pixel.Vec, HP float64, MaxHP float64, pos pixel.Vec) {
	rectWidth := scaledSize.X * ((MaxHP - (MaxHP - HP)) * 0.01)
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

func SetEnemyAnimation() []*pixel.Sprite {
	imagePaths := []string{"assets/monster/Slime/スライムA_待機000.png", "assets/monster/Slime/スライムA_待機001.png", "assets/monster/Slime/スライムA_待機002.png"}
	sprites, err := myIo.LoadSpriteSheet(imagePaths)
	if err != nil {
		panic(err)
	}
	return sprites
}
