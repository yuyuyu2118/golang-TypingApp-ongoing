package event

import (
	"strconv"
)

type WeaponPurchaseEvent struct {
	Weapon1 bool
	Weapon2 bool
	Weapon3 bool
	Weapon4 bool
	Weapon5 bool
	Weapon6 bool
	Weapon7 bool
	Weapon8 bool
	Weapon9 bool
	Weapon0 bool
}

var WeaponPurchaseEventInstance WeaponPurchaseEvent

func CreateWeaponPurchaseEvent(value []string) {
	//temp := myIo.CsvToSliceAll("player\\save\\save.csv")

	var weaponPurchaseBool []bool
	for _, b := range value {
		tempInt, _ := strconv.Atoi(b)
		if tempInt == 1 {
			weaponPurchaseBool = append(weaponPurchaseBool, true)
		} else {
			weaponPurchaseBool = append(weaponPurchaseBool, false)
		}
	}
	WeaponPurchaseEventInstance = WeaponPurchaseEvent{
		Weapon1: weaponPurchaseBool[0],
		Weapon2: weaponPurchaseBool[1],
		Weapon3: weaponPurchaseBool[2],
		Weapon4: weaponPurchaseBool[3],
		Weapon5: weaponPurchaseBool[4],
		Weapon6: weaponPurchaseBool[5],
		Weapon7: weaponPurchaseBool[6],
		Weapon8: weaponPurchaseBool[7],
		Weapon9: weaponPurchaseBool[8],
		Weapon0: weaponPurchaseBool[9],
	}
}
