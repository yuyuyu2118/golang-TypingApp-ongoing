package myUtil

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"golang.org/x/image/colornames"
)

func ShopInitAndText(win *pixelgl.Window, Txt *text.Text, botText string) (float64, float64, pixel.Vec) {
	xOffSet := 30.0
	yOffSet := myPos.TopLefPos(win, Txt).Y - 50
	txtPos := pixel.V(0, 0)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, botText)
	myPos.DrawPos(win, Txt, myPos.BotCenPos(win, Txt))
	return xOffSet, yOffSet, txtPos
}

func DisplayShopLineup(win *pixelgl.Window, lineUp []string, buttonRect []pixel.Rect, LineHeight float64, textColor color.RGBA, Txt *text.Text, xOffSet float64, yOffSet float64, txtPos pixel.Vec) []pixel.Rect {
	for _, tempName := range lineUp {
		Txt.Clear()
		Txt.Color = colornames.White

		fmt.Fprintln(Txt, tempName)

		yOffSet -= Txt.LineHeight + LineHeight
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)

		//TODO: buttonSliceの初期化処理をすれば、これも関数内に配置できる
		buttonRect = append(buttonRect, Txt.Bounds().Moved(txtPos))
	}
	return buttonRect
}
