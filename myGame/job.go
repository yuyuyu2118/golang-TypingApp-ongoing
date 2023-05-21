package myGame

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/player"
	"golang.org/x/image/colornames"
)

var (
	job1Button = pixel.Rect{}
	job2Button = pixel.Rect{}
	job3Button = pixel.Rect{}
	// job4Button = pixel.Rect{}
	// job5Button = pixel.Rect{}
	// job6Button = pixel.Rect{}
)

func InitJob(win *pixelgl.Window, Txt *text.Text) {

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "Select your job")
	tempPosition = myPos.TopCenterPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "1. Warrior")
	tempPosition = myPos.CenterLeftPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
	job1Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "2. Priest")
	tempPosition = myPos.CenterPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
	job2Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "3. Wizard")
	tempPosition = myPos.CenterRightPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
	job3Button = Txt.Bounds().Moved(tempPosition)
}

func JobClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *player.PlayerStatus) GameState {

	if job1Button.Contains(mousePos) || win.JustPressed(pixelgl.Key1) {
		CurrentGS = GoToScreen
		player.Job = "Warrior"
	} else if job2Button.Contains(mousePos) || win.JustPressed(pixelgl.Key2) {
		CurrentGS = GoToScreen
		player.Job = "Priest"
	} else if job3Button.Contains(mousePos) || win.JustPressed(pixelgl.Key3) {
		CurrentGS = GoToScreen
		player.Job = "Wizard"
	}
	log.Println("YourJob is", player.Job)
	return CurrentGS
}

func InitPlayerJob(win *pixelgl.Window, Txt *text.Text, player *player.PlayerStatus) {
	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, player.Job)
	xOffSet := 0.0
	yOffSet := win.Bounds().H() / 3
	txtPos := pixel.V(xOffSet, yOffSet)
	tempPosition := pixel.IM.Moved(txtPos)
	Txt.Draw(win, tempPosition)
}
