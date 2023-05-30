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

type JobState int

const (
	jobNil JobState = iota
	job1
	job2
	job3
	job4
	job5
)

var keyToJob = map[pixelgl.Button]JobState{
	pixelgl.Key1: job1,
	pixelgl.Key2: job2,
	pixelgl.Key3: job3,
	pixelgl.Key4: job4,
	pixelgl.Key5: job5,
}

var jobSlice = []string{"1. ???", "2. ???", "3. ???", "4. ???", "5. ???"}
var jobNum = []string{"job0", "job1", "job2", "job3", "job4"}
var jobName = []string{"見習い剣士", "狩人", "モンク", "魔法使い", "化け物"}

var currentjobState JobState

func InitJob(win *pixelgl.Window, Txt *text.Text) {
	xOffSet := myPos.TopLefPos(win, Txt).X + 400
	yOffSet := myPos.TopLefPos(win, Txt).Y - 50
	txtPos := pixel.V(0, 0)

	for i, v := range jobName {
		if event.UnlockNewJobEventInstance.Jobs[i] {
			jobSlice[i] = strconv.Itoa(i+1) + ". " + v
		}
	}

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

	for i := 0; i < len(keyToJob)-1; i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if win.Pressed(key) && event.UnlockNewJobEventInstance.Jobs[i] {
			currentjobState = JobState(i + 1)
			break
		}
	}
	//TODO: ジョブ説明を追加する
	// if currentdungeonState >= dungeon1 && currentdungeonState <= dungeon10 {
	// 	DescriptionWeapon(win, descWeapon, int(currentdungeonState)-1)
	// }
}

func JobClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *player.PlayerStatus) myState.GameState {

	for i := 0; i < len(keyToJob); i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if (jobButtonSlice[i].Contains(mousePos) || win.Pressed(key)) && event.UnlockNewJobEventInstance.Jobs[i] && myState.CurrentGS == myState.JobSelect {
			currentjobState = JobState(i + 1)
			log.Println("ジョブ選択", i+1, jobName[i])
			player.Job = jobName[i]
			myState.CurrentGS = myState.GoToScreen
			break
		}
	}

	if myState.CurrentGS == myState.JobSelect && (win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentGS = myState.GoToScreen
		log.Println("jobScreen -> GoToScreen")
	}
	log.Println("YourJob is", player.Job)
	return myState.CurrentGS
}
