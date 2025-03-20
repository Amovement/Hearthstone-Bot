package main

import (
	"github.com/Amovement/Hearthstone-Bot/internal"
	"github.com/Amovement/Hearthstone-Bot/pkg/win"
)

const title = "炉石传说"

func main() {
	bot := internal.NewBot(win.GetHwndByTitle(title))
	bot.Start()
}
