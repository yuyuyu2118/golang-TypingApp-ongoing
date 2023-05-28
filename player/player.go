package player

import (
	"fmt"
	"log"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"golang.org/x/image/colornames"
)

type PlayerStatus struct {
	Name            string
	MaxHP           float64
	HP              float64
	OP              float64
	DP              float64
	MaxSP           float64
	SP              float64
	BaseSP          float64
	Gold            int
	Job             string
	AP              int
	Language        string
	AttackTimer     float64
	PossessedWeapon []string
}

var (
	skillRect = pixel.R(0, 0, 0, 0)
)

var PlayerStatusInstance *PlayerStatus

func NewPlayerStatus(value []string, possess []string) *PlayerStatus {
	Name := value[0]
	MaxHP, _ := strconv.ParseFloat(value[1], 64)
	HP, _ := strconv.ParseFloat(value[2], 64)
	OP, _ := strconv.ParseFloat(value[3], 64)
	DP, _ := strconv.ParseFloat(value[4], 64)
	MaxSP, _ := strconv.ParseFloat(value[5], 64)
	SP, _ := strconv.ParseFloat(value[6], 64)
	BaseSP, _ := strconv.ParseFloat(value[7], 64)
	Gold, _ := strconv.Atoi((value[8]))
	Job := value[9]
	AP, _ := strconv.Atoi((value[10]))
	Language := value[11]
	AttackTimer, _ := strconv.ParseFloat(value[12], 64)
	PossessedWeapon := possess

	PlayerStatusInstance := &PlayerStatus{
		Name:            Name,
		MaxHP:           MaxHP,
		HP:              HP,
		OP:              OP,
		DP:              DP,
		MaxSP:           MaxSP,
		SP:              SP,
		BaseSP:          BaseSP,
		Gold:            Gold,
		Job:             Job,
		AP:              AP,
		Language:        Language,
		AttackTimer:     AttackTimer,
		PossessedWeapon: PossessedWeapon,
	}
	return PlayerStatusInstance
}

