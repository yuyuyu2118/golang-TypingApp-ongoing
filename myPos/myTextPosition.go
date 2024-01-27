package myPos

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

func DrawPos(win *pixelgl.Window, txt *text.Text, pos pixel.Vec) {
	txt.Draw(win, pixel.IM.Moved(pos))
}

func CenterPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	centerPos := pixel.V(
		win.Bounds().Center().Sub(txt.Bounds().Center()).X,
		win.Bounds().Center().Sub(txt.Bounds().Center()).Y,
	)
	return centerPos
}

// 画面中央の右隅にテキストを描画
func CenterRightPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	centerLeftPos := pixel.V(
		win.Bounds().Center().Sub(txt.Bounds().Center()).X+win.Bounds().Max.X/3,
		win.Bounds().Center().Sub(txt.Bounds().Center()).Y,
	)
	return centerLeftPos
}

// 画面中央の左隅にテキストを描画
func CenterLeftPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	centerLeftPos := pixel.V(
		win.Bounds().Center().Sub(txt.Bounds().Center()).X-win.Bounds().Max.X/3,
		win.Bounds().Center().Sub(txt.Bounds().Center()).Y,
	)
	return centerLeftPos
}

func TopCenterPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	TopCenterPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X,
		win.Bounds().Min.Y+float64(WinHSize/6),
	)
	return TopCenterPos
}

func TopRightPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	TopCenterPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X+win.Bounds().Max.X/3,
		win.Bounds().Min.Y+float64(WinHSize/6),
	)
	return TopCenterPos
}

func TopLeftPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	centerLeftPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X-win.Bounds().Max.X/3,
		win.Bounds().Min.Y+float64(WinHSize/6),
	)
	return centerLeftPos
}

func BottleCenterPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bottleCenterPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X,
		-float64(WinHSize/3),
	)
	return bottleCenterPos
}

func BottleRightPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bottleLeftPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X+win.Bounds().Max.X/3,
		-float64(WinHSize/3),
	)
	return bottleLeftPos
}

func BottleLeftPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bottleLeftPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X-win.Bounds().Max.X/3,
		-float64(WinHSize/3),
	)
	return bottleLeftPos
}

func LineCenterAlign(win *pixelgl.Window, lines []string, txt *text.Text, position string) {
	for _, line := range lines {
		centerX := win.Bounds().Center().Sub(txt.BoundsOf(line).Center()).X
		txt.Dot.X = centerX
		fmt.Fprintln(txt, line)
	}
	if position == "center" {
		DrawPos(win, txt, CenterPos(win, txt))
	}
}

func BottleRoundCenterPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bottleCenterPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X,
		-win.Bounds().Center().Y/2,
	)
	return bottleCenterPos
}

// ウィンドウの幅と高さに基づいて、相対的な位置を計算する関数
func RelativePos(win *pixelgl.Window, relX float64, relY float64) pixel.Vec {
	return pixel.V(
		win.Bounds().Min.X+(win.Bounds().W()*relX),
		win.Bounds().Min.Y+(win.Bounds().H()*relY),
	)
}

// ウィンドウの幅と高さに基づいて、相対的な位置を計算する関数
// テキストオブジェクトの中心を基準に位置を決める
func RelativeCenterPos(win *pixelgl.Window, txt *text.Text, relX float64, relY float64) pixel.Vec {
	txtBounds := txt.Bounds()
	txtWidth := txtBounds.W()
	txtHeight := txtBounds.H()

	return pixel.V(
		win.Bounds().Min.X+(win.Bounds().W()*relX)-txtWidth/2,
		win.Bounds().Min.Y+(win.Bounds().H()*relY)-txtHeight/2,
	)
}

// DrawCenteredText関数を修正して、*text.Text型のオブジェクトと文字列を受け取るようにします。
func RelativeDraw(win *pixelgl.Window, txt *text.Text, txtStr string, relX float64, relY float64) {
	// テキストオブジェクトをクリアして新しい文字列を設定
	txt.Clear()
	fmt.Fprintln(txt, txtStr)

	// テキストオブジェクトの中心を基準にした相対位置を計算
	pos := RelativeCenterPos(win, txt, relX, relY)

	// 計算された位置にテキストを描画
	DrawPos(win, txt, pos)
}

