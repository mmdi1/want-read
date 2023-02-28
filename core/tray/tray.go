package tray

import (
	"github.com/getlantern/systray"
)

type Tray struct {
	Title string
}

func InitTray() *Tray {
	systray.SetTitle("app")

	return &Tray{}
}
func (sef *Tray) AddMenuItem(name, tip string, cb func()) {
	// item := systray.AddMenuItem(name, tip)
}
func onExit()  {}
func onReady() {}
func Run() {
	systray.Run(onReady, onExit)
}
