package main

import (
	"github.com/Amovement/Hearthstone-Bot/internal"
	"github.com/Amovement/Hearthstone-Bot/pkg/win"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"log"
)

const title = "炉石传说"

func main() {
	var (
		mw          *walk.MainWindow
		startButton *walk.PushButton
	)
	// 创建主窗口
	if err := (declarative.MainWindow{
		AssignTo: &mw,
		Title:    "Hearth Bot @Amove",
		Size:     declarative.Size{Width: 300, Height: 264},
		Layout:   declarative.VBox{},
		Children: []declarative.Widget{
			// 第一个按钮
			declarative.PushButton{
				Text: "显示炉石",
				OnClicked: func() {
					hwnd := win.GetHwndByTitle(title)
					win.ShowAndTopWindow(hwnd)
				},
				MinSize: declarative.Size{
					Width: 200, Height: 65,
				},
			},
			// 第二个按钮
			declarative.PushButton{
				Text: "隐藏炉石",
				OnClicked: func() {
					hwnd := win.GetHwndByTitle(title)
					win.HideWindow(hwnd)
				},
				MinSize: declarative.Size{
					Width: 200, Height: 65,
				},
			},
			// 第三个按钮
			declarative.PushButton{
				Text:     "开始挂机",
				AssignTo: &startButton,
				OnClicked: func() {
					startButton.SetEnabled(false)
					bot := internal.NewBot(win.GetHwndByTitle(title))
					go bot.Start()
				},
				MinSize: declarative.Size{
					Width: 200, Height: 65,
				},
			},
		},
	}).Create(); err != nil {
		log.Fatal(err)
	}

	// 运行主窗口
	mw.Run()
	//bot := internal.NewBot(win.GetHwndByTitle(title))
	//bot.Start()
}
