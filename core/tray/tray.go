package tray

import (
	"log"
	"os"
	"want-read/configs"
	"want-read/core/db"
	"want-read/server/api/ws"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func onReady() {
	systray.SetTitle("Awesome App")
	bt, err := os.ReadFile("./icon.ico")
	if err != nil {
		log.Println("read file error: ", err)
	}
	systray.SetIcon(bt)
	openMain := systray.AddMenuItem("设置", "设置界面")
	quitMenu := systray.AddMenuItem("退出", "退出程序")
	go func() {
		for {
			select {
			case <-quitMenu.ClickedCh:
				ws.Conn.SendMsg(db.WsMsgModel{
					Id: db.WsID_Exit,
				})
				runtime.Quit(configs.APP_CTX)
			case <-openMain.ClickedCh:
				runtime.WindowSetSize(configs.APP_CTX, configs.APP_Width, configs.APP_Height)
				runtime.WindowCenter(configs.APP_CTX)
				ws.Conn.SendMsg(db.WsMsgModel{
					Id: db.WsID_Setting,
				})
			}
		}
	}()
}
func onExit() {
	println("click exit")
}

func Run() {
	systray.Run(onReady, onExit)
}
