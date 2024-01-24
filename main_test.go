package main

import (
	"testing"

	"github.com/faiface/pixel/pixelgl"
	"github.com/stretchr/testify/assert"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"golang.org/x/image/colornames"
)

func TestMain(m *testing.M) {
	pixelgl.Run(func() {
		// テストを実行します。
		m.Run()
	})
}

func TestInitializeWindow(t *testing.T) {
	win, cfg := initializeWindow()

	// ウィンドウインスタンスがnilではないことを確認
	assert.NotNil(t, win, "initializeWindow()はウィンドウインスタンスを返すべきです")

	// ウィンドウ設定が期待どおりであることを確認
	assert.Equal(t, "TypingRPG", cfg.Title, "ウィンドウのタイトルが正しく設定されているべきです")
	assert.Equal(t, true, cfg.VSync, "VSyncが有効になっているべきです")
	assert.Equal(t, false, cfg.Resizable, "ウィンドウはリサイズ不可であるべきです")

	// ウィンドウのサイズが期待どおりであることを確認
	expectedWidth := float64(winHSize)
	expectedHeight := float64((winHSize / 16) * 9)
	assert.Equal(t, expectedWidth, cfg.Bounds.Max.X, "ウィンドウの幅が正しく設定されているべきです")
	assert.Equal(t, expectedHeight, cfg.Bounds.Max.Y, "ウィンドウの高さが正しく設定されているべきです")
}

func TestSetCfg(t *testing.T) {
	// myPos.SetCfgを呼び出して、WinHSizeを設定します。
	myPos.SetCfg(winHSize)

	assert.Equal(t, winHSize, myPos.WinHSize, "ウィンドウの高さが正しく設定されているべきです")
	t.Log("ウィンドウの高さが正しく設定されています")
}

func TestInitTxtFontLoading(t *testing.T) {
	// InitTxtFontLoading関数を呼び出します。
	myUtil.InitTxtFontLoading()

	// グローバル変数が適切に設定されているか検証します。
	// 例えば、BasicTxtがnilでないことを確認します。
	assert.NotNil(t, myUtil.BasicTxt, "BasicTxtが初期化されていなければなりません")
	assert.NotNil(t, myUtil.StartTxt, "StartTxtが初期化されていなければなりません")
	assert.NotNil(t, myUtil.ScreenTxt, "ScreenTxtが初期化されていなければなりません")
	assert.NotNil(t, myUtil.DescriptionTxt, "DescriptionTxtが初期化されていなければなりません")
	assert.NotNil(t, myUtil.HunterBulletTxt, "HunterBulletTxtが初期化されていなければなりません")
	assert.NotNil(t, myUtil.CompletedTxt, "CompletedTxtが初期化されていなければなりません")
	assert.NotNil(t, myUtil.StatusTxt, "StatusTxtが初期化されていなければなりません")

	// 以下は、色やフォントサイズが期待通りであることを検証する例です。
	assert.Equal(t, colornames.White, myUtil.BasicTxt.Color, "BasicTxt should have the color white")
}
