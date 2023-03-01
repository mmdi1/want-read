package main

import (
	"context"
	"fmt"
	"log"
	"syscall"
	"want-read/configs"

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
	str, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		log.Fatalln("err", err)
		return
	}
	configs.APP_HUND = win.FindWindow(nil, str)
	win.SetWindowLong(configs.APP_HUND, win.GWL_EXSTYLE, win.GetWindowLong(configs.APP_HUND, win.GWL_EXSTYLE)|win.WS_EX_LAYERED)
	// monitor.KeyBoardHandler()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
