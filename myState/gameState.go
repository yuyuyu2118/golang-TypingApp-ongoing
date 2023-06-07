package myState

var CurrentGS GameState

type GameState int

const (
	FadeScreen GameState = iota
	StartScreen
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
	BlackSmithScreen
)

var CurrentBelong BelongState

type BelongState int

const (
	WeaponBelong BelongState = iota
	ArmorBelong
	AccessoryBelong
	MaterialsBelong
)

var CurrentBlackSmith BlackSmithState

type BlackSmithState int

const (
	WeaponBlackSmith BlackSmithState = iota
	ArmorBlackSmith
	AccessoryBlackSmith
	MaterialsBlackSmith
)
