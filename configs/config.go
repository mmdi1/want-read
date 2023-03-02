package configs

import "github.com/lxn/win"

var (
	OLD_GWL_EXSTYLE int32 = 327937
	APP_HUND        win.HWND
	IS_HIDE         = false
	//当前阅读的书
	CurrentReadBook [][]rune
	CurrentPage     int
)
