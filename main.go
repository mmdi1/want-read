package main

import (
	"embed"
	"fmt"
	"os"
	"want-read/core/db"
	"want-read/external/read"
	"want-read/server/api"
	"want-read/server/api/ws"

	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"github.com/getlantern/systray"
	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var (
	//go:embed all:frontend/dist
	assets embed.FS
	title  = "app"
)

func onReady() {
	systray.SetTitle("Awesome App")
	bt, err := os.ReadFile("./favicon.ico")
	if err != nil {
		fmt.Println("read file error: ", err)
	}
	systray.SetIcon(bt)
	fmt.Println(".................", len(bt))
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
	db.InitLevelDb("./data")
	// Create an instance of the app structure
	app := NewApp()
	readApp := read.NewApp()
	// 设置托盘提示信息
	go systray.Run(onReady, onExit)
	r := gin.Default()
	r.Use(api.Cors())
	api.LocalUrl(r)
	wsGin := gin.New()
	wsGin.Use(api.Cors())
	wsGin.GET("/ws", ws.WsHandler)
	go wsGin.Run("127.0.0.1:8899")
	err := wails.Run(&options.App{
		Title:     title,
		Width:     600,
		Height:    400,
		Frameless: true, //边框
		// AlwaysOnTop: true, //是否最顶层
		AssetServer: &assetserver.Options{
			// Assets:  nil,
			Handler: r,
			Assets:  assets,
		},
		// BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},

		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  false,
			// DisableWindowIcon:                 true,
			// DisableFramelessWindowDecorations: true,
			// WebviewUserDataPath:               "",
			// Theme:                             windows.SystemDefault,
			// CustomTheme: &windows.ThemeSettings{
			// 	DarkModeTitleBar:   windows.RGB(20, 20, 20),
			// 	DarkModeTitleText:  windows.RGB(200, 200, 200),
			// 	DarkModeBorder:     windows.RGB(20, 0, 20),
			// 	LightModeTitleBar:  windows.RGB(200, 200, 200),
			// 	LightModeTitleText: windows.RGB(20, 20, 20),
			// 	LightModeBorder:    windows.RGB(200, 200, 200),
			// },
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
			readApp,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
