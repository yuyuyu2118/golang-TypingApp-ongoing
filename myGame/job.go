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

var (
	jobButtonSlice = []pixel.Rect{}
)

func InitJob(win *pixelgl.Window, Txt *text.Text) {
	xOffSet := myPos.TopLefPos(win, Txt).X + 400
	yOffSet := myPos.TopLefPos(win, Txt).Y - 50
	txtPos := pixel.V(0, 0)

	//gotoSlice := []string{"1. Dungeon", "2. Town", "3. Equipment", "4. Job", "5. Save", "6. EXIT"}
	jobSlice := []string{"1. 見習い剣士", "2. 狩人", "3. モンク", "4. 魔法使い", "5. 化け物", "BackSpace. 戻る"}

	for _, jobName := range jobSlice {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, jobName)
		yOffSet -= Txt.LineHeight + 25
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
		jobButtonSlice = append(jobButtonSlice, Txt.Bounds().Moved(txtPos))
	}
}

func JobClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *player.PlayerStatus) GameState {

	if CurrentGS == JobSelect && (jobButtonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1)) {
		CurrentGS = GoToScreen
		player.Job = "見習い剣士"
	} else if CurrentGS == JobSelect && (jobButtonSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.Key2)) {
		CurrentGS = GoToScreen
		player.Job = "狩人"
	} else if CurrentGS == JobSelect && (jobButtonSlice[2].Contains(mousePos) || win.JustPressed(pixelgl.Key3)) {
		CurrentGS = GoToScreen
		player.Job = "モンク"
	} else if CurrentGS == JobSelect && (jobButtonSlice[3].Contains(mousePos) || win.JustPressed(pixelgl.Key4)) {
		CurrentGS = GoToScreen
		player.Job = "魔法使い"
	} else if CurrentGS == JobSelect && (jobButtonSlice[4].Contains(mousePos) || win.JustPressed(pixelgl.Key5)) {
		CurrentGS = GoToScreen
		player.Job = "化け物"
	} else if CurrentGS == JobSelect && (jobButtonSlice[5].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace)) {
		CurrentGS = GoToScreen
		log.Println("jobScreen -> GoToScreen")
	}
	log.Println("YourJob is", player.Job)
	return CurrentGS
}

// TODO: 不要
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
