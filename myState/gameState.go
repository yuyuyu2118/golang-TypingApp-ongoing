package myState

var CurrentGS GameState

type GameState int

const (
	StartScreen GameState = iota
	GoToScreen

	StageSelect
	TownScreen
	EquipmentScreen
	JobSelect
	SaveScreen

	PlayingScreen
	BattleEnemyScreen
	SkillScreen
	EndScreen
	TestState

	WeaponShop
	ArmorShop
	AccessoryShop
	BlackSmith
)
