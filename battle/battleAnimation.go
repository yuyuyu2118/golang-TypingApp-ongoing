package battle

import (
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
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
			anim.Offset = pixel.V(rand.Float64(), rand.Float64())
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

func PlayerAttack(win *pixelgl.Window, damage float64, position pixel.Vec) {
	anim := &DamageAnimation{
		Text:       strconv.FormatFloat(damage, 'f', 2, 64),
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

const (
	uniqueSkillAnimationSpeed    = 0.04
	uniqueSkillAnimationFPS      = 180
	uniqueSkillAnimationDuration = 5 * time.Second
)

type UniqueSkillAnimation struct {
	Text       string
	Position   pixel.Vec
	Offset     pixel.Vec
	Progress   float64
	RemoveFlag bool
	Done       chan struct{}
}

var UniqueSkillAnimations []*UniqueSkillAnimation

func RunUniqueSkillAnimation(win *pixelgl.Window, anim *UniqueSkillAnimation, txtColor color.Color) {
	for anim.Progress <= 1.0 {
		anim.Progress += uniqueSkillAnimationSpeed
		time.Sleep(time.Second / uniqueSkillAnimationFPS)
		if anim.RemoveFlag {
			close(anim.Done)
			return
		}
		DrawUniqueSkillAnimation(win, anim, txtColor)
	}
	RemoveUniqueSkillAnimation(anim)
	close(anim.Done)
}

func DrawUniqueSkillAnimation(win *pixelgl.Window, anim *UniqueSkillAnimation, txtColor color.Color) {
	//TODO: いちいち読み込まないように
	fontPath := "assets\\fonts\\NotoSans-Black.ttf"
	face, _ := myUtil.LoadTTF(fontPath, float64(60))
	basicAtlas := text.NewAtlas(face, text.ASCII)
	txt := text.New(pixel.V(0, 0), basicAtlas)
	txt.Color = txtColor

	for _, anim := range UniqueSkillAnimations {
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
			anim.Offset = pixel.V(rand.Float64(), rand.Float64())
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

func RemoveUniqueSkillAnimation(anim *UniqueSkillAnimation) {
	for i, a := range UniqueSkillAnimations {
		if a == anim {
			UniqueSkillAnimations = append(UniqueSkillAnimations[:i], UniqueSkillAnimations[i+1:]...)
			break
		}
	}
}

var tempTxt string

func UniqueSkill(win *pixelgl.Window, tempPoint float64, position pixel.Vec, txtColor color.Color, player *player.PlayerStatus, assignTxt string) {

	if player.EquipmentWeapon[0] == weaponName[3] {
		tempTxt = assignTxt
	} else if player.EquipmentWeapon[0] == weaponName[4] {
		tempTxt = assignTxt
	} else if player.EquipmentWeapon[0] == weaponName[5] {
		tempTxt = assignTxt
	} else if player.EquipmentWeapon[0] == weaponName[6] {
		tempTxt = assignTxt
	} else if player.EquipmentWeapon[0] == weaponName[7] {
		tempTxt = assignTxt
	} else if player.EquipmentWeapon[0] == weaponName[8] {
		tempTxt = assignTxt
	} else if player.EquipmentWeapon[0] == weaponName[9] {
		tempTxt = assignTxt
	}

	if player.EquipmentArmor[0] == armorName[3] {
		tempTxt = assignTxt
	} else if player.EquipmentArmor[0] == armorName[4] {
		tempTxt = assignTxt
	} else if player.EquipmentArmor[0] == armorName[5] {
		tempTxt = assignTxt
	} else if player.EquipmentArmor[0] == armorName[6] {
		tempTxt = assignTxt
	} else if player.EquipmentArmor[0] == armorName[7] {
		tempTxt = assignTxt
	} else if player.EquipmentArmor[0] == armorName[8] {
		tempTxt = assignTxt
	} else if player.EquipmentArmor[0] == armorName[9] {
		tempTxt = assignTxt
	}

	if player.EquipmentAccessory[0] == accessoryName[3] {
		tempTxt = assignTxt
	} else if player.EquipmentAccessory[0] == accessoryName[4] {
		tempTxt = assignTxt
	} else if player.EquipmentAccessory[0] == accessoryName[5] {
		tempTxt = assignTxt
	} else if player.EquipmentAccessory[0] == accessoryName[6] {
		tempTxt = assignTxt
	} else if player.EquipmentAccessory[0] == accessoryName[7] {
		tempTxt = assignTxt
	} else if player.EquipmentAccessory[0] == accessoryName[8] {
		tempTxt = assignTxt
	} else if player.EquipmentAccessory[0] == accessoryName[9] {
		tempTxt = assignTxt
	}

	anim := &UniqueSkillAnimation{
		Text:       tempTxt + strconv.FormatFloat(tempPoint, 'f', 2, 64),
		Position:   position,
		Progress:   0.0,
		RemoveFlag: false,
		Done:       make(chan struct{}),
	}
	UniqueSkillAnimations = append(UniqueSkillAnimations, anim)
	go RunUniqueSkillAnimation(win, anim, txtColor)

	removeChan := make(chan *UniqueSkillAnimation)

	go func() {
		<-anim.Done
		removeChan <- anim
	}()

	go func() {
		select {
		case <-time.After(uniqueSkillAnimationDuration):
			removeChan <- anim
		case animToRemove := <-removeChan:
			RemoveUniqueSkillAnimation(animToRemove)
		}
	}()
}
