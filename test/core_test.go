package test

import (
	"github.com/Amovement/Hearthstone-Bot/internal"
	"github.com/Amovement/Hearthstone-Bot/pkg/win"
	"testing"
)

func TestBotCore(t *testing.T) {
	hwndByTitle := win.GetHwndByTitle("炉石传说")
	bot := internal.NewBot(hwndByTitle)
	bot.Start()
}

func TestSurrender(t *testing.T) {
	hwndByTitle := win.GetHwndByTitle("炉石传说")
	bot := internal.NewBot(hwndByTitle)
	bot.Surrender()
}
