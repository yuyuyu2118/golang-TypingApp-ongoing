package battle

import (
	"log"
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myPlayer"
	"golang.org/x/image/colornames"
)

var weaponName = []string{"木の棒", "果物ナイフ", "木刀", "ドレインソード", "スタンハンマー", "鉄の剣", "隼の剣", "勇者の剣", "名刀村正", "死神の大鎌"}
var armorName = []string{"草織りのローブ", "フルーツアーマー", "木の鎧", "ソウルバインドプレート", "スタンプレート", "鉄の鎧", "飛翔のマント", "勇者の鎧", "刃舞の衣", "冥界の鎧"}
var accessoryName = []string{"樹木のペンダント", "フルーツブレスレット", "平和のバンド", "ライフリンクのリング", "ショックウェーブリング", "鉄のブレスレット", "疾走のリング", "勇者のペンダント", "刀匠の指輪", "霊魂のイヤリング"}

func AttackRelationWeaponSkill(win *pixelgl.Window, player *myPlayer.PlayerStatus, enemy *enemy.EnemyStatus, tempWordDamage *float64) {
	if player.EquipmentWeapon[0] == weaponName[5] {
		if rand.Float64() <= 0.1 {
			UniqueSkill(win, (*tempWordDamage)*1.3, win.Bounds().Center().Sub(pixel.V(0, 150)), colornames.Aqua, player, "CriticalSlash! ")
			enemy.HP += (*tempWordDamage) * 1.3
		}
	} else if player.EquipmentWeapon[0] == weaponName[6] {
		for i := 0; i < rand.Intn(3); i++ {
			randomPotision := 40.0 * float64(i)
			UniqueSkill(win, (*tempWordDamage)*0.3, win.Bounds().Center().Sub(pixel.V(0, 150+randomPotision)), colornames.Yellow, player, "SwiftSlash! ")
			enemy.HP += (*tempWordDamage) * 0.3
		}
	} else if player.EquipmentWeapon[0] == weaponName[7] {
		log.Println(myGame.StageNum)
		if rand.Float64() <= 0.15 {
			UniqueSkill(win, (*tempWordDamage)*(float64(myGame.StageNum)*0.25), win.Bounds().Center().Sub(pixel.V(0, 150)), colornames.Darkorange, player, "Holiness! ")
			enemy.HP += (*tempWordDamage) * (float64(myGame.StageNum) * 0.25)
		}
	} else if player.EquipmentWeapon[0] == weaponName[8] {
		log.Println(myGame.StageNum)
		if rand.Float64() <= 0.2 {
			UniqueSkill(win, (*tempWordDamage)*1.3, win.Bounds().Center().Sub(pixel.V(0, 150)), colornames.Aqua, player, "Mind'sEye! ")
			enemy.HP += (*tempWordDamage) * 1.3
		}
	}
	if player.EquipmentAccessory[0] == accessoryName[4] {
		log.Println(myGame.StageNum)
		if rand.Float64() <= 0.05 {
			UniqueSkill(win, (*tempWordDamage)*0.25, win.Bounds().Center().Sub(pixel.V(0, 300)), colornames.Red, player, "ShockWave! ")
			enemy.HP += (*tempWordDamage) * 0.25
		}
	} else if player.EquipmentAccessory[0] == accessoryName[8] {
		log.Println(myGame.StageNum)
		if rand.Float64() <= 0.05 {
			UniqueSkill(win, 1.25, win.Bounds().Center().Sub(pixel.V(0, 300)), colornames.Purple, player, "Sharpen! OP*")
			player.OP *= 1.25
		}
	} else if player.EquipmentAccessory[0] == accessoryName[9] {
		log.Println(myGame.StageNum)
		if rand.Float64() <= 0.05 {
			UniqueSkill(win, 1.25, win.Bounds().Center().Sub(pixel.V(0, 300)), colornames.Silver, player, "SoulProtect! DP*")
			player.DP *= 1.25
		}
	}
}

func RecoveryRelationWeaponSkill(win *pixelgl.Window, player *myPlayer.PlayerStatus, enemy *enemy.EnemyStatus, tempWordDamage *float64) {
	if player.EquipmentWeapon[0] == weaponName[3] {
		if rand.Float64() <= 0.1 {
			UniqueSkill(win, -(*tempWordDamage)*0.01, win.Bounds().Center().Sub(pixel.V(0, 150)), colornames.Green, player, "LifeDrain! ")
			player.HP += -(*tempWordDamage) * 0.01
			if player.HP >= player.MaxHP {
				player.HP = player.MaxHP
			}
		}
	}
	if player.EquipmentArmor[0] == armorName[9] {
		if rand.Float64() <= 0.1 {
			UniqueSkill(win, -(*tempWordDamage)*0.01, win.Bounds().Center().Sub(pixel.V(0, 200)), colornames.Green, player, "UnderWorld! ")
			player.HP += -(*tempWordDamage) * 0.01
			if player.HP >= player.MaxHP {
				player.HP = player.MaxHP
			}
		}
	}
	if player.EquipmentAccessory[0] == accessoryName[3] {
		if rand.Float64() <= 0.05 {
			UniqueSkill(win, -(*tempWordDamage)*0.01, win.Bounds().Center().Sub(pixel.V(0, 300)), colornames.Green, player, "LifeLink! ")
			player.HP += -(*tempWordDamage) * 0.01
			if player.HP >= player.MaxHP {
				player.HP = player.MaxHP
			}
		}
	}
}

func TimerRelationWeaponSkill(win *pixelgl.Window, player *myPlayer.PlayerStatus, tempTimer *float64) float64 {
	if player.EquipmentWeapon[0] == weaponName[4] {
		if rand.Float64() <= 0.1 {
			UniqueSkill(win, 1.0, win.Bounds().Center().Sub(pixel.V(0, 150)), colornames.Blue, player, "StunAttack! ")
			return 1.0
		}
	}
	return 0.0
}

