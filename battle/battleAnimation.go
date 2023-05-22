package battle

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"golang.org/x/image/colornames"
)

const (
	animationSpeed    = 0.1
	animationFPS      = 60
	animationDuration = 2 * time.Second
)

type DamageAnimation struct {
	Text       string
	Position   pixel.Vec
	Offset     pixel.Vec
	Progress   float64
	RemoveFlag bool
	Done       chan struct{}
}

var DamageAnimations []*DamageAnimation

func RunDamageAnimation(win *pixelgl.Window, anim *DamageAnimation) {
	for anim.Progress <= 1.0 {
		anim.Progress += animationSpeed
		time.Sleep(time.Second / animationFPS)
		if anim.RemoveFlag {
			close(anim.Done)
			return
		}
		DrawDamageAnimation(win, anim)
	}
	RemoveAnimation(anim)
	close(anim.Done)
}

func DrawDamageAnimation(win *pixelgl.Window, anim *DamageAnimation) {
	fontPath := "assets\\fonts\\NotoSans-Black.ttf"
	face, _ := myUtil.LoadTTF(fontPath, float64(60))
	basicAtlas := text.NewAtlas(face, text.ASCII)
	txt := text.New(pixel.V(0, 0), basicAtlas)
	txt.Color = colornames.White

	for _, anim := range DamageAnimations {
		text := anim.Text
		txt.Clear()
		txt.WriteString(text) // テキストを書き込む
		// テキストの中心位置を計算
		bounds := txt.Bounds()
		txtPos := anim.Position.Sub(pixel.V(bounds.W()/2, bounds.H()/2))

		progress := anim.Progress
		if progress > 1.0 {
			progress = 1.0
		}

		// ランダムなオフセットを生成（初回のみ）
		if anim.Offset == pixel.ZV {
			anim.Offset = pixel.V(rand.Float64()*150-20, rand.Float64()*50-20)
		}

		alpha := 1.0
		scale := 1.0
		// テキストのアニメーションを描画
		// アニメーションの進行状況に応じて拡大縮小と位置を設定し、テキストの色に透明度を適用
		txt.DrawColorMask(win, pixel.IM.
			Scaled(txt.Orig, scale).        // アニメーションの進行状況に応じて拡大縮小
			Moved(txtPos.Add(anim.Offset)), // テキストの位置を移動
			pixel.Alpha(alpha), // 透明度を適用
		)
	}
}

func RemoveAnimation(anim *DamageAnimation) {
	for i, a := range DamageAnimations {
		if a == anim {
			DamageAnimations = append(DamageAnimations[:i], DamageAnimations[i+1:]...)
			break
		}
	}
}

func PlayerAttack(win *pixelgl.Window, damage int, position pixel.Vec) {
	anim := &DamageAnimation{
		Text:       strconv.Itoa(damage),
		Position:   position,
		Progress:   0.0,
		RemoveFlag: false,
		Done:       make(chan struct{}),
	}
	DamageAnimations = append(DamageAnimations, anim)
	go RunDamageAnimation(win, anim)

	removeChan := make(chan *DamageAnimation)

	go func() {
		<-anim.Done
		removeChan <- anim
	}()

	go func() {
		select {
		case <-time.After(animationDuration):
			removeChan <- anim
		case animToRemove := <-removeChan:
			RemoveAnimation(animToRemove)
		}
	}()
}
