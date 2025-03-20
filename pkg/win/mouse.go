package win

import (
	"fmt"
	"github.com/Amovement/Hearthstone-Bot/utils"
	"github.com/lxn/win"
	"math"
	"syscall"
	"time"
)

const (
	HearthstoneWindowWidth  = 700
	HearthstoneWindowHeight = 400
	MouseEventLeftDown      = 0x0002
	MouseEventLeftUp        = 0x0004
)

// LeftMouseClick 不清楚为什么没有生效, 可能是窗口做了过滤?
func LeftMouseClick(hwnd win.HWND, x, y int, cnt int) {
	width, height, err := GetWindowSize(uintptr(hwnd))
	if err != nil {
		fmt.Printf("获取窗口大小失败：%v\n", err)
		return
	}

	// 计算要点击的位置
	clickX := int32(float64(width) * float64(x) / HearthstoneWindowWidth)
	clickY := int32(float64(height) * float64(y) / HearthstoneWindowHeight)

	lParam := win.MAKELONG(uint16(clickX), uint16(clickY))
	count := 0
	for count < cnt {
		win.SendMessage(hwnd, win.WM_LBUTTONDOWN, win.MK_LBUTTON, uintptr(lParam))
		time.Sleep(200 * time.Millisecond)
		win.SendMessage(hwnd, win.WM_LBUTTONUP, 0, uintptr(lParam))
		count++
		time.Sleep(200 * time.Millisecond)
	}
}

// MockClick 带偏移的点击
func MockClick(hwndAddress uintptr, x, y int) {
	time.Sleep(time.Second)
	hwnd := syscall.Handle(hwndAddress)

	// 获取窗口的位置和大小
	rect, err := GetWindowRectByHandle(hwnd)
	if err != nil {
		fmt.Printf("获取窗口位置和大小失败：%+v\n", err)
		return
	}

	width, height, err := GetWindowSize(hwndAddress)
	if err != nil {
		fmt.Printf("获取窗口大小失败：%+v\n", err)
	}

	// 生成符合正态分布的偏移量
	mean := 0.0
	stdDev := 3.0
	offsetX := int32(math.Round(utils.RandNormal(mean, stdDev)))
	offsetY := int32(math.Round(utils.RandNormal(mean, stdDev)))

	// 计算要点击的位置
	clickX := rect.Left + int32(float64(width)*float64(x)/HearthstoneWindowWidth) + offsetX
	clickY := rect.Top + int32(float64(height)*float64(y)/HearthstoneWindowHeight) + offsetY

	// 移动鼠标到指定位置
	_, _, _ = setCursorPos.Call(uintptr(clickX), uintptr(clickY))

	// 模拟鼠标点击
	_, _, _ = mouseEvent.Call(
		uintptr(MouseEventLeftDown),
		0,
		0,
		0,
		0,
	)
	waitTime := utils.Rand(100, 900)
	time.Sleep(time.Duration(waitTime) * time.Millisecond)
	_, _, _ = mouseEvent.Call(
		uintptr(MouseEventLeftUp),
		0,
		0,
		0,
		0,
	)

	waitTime = utils.Rand(100, 900)
	time.Sleep(time.Duration(waitTime) * time.Millisecond)
	// fmt.Printf("点击 %d, %d\n", clickX, clickY)
}

// MockClickWithoutOffset 精确点击
func MockClickWithoutOffset(hwndAddress uintptr, x, y int) {
	time.Sleep(time.Second)
	hwnd := syscall.Handle(hwndAddress)

	// 获取窗口的位置和大小
	rect, err := GetWindowRectByHandle(hwnd)
	if err != nil {
		fmt.Printf("获取窗口位置和大小失败：%+v\n", err)
		return
	}

	width, height, err := GetWindowSize(hwndAddress)
	if err != nil {
		fmt.Printf("获取窗口大小失败：%+v\n", err)
	}

	// 计算要点击的位置
	clickX := rect.Left + int32(float64(width)*float64(x)/HearthstoneWindowWidth)
	clickY := rect.Top + int32(float64(height)*float64(y)/HearthstoneWindowHeight)

	// 移动鼠标到指定位置
	_, _, _ = setCursorPos.Call(uintptr(clickX), uintptr(clickY))

	// 模拟鼠标点击
	_, _, _ = mouseEvent.Call(
		uintptr(MouseEventLeftDown),
		0,
		0,
		0,
		0,
	)
	waitTime := utils.Rand(100, 900)
	time.Sleep(time.Duration(waitTime) * time.Millisecond)
	_, _, _ = mouseEvent.Call(
		uintptr(MouseEventLeftUp),
		0,
		0,
		0,
		0,
	)

	waitTime = utils.Rand(100, 900)
	time.Sleep(time.Duration(waitTime) * time.Millisecond)
	// fmt.Printf("点击 %d, %d\n", clickX, clickY)
}