func DebuffRelationWeaponSkill(win *pixelgl.Window, player *myPlayer.PlayerStatus, enemy *enemy.EnemyStatus) {
	if player.EquipmentWeapon[0] == weaponName[9] {
		if rand.Float64() <= 0.2 {
			UniqueSkill(win, -1.0, win.Bounds().Center().Sub(pixel.V(0, 150)), colornames.Red, player, "ShadowCurse! ")
			enemy.OP -= 1.0
		}
	}
}

func AttackRelationArmorSkill(win *pixelgl.Window, player *myPlayer.PlayerStatus, enemy *enemy.EnemyStatus, tempWordDamage *float64) {
	if player.EquipmentArmor[0] == armorName[8] {
		for i := 0; i < rand.Intn(3); i++ {
			randomPotision := 40.0 * float64(i)
			UniqueSkill(win, (*tempWordDamage)*0.3, win.Bounds().Center().Sub(pixel.V(0, 200+randomPotision)), colornames.Yellowgreen, player, "ShadowSwift! ")
			enemy.HP += (*tempWordDamage) * 0.3
		}
	}
}

func RecoveryRelationArmorSkill(win *pixelgl.Window, player *myPlayer.PlayerStatus, enemy *enemy.EnemyStatus, tempWordDamage *float64) {
	if player.EquipmentArmor[0] == armorName[3] {
		if rand.Float64() <= 0.1 {
			UniqueSkill(win, enemy.MaxOP*0.05, win.Bounds().Center().Sub(pixel.V(0, 200)), colornames.Green, player, "SoulBind! ")
			player.HP += enemy.MaxOP * 0.05
			if player.HP >= player.MaxHP {
				player.HP = player.MaxHP
			}
		}
	} else if player.EquipmentArmor[0] == armorName[7] {
		if rand.Float64() <= 0.1 {
			randomValue := 0.5 + (rand.Float64() * 1)
			roundedRandomValue := float64(int(randomValue*10)) / 10
			UniqueSkill(win, roundedRandomValue, win.Bounds().Center().Sub(pixel.V(0, 200)), colornames.Green, player, "BraveRecovery! ")
			player.HP += roundedRandomValue
			if player.HP >= player.MaxHP {
				player.HP = player.MaxHP
			}
		}
	}
}

func TimerRelationArmorSkill(win *pixelgl.Window, player *myPlayer.PlayerStatus, tempTimer *float64) float64 {
	if player.EquipmentArmor[0] == armorName[4] {
		if rand.Float64() <= 0.1 {
			UniqueSkill(win, 1.0, win.Bounds().Center().Sub(pixel.V(0, 200)), colornames.Blue, player, "StampMaster! ")
			return 1.0
		}
	}
	return 0.0
}

func DefenceRelationArmorSkill(win *pixelgl.Window, player *myPlayer.PlayerStatus, enemy *enemy.EnemyStatus) {
	if player.EquipmentArmor[0] == armorName[5] {
		if rand.Float64() <= 0.1 {
			UniqueSkill(win, player.DP*1.2, win.Bounds().Center().Sub(pixel.V(0, 200)), colornames.Silver, player, "IronWall! DP*")
			player.DP *= 1.2
		}
	} else if player.EquipmentArmor[0] == armorName[6] {
		if rand.Float64() <= 0.1 {
			UniqueSkill(win, -99, win.Bounds().Center().Sub(pixel.V(0, 200)), colornames.Silver, player, "Floating! ")
			enemy.OP -= 99
		}
	} else if player.EquipmentArmor[0] == armorName[7] {
		if rand.Float64() <= 0.1 {
			UniqueSkill(win, player.DP*1.2, win.Bounds().Center().Sub(pixel.V(0, 250)), colornames.Silver, player, "BraveDefence! DP*")
			player.DP *= 1.2
		}
	} else if player.EquipmentArmor[0] == armorName[8] {
		if rand.Float64() <= 0.1 {
			UniqueSkill(win, -99, win.Bounds().Center().Sub(pixel.V(0, 200)), colornames.Silver, player, "ShadowEscape! ")
			enemy.OP -= 99
		}
	}

	if player.EquipmentAccessory[0] == accessoryName[5] {
		if rand.Float64() <= 0.05 {
			UniqueSkill(win, player.DP*1.1, win.Bounds().Center().Sub(pixel.V(0, 300)), colornames.Silver, player, "Protect! DP*")
			player.DP *= 1.1
		}
	} else if player.EquipmentAccessory[0] == accessoryName[6] {
		if rand.Float64() <= 0.05 {
			UniqueSkill(win, -99, win.Bounds().Center().Sub(pixel.V(0, 300)), colornames.Silver, player, "Sprint! ")
			enemy.OP -= 99
		}
	}
}

func BuffRelationArmorSkill(win *pixelgl.Window, player *myPlayer.PlayerStatus, enemy *enemy.EnemyStatus) {
	if player.EquipmentArmor[0] == armorName[9] {
		if rand.Float64() <= 0.1 {
			UniqueSkill(win, 1.3, win.Bounds().Center().Sub(pixel.V(0, 250)), colornames.Purple, player, "UnderWorld! OP*")
			player.OP *= 1.3
		}
	}
	if player.EquipmentAccessory[0] == accessoryName[7] {
		if rand.Float64() <= 0.05 {
			UniqueSkill(win, 1.1, win.Bounds().Center().Sub(pixel.V(0, 300)), colornames.Purple, player, "Brave! OP&DP*")
			player.OP *= 1.1
			player.DP *= 1.1
		}
	}
}
