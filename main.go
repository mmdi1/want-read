package main

import (
	"embed"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func onReady() {
	systray.SetTitle("Awesome App")
	quitMenu := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		select {
		case <-quitMenu.ClickedCh:
			systray.Quit()
			return
		}
	}()
}
func onExit() {
	println("click exit")
}
func main() {
	// Create an instance of the app structure
	app := NewApp()
	// 设置托盘提示信息
	go systray.Run(onReady, onExit)
	// Sets the icon of a menu item. Only available on Mac and Windows.
	// Create application with options
	err := wails.Run(&options.App{
		Title:       "quiet-read",
		Width:       1024,
		Height:      100,
		Frameless:   true,  //边框
		AlwaysOnTop: false, //是否最顶层
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		Windows: &windows.Options{

			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			Theme:                             windows.SystemDefault,
			CustomTheme: &windows.ThemeSettings{
				DarkModeTitleBar:   windows.RGB(20, 20, 20),
				DarkModeTitleText:  windows.RGB(200, 200, 200),
				DarkModeBorder:     windows.RGB(20, 0, 20),
				LightModeTitleBar:  windows.RGB(200, 200, 200),
				LightModeTitleText: windows.RGB(20, 20, 20),
				LightModeBorder:    windows.RGB(200, 200, 200),
			},
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
