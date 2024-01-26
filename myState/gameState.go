package myState

// CurrentGS は現在のゲームの状態を保持するグローバル変数です。
var CurrentGS GameState

// GameState はゲームの状態を表す整数型です。
type GameState int

// 以下はGameStateの定数で、ゲームの異なる画面や状態を表しています。
// iotaは連続する整数値を自動で割り当てるために使われます。
const (
	FadeScreen  GameState = iota // 画面のフェード処理
	StartScreen                  // スタート画面
	GoToScreen                   // 画面遷移

	StageSelect     // ステージ選択
	TownScreen      // 街の画面
	EquipmentScreen // 装備画面
	JobSelect       // 職業選択画面
	SaveScreen      // セーブ画面

	PlayingScreen     // プレイ中の画面
	BattleEnemyScreen // 敵との戦闘画面
	SkillScreen       // スキル画面
	EndScreen         // エンド画面
	TestState         // テスト状態

	WeaponShop       // 武器屋
	ArmorShop        // 鎧屋
	AccessoryShop    // アクセサリー屋
	BlackSmithScreen // 鍛冶屋の画面
)

// CurrentBelong は現在の所属状態を保持するグローバル変数です。
var CurrentBelong BelongState

// BelongState はプレイヤーの所属状態を表す整数型です。
type BelongState int

// 以下はBelongStateの定数で、プレイヤーが所属する可能性のあるカテゴリを表しています。
const (
	WeaponBelong    BelongState = iota // 武器に所属
	ArmorBelong                        // 鎧に所属
	AccessoryBelong                    // アクセサリーに所属
	MaterialsBelong                    // 材料に所属
)

// CurrentBlackSmith は現在の鍛冶屋の状態を保持するグローバル変数です。
var CurrentBlackSmith BlackSmithState

// BlackSmithState は鍛冶屋の状態を表す整数型です。
type BlackSmithState int

// 以下はBlackSmithStateの定数で、鍛冶屋で扱うアイテムのカテゴリを表しています。
const (
	WeaponBlackSmith    BlackSmithState = iota // 武器の鍛冶
	ArmorBlackSmith                            // 鎧の鍛冶
	AccessoryBlackSmith                        // アクセサリーの鍛冶
	MaterialsBlackSmith                        // 材料の鍛冶
)
