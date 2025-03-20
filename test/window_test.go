package test

import (
	"github.com/Amovement/Hearthstone-Bot/pkg/win"
	lxnWin "github.com/lxn/win"
	"testing"
)

func TestSetWindowSize(t *testing.T) {
	hwndByTitle := win.GetHwndByTitle("炉石传说")
	win.SetWindowSize(hwndByTitle, 700, 400)
}

func TestGetHwndByTitle(t *testing.T) {
	hwndByTitle := win.GetHwndByTitle("炉石传说")
	t.Log(hwndByTitle)
}

func TestGetWindowSize(t *testing.T) {
	hwndByTitle := win.GetHwndByTitle("炉石传说")
	width, height, err := win.GetWindowSize(hwndByTitle)
	if err != nil {
		t.Error(err)
	}
	t.Log(width, height)
}

func TestLeftMouseClick(t *testing.T) {
	hwndByTitle := win.GetHwndByTitle("炉石传说")
	win.LeftMouseClick(lxnWin.HWND(hwndByTitle), 495, 327, 1)
}

func TestMockClick(t *testing.T) {
	hwndByTitle := win.GetHwndByTitle("炉石传说")
	win.MockClick(hwndByTitle, 495, 327)
}

func TestShowWindow(t *testing.T) {
	hwndByTitle := win.GetHwndByTitle("炉石传说")
	win.ShowWindow(hwndByTitle)
}

func TestHideWindow(t *testing.T) {
	hwndByTitle := win.GetHwndByTitle("炉石传说")
	win.HideWindow(hwndByTitle)
}

func TestSetTopWindow(t *testing.T) {
	hwndByTitle := win.GetHwndByTitle("炉石传说")
	win.SetTopWindow(hwndByTitle)
}

func TestShowAndTopWindow(t *testing.T) {
	hwndByTitle := win.GetHwndByTitle("炉石传说")
	win.ShowAndTopWindow(hwndByTitle)
}