func SetPlayerSkillBarVertical(win *pixelgl.Window, player *PlayerStatus) {
	// log.Println("win.Bounds().Max.X=", win.Bounds().Max.X)
	// log.Println("win.Bounds().Max.Y=", win.Bounds().Max.Y)
	rect := pixel.R(
		win.Bounds().Max.X-win.Bounds().Max.X/15,
		win.Bounds().Min.Y+win.Bounds().Max.Y/9,
		win.Bounds().Max.X-win.Bounds().Max.X/36,
		win.Bounds().Max.Y-win.Bounds().Max.Y/9,
	)
	imd := imdraw.New(nil)
	imd.Color = colornames.Orange
	imd.Push(rect.Min, rect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func SetPlayerSkillBarOutVertical(win *pixelgl.Window, player *PlayerStatus) {
	//log.Println("maxsp", player.MaxSP, "SP=", player.SP, "Base=", player.BaseSP)
	rMinX := win.Bounds().Max.X - win.Bounds().Max.X/15
	rMinY := win.Bounds().Min.Y + win.Bounds().Max.Y/9
	rMaxX := win.Bounds().Max.X - win.Bounds().Max.X/36
	rMaxY := win.Bounds().Max.Y - win.Bounds().Max.Y/9
	if player.MaxSP <= player.SP {
		player.SP = 50
		skillRect = pixel.R(
			rMinX,
			rMinY+((win.Bounds().Max.Y-win.Bounds().Max.Y/9)-(win.Bounds().Min.Y+win.Bounds().Max.Y/9)),
			rMaxX,
			rMaxY,
		)
	} else if player.SP == 0 {
		skillRect = pixel.R(
			rMinX,
			rMinY,
			rMaxX,
			rMaxY,
		)
	} else if player.SP < player.MaxSP {
		skillRect = pixel.R(
			rMinX,
			rMinY+((win.Bounds().Max.Y-win.Bounds().Max.Y/9)-(win.Bounds().Min.Y+win.Bounds().Max.Y/9))*((player.MaxSP-(player.MaxSP-player.SP))/player.MaxSP),
			rMaxX,
			rMaxY,
		)
	}
	imd := imdraw.New(nil)
	imd.Color = colornames.Red
	imd.Push(skillRect.Min, skillRect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func SetPlayerHPBarVertical(win *pixelgl.Window, player *PlayerStatus) {
	rMinX := win.Bounds().Max.X - win.Bounds().Max.X/9
	rMinY := win.Bounds().Min.Y + win.Bounds().Max.Y/9
	rMaxX := win.Bounds().Max.X - win.Bounds().Max.X/15
	rMaxY := win.Bounds().Max.Y - win.Bounds().Max.Y/9

	hpRect := pixel.R(
		rMinX,
		rMinY,
		rMaxX,
		rMinY+(rMaxY-rMinY)*(player.HP/player.MaxHP), //-rMaxY*((player.HP)/player.MaxHP),
	)

	if player.HP <= 0 {
		hpRect = pixel.R(
			rMinX,
			rMinY,
			rMinX,
			rMinY,
		)
	}

	imd := imdraw.New(nil)
	imd.Color = colornames.Green
	imd.Push(hpRect.Min, hpRect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func SetPlayerHPBarOutVertical(win *pixelgl.Window, player *PlayerStatus) {
	hpRect := pixel.R(
		win.Bounds().Max.X-win.Bounds().Max.X/9,
		win.Bounds().Min.Y+win.Bounds().Max.Y/9,
		win.Bounds().Max.X-win.Bounds().Max.X/15,
		win.Bounds().Max.Y-win.Bounds().Max.Y/9,
	)
	imd := imdraw.New(nil)
	imd.Color = colornames.Red
	imd.Push(hpRect.Min, hpRect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func SetPlayerHPBarHorizontal(win *pixelgl.Window, player *PlayerStatus) {
	rMinX := win.Bounds().Center().X / 4
	rMinY := win.Bounds().Min.Y + win.Bounds().Max.Y/12
	rMaxX := win.Bounds().Center().X / 2
	rMaxY := win.Bounds().Min.Y + win.Bounds().Max.Y/12 + 30

	hpRect := pixel.R(
		rMinX,
		rMinY,
		rMinX+(rMaxX-rMinX)*(player.HP/player.MaxHP),
		rMaxY, //-rMaxY*((player.HP)/player.MaxHP)s,
	)

	if player.HP <= 0 {
		hpRect = pixel.R(
			rMinX,
			rMinY,
			rMinX,
			rMinY,
		)
	}

	imd := imdraw.New(nil)
	imd.Color = colornames.Green
	imd.Push(hpRect.Min, hpRect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func SetPlayerHPBarOutHorizontal(win *pixelgl.Window, player *PlayerStatus) {
	rMinX := win.Bounds().Center().X / 4
	rMinY := win.Bounds().Min.Y + win.Bounds().Max.Y/12
	rMaxX := win.Bounds().Center().X / 2
	rMaxY := win.Bounds().Min.Y + win.Bounds().Max.Y/12 + 30
	hpRect := pixel.R(
		rMinX,
		rMinY,
		rMaxX,
		rMaxY,
	)
	imd := imdraw.New(nil)
	imd.Color = colornames.Red
	imd.Push(hpRect.Min, hpRect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func SetPlayerSkillBarHorizontal(win *pixelgl.Window, player *PlayerStatus) {
	rMinX := win.Bounds().Center().X / 4
	rMinY := win.Bounds().Min.Y + win.Bounds().Max.Y/12
	rMaxX := win.Bounds().Center().X / 2
	rMaxY := win.Bounds().Min.Y + win.Bounds().Max.Y/12 - 30
	rect := pixel.R(
		rMinX,
		rMinY,
		rMaxX,
		rMaxY,
	)
	imd := imdraw.New(nil)
	imd.Color = colornames.Orange
	imd.Push(rect.Min, rect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func SetPlayerSkillBarOutHorizontal(win *pixelgl.Window, player *PlayerStatus) {
	//log.Println("maxsp", player.MaxSP, "SP=", player.SP, "Base=", player.BaseSP)
	rMinX := win.Bounds().Center().X / 4
	rMinY := win.Bounds().Min.Y + win.Bounds().Max.Y/12
	rMaxX := win.Bounds().Center().X / 2
	rMaxY := win.Bounds().Min.Y + win.Bounds().Max.Y/12 - 30
	if player.MaxSP <= player.SP {
		player.SP = 50
		skillRect = pixel.R(
			rMinX+(rMaxX-rMinX),
			rMinY,
			rMaxX,
			rMaxY,
		)
	} else if player.SP == 0 {
		skillRect = pixel.R(
			rMinX,
			rMinY,
			rMaxX,
			rMaxY,
		)
	} else if player.SP < player.MaxSP {
		skillRect = pixel.R(
			rMinX+(rMaxX-rMinX)*(player.SP/player.MaxSP),
			rMinY,
			rMaxX,
			rMaxY,
		)
	}
	imd := imdraw.New(nil)
	imd.Color = colornames.Red
	imd.Push(skillRect.Min, skillRect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func (player *PlayerStatus) SetPlayerBattleInf(win *pixelgl.Window, Txt *text.Text) {
	// SetPlayerSkillBarVertical(win, player)
	// SetPlayerSkillBarOutVertical(win, player)
	// SetPlayerHPBarOutVertical(win, player)
	// SetPlayerHPBarVertical(win, player)

	SetPlayerSkillBarHorizontal(win, player)
	SetPlayerSkillBarOutHorizontal(win, player)
	SetPlayerHPBarOutHorizontal(win, player)
	SetPlayerHPBarHorizontal(win, player)

	InitPlayerHPSP(win, Txt, player)
}

func (player *PlayerStatus) InitPlayerStatus(win *pixelgl.Window, Txt *text.Text) {
	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, player.Job, " Gold:", player.Gold, "S")
	fmt.Fprintln(Txt, "アタックタイマー: ", player.AttackTimer)
	//TODO: 攻撃力防御力の表示をする
	// fmt.Fprintln(Txt, "攻撃力: ", player.OP)
	// fmt.Fprintln(Txt, "防御力: ", player.DP)
	tempPosition := myPos.TopLefPos(win, Txt).Add(pixel.V(0, 30))
	myPos.DrawPos(win, Txt, tempPosition)
}

func InitPlayerHPSP(win *pixelgl.Window, Txt *text.Text, player *PlayerStatus) {
	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, player.HP, "\n", player.SP)
	xOffSet := 30.0
	yOffSet := 50.0
	txtPos := pixel.V(xOffSet, yOffSet)
	tempPosition := pixel.IM.Moved(myPos.BottleLeftPos(win, Txt).Add(txtPos))
	Txt.Draw(win, tempPosition)
}

func TempFunc() {
	log.Println("temp")
}
