package internal

import (
	"fmt"
	"github.com/Amovement/Hearthstone-Bot/pkg/win"
	"github.com/Amovement/Hearthstone-Bot/utils"
	"time"
)

type Bot struct {
	hwnd uintptr
}

func NewBot(hwnd uintptr) *Bot {
	return &Bot{hwnd: hwnd}
}

func (b *Bot) Start() {
	fmt.Println("bot 将在 2 秒后启动...")
	time.Sleep(2 * time.Second)

	// 初始化设置窗体
	win.ShowAndTopWindow(b.hwnd)
	win.SetWindowSize(b.hwnd, win.HearthstoneWindowWidth, win.HearthstoneWindowHeight)
	cnt := 0
	// 开始循环
	for {
		win.ShowAndTopWindow(b.hwnd)
		fmt.Println("选中巫妖王..")
		win.MockClick(b.hwnd, 260, 100)

		fmt.Println("开始比赛..")
		win.MockClick(b.hwnd, 495, 327)

		fmt.Println("静待 10 秒..")
		time.Sleep(time.Second * 10)
		fmt.Println("选择角色..")
		win.ShowAndTopWindow(b.hwnd)
		win.MockClick(b.hwnd, 500, 329)

		fmt.Println("静待 20 秒..")
		time.Sleep(time.Second * 20)

		fmt.Println("结束选牌")
		win.MockClick(b.hwnd, 350, 317)

		b.heartStop() // 隐藏界面 开始挂机

		myTurn := utils.Rand(1, 4)
		for ; myTurn > 0; myTurn-- {
			// 随机过一回合
			fmt.Println("本次剩余回合数:", myTurn)
			fmt.Println("随机过一回合")
			win.MockClick(b.hwnd, 550, 195)
			b.heartStop()
		}

		b.Surrender()
		cnt++
		fmt.Println("第 ", cnt, " 局结束 ", time.Now().Format("2006-01-02 15:04:05"))
	}
}

// heartStop 后台隐藏挂机
func (b *Bot) heartStop() {
	fmt.Println("隐藏界面开始挂机..", time.Now().Format("2006-01-02 15:04:05"))
	win.HideWindow(b.hwnd)
	time.Sleep(3 * time.Minute)
	win.ShowAndTopWindow(b.hwnd)
	time.Sleep(3 * time.Second)
}

func (b *Bot) Surrender() {
	win.ShowAndTopWindow(b.hwnd)
	fmt.Println("认输组合拳")
	win.MockClickWithoutOffset(b.hwnd, 677, 384)
	// b.heartStop()
	win.MockClickWithoutOffset(b.hwnd, 350, 160)
	fmt.Println("静待五秒...")
	time.Sleep(5)
	// 随便找两个地方点两下
	fmt.Println("随便找两个地方点两下.. ")
	x := utils.Rand(100, 700)
	y := utils.Rand(100, 400)
	win.MockClick(b.hwnd, x, y)
	win.HideWindow(b.hwnd)
	fmt.Println("静待十秒...")
	time.Sleep(10)
}
