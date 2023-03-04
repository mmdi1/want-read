package main

import (
	"context"
	"log"
	"syscall"
	"want-read/configs"
	"want-read/core/monitor"

	"github.com/getlantern/systray"

	"github.com/lxn/win"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	configs.APP_CTX = ctx
	str, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		log.Println("app err:", err)
		return
	}
	configs.APP_HUND = win.FindWindow(nil, str)
	win.SetWindowLong(configs.APP_HUND, win.GWL_EXSTYLE, win.GetWindowLong(configs.APP_HUND, win.GWL_EXSTYLE)|win.WS_EX_LAYERED)
	monitor.KeyBoardHandler()
}

func (a *App) shutdown(ctx context.Context) {
	log.Println("app shutdown")
	systray.Quit()
}
