package myUtil

import (
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

func AnimateText(win *pixelgl.Window, Txt, completedText *text.Text, lines []string, start time.Time, startPos, endPos pixel.Vec, duration float64) {
	elapsedTime := time.Since(start).Seconds()

	completedText.Clear()
	paragraphIdx := 0
	for idx, line := range lines {
		if elapsedTime > duration*float64(idx+1) {
			completedText.Dot.X = endPos.X
			completedText.Dot.Y = endPos.Y - completedText.LineHeight*float64(idx-paragraphIdx)
			completedText.WriteString(line)

			if line == "\n" {
				paragraphIdx++
			}

			continue // この行のアニメーションが終了した場合、次の行に移動
		}

		// 各行のアニメーション時間を計算
		lineElapsedTime := math.Max(elapsedTime-duration*float64(idx), 0)

		// 補間処理
		progress := lineElapsedTime / duration
		currentPosX := lerp(startPos.X, endPos.X, progress)
		currentPosY := lerp(startPos.Y, endPos.Y-Txt.LineHeight*float64(idx-paragraphIdx), progress)

		// テキストの描画
		if line == "\n" {
			paragraphIdx++
			continue // 改行文字の場合、次の行に移動
		}
		Txt.Dot.X = currentPosX
		Txt.Dot.Y = currentPosY
		Txt.WriteString(line)
	}
	Txt.Draw(win, pixel.IM.Scaled(Txt.Orig, 1))
	completedText.Draw(win, pixel.IM.Scaled(completedText.Orig, 1)) // 描画済みテキストを描画
}

func lerp(start, end float64, t float64) float64 {
	return start + (end-start)*t
}

func FadeScreen(win *pixelgl.Window, imd *imdraw.IMDraw, setTime time.Time, fadeDuration float64) (float64, float64) {
	win.Clear(colornames.Darkcyan)
	elapsedTime := time.Since(setTime).Seconds()
	cycleTime := math.Mod(elapsedTime, fadeDuration) // サイクル内の経過時間を計算
	imd.Clear()                                      // ここでIMDrawの内容をクリアする
	return cycleTime / fadeDuration, elapsedTime
}

func DrawFadingRectangleInOut(win *pixelgl.Window, imd *imdraw.IMDraw, alpha float64, fadeIn bool) {
	if fadeIn { // フェードアウトの場合はアルファ値を反転
		alpha = 1 - alpha
	}

	imd.Color = pixel.RGBA{R: 0, G: 0, B: 0, A: alpha}
	imd.Push(win.Bounds().Min, win.Bounds().Max)
	imd.Rectangle(0)
}