// ウィンドウの幅と高さに基づいて、相対的な位置を計算する関数
// テキストオブジェクトの左下隅を基準に位置を決める
func RelativeDrawFromCorner(win *pixelgl.Window, txt *text.Text, txtStr string, relX float64, relY float64) {
	txt.Clear()
	fmt.Fprintln(txt, txtStr)

	// テキストオブジェクトの左下隅を基準にした相対位置を計算
	pos := pixel.V(
		win.Bounds().Min.X+(win.Bounds().W()*relX),
		win.Bounds().Min.Y+(win.Bounds().H()*relY),
	)

	// 計算された位置にテキストを描画
	DrawPos(win, txt, pos)
}

// DrawRectBorder は、指定された相対位置に枠線を描画する関数です。
func DrawRectBorder(win *pixelgl.Window, relX1, relY1, relX2, relY2 float64, thickness float64, borderColor color.RGBA) {
	imd := imdraw.New(nil) // imdrawのインスタンスを作成

	// ウィンドウのサイズに基づいて、四角形の各頂点の絶対座標を計算
	x1 := win.Bounds().Min.X + (win.Bounds().W() * relX1)
	y1 := win.Bounds().Min.Y + (win.Bounds().H() * relY1)
	x2 := win.Bounds().Min.X + (win.Bounds().W() * relX2)
	y2 := win.Bounds().Min.Y + (win.Bounds().H() * relY2)

	// 枠線の色と太さを設定
	imd.Color = borderColor
	imd.EndShape = imdraw.SharpEndShape // 線の端を鋭角に
	imd.Push(pixel.V(x1, y1))           // 左下の頂点
	imd.Push(pixel.V(x2, y1))           // 右下の頂点
	imd.Push(pixel.V(x2, y2))           // 右上の頂点
	imd.Push(pixel.V(x1, y2))           // 左上の頂点
	imd.Polygon(thickness)              // 四角形の枠線を描画

	imd.Draw(win) // ウィンドウに描画
}

type MessageBox struct {
	win          *pixelgl.Window
	txt          *text.Text
	borderColor  color.RGBA
	textColor    color.RGBA // テキストの色を追加
	borderThick  float64
	relX1, relY1 float64
	relX2, relY2 float64
}

var (
	PADDING_X = 0.005
	PADDING_Y = 0.005
)

func NewMessageBox(win *pixelgl.Window, txt *text.Text, borderColor color.RGBA, textColor color.RGBA, borderThick float64, relX1, relY1, relX2, relY2 float64) *MessageBox {
	return &MessageBox{
		win:         win,
		txt:         txt,
		borderColor: borderColor,
		textColor:   textColor,
		borderThick: borderThick,
		relX1:       relX1 + PADDING_X,
		relY1:       relY1 + PADDING_Y,
		relX2:       relX2 - PADDING_X,
		relY2:       relY2 - PADDING_Y,
	}
}

func (mb *MessageBox) DrawMessageBox() {
	DrawRectBorder(mb.win, mb.relX1, mb.relY1, mb.relX2, mb.relY2, mb.borderThick, mb.borderColor)
}

func (mb *MessageBox) DrawMessageTxt(message string) {
	mb.txt.Clear()
	mb.txt.Color = mb.textColor // テキストの色を設定
	fmt.Fprintln(mb.txt, message)

	// メッセージボックスの利用可能な幅と高さを計算
	availableWidth := mb.win.Bounds().W() * (mb.relX2 - mb.relX1)
	availableHeight := mb.win.Bounds().H() * (mb.relY2 - mb.relY1)

	// パディングをメッセージボックスのサイズに依存させる
	paddingX := availableWidth / 128
	paddingY := availableHeight / 128

	// テキスト内の改行の数を数える
	lineCount := strings.Count(message, "\n") + 1

	// X座標はメッセージボックスの左端からパディングを加えた位置
	startX := mb.win.Bounds().Min.X + (mb.win.Bounds().W() * mb.relX1) + paddingX
	// Y座標はメッセージボックスの上端からテキストの高さとパディングを引いた位置
	// 1行分の高さを計算
	singleLineHeight := mb.txt.Bounds().H() / float64(lineCount)
	startY := mb.win.Bounds().Min.Y + (mb.win.Bounds().H() * mb.relY2) - singleLineHeight - paddingY

	DrawPos(mb.win, mb.txt, pixel.V(startX, startY))
}
