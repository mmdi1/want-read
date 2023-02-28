package main

import (
	"context"
	"fmt"
	"syscall"

	"github.com/lxn/win"
)

// App struct
type App struct {
	ctx context.Context
}

var (
	OLD_GWL_EXSTYLE int32 = 327937
	hwnd            win.HWND
)

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}
func settingWinDefault() {
	win.SetWindowLong(hwnd, win.GWL_EXSTYLE, win.WS_EX_LAYERED)
}
func settingWinTransfer() {
	win.SetWindowLong(hwnd, win.GWL_EXSTYLE, win.WS_EX_LAYERED)
}

// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	str, err := syscall.UTF16PtrFromString(title)
	if err == nil {
		hwnd = win.FindWindow(nil, str)
	}
	settingWinTransfer()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ReceiveFilePath(path string) {

}
func (a *App) SettingDefaultWindows() {
	settingWinDefault()
}
