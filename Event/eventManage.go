package event

import (
	"strconv"
)

func InitializeEventInstance(loadContent [][]string) {
	CreateWeaponPurchaseEvent(loadContent[2])
	CreateArmorPurchaseEvent(loadContent[2])
	CreateAccessoryPurchaseEvent(loadContent[2])
	CreateUnlockNewJobEvent(loadContent[2])
	CreateUnlockNewDungeonEvent(loadContent[2])
}

type WeaponPurchaseEvent struct {
	Weapons map[int]bool
}

var WeaponPurchaseEventInstance *WeaponPurchaseEvent

// TODO: 同じような変数のインスタンスを生成するものはすべてマップに置き換えることで、for文に変換できる
// weaponが購入されていればweaponPurchaseBoolに対応した武器がtrueになる
func CreateWeaponPurchaseEvent(value []string) {
	var weaponPurchaseBool = make(map[int]bool)
	for i, b := range value {
		tempInt, _ := strconv.Atoi(b)
		if tempInt >= 1 {
			weaponPurchaseBool[i] = true
		} else {
			weaponPurchaseBool[i] = false
		}
	}
	WeaponPurchaseEventInstance = &WeaponPurchaseEvent{
		Weapons: weaponPurchaseBool,
	}
}

type ArmorPurchaseEvent struct {
	Armors map[int]bool
}

var ArmorPurchaseEventInstance *ArmorPurchaseEvent

func CreateArmorPurchaseEvent(value []string) {
	var armorPurchaseBool = make(map[int]bool)
	for i, b := range value {
		tempInt, _ := strconv.Atoi(b)
		if tempInt >= 1 {
			armorPurchaseBool[i] = true
		} else {
			armorPurchaseBool[i] = false
		}
	}
	ArmorPurchaseEventInstance = &ArmorPurchaseEvent{
		Armors: armorPurchaseBool,
	}
}

type AccessoryPurchaseEvent struct {
	Accessorys map[int]bool
}

var AccessoryPurchaseEventInstance *AccessoryPurchaseEvent

func CreateAccessoryPurchaseEvent(value []string) {
	var accessoryPurchaseBool = make(map[int]bool)
	for i, b := range value {
		tempInt, _ := strconv.Atoi(b)
		if tempInt >= 1 {
			accessoryPurchaseBool[i] = true
		} else {
			accessoryPurchaseBool[i] = false
		}
	}
	AccessoryPurchaseEventInstance = &AccessoryPurchaseEvent{
		Accessorys: accessoryPurchaseBool,
	}
}

type UnlockNewJobEvent struct {
	Jobs map[int]bool
}

var UnlockNewJobEventInstance *UnlockNewJobEvent

func CreateUnlockNewJobEvent(value []string) {
	var unlockNewJobBool = make(map[int]bool)
	for i, b := range value {
		tempInt, _ := strconv.Atoi(b)
		if tempInt >= 1 {
			unlockNewJobBool[i] = true
		} else {
			unlockNewJobBool[i] = false
		}
	}
	UnlockNewJobEventInstance = &UnlockNewJobEvent{
		Jobs: unlockNewJobBool,
	}
}

type UnlockNewDungeonEvent struct {
	Dungeons map[int]bool
}

var UnlockNewDungeonEventInstance *UnlockNewDungeonEvent

func CreateUnlockNewDungeonEvent(value []string) {
	var unlockNewDungeonBool = make(map[int]bool)
	unlockNewDungeonBool[0] = true
	for i, b := range value {
		tempInt, _ := strconv.Atoi(b)
		if tempInt >= 1 {
			unlockNewDungeonBool[i+1] = true
		} else {
			unlockNewDungeonBool[i+1] = false
		}
	}
	UnlockNewDungeonEventInstance = &UnlockNewDungeonEvent{
		Dungeons: unlockNewDungeonBool,
	}
}
