package read

import (
	"fmt"
	"want-read/configs"

	"github.com/lxn/win"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

// GWL_EXSTYLE       //设置一个新的扩展窗口风格
// GWL_HINSTANCE //设置一个新的应用程序实例句柄
// GWL_ID //为窗口设置一个新的标识
// GWL_STYLE   //设置一个新的窗口风格
// GWL_USERDATA  //设置与窗口相关的32位值
// GWL_WNDPROC          //设置一个新的窗口过程
func (sef *App) ReadMod() {
	configs.IS_READ_MOD = !configs.IS_READ_MOD
	fmt.Println("aaaaaaaaaaaaaa", configs.IS_READ_MOD)
	if configs.IS_READ_MOD {
		num := win.SetWindowLong(configs.APP_HUND, win.GWL_EXSTYLE, win.GetWindowLong(configs.APP_HUND, win.GWL_EXSTYLE)|win.WS_EX_LAYERED)
		fmt.Println("zzzzzzzzzzzzzz", configs.APP_HUND, num)
		return
	}
	win.SetWindowLong(configs.APP_HUND, win.GWL_EXSTYLE, configs.OLD_GWL_EXSTYLE)
}
