package myGame

import (
	"log"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	event "github.com/yuyuyu2118/typingGo/Event"
	"github.com/yuyuyu2118/typingGo/myPlayer"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
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
	jobMessageBox := myPos.NewMessageBox(win, myUtil.MessageTxt, colornames.White, colornames.White, 5, 0, 0, 1, 0.4)
	jobMessageBox.DrawMessageBox()

	var jobOptions string
	for i, job := range jobName {
		if event.UnlockNewJobEventInstance.Jobs[i] {
			jobOptions += strconv.Itoa(i+1) + ". " + job + "\n"
		} else {
			jobOptions += strconv.Itoa(i+1) + ". ???\n"
		}
	}

	jobMessageBox.DrawMessageTxt("どのジョブを選びますか？キーボードに対応する数字を入力してください。\n" + jobOptions + "\nBackSpaceキーでタイトルに戻る")

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

func JobClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *myPlayer.PlayerStatus) myState.GameState {

	for i := 0; i < len(keyToJob); i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if (win.Pressed(key)) && event.UnlockNewJobEventInstance.Jobs[i] && myState.CurrentGS == myState.JobSelect {
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
