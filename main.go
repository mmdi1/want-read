package main

import (
	"embed"
	"log"
	"os"
	"want-read/core/db"
	"want-read/core/tray"
	"want-read/external/read"
	"want-read/external/setting"
	"want-read/server/api"
	"want-read/server/api/ws"

	"github.com/wailsapp/wails/v2/pkg/options/windows"

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

func main() {
	f, errs := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errs != nil {
		log.Fatal(errs)
	}
	defer f.Close()
	log.SetOutput(f)
	db.InitLevelDb("./data")
	app := NewApp()
	readApp := read.NewApp()
	settingApp := setting.NewApp()
	// 设置托盘提示信息
	go tray.Run()
	r := gin.Default()
	r.Use(api.Cors())
	api.LocalUrl(r)
	wsGin := gin.New()
	wsGin.Use(api.Cors())
	wsGin.GET("/ws", ws.WsHandler)
	go wsGin.Run("127.0.0.1:8899")
	err := wails.Run(&options.App{
		Title:       title,
		Width:       600,
		Height:      396,
		Frameless:   true, //边框
		AlwaysOnTop: true, //是否最顶层
		AssetServer: &assetserver.Options{
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
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
			readApp,
			settingApp,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
